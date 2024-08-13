package auth

import (
	"connectrpc.com/connect"
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (h *Handler) SignOut(ctx context.Context, req *connect.Request[authv1.SignOutRequest]) (*connect.Response[authv1.SignOutResponse], error) {
	//TODO implement me
	panic("implement me")
}
