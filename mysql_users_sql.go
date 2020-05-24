package blog

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	dbUtils "github.com/stevepartridge/db/utils"

	"github.com/stevepartridge/blogorama/pkg/log"
)

const (
	usersTableName = "users"

	userSelectSQL = `
    ` + usersTableName + `.id,
    ` + usersTableName + `.api_key,
    ` + usersTableName + `.name,
		` + usersTableName + `.active,
    ` + usersTableName + `.updated_by,
    ` + usersTableName + `.updated_at,
    ` + usersTableName + `.created_by,
    ` + usersTableName + `.created_at
  `
)

type userResult struct {
	ID          int            `db:"id"`
	UUID        string         `db:"api_key"`
	Name        string         `db:"name"`
	Active      bool           `db:"active"`
	UpdatedByID sql.NullInt64  `db:"updated_by"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	CreatedByID int            `db:"created_by"`
	CreatedAt   sql.NullString `db:"created_at"`
}

func (u userResult) ToUser() User {

	user := User{
		ID:          u.ID,
		Name:        u.Name,
		Active:      u.Active,
		UpdatedByID: int(u.UpdatedByID.Int64),
		UpdatedAt:   dbUtils.NullStringToTime(&u.UpdatedAt),
		CreatedByID: u.CreatedByID,
		CreatedAt:   dbUtils.NullStringToTime(&u.CreatedAt),
	}

	return user
}

func parseUserFromRows(rows *sqlx.Rows) User {

	var result userResult

	for rows.Next() {
		err := rows.StructScan(&result)
		if !log.IfError(err) {
			return result.ToUser()
		}
	}

	return User{}
}

func parseUsersFromRows(rows *sqlx.Rows) []User {

	list := []User{}

	for rows.Next() {
		var result userResult
		err := rows.StructScan(&result)
		if !log.IfError(err) {
			list = append(list, result.ToUser())
		}
	}

	return list
}
