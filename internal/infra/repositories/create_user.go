package repositories

func (u *User) CreateUser(username, password string) error {
	_, err := u.connection.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}
