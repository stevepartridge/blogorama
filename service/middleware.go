package service

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/hlog"
	zerolog "github.com/rs/zerolog/log"
	"github.com/stevepartridge/blogorama/pkg/middleware"
)

func (s *Service) middleware() {
	s.Server.Grpc.AddUnaryInterceptors(
		middleware.RequestInterceptor(),
		middleware.TelemetryInterceptor(),
		s.FauxAuth.Interceptor(),
	)
}

// httpMiddleware is specific to http1 requests only, they may not be routed to a GRPC method
func (s *Service) httpMiddleware() {

	_log := zerolog.With().
		Str("server", "http1").
		Logger()

	s.Server.AddHttpMiddlware(hlog.NewHandler(_log))

	s.Server.AddHttpMiddlware(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	s.Server.AddHttpMiddlware(hlog.RemoteAddrHandler("ip"))
	s.Server.AddHttpMiddlware(hlog.UserAgentHandler("user_agent"))
	s.Server.AddHttpMiddlware(hlog.RefererHandler("referer"))
	s.Server.AddHttpMiddlware(hlog.RequestIDHandler("req_id", "Request-Id"))

	s.Server.AddHttpMiddlware(middleware.AddGatewayRequestIdHeader) // Pass Requesting ID to GRPC context
	s.Server.AddHttpMiddlware(middleware.AddGatewayAPIKeyHeader)    // Pass X-API-Key to GRPC context

	s.Server.AddHttpMiddlware(s.FauxAuth.HTTPMiddleware)
}
