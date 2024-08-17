package auth

import (
	"connectrpc.com/connect"
	"context"
	"net/http"
	authv1 "sudoku/gen/sudoku/auth/v1"
	"sudoku/service/auth"
)

type Handler struct {
	authService auth.IAuthService
}

func NewHandler(authService auth.IAuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) SignIn(ctx context.Context, req *connect.Request[authv1.SignInRequest]) (*connect.Response[authv1.SignInResponse], error) {
	output, err := h.authService.SignIn(auth.SignInInput{
		Provider: auth.OAuthProvider(req.Msg.Provider),
	})
	if err != nil {
		return nil, err
	}

	// TODO: 属性を追加
	cookie := http.Cookie{
		Name:  "state_jwt",
		Value: output.StateJWT,
	}
	// Cookie として返していいのか?
	req.Header().Set("Set-Cookie", cookie.String())

	res := connect.NewResponse(&authv1.SignInResponse{
		AuthorizationUrl: output.AuthorizationURL,
	})

	return res, nil
}

func (h *Handler) SignOut(ctx context.Context, req *connect.Request[authv1.SignOutRequest]) (*connect.Response[authv1.SignOutResponse], error) {
	//TODO implement me
	panic("implement me")
}
