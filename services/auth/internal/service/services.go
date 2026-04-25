package service

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(email, password string) error {
	return nil
}
