package blog

import (
	"database/sql"

	"github.com/stevepartridge/blogorama/pkg/log"
	"github.com/stevepartridge/blogorama/pkg/utils"
	"github.com/stevepartridge/db"
)

// GetUsers retrieves a list of Users with optional options
//   (optional)
//   opts[0] = limit, default: 25
//   opts[1] = offset, default: 0
func (store MySQL) GetUsers(opts ...int) ([]User, ResultsInfo, error) {
	log.Debug().Interface("options", opts).Msg("GetUsers")

	var (
		resultsInfo = ResultsInfo{
			Limit:  25,
			Offset: 0,
		}
	)

	if len(opts) > 0 {
		resultsInfo.Limit = opts[0]
		if len(opts) > 1 {
			resultsInfo.Offset = opts[1]
		}
	}

	conn := db.Connx(mysqlDbID)

	query := `
    SELECT
      ` + userSelectSQL + `
    FROM
      ` + usersTableName + `
    `

	params := []interface{}{}

	if resultsInfo.Limit > 0 {
		query = query + ` LIMIT `
		if resultsInfo.Offset > 0 {
			query = query + ` ?, `
			params = append(params, resultsInfo.Offset)
		}
		params = append(params, resultsInfo.Limit)
		query = query + ` ? `
	}

	rows, err := conn.Queryx(query, params...)
	log.IfError(err)

	if err != nil {
		return []User{}, resultsInfo, err
	}

	defer rows.Close()

	resultsInfo.Total, err = store.GetTotalUsers()
	if err != nil {
		return []User{}, resultsInfo, err
	}

	return parseUsersFromRows(rows), resultsInfo, nil

}

// GetUsersByIDs retrieves a list of User/s based on provided IDs
func (store MySQL) GetUsersByIDs(ids ...int) ([]User, error) {
	log.Debug().Ints("ids", ids).Msg("GetUsersByIDs")

	if len(ids) == 0 {
		return []User{}, nil
	}

	idList := utils.IntSliceToSeparatedString(ids, ",")

	conn := db.Connx(mysqlDbID)

	query := `
    SELECT
      ` + userSelectSQL + `
    FROM
			` + usersTableName + `
		WHERE
			id IN (` + idList + `)
    `

	rows, err := conn.Queryx(query)
	log.IfError(err)

	if err != nil {
		return []User{}, err
	}

	defer rows.Close()

	return parseUsersFromRows(rows), nil

}

// GetTotalUsers retreives the total users count
func (store MySQL) GetTotalUsers() (int, error) {
	log.Debug().Msg("GetTotalUsers")

	conn := db.Connx(mysqlDbID)

	rows, err := conn.Queryx(`
    SELECT
      COUNT(id)
    FROM
			` + usersTableName + `
		`)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var count sql.NullInt64
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	if !count.Valid {
		return 0, nil
	}

	if count.Int64 > 0 {
		return int(count.Int64), nil
	}

	return 0, nil

}
