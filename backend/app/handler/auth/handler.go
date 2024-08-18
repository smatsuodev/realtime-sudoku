package auth

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"net/http"
	authv1 "sudoku/gen/sudoku/auth/v1"
	authS "sudoku/service/auth"
)

type Handler struct {
	authService authS.IAuthService
}

func NewHandler(authService authS.IAuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func convertProvider(provider authv1.OAuthProvider) (authS.OAuthProvider, error) {
	switch provider {
	case authv1.OAuthProvider_OAUTH_PROVIDER_GITHUB:
		return authS.OAuthProviderGitHub, nil
	default:
		return "", errors.New("unsupported provider")
	}
}

func (h *Handler) SignIn(ctx context.Context, req *connect.Request[authv1.SignInRequest]) (*connect.Response[authv1.SignInResponse], error) {
	provider, err := convertProvider(req.Msg.Provider)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	output, err := h.authService.SignIn(authS.SignInInput{
		Provider: provider,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&authv1.SignInResponse{
		AuthorizationUrl: output.AuthorizationURL,
	})

	// TODO: セキュリティを考慮して属性を追加
	// Secure を付けたいが, 開発環境で cookie が送られなくて困りそう
	// SameSite: strict だとコールバックに cookie が送られなさそう
	cookie := http.Cookie{
		Name:     "state_jwt",
		Value:    output.StateJWT,
		HttpOnly: true,
	}
	// Cookie として返していいのか?
	res.Header().Set("Set-Cookie", cookie.String())

	return res, nil
}

func (h *Handler) SignOut(ctx context.Context, req *connect.Request[authv1.SignOutRequest]) (*connect.Response[authv1.SignOutResponse], error) {
	//TODO implement me
	panic("implement me")
}
