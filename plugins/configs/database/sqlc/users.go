package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO users (
  name,
  date_of_birth,
  description
) VALUES (
  $1, $2, $3
) RETURNING id, name, date_of_birth, created_at, description
`

type CreateAccountParams struct {
	Name        string
	DateOfBirth sql.NullTime
	Description sql.NullString
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Name, arg.DateOfBirth, arg.Description)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DateOfBirth,
		&i.CreatedAt,
		&i.Description,
	)
	return i, err
}
