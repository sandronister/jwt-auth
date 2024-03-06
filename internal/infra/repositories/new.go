package repositories

import (
	"database/sql"
)

type User struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *User {
	return &User{connection: connection}
}
