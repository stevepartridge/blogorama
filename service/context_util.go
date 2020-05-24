package service

import (
	"context"
	"strconv"

	"github.com/stevepartridge/blogorama/pkg/middleware"
)

func extractUserIDFromContext(ctx context.Context) int {
	userID, _ := strconv.Atoi(middleware.GetStringFromContext(ctx, middleware.UserIDKey))
	return userID
}
