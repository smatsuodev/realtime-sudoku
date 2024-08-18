//go:generate mockgen -source=$GOFILE -destination=mock_service.go -package=$GOPACKAGE
package auth

type IAuthService interface {
	SignIn(SignInInput) (SignInOutput, error)
	SignOut(SignOutInput) (SignOutOutput, error)
	OAuthCallback(OAuthCallbackInput) (OAuthCallbackOutput, error)
}

type OAuthClient interface {
	GetAccessToken(code string, redirectURI string) (string, error)
}
