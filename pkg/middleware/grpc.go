package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/stevepartridge/blogorama/pkg/log"
)

// RequestInterceptor is a basic middleware to log each request
func RequestInterceptor() func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		log.Info().
			Str("request_method", info.FullMethod).
			Str("user_agent", GetUserAgentFromContext(ctx)).
			Str("ip", GetForwardedForFromContext(ctx)).
			Str("referer", GetForwardedHostFromContext(ctx)).
			Str("request_id", GetRequestIdFromContext(ctx)).
			Int("user_id", GetUserIdFromContext(ctx)).
			Msg("Request")

		resp, err := handler(ctx, req)
		if err != nil {
			log.Debug().Err(err).Msg("RequestInterceptor")
		}

		return resp, err
	}
}

// TelemetryInterceptor is a basic middleware to log call durations
func TelemetryInterceptor() func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		started := time.Now().UTC()

		resp, err := handler(ctx, req)
		if err != nil {
			log.Debug().Err(err).Msg("TelemetryInterceptor")
		}

		finished := time.Now().UTC()

		log.Info().
			Str("request_method", info.FullMethod).
			Str("start", started.Format(time.RFC3339Nano)).
			Str("finish", finished.Format(time.RFC3339Nano)).
			Str("duration", finished.Sub(started).String()).
			Str("duration", finished.Sub(started).String()).
			Str("request_id", GetRequestIdFromContext(ctx)).
			Int("user_id", GetUserIdFromContext(ctx)).
			Msg("Telemetry")

		return resp, err
	}
}
