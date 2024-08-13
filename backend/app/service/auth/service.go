package auth

var _ IAuthService = (*Service)(nil)

type IAuthService interface {
	SignIn(SignInInput) (SignInOutput, error)
	SignOut(SignOutInput) (SignOutOutput, error)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SignIn(input SignInInput) (SignInOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SignOut(input SignOutInput) (SignOutOutput, error) {
	//TODO implement me
	panic("implement me")
}
