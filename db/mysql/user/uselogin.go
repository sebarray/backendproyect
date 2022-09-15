package userdb

import (
	"github.com/sebarray/backendproyect/model"
	"github.com/sebarray/backendproyect/service"
	"golang.org/x/crypto/bcrypt"
)

func (u Userdb) Login(user model.User) (string, error) {

	userdb, err := UserExist(user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdb.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	t, err := service.CreateToken(userdb)

	if err != nil {
		return "", err
	}

	return t, nil
}
