package fauxauth

import blog "github.com/stevepartridge/blogorama"

// FauxAuth is a simulated Auth mechinism. Ideally it'd be something more robust with OAuth2
// credentials and token to better support both humans and machines with a common standard
type FauxAuth struct{}

// New sets up this basic faux auth.
func New() (FauxAuth, error) {
	return FauxAuth{}, nil
}

// ValidateKey does a very basic validation of the key. It makes the assumption
// that if a user is found with the key then it's valid and in good standing.
// It also makes the basic assumption that if the user is not active then it's invalid.
func (f *FauxAuth) ValidateKey(key string) (blog.User, error) {

	if blog.Store == nil {
		return blog.User{}, ErrUnableToValidate
	}

	// There should be a caching layer to maintain an active list of keys and users
	// somwthing that could listen to system events to effectively update users when they're
	// deactivated, deleted, etc. without having to check at middleware request time
	// Instead, this is a per request check to keep this simplified for this service's use case
	user, err := blog.Store.GetUserByAPIKey(key)
	if err != nil {
		return blog.User{}, err
	}

	if !user.Active {
		return blog.User{}, ErrAPIKeyNotValid
	}

	return user, nil

}
