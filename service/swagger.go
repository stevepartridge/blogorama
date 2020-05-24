package service

import (
	"github.com/go-openapi/spec"
	"github.com/stevepartridge/blogorama/pkg/log"
	"github.com/stevepartridge/service/swagger"
)

func (s *Service) swagger() {

	swag, err := swagger.New(s.Server.Router)
	if err != nil {
		log.Warn().Err(err).Msg("Error setting up swagger")
		return
	}

	// Swagger Docs
	//

	swag.Title = Name
	swag.Version = Version
	swag.Schemes = []string{"https"}

	// swag.Path = "/docs"

	swag.JSONData, err = Asset("protos/swagger/service.swagger.json")
	if err != nil {
		log.Warn().Err(err).Msg("Error loading swagger json")
		return
	}

	// oauthSecurityScheme := spec.SecurityScheme{}
	// oauthSecurityScheme.Type = "oauth2"
	// oauthSecurityScheme.AuthorizationURL = fmt.Sprintf("%s/oauth2/auth", viper.GetString("PUBLIC_OAUTH_URL"))
	// oauthSecurityScheme.TokenURL = fmt.Sprintf("%s/oauth2/token", viper.GetString("PUBLIC_OAUTH_URL"))
	// oauthSecurityScheme.Flow = "accessCode"

	// oauthSecurityScheme.Scopes = map[string]string{
	// 	"offline": "",
	// 	"openid":  "",
	// }

	// swag.AddSecurityScheme("consent", oauthSecurityScheme)

	// credsSecurityScheme := spec.SecurityScheme{}
	// credsSecurityScheme.Type = "oauth2"
	// credsSecurityScheme.AuthorizationURL = fmt.Sprintf("%s/oauth2/auth", viper.GetString("PUBLIC_OAUTH_URL"))
	// credsSecurityScheme.TokenURL = fmt.Sprintf("%s/oauth2/token", viper.GetString("PUBLIC_OAUTH_URL"))
	// credsSecurityScheme.Flow = "application"

	// credsSecurityScheme.Scopes = map[string]string{
	// 	"offline": "",
	// 	"openid":  "",
	// }

	// swag.AddSecurityScheme("credentials", credsSecurityScheme)

	apiKeySecurityScheme := spec.SecurityScheme{}
	apiKeySecurityScheme.Type = "apiKey"
	apiKeySecurityScheme.In = "header"
	apiKeySecurityScheme.Name = "X-API-Key"

	swag.AddSecurityScheme("bearer", apiKeySecurityScheme)

	swag.Serve()
}
