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

func (ud *userDomain) GetDisplayName() string {
	return ud.displayName
}

func (ud *userDomain) GetUsername() string {
	return ud.username
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()

	defer hash.Reset()

	hash.Write([]byte(ud.password))

	ud.password = hex.EncodeToString(hash.Sum(nil))
}
