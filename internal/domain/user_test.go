package domain_test

import (
	"glacier/internal/domain"
	"testing"
)

func TestUserValidation(t *testing.T)  {
	_, err := domain.NewUser("", "test@email.com")
	if err == nil {
		t.Errorf("Expected nil, Got %v", err)
	}

	u, err := domain.NewUser("zk", "zk@zki.com")
	if err != nil {
		t.Error("Got Error")
	}

	if u.Email != "zk@zki.com" {
		t.Errorf("Expected 'zk@zki.com', Got %v", u.Email)
	}

	_, err = domain.NewUser("ken", "ken")
	if err.Error() != "invalid email format" {
		t.Errorf("Expected 'invalid email format', Got %v", err.Error())
	}
}
