package middleware

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

func AddGatewayRequestIdHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, exists := hlog.IDFromRequest(r)
		if !exists {
			fmt.Println("not exists")
		}

		// fmt.Println("before", id.String())
		r.Header.Add("Grpc-Metadata-Request-Id", id.String())
		// defer fmt.Println("after")
		h.ServeHTTP(w, r)
	})
}

func AddGatewayAPIKeyHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key != "" {
			r.Header.Add("Grpc-Metadata-X-Api-Key", key)
		}
		// defer fmt.Println("after")
		h.ServeHTTP(w, r)
	})
}
