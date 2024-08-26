package auth

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/google/uuid"
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
	var cookies []*http.Cookie

	sessionCookie := req.Header().Get(SessionCookieName)
	if sessionCookie != "" {
		// 無効なセッションクッキーを削除させる用
		c := http.Cookie{
			Name:   SessionCookieName,
			MaxAge: -1,
		}

		sessionID, err := uuid.Parse(sessionCookie)
		if err != nil {
			// サインイン処理は継続
			cookies = append(cookies, &c)
		} else {
			output, err := h.authService.ValidateSession(authS.ValidateSessionInput{
				SessionID: sessionID,
			})
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, err)
			}

			if output.IsValid {
				return connect.NewResponse(&authv1.SignInResponse{
					Status: authv1.SignInStatus_SIGN_IN_STATUS_ALREADY_SIGNED_IN,
				}), nil
			}

			// サインイン処理は継続
			cookies = append(cookies, &c)
		}
	}

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
		Status:           authv1.SignInStatus_SIGN_IN_STATUS_REQUIRES_SIGNING_IN,
	})

	// Cookie として返していいのか?
	// TODO: セキュリティを考慮して属性を追加
	// Secure を付けたいが, 開発環境で cookie が送られなくて困りそう
	// SameSite: strict だとコールバックに cookie が送られなさそう
	cookies = append(cookies, &http.Cookie{
		Name:     "state_jwt",
		Value:    output.StateJWT,
		Path:     "/",
		HttpOnly: true,
	})

	for _, cookie := range cookies {
		res.Header().Set("Set-Cookie", cookie.String())
	}

	return res, nil
}

func (h *Handler) SignOut(ctx context.Context, req *connect.Request[authv1.SignOutRequest]) (*connect.Response[authv1.SignOutResponse], error) {
	//TODO implement me
	panic("implement me")
}
