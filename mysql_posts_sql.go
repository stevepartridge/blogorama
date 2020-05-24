package blog

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	dbUtils "github.com/stevepartridge/db/utils"

	"github.com/stevepartridge/blogorama/pkg/log"
)

const (
	postsTableName = "posts"

	postSelectSQL = `
    ` + postsTableName + `.id,
    ` + postsTableName + `.title,
    ` + postsTableName + `.content,
		` + postsTableName + `.private,
    ` + postsTableName + `.updated_by,
    ` + postsTableName + `.updated_at,
    ` + postsTableName + `.created_by,
    ` + postsTableName + `.created_at
  `
)

type postResult struct {
	ID          int            `db:"id"`
	Title       string         `db:"title"`
	Content     string         `db:"content"`
	Private     bool           `db:"private"`
	UpdatedByID sql.NullInt64  `db:"updated_by"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	CreatedByID int            `db:"created_by"`
	CreatedAt   sql.NullString `db:"created_at"`
}

func (u postResult) ToPost() Post {

	post := Post{
		ID:          u.ID,
		Title:       u.Title,
		Content:     u.Content,
		Private:     u.Private,
		UpdatedByID: int(u.UpdatedByID.Int64),
		UpdatedAt:   dbUtils.NullStringToTime(&u.UpdatedAt),
		CreatedByID: u.CreatedByID,
		CreatedAt:   dbUtils.NullStringToTime(&u.CreatedAt),
	}

	return post
}

func parsePostFromRows(rows *sqlx.Rows) Post {

	var result postResult

	for rows.Next() {
		err := rows.StructScan(&result)
		if !log.IfError(err) {
			return result.ToPost()
		}
	}

	return Post{}
}

func parsePostsFromRows(rows *sqlx.Rows) []Post {

	list := []Post{}

	for rows.Next() {
		var result postResult
		err := rows.StructScan(&result)
		if !log.IfError(err) {
			list = append(list, result.ToPost())
		}
	}

	return list
}
