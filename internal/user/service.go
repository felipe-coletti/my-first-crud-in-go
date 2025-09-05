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

type userService struct {

}

func NewUserService() UserServiceInterface {
	return &userService{}
}

func (us *userService) FindUser(userId string) (*UserDomainInterface, *errors.RestErr) {
	return nil, nil
}

func (us *userService) CreateUser(userDomain UserDomainInterface) *errors.RestErr {
	logger.Info("Init CreateUser in user domain",
		zap.String("journey", "createUser"),
	)

	userDomain.EncryptPassword()

	return nil
}

func (us *userService) UpdateUser(userId string, userDomain UserDomainInterface) *errors.RestErr {
	return nil
}

func (us *userService) DeleteUser(userId string) *errors.RestErr {
	return nil
}
