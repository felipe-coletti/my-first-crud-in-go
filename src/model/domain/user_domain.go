package domain

import (
	"crypto/md5"
	"encoding/hex"
)

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
