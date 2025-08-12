package domain

import (
	"errors"
	"regexp"
)

// User represents a user entity in the system
type User struct {
	ID    uint64
	Name  string
	Email string
}

func NewUser(name, email string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	// Example of domain-level business rule
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	return &User{
		Name:  name,
		Email: email,
	}, nil
}

// The domain layer contains its own logic. It doesn't need to import any other package
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}
