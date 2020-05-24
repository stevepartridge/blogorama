package service

import (
	"errors"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	blog "github.com/stevepartridge/blogorama"
	"github.com/stevepartridge/blogorama/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrRequestIsNil   = errors.New("blog.request.is_nil")
	ErrInvalidRequest = errors.New("blog.bad_request")
	ErrNotFound       = errors.New("blog.not_found")

	ErrResourceConnectionFailure = errors.New("blog.resource.connection_failed")
)

func handleError(err error) error {
	if strings.Contains(err.Error(), "dial tcp") {
		log.IfError(err)
		err = ErrResourceConnectionFailure
	}
	return status.Error(codeFromError(err), err.Error())
}

func codeFromError(err error) codes.Code {

	code := codes.Unknown

	switch err {
	case
		blog.ErrDatastoreUnsupportedOrMissingType,
		blog.ErrCreateUserIDPresent,
		blog.ErrCreateUserNameMissing,
		blog.ErrUpdateUserIDMissing,
		blog.ErrUpdateUserInvalidID,
		blog.ErrUpdateUserNameMissing,
		blog.ErrUpdateUserUpdatedByMissing,
		blog.ErrDeleteUserMissingID,
		blog.ErrDeleteUserMissingDeletedByID,
		blog.ErrGetUserByAPIKeyMissingAPIKey,
		blog.ErrGetUserIDByAPIKeyMissingAPIKey,
		blog.ErrGetAPIKeyForUserIDMissingUserID,
		blog.ErrGetPostsByIDsMissingMinimumIDs,
		blog.ErrCreatePostIDPresent,
		blog.ErrCreatePostTitleMissing,
		blog.ErrCreatePostContentMissing,
		blog.ErrUpdatePostIDMissing,
		blog.ErrUpdatePostTitleMissing,
		blog.ErrUpdatePostContentMissing,
		blog.ErrUpdatePostUpdatedByMissing,
		blog.ErrDeletePostMissingID,
		blog.ErrDeletePostMissingDeletedByID,
		ErrRequestIsNil:

		code = codes.InvalidArgument

	case
		blog.ErrGetPostNotFound,
		blog.ErrGetUserNotFound,
		ErrNotFound:

		code = codes.NotFound
	}

	if code == codes.Unknown {
		if strings.Contains(err.Error(), "not_found") {
			code = codes.NotFound
		}

	}

	return code

}

func httpStatusFromError(err error) int {
	return runtime.HTTPStatusFromCode(codeFromError(err))
}
