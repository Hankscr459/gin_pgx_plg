package db

import (
	"database/sql"
)

type User struct {
	ID          int32
	Name        string
	DateOfBirth sql.NullTime
	CreatedAt   sql.NullTime
	Description sql.NullString
}
