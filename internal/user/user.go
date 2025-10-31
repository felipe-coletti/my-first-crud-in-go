package user

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetDisplayName() string
	GetUsername() string
	GetEmail() string
	GetPassword() string
	EncryptPassword()
}

type userDomain struct {
	displayName string
	username    string
	email       string
	password    string
}

func NewUserDomain(
	displayName string,
	username string,
	email string,
	password string,
) UserDomainInterface {
	return &userDomain{
		displayName,
		username,
		email,
		password,
	}
}

func (userDomain *userDomain) GetDisplayName() string {
	return userDomain.displayName
}

func (userDomain *userDomain) GetUsername() string {
	return userDomain.username
}

func (userDomain *userDomain) GetEmail() string {
	return userDomain.email
}

func (userDomain *userDomain) GetPassword() string {
	return userDomain.password
}

func (userDomain *userDomain) EncryptPassword() {
	hash := md5.New()

	defer hash.Reset()

	hash.Write([]byte(userDomain.password))

	userDomain.password = hex.EncodeToString(hash.Sum(nil))
}
