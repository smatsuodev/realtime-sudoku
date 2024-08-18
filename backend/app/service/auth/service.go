package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const (
	StateJWTIssuer = "sudoku"
)

var _ IAuthService = (*Service)(nil)

//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock
type IAuthService interface {
	SignIn(SignInInput) (SignInOutput, error)
	SignOut(SignOutInput) (SignOutOutput, error)
	OAuthCallback(OAuthCallbackInput) (OAuthCallbackOutput, error)
}

type StateClaims struct {
	State string `json:"state"`
	jwt.RegisteredClaims
}

type ServiceConfig struct {
	JWTSecret        string
	OAuthClientID    string
	OAuthRedirectURI string
}

type Service struct {
	config ServiceConfig
}

func NewService(config ServiceConfig) *Service {
	return &Service{
		config: config,
	}
}

func (s *Service) SignIn(input SignInInput) (SignInOutput, error) {
	state := uuid.NewString()

	claims := StateClaims{
		State: state,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Issuer:    StateJWTIssuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
	token, err := s.generateJWT(claims)
	if err != nil {
		return SignInOutput{}, err
	}

	// NOTE: GitHub 以外は未対応
	if input.Provider != OAuthProviderGitHub {
		return SignInOutput{}, errors.New("unsupported provider")
	}
	url := generateGitHubAuthURL(generateURLParams{
		redirectURI: s.config.OAuthRedirectURI,
		state:       state,
		scope:       []string{}, // no scope
		clientID:    s.config.OAuthClientID,
	})

	return SignInOutput{
		StateJWT:         token,
		AuthorizationURL: url,
	}, nil
}

func (s *Service) generateJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *Service) SignOut(input SignOutInput) (SignOutOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) OAuthCallback(input OAuthCallbackInput) (OAuthCallbackOutput, error) {
	// TODO: state の検証

	// TODO: アクセストークンの取得

	// TODO: ユーザ情報の取得

	// TODO: ユーザーの作成

	// TODO: セッションの作成

	panic("implement me")
}
