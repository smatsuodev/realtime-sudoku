package route

import (
	"net/http"
	"sudoku/config"
	"sudoku/gen/sudoku/auth/v1/authv1connect"
	authH "sudoku/handler/auth"
	authI "sudoku/infra/auth"
	authS "sudoku/service/auth"
)

func Register(mux *http.ServeMux) {
	cfg := config.NewEnvConfig()

	// construct services
	authService := authS.NewService(authS.ServiceConfig{
		JWTSecret:        cfg.JWTSecret,
		OAuthClientID:    cfg.GitHubClientID,
		OAuthRedirectURI: cfg.OAuthRedirectURI,
	}, authI.NewOAuthClientImpl(cfg.GitHubClientID, cfg.GitHubClientSecret))

	// construct handlers
	authHandler := authH.NewHandler(authService)

	// register handlers
	mux.Handle(authv1connect.NewAuthServiceHandler(authHandler))
}
