package fauxauth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/stevepartridge/blogorama/pkg/middleware"
	"github.com/stevepartridge/blogorama/pkg/whitelist"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Interceptor handles the validation of all methods that are not previously whitelisted
// with the use of a basic X-API-Key header check
func (f *FauxAuth) Interceptor() func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		if !whitelist.MethodAllowed(info.FullMethod) {

			key := middleware.GetAPIKeyFromContext(ctx)
			if key == "" {
				return ctx, status.Error(codes.Unauthenticated, ErrMissingAPIKey.Error())
			}

			user, err := f.ValidateKey(key)
			if err != nil {
				return ctx, status.Error(codes.Unauthenticated, err.Error())
			}

			// Add the User.ID to the the context for easy reference down the chain
			ctx = middleware.MergeStringValueIntoContextFromMetadata(ctx, middleware.UserIDKey, strconv.Itoa(user.ID))

		}

		return handler(ctx, req)
	}
}

// HTTPMiddleware is here to show the optional method to leverage for http/1 only
func (f *FauxAuth) HTTPMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		/*
			NOTE:
			This is where restrictions and auth checks outside of the GRPC methods would be
			typically.  Likely a reusable method call once the necessary info is pulled from
			the http.Request info
		*/

		h.ServeHTTP(w, r)
	})
}
