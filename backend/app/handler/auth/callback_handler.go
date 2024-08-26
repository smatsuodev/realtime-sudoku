package auth

import (
	"errors"
	"net/http"
	"sudoku/service/auth"
)

const (
	SessionCookieName = "sudoku_session_id"
)

type CallbackHandler struct {
	authService auth.IAuthService
}

func NewCallbackHandler(authService auth.IAuthService) *CallbackHandler {
	return &CallbackHandler{authService: authService}
}

func (h *CallbackHandler) Handle(w http.ResponseWriter, r *http.Request) {
	queryParams, err := getQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stateCookie, err := r.Cookie("state_jwt")
	if err != nil {
		http.Error(w, "state cookie is required", http.StatusBadRequest)
		return
	}

	output, err := h.authService.OAuthCallback(auth.OAuthCallbackInput{
		Code:     queryParams.Code,
		State:    queryParams.State,
		StateJWT: stateCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    output.SessionID,
		Secure:   false, // NOTE: true が望ましいが, 開発環境で支障が出そう
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	// 成功レスポンスはこれでいいのか?
	w.WriteHeader(http.StatusNoContent)
}

type callbackQueryParams struct {
	Code  string
	State string
}

func getQueryParams(r *http.Request) (callbackQueryParams, error) {
	if !(r.URL.Query().Has("code") && r.URL.Query().Has("state")) {
		return callbackQueryParams{}, errors.New("code and state are required")
	}

	return callbackQueryParams{
		Code:  r.URL.Query().Get("code"),
		State: r.URL.Query().Get("state"),
	}, nil
}
