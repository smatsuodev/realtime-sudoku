package route

import (
	"connectrpc.com/connect"
	"gorm.io/gorm"
	"net/http"
	"sudoku/config"
	"sudoku/gen/sudoku/auth/v1/authv1connect"
	authH "sudoku/handler/auth"
	"sudoku/handler/interceptor"
	authI "sudoku/infra/auth"
	gormRepo "sudoku/infra/gorm"
	authS "sudoku/service/auth"
)

func Register(mux *http.ServeMux, cfg config.EnvConfig, db *gorm.DB) {
	// construct services
	authService := authS.NewService(
		authS.ServiceConfig{
			JWTSecret:        cfg.JWTSecret,
			OAuthClientID:    cfg.GitHubClientID,
			OAuthRedirectURI: cfg.OAuthRedirectURI,
		},
		authI.NewOAuthClientImpl(cfg.GitHubClientID, cfg.GitHubClientSecret),
		authI.NewGitHubAPIImpl(),
		gormRepo.NewUserRepository(db),
		gormRepo.NewSessionRepository(db),
	)

	// construct handlers
	authHandler := authH.NewHandler(authService)
	authCallbackHandler := authH.NewCallbackHandler(authService)

	// register handlers
	interceptors := connect.WithInterceptors(interceptor.NewAuthInterceptor(authService))

	mux.Handle(authv1connect.NewAuthServiceHandler(authHandler, interceptors))
	mux.HandleFunc("/auth/github/callback", authCallbackHandler.Handle)
}
