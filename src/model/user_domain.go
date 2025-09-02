package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/felipe-coletti/my-first-crud-in-go/src/config/logger"
	"github.com/felipe-coletti/my-first-crud-in-go/src/config/rest_err"
	"go.uber.org/zap"
)

type UserDomain struct {
	DisplayName string
	Username    string
	Email       string
	Password    string
}

func (ud *UserDomain) encryptPassword() {
	hash := md5.New()

	defer hash.Reset()

	hash.Write([]byte(ud.Password))

	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}

func NewUserDomain(
	displayName string,
	username string,
	email string,
	password string,
) UserDomainInterface {
	return &UserDomain{
		displayName,
		username,
		email,
		password,
	}
}

func (ud *UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr) {
	panic("unimplemented")
}

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser in user domain",
		zap.String("journey", "createUser"),
	)

	ud.encryptPassword()

	return nil
}

func (ud *UserDomain) UpdateUser(string) *rest_err.RestErr {
	panic("unimplemented")
}

func (ud *UserDomain) DeleteUser(string) *rest_err.RestErr {
	return nil
}
