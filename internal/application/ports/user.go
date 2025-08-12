package ports

import "glacier/internal/domain"

// UserRepositoryPort is an outer port for the application layer.
// It defines the contract for saving and retrieving users.
// The application layer only knows about this interface, not the concrete implementation.
type UserRepositoryPort interface {
	Save(user *domain.User) error
	FindByID(id uint64) (*domain.User, error)
}
