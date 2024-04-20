package auth

import "time"

type User struct {
	ID        int64     `db:"id"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
