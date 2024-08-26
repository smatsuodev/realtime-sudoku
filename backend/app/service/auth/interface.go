//go:generate mockgen -source=$GOFILE -destination=mock_service.go -package=$GOPACKAGE
package auth

type IAuthService interface {
	SignIn(SignInInput) (SignInOutput, error)
	SignOut(SignOutInput) (SignOutOutput, error)
	OAuthCallback(OAuthCallbackInput) (OAuthCallbackOutput, error)
	ValidateSession(ValidateSessionInput) (ValidateSessionOutput, error)
}

type OAuthClient interface {
	GetAccessToken(code string, redirectURI string) (string, error)
}

type GitHubAPI interface {
	// access token を引数でもらうのが微妙
	GetUser(accessToken string) (*GitHubUser, error)
}

// ここに定義していいのか?
type GitHubUser struct {
	ID    uint
	Login string
}
