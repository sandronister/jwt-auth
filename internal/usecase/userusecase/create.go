package userusecase

import "golang.org/x/crypto/bcrypt"

func (u *User) CreateUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.repository.CreateUser(username, string(hashedPassword))
}
