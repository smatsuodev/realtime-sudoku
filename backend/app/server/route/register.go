package route

import (
	"net/http"
	"sudoku/config"
	"sudoku/gen/sudoku/auth/v1/authv1connect"
	authH "sudoku/handler/auth"
	authS "sudoku/service/auth"
)

func Register(mux *http.ServeMux) {
	cfg := config.NewEnvConfig()

	// construct services
	// TODO: OAuthClient を差し替える
	authService := authS.NewService(authS.ServiceConfig{
		JWTSecret:        cfg.JWTSecret,
		OAuthClientID:    cfg.GitHubClientID,
		OAuthRedirectURI: cfg.OAuthRedirectURI,
	}, new(authS.MockOAuthClient))

	// construct handlers
	authHandler := authH.NewHandler(authService)

	// register handlers
	mux.Handle(authv1connect.NewAuthServiceHandler(authHandler))
}
