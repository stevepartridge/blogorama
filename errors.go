package blog

import (
	"errors"
	"fmt"
)

var (
	// ErrDatastoreUnsupportedOrMissingType invalid datastore type error
	ErrDatastoreUnsupportedOrMissingType = errors.New("blog.datastore.type.missing_or_unsupported.%s")

	// Users
	//

	// ErrCreateUserIDPresent user missing ID error
	ErrCreateUserIDPresent = errors.New("blog.users.user.create.user_id.present")
	// ErrCreateUserNameMissing user missing Name error
	ErrCreateUserNameMissing = errors.New("blog.users.user.create.missing.name")

	// ErrUpdateUserIDMissing user missing ID error
	ErrUpdateUserIDMissing = errors.New("blog.users.user.update.missing.id")
	// ErrUpdateUserInvalidID user invalid ID error
	ErrUpdateUserInvalidID = errors.New("blog.users.user.update.invalid.id")
	// ErrUpdateUserNameMissing user missing Name error
	ErrUpdateUserNameMissing = errors.New("blog.users.user.update.missing.name")
	// ErrUpdateUserUpdatedByMissing user missing UpdatedByID error
	ErrUpdateUserUpdatedByMissing = errors.New("blog.users.user.update.missing.updated_by")

	// ErrDeleteUserMissingID user missing ID error
	ErrDeleteUserMissingID = errors.New("blog.users.user.delete.missing.id")
	// ErrDeleteUserMissingDeletedByID user missing deletedBy error
	ErrDeleteUserMissingDeletedByID = errors.New("blog.users.user.delete.missing.deleted_by")

	// ErrGetUserByAPIKeyMissingAPIKey
	ErrGetUserByAPIKeyMissingAPIKey = errors.New("blog.users.user.get_by_api_key.missing.api_key")

	// ErrGetUserIDByAPIKeyMissingAPIKey
	ErrGetUserIDByAPIKeyMissingAPIKey = errors.New("blog.users.user.get_user_id_by_api_key.missing.api_key")

	// ErrGetAPIKeyForUserIDMissingUserID
	ErrGetAPIKeyForUserIDMissingUserID = errors.New("blog.users.user.get_api_key_for_user_id.missing.user_id")

	// ErrGetUserNotFound error
	ErrGetUserNotFound = errors.New("blog.users.user.not_found")

	// Posts
	//

	// ErrGetPostsByIDsMissingMinimumIDs
	ErrGetPostsByIDsMissingMinimumIDs = errors.New("blog.posts.post.get_by_ids.missing.minimum_one_id")

	// ErrCreatePostIDPresent
	ErrCreatePostIDPresent = errors.New("blog.posts.create.post_id.present")

	// ErrCreatePostTitleMissing
	ErrCreatePostTitleMissing = errors.New("blog.posts.create.missing.title")

	// ErrCreatePostContentMissing
	ErrCreatePostContentMissing = errors.New("blog.posts.create.missing.content")

	// ErrUpdatePostIDMissing
	ErrUpdatePostIDMissing = errors.New("blog.posts.update.missing.id")

	// ErrUpdatePostTitleMissing
	ErrUpdatePostTitleMissing = errors.New("blog.posts.update.missing.title")

	// ErrUpdatePostContentMissing
	ErrUpdatePostContentMissing = errors.New("blog.posts.update.missing.content")

	// ErrUpdatePostUpdatedByMissing
	ErrUpdatePostUpdatedByMissing = errors.New("blog.posts.update.missing.updated_by")

	// ErrDeletePostMissingID
	ErrDeletePostMissingID = errors.New("blog.posts.delete.missing.id")

	// ErrDeletePostMissingDeletedByID
	ErrDeletePostMissingDeletedByID = errors.New("blog.posts.delete.missing.deleted_by_id")

	// ErrGetPostNotFound error
	ErrGetPostNotFound = errors.New("blog.posts.post.not_found")
)

// ErrReplacer allows for static errors to have dynamic values
func ErrReplacer(err error, replacers ...interface{}) error {
	return fmt.Errorf(err.Error(), replacers...)
}
