package userdb

import "github.com/sebarray/backendproyect/model"

type Userdb struct{}

type IUser interface {
	CreateUser(user model.User) (model.User, error)
	Login(user model.User) (string, error)
}

var providers = map[string]IUser{}

func init() {
	providers["mysql"] = Userdb{}
}

func GetProvider(typeDb string) IUser {
	return providers[typeDb]
}
