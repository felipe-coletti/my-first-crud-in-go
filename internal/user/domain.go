package user

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/felipe-coletti/my-first-crud-in-go/pkg/errors"
	"github.com/felipe-coletti/my-first-crud-in-go/pkg/logger"
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
	FindUser(string) (*UserDomain, *errors.RestErr)
	CreateUser() *errors.RestErr
	UpdateUser(string) *errors.RestErr
	DeleteUser(string) *errors.RestErr
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

func (ud *UserDomain) FindUser(string) (*UserDomain, *errors.RestErr) {
	panic("unimplemented")
}

func (ud *UserDomain) CreateUser() *errors.RestErr {
	logger.Info("Init CreateUser in user domain",
		zap.String("journey", "createUser"),
	)

	ud.encryptPassword()

	return nil
}

func (ud *UserDomain) UpdateUser(string) *errors.RestErr {
	panic("unimplemented")
}

func (ud *UserDomain) DeleteUser(string) *errors.RestErr {
	return nil
}
