package user

import (
	"github.com/felipe-coletti/my-first-crud-in-go/pkg/errors"
	"github.com/felipe-coletti/my-first-crud-in-go/pkg/logger"
	"go.uber.org/zap"
)

type UserServiceInterface interface {
	FindUser(string) (*UserDomainInterface, *errors.RestErr)
	CreateUser(UserDomainInterface) *errors.RestErr
	UpdateUser(string, UserDomainInterface) *errors.RestErr
	DeleteUser(string) *errors.RestErr
}

type userService struct {}

func NewUserService() UserServiceInterface {
	return &userService{}
}

func (userService *userService) FindUser(userId string) (*UserDomainInterface, *errors.RestErr) {
	return nil, nil
}

func (userService *userService) CreateUser(domainInterface UserDomainInterface) *errors.RestErr {
	logger.Info("Init CreateUser in user domain",
		zap.String("journey", "createUser"),
	)

	domainInterface.EncryptPassword()

	return nil
}

func (userService *userService) UpdateUser(userId string, domainInterface UserDomainInterface) *errors.RestErr {
	return nil
}

func (userService *userService) DeleteUser(userId string) *errors.RestErr {
	return nil
}
