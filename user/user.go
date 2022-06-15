package user

import "errors"

type User struct {
	username, password string
	isActive           int
}

func AddUser(username, password string, isActive int) (User, error) {
	if username == "" || password == "" {
		return User{}, errors.New("username or password is not entered !!!")
	}
	return User{username: username, password: password, isActive: isActive}, nil
}

func (user *User) GetFieldsFromUser() (string, string, int) {
	return user.username, user.password, user.isActive
}
