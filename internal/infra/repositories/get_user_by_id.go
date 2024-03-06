package repositories

import "github.com/sandronister/jwt-auth/internal/entity"

func (r *User) GetUserById(id int) (entity.User, error) {
	var user entity.User
	err := r.connection.QueryRow("SELECT id,username FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username)
	return user, err
}
