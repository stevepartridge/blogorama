package blog

import (
	"database/sql"
	"strings"

	"github.com/stevepartridge/blogorama/pkg/log"
	"github.com/stevepartridge/blogorama/pkg/uuid"
	"github.com/stevepartridge/db"
)

// Create a new User by passing in a reference that excludes an ID
//  required:
//    Name
func (store MySQL) CreateUser(user *User) error {
	log.Debug().Interface("user", user).Msg("CreateUser")

	if user.ID > 0 {
		return ErrCreateUserIDPresent
	}

	user.Name = strings.Trim(user.Name, " ")

	if user.Name == "" {
		return ErrCreateUserNameMissing
	}

	user.APIKey = uuid.Generate()

	conn := db.Connx(mysqlDbID)

	result, err := conn.Exec(`
    INSERT INTO `+usersTableName+` SET
      api_key      = ?,
      name         = ?,
      active       = 1,
      created_by   = ?,
      created_at   = NOW(6)
    `,
		user.APIKey,
		user.Name,
		user.CreatedByID,
	)

	if err != nil {
		return err
	}

	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	u, err := store.GetUserByID(int(id))
	if err != nil {
		return err
	}

	if u.ID > 0 {
		*user = u
	}

	// Emit event for user created here

	return err

}

// Update an existing User and modify the reference to reflect the updated info in the database
//  required:
//    ID
//    Name
//    UpdatedByID
func (store MySQL) UpdateUser(user *User) error {
	log.Debug().Interface("user", user).Msg("UpdateUser")

	if user.ID == 0 {
		return ErrUpdateUserIDMissing
	}

	u, err := store.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	if u.ID == 0 {
		return ErrUpdateUserInvalidID
	}

	user.Name = strings.Trim(user.Name, " ")

	if user.Name == "" {
		return ErrUpdateUserNameMissing
	}

	if user.UpdatedByID == 0 {
		return ErrUpdateUserUpdatedByMissing
	}

	conn := db.Connx(mysqlDbID)

	_, err = conn.Exec(`
    UPDATE `+usersTableName+` SET
      name        = ?,
      active      = ?,
      updated_by  = ?,
      updated_at  = NOW(6)
    WHERE
      id = ?
    `,
		user.Name,
		user.Active,
		user.UpdatedByID,
		user.ID,
	)

	if err != nil {
		return err
	}

	u, err = store.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	if u.ID > 0 {
		*user = u
	}

	// Emit event for user updated here

	return err

}

// GetUserByID retreives a user by their ID value
func (store MySQL) GetUserByID(id int) (User, error) {
	log.Debug().Int("user_id", id).Msg("GetUserByID")

	conn := db.Connx(mysqlDbID)

	rows, err := conn.Queryx(`
    SELECT
      `+userSelectSQL+`
    FROM
      `+usersTableName+`
    WHERE
      id = ?
    LIMIT 1`,
		id)

	if err != nil {
		return User{}, err
	}

	defer rows.Close()

	return parseUserFromRows(rows), err

}

// GetUserByAPIKey retreives a user by their APIKey value
func (store MySQL) GetUserByAPIKey(apiKey string) (User, error) {
	log.Debug().Str("api_key", apiKey).Msg("GetUserByAPIKey")

	if strings.Trim(apiKey, " ") == "" {
		return User{}, ErrGetUserByAPIKeyMissingAPIKey
	}

	conn := db.Connx(mysqlDbID)

	rows, err := conn.Queryx(`
    SELECT
      `+userSelectSQL+`
    FROM
      `+usersTableName+`
    WHERE
      api_key = ?
    LIMIT 1`,
		apiKey)

	if err != nil {
		return User{}, err
	}

	defer rows.Close()

	return parseUserFromRows(rows), err

}

// GetAPIKeyForUserID retreives the APIKey for User ID
func (store MySQL) GetAPIKeyForUserID(userID int) (string, error) {
	log.Debug().Int("user_id", userID).Msg("GetAPIKeyForUserID")

	if userID == 0 {
		return "", ErrGetAPIKeyForUserIDMissingUserID
	}

	conn := db.Connx(mysqlDbID)

	rows, err := conn.Queryx(`
    SELECT
      api_key
    FROM
      `+usersTableName+`
    WHERE
      id = ?
    LIMIT 1`,
		userID)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var key sql.NullString
	for rows.Next() {
		err = rows.Scan(&key)
		if err != nil {
			return "", err
		}
	}

	if !key.Valid {
		return "", nil
	}

	if key.String != "" {
		return key.String, nil
	}

	return "", nil

}

// GetUserIDByAPIKey retreives a user's ID by their APIKey value
func (store MySQL) GetUserIDByAPIKey(apiKey string) (int, error) {
	log.Debug().Str("api_key", apiKey).Msg("GetUserIDByAPIKey")

	if strings.Trim(apiKey, " ") == "" {
		return 0, ErrGetUserIDByAPIKeyMissingAPIKey
	}

	conn := db.Connx(mysqlDbID)

	rows, err := conn.Queryx(`
    SELECT
      id
    FROM
      `+usersTableName+`
    WHERE
      api_key = ?
    LIMIT 1`,
		apiKey)
	defer rows.Close()

	if err != nil {
		return 0, err
	}

	var id sql.NullInt64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	if !id.Valid {
		return 0, nil
	}

	if id.Int64 > 0 {
		return int(id.Int64), nil
	}

	return 0, nil
}

// DeleteUser removes a user from the database
func (store MySQL) DeleteUser(id, deletedBy int) error {
	log.Debug().Int("id", id).Int("deleted_by", deletedBy).Msg("DeleteUser")

	if id == 0 {
		return ErrDeleteUserMissingID
	}

	if deletedBy == 0 {
		return ErrDeleteUserMissingDeletedByID
	}

	conn := db.Connx(mysqlDbID)

	// Issue an update with the deleted by in order to have an audit log in the history table
	_, err := conn.Exec(`
    UPDATE `+usersTableName+` SET
      active     = 0,
      updated_by = ?
    WHERE
      id   = ?
    `,
		id,
		deletedBy,
	)

	if err != nil {
		return err
	}

	// Actually remove the user from the database
	_, err = conn.Exec(`
    DELETE FROM
      `+usersTableName+`
    WHERE
      id   = ?
    `,
		id,
	)

	if err != nil {
		return err
	}

	return nil

}
