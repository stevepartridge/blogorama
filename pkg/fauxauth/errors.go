package fauxauth

import "errors"

var (
	// ErrNotAllowed message for not allowed
	ErrNotAllowed = errors.New("auth.not_allowed")

	// ErrMissingAPIKey missing api key error
	ErrMissingAPIKey = errors.New("auth.missing_api_key")

	// ErrUnableToValidate unable to validate key error, can be ambiguous to the actual cause
	ErrUnableToValidate = errors.New("auth.unable_to_validate")

	// ErrAPIKeyNotValid api key not valid error
	ErrAPIKeyNotValid = errors.New("auth.api_key_not_valid")
)
