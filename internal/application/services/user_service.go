package services

import (
	"glacier/internal/application/ports"
	"glacier/internal/domain"
)

// UserService is a use case that orchestrates domain objects.
// It depends on the UserRepositoryPort interface.
type UserService struct {
	repo ports.UserRepositoryPort
	log  ports.LoggerPort
}

func NewUserService(repo ports.UserRepositoryPort) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name, email string) (*domain.User, error) {
	s.log.Info("Attempting to create a new user", "email", email)
	user, err := domain.NewUser(name, email)
	if err != nil {
		s.log.Error("Failed to create user due to invalid data", "error", err)
		return nil, err
	}
	if err := s.repo.Save(user); err != nil {
		s.log.Error("Failed to save user to repository", "error", err)
		return nil, err
	}
	s.log.Info("Successfully created user", "userID", user.ID)
	return user, nil
}
