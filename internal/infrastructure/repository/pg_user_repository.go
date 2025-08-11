package repository

import (
	"database/sql"
	"glacier/internal/domain"
)

type PgUserRepository struct {
	db *sql.DB
}

func NewPgUserRepository(db *sql.DB) *PgUserRepository {
	return &PgUserRepository{
		db: db,
	}
}

func (r *PgUserRepository) Save(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (r *PgUserRepository) FindByID(id uint64) (*domain.User, error) {
	return nil, nil
}
