package blog

import (
	"database/sql"

	"github.com/stevepartridge/blogorama/pkg/log"
	"github.com/stevepartridge/blogorama/pkg/utils"
	"github.com/stevepartridge/db"
)

// GetPosts retrieves a list of Posts with optional options
//   (optional)
//   opts[0] = limit, default: 25
//   opts[1] = offset, default: 0
//   opts[2] = created_by, default: 0
func (store MySQL) GetPosts(opts ...int) ([]Post, ResultsInfo, error) {
	log.Debug().Interface("options", opts).Msg("GetPosts")

	var (
		resultsInfo = ResultsInfo{
			Limit:  25,
			Offset: 0,
		}
		createdBy = 0
		whereSQL  = ""
	)

	if len(opts) > 0 {
		resultsInfo.Limit = opts[0]
		if len(opts) > 1 {
			resultsInfo.Offset = opts[1]

			if len(opts) > 2 {
				if opts[2] > 0 {
					createdBy = opts[2]
					whereSQL = " WHERE created_by = ? "
				}
			}
		}

	}

	conn := db.Connx(mysqlDbID)

	query := `
    SELECT
      ` + postSelectSQL + `
    FROM
			` + postsTableName + `
		` + whereSQL + `
    `

	params := []interface{}{}

	if createdBy > 0 {
		params = append(params, createdBy)
	}

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
		return []Post{}, resultsInfo, err
	}

	defer rows.Close()

	// Handle createdBy if present in total lookup
	ids := []int{}
	if createdBy > 0 {
		ids = append(ids, createdBy)
	}
	resultsInfo.Total, err = store.GetTotalPosts(ids...)

	return parsePostsFromRows(rows), resultsInfo, nil

}

// GetPostsByCreatedByID is a helper method to retrieve posts by specific Post ID
//   (optional)
//   opts[0] = limit, default: 25
//   opts[1] = offset, default: 0
func (store MySQL) GetPostsByCreatedByID(createdBy int, opts ...int) ([]Post, ResultsInfo, error) {

	// If opts is empty pad with zero values
	if len(opts) == 0 {
		opts = []int{0, 0, createdBy}
	}

	// If only limit is present pad with offset as zero
	if len(opts) == 1 {
		opts = append(opts, []int{0, createdBy}...)
	}

	return store.GetPosts(opts...)
}

// GetPostsByIDs retrieves a list of Post/s based on provided IDs
func (store MySQL) GetPostsByIDs(ids ...int) ([]Post, error) {
	log.Debug().Ints("ids", ids).Msg("GetPostsByIDs")

	if len(ids) == 0 {
		return []Post{}, ErrGetPostsByIDsMissingMinimumIDs
	}

	idList := utils.IntSliceToSeparatedString(ids, ",")

	conn := db.Connx(mysqlDbID)

	query := `
    SELECT
      ` + postSelectSQL + `
    FROM
			` + postsTableName + `
		WHERE
			id IN (` + idList + `)
    `

	rows, err := conn.Queryx(query)
	log.IfError(err)

	if err != nil {
		return []Post{}, err
	}

	defer rows.Close()

	return parsePostsFromRows(rows), nil

}

// GetTotalPosts retreives the total posts count
// optional
//	CreatedByID(s)
func (store MySQL) GetTotalPosts(createdBy ...int) (int, error) {
	log.Debug().Ints("created_by", createdBy).Msg("GetTotalPosts")

	whereSQL := ""

	if len(createdBy) > 0 {
		ids := utils.IntSliceToSeparatedString(createdBy, ",")
		whereSQL = ` WHERE created_by IN (` + ids + `) `
	}

	conn := db.Connx(mysqlDbID)

	rows, err := conn.Queryx(`
    SELECT
      COUNT(id)
    FROM
			` + postsTableName + `
		` + whereSQL + `
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
