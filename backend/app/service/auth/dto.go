package auth

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
	SessionID string
}

type SignOutOutput struct{}
