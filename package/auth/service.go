package auth

type Service interface {
	HashPassword(password string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) HashPassword(password string) (string, error) {
	return s.repository.HashPassword(password)
}
