package service

import (
	"fmt"

	"github.com/felipe-coletti/my-first-crud-in-go/src/config/logger"
	"github.com/felipe-coletti/my-first-crud-in-go/src/config/rest_err"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/domain"
	"go.uber.org/zap"
)

func (userService *userService) FindUser(userId string) (*domain.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (userService *userService) CreateUser(domainInterface domain.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateUser in user domain",
		zap.String("journey", "createUser"),
	)

	domainInterface.EncryptPassword()

	fmt.Println(domainInterface.GetPassword())

	return nil
}

func (userService *userService) UpdateUser(userId string, domainInterface domain.UserDomainInterface) *rest_err.RestErr {
	return nil
}

func (userService *userService) DeleteUser(userId string) *rest_err.RestErr {
	return nil
}
