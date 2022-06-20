package user_service

import (
	user "CRUDRestApiDatabase/model"
	"errors"
	"strings"
)

type UserInterface interface {
	Register(user *user.User) (*user.User, error)
}

type UserSrvc struct{}

func NewUserSvc() UserInterface {
	return &UserSrvc{}
}

func (u *UserSrvc) Register(user *user.User) (*user.User, error) {
	//emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if user.Email == "" {
		return nil, errors.New("Email harus di isi")
	}
	// if emailRegex.MatchString(user.Email) {
	// 	return nil, errors.New("Email harus sesuai")
	// }
	if !strings.Contains(user.Email, "@gmail.com") {
		return nil, errors.New("Must contain @gmail.com")
	}
	if user.Username == "" {
		return nil, errors.New("Password harus di isi")
	}
	if user.Password == "" || len(user.Password) < 6 {
		return nil, errors.New("Password harus di isi dan harus lebih dari 6 huruf")
	}
	if user.Age_user <= 8 {
		return nil, errors.New("Umur harus diatas 8 tahun")
	}
	return user, nil
}
