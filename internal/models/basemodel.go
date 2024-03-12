package models

import (
	"database/sql"
	"time"
)

type Metadata struct {
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Person struct {
	Metadata

	ID    int `db:"id"`
	Email string
}
