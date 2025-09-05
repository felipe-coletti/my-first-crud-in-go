package service

import (
	"fmt"

	"github.com/felipe-coletti/my-first-crud-in-go/src/config/logger"
	"github.com/felipe-coletti/my-first-crud-in-go/src/config/rest_err"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model"
	"go.uber.org/zap"
)

func (us *userService) FindUser(userId string) (*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (us *userService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateUser in user domain",
		zap.String("journey", "createUser"),
	)

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}

func (us *userService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}

func (us *userService) DeleteUser(userId string) *rest_err.RestErr {
	return nil
}