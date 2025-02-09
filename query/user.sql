-- name: CreateUser :exec
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;