package service

import (
	"github.com/ory/viper"
	"github.com/rs/cors"
)

const (
	// EnvCorsEnable env var
	EnvCorsEnable = "CORS_ENABLE"
	// EnvCorsDebug env var
	EnvCorsDebug = "CORS_DEBUG"
)

// CORS sets up the Cross-Origin settings if the current configuration asks for it
func (s *Service) CORS() {

	if viper.GetBool(EnvCorsEnable) {

		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{
				"Access-Control-Allow-Origin",
				"Content-Type",
				"Authorization",
			},
			Debug: viper.GetBool(EnvCorsDebug),
		})
		s.Server.AddHttpMiddlware(c.Handler)
	}

}
