package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"sudoku/model"
	"sudoku/repository"
	"time"
)

const (
	StateJWTIssuer = "sudoku"
)

var _ IAuthService = (*Service)(nil)

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
	config      ServiceConfig
	oauthClient OAuthClient
	githubAPI   GitHubAPI
	userRepo    repository.IUserRepository
}

func NewService(
	config ServiceConfig,
	oauthClient OAuthClient,
	githubAPI GitHubAPI,
	userRepo repository.IUserRepository,
) *Service {
	return &Service{
		config:      config,
		oauthClient: oauthClient,
		githubAPI:   githubAPI,
		userRepo:    userRepo,
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
	// 今の実装では false ならエラーが返るので, そのまま返す
	if ok, err := s.isValidStateJWT(input.StateJWT, input.State); !ok {
		return OAuthCallbackOutput{}, err
	}

	accessToken, err := s.oauthClient.GetAccessToken(input.Code, s.config.OAuthRedirectURI)
	if err != nil {
		return OAuthCallbackOutput{}, err
	}

	githubUser, err := s.githubAPI.GetUser(accessToken)
	if err != nil {
		return OAuthCallbackOutput{}, err
	}

	_, err = s.createUserIfNotExists(githubUser)
	if err != nil {
		return OAuthCallbackOutput{}, err
	}

	// TODO: セッションの作成

	panic("implement me")
}

func (s *Service) createUserIfNotExists(githubUser *GitHubUser) (*model.User, error) {
	maybeUser, err := s.userRepo.FindByGitHubUserID(githubUser.ID)
	if err != nil {
		return nil, err
	}

	if maybeUser.IsPresent() {
		return maybeUser.MustGet(), nil
	}

	// ユーザーが存在しないので作成する
	user := &model.User{
		ID:   0, // TODO: ID はどうする?
		Name: githubUser.Login,
	}
	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) isValidStateJWT(token string, state string) (bool, error) {
	claims, err := s.parseJWT(token)
	if err != nil {
		return false, err
	}

	if exp, err := claims.GetExpirationTime(); err != nil {
		if exp == nil {
			return false, errors.New("no expiration time")
		}
		if exp.Before(time.Now()) {
			return false, errors.New("expired state")
		}
	}

	// URL 発行時に生成した state と, callback へのリダイレクト時に受け取る state が一致するか
	if claims.State != state {
		return false, errors.New("invalid state")
	}

	return true, nil
}

func (s *Service) parseJWT(token string) (*StateClaims, error) {
	t, err := jwt.ParseWithClaims(token, &StateClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*StateClaims)
	if !ok {
		return nil, errors.New("unexpected claims")
	}

	return claims, nil
}
