package repositories

import "github.com/sandronister/jwt-auth/internal/entity"

func (u *User) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := u.connection.QueryRow("SELECT id,username,password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}
