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

	// ErrCreatePostIsNil error
	ErrCreateUserIsNil = errors.New("blog.users.user.create.user.is_nil")

	// ErrCreateUserIDPresent user missing ID error
	ErrCreateUserIDPresent = errors.New("blog.users.user.create.user_id.present")

	// ErrCreateUserNameMissing user missing Name error
	ErrCreateUserNameMissing = errors.New("blog.users.user.create.missing.name")

	// ErrUpdatePostIsNil error
	ErrUpdateUserIsNil = errors.New("blog.users.user.update.user.is_nil")

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

	// ErrGetUserByAPIKeyMissingAPIKey error message
	ErrGetUserByAPIKeyMissingAPIKey = errors.New("blog.users.user.get_by_api_key.missing.api_key")

	// ErrGetUserIDByAPIKeyMissingAPIKey error message
	ErrGetUserIDByAPIKeyMissingAPIKey = errors.New("blog.users.user.get_user_id_by_api_key.missing.api_key")

	// ErrGetAPIKeyForUserIDMissingUserID error message
	ErrGetAPIKeyForUserIDMissingUserID = errors.New("blog.users.user.get_api_key_for_user_id.missing.user_id")

	// ErrGetUserNotFound  error message
	ErrGetUserNotFound = errors.New("blog.users.user.not_found")

	// Posts
	//

	// ErrCreatePostIsNil error
	ErrCreatePostIsNil = errors.New("blog.users.user.create.user.is_nil")

	// ErrGetPostsByIDsMissingMinimumIDs error message
	ErrGetPostsByIDsMissingMinimumIDs = errors.New("blog.posts.post.get_by_ids.missing.minimum_one_id")

	// ErrCreatePostIDPresent error message
	ErrCreatePostIDPresent = errors.New("blog.posts.create.post_id.present")

	// ErrCreatePostTitleMissing error message
	ErrCreatePostTitleMissing = errors.New("blog.posts.create.missing.title")

	// ErrCreatePostContentMissing error message
	ErrCreatePostContentMissing = errors.New("blog.posts.create.missing.content")

	// ErrUpdatePostIsNil error
	ErrUpdatePostIsNil = errors.New("blog.posts.update.post.is_nil")

	// ErrUpdatePostIDMissing error message
	ErrUpdatePostIDMissing = errors.New("blog.posts.update.missing.id")

	// ErrUpdatePostTitleMissing error message
	ErrUpdatePostTitleMissing = errors.New("blog.posts.update.missing.title")

	// ErrUpdatePostContentMissing error message
	ErrUpdatePostContentMissing = errors.New("blog.posts.update.missing.content")

	// ErrUpdatePostUpdatedByMissing error message
	ErrUpdatePostUpdatedByMissing = errors.New("blog.posts.update.missing.updated_by")

	// ErrUpdatePostInvalidUser error message
	ErrUpdatePostInvalidUser = errors.New("blog.posts.update.invalid.updated_by")

	// ErrDeletePostInvalidUser error message
	ErrDeletePostInvalidUser = errors.New("blog.posts.post.delete.user_not_permitted")

	// ErrDeletePostMissingID error message
	ErrDeletePostMissingID = errors.New("blog.posts.delete.missing.id")

	// ErrDeletePostMissingDeletedByID error message
	ErrDeletePostMissingDeletedByID = errors.New("blog.posts.delete.missing.deleted_by_id")

	// ErrDeletePostInvalidID
	ErrDeletePostInvalidID = errors.New("blog.posts.delete.invalid.id")

	// ErrGetPostNotFound error message
	ErrGetPostNotFound = errors.New("blog.posts.post.not_found")
)

// ErrReplacer allows for static errors to have dynamic values
func ErrReplacer(err error, replacers ...interface{}) error {
	return fmt.Errorf(err.Error(), replacers...)
}
