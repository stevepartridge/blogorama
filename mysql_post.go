package blog

import (
	"strings"

	"github.com/stevepartridge/blogorama/pkg/log"
	"github.com/stevepartridge/db"
)

// Create a new Post by passing in a reference that excludes an ID
//  required:
//    Title
//    Content
//    CreatedByID
func (store MySQL) CreatePost(post *Post) error {
	log.Debug().Interface("post", post).Msg("CreatePost")

	if post.ID > 0 {
		return ErrCreatePostIDPresent
	}

	post.Title = strings.Trim(post.Title, " ")

	if post.Title == "" {
		return ErrCreatePostTitleMissing
	}

	if post.Content == "" {
		return ErrCreatePostContentMissing
	}

	conn := db.Connx(mysqlDbID)

	result, err := conn.Exec(`
    INSERT INTO `+postsTableName+` SET
      title        = ?,
      content      = ?,
      active       = 1,
      created_by   = ?,
      created_at   = NOW(6)
    `,
		post.Title,
		post.Content,
		post.CreatedByID,
	)

	if err != nil {
		return err
	}

	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	u, err := store.GetPostByID(int(id))
	if err != nil {
		return err
	}

	if u.ID > 0 {
		*post = u
	}

	// Emit event for post created here

	return err

}

// Update an existing Post and modify the reference to reflect the updated info in the database
//  required:
//    ID
//    Title
//    Content
//    UpdatedByID
func (store MySQL) UpdatePost(post *Post) error {
	log.Debug().Interface("post", post).Msg("UpdatePost")

	if post.ID == 0 {
		return ErrUpdatePostIDMissing
	}

	post.Title = strings.Trim(post.Title, " ")

	if post.Title == "" {
		return ErrUpdatePostTitleMissing
	}

	if post.Content == "" {
		return ErrUpdatePostContentMissing
	}

	if post.UpdatedByID == 0 {
		return ErrUpdatePostUpdatedByMissing
	}

	conn := db.Connx(mysqlDbID)

	_, err := conn.Exec(`
    UPDATE `+postsTableName+` SET
      title       = ?,
      content     = ?,
      private     = ?,
      updated_by  = ?,
      updated_at  = NOW(6)
    WHERE
      id = ?
    `,
		post.Title,
		post.Content,
		post.Private,
		post.UpdatedByID,
		post.ID,
	)

	if err != nil {
		return err
	}

	p, err := store.GetPostByID(post.ID)
	if err != nil {
		return err
	}

	if p.ID > 0 {
		*post = p
	}

	// Emit event for post updated here

	return err

}

// GetPostByID retreives a post by their ID value
func (store MySQL) GetPostByID(id int) (Post, error) {
	log.Debug().Int("post_id", id).Msg("GetPostByID")

	conn := db.Connx(mysqlDbID)

	rows, err := conn.Queryx(`
    SELECT
      `+postSelectSQL+`
    FROM
      `+postsTableName+`
    WHERE
      id = ?
    LIMIT 1`,
		id)

	if err != nil {
		return Post{}, err
	}

	defer rows.Close()

	return parsePostFromRows(rows), err

}

// DeletePost removes a post from the database
func (store MySQL) DeletePost(id, deletedBy int) error {

	if id == 0 {
		return ErrDeletePostMissingID
	}

	if deletedBy == 0 {
		return ErrDeletePostMissingDeletedByID
	}

	conn := db.Connx(mysqlDbID)

	// Issue an update with the deleted by in order to have an audit log in the history table
	_, err := conn.Exec(`
    UPDATE `+postsTableName+` SET
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

	// Actually remove the post from the database
	_, err = conn.Exec(`
    DELETE FROM
      `+postsTableName+`
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
