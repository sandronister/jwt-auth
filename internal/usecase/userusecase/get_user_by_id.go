package userusecase

import "github.com/sandronister/jwt-auth/internal/entity"

func (u *User) GetUserById(id int) (entity.User, error) {
	return u.repository.GetUserById(id)
}
