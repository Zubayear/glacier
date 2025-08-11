package services

import (
	"glacier/internal/application/ports"
	"glacier/internal/domain"
)

// UserService is a use case that orchestrates domain objects.
// It depends on the UserRepositoryPort interface.
type UserService struct {
	repo ports.UserRepositoryPort
}

func NewUserService(repo ports.UserRepositoryPort) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name, email string) (*domain.User, error) {
	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}
