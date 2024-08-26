package auth

import "github.com/google/uuid"

type OAuthProvider string

const (
	OAuthProviderGitHub OAuthProvider = "github"
)

type SignInInput struct {
	Provider OAuthProvider
}

type SignInOutput struct {
	AuthorizationURL string
	StateJWT         string
}

type SignOutInput struct {
	SessionID uuid.UUID
}

type SignOutOutput struct{}

type OAuthCallbackInput struct {
	Code     string
	State    string
	StateJWT string
}

type OAuthCallbackOutput struct {
	SessionID uuid.UUID
}

type ValidateSessionInput struct {
	SessionID uuid.UUID
}

type ValidateSessionOutput struct {
	IsValid bool
}
