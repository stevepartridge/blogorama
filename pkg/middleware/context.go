package middleware

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

// MergeStringValueIntoContextFromMetadata
func MergeStringValueIntoContextFromMetadata(ctx context.Context, key string, value string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if val, ok := md[key]; ok {
			if len(val) > 0 {
				ctx = context.WithValue(ctx, key, val[0])
			}
		}
	}
	return ctx
}

// MergeIntValueIntoContextFromMetadata
func MergeIntValueIntoContextFromMetadata(ctx context.Context, key string, value int) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if keyValue, ok := md[key]; ok {
			if len(keyValue) > 0 {
				ctx = context.WithValue(ctx, key, keyValue[0])
			}
		}
	}
	return ctx
}

// MergeRequestingUserIdIntoContextFromMetadata
func MergeRequestingUserIdIntoContextFromMetadata(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if reqClientId, ok := md[RequestingUserIdKey]; ok {
			if len(reqClientId) > 0 {
				ctx = context.WithValue(ctx, RequestingUserIdKey, reqClientId[0])
			}
		}
	}
	return ctx
}

// GetUserAgentFromContext
func GetUserAgentFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["grpcgateway-user-agent"]; ok && len(key) > 0 {
			return key[0]
		}

		if key, ok := md["user-agent"]; ok && len(key) > 0 {
			return key[0]
		}
	}

	return ""
}

// GetForwardedForFromContext
func GetForwardedForFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["x-forwarded-for"]; ok && len(key) > 0 {
			return key[0]
		}
	}

	return ""
}

// GetForwardedHostFromContext
func GetForwardedHostFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["x-forwarded-host"]; ok && len(key) > 0 {
			return key[0]
		}
	}
	return ""
}

// GetRequestIdFromContext
func GetRequestIdFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["request-id"]; ok && len(key) > 0 {
			return key[0]
		}
	}
	return ""
}

// GetAPIKeyFromContext
func GetAPIKeyFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["x-api-key"]; ok && len(key) > 0 {
			return key[0]
		}
	}
	return ""
}

// GetAuthorizationFromContext
func GetAuthorizationFromContext(ctx context.Context) string {

	token := ""

	if ctx.Value(GRPCAuthHeaderKey) != nil {
		token = ctx.Value(GRPCAuthHeaderKey).(string)
	}

	if token == "" {

		if md, ok := metadata.FromIncomingContext(ctx); ok {

			if key, ok := md["grpcgateway-authorization"]; ok && len(key) > 0 {
				token = key[0]
			}

			if key, ok := md["authorization"]; ok && len(key) > 0 {
				token = key[0]
			}

		}

	}

	if token != "" {
		if strings.Index(strings.ToLower(token), "bearer") == 0 {
			token = token[7:]
		}
	}

	return token
}

// ContextWithRequestingUserId
func ContextWithRequestingUserId(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, RequestingUserIdKey, userId)
}

// ContextWithRequestingMachineId
func ContextWithRequestingMachineId(ctx context.Context, machineId string) context.Context {
	return context.WithValue(ctx, RequestingMachineIdKey, machineId)
}

// ContextWithRequestingServiceId
func ContextWithRequestingServiceId(ctx context.Context, serviceId int) context.Context {
	return context.WithValue(ctx, RequestingServiceIdKey, serviceId)
}

// GetUserIdFromCont
func GetUserIdFromContext(ctx context.Context) int {
	val := ctx.Value(RequestingUserIdKey)
	if val != nil {
		return val.(int)
	}
	return 0
}

// GetMachineIdFromContext
func GetMachineIdFromContext(ctx context.Context) string {
	val := ctx.Value(RequestingMachineIdKey)
	if val != nil {
		return val.(string)
	}
	return ""
}

// GetServiceIdFromCont
func GetServiceIdFromContext(ctx context.Context) int {
	val := ctx.Value(RequestingServiceIdKey)
	if val != nil {
		return val.(int)
	}
	return 0
}

// GetStringFromContext
func GetStringFromContext(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val != nil {
		return val.(string)
	}
	return ""
}

// RequestIsFromService
func RequestIsFromService(ctx context.Context) bool {
	return GetServiceIdFromContext(ctx) > 0
}

// RequestIsFromMachine
func RequestIsFromMachine(ctx context.Context) bool {
	return GetMachineIdFromContext(ctx) != ""
}
