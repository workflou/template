-- name: CountAdmins :one
SELECT COUNT(*) FROM users
WHERE is_admin = true;

-- name: CreateAdmin :exec
INSERT INTO users (name, email, password, is_admin)
VALUES (?, ?, ?, true);