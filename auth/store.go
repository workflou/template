package auth

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Save(ctx context.Context, u User) (int, error) {
	sql := `INSERT INTO users (email, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := s.db.QueryRow(sql, u.Email, u.CreatedAt, u.UpdatedAt).Scan(&id)
	return id, err
}

func (s *Store) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := s.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}
