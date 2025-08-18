package repository

import (
	"context"
	"glacier/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PgUserRepository (outer layer) provides the implementation contract defined in inner layer
type PgUserRepository struct {
	pool *pgxpool.Pool
}

func NewPgUserRepository(pool *pgxpool.Pool) *PgUserRepository {
	return &PgUserRepository{
		pool,
	}
}

func (r *PgUserRepository) Save(ctx context.Context, user *domain.User) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (r *PgUserRepository) FindByID(id uint64) (*domain.User, error) {
	return nil, nil
}
