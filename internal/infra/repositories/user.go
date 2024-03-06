package repositories

import (
	"database/sql"

	"github.com/sandronister/jwt-auth/internal/entity"
)

type User struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *User {
	return &User{connection: connection}
}

func (u *User) CreateUser(username, password string) error {
	_, err := u.connection.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

func (u *User) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := u.connection.QueryRow("SELECT id,username,password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}
