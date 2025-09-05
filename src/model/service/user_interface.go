package service

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/config/rest_err"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model"
)

type UserServiceInterface interface {
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}

type userService struct {
	
}

func NewUserService() UserServiceInterface {
	return &userService{}
}