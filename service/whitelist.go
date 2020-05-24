package service

import "github.com/stevepartridge/blogorama/pkg/whitelist"

// setWhitelist establishes the fixed methods that are open
func setWhitelist() {

	whitelist.AddMethod("/blog.Blog/Info")

	// Normally this wouldn't be open, but for this service we'll do it for easy user creation
	whitelist.AddMethod("/blog.Blog/CreateUser")

}
