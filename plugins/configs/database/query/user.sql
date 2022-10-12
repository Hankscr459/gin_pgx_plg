-- name: CreateAccount :one
INSERT INTO users (
  name,
  date_of_birth,
  description
) VALUES (
  $1, $2, $3
) RETURNING *;