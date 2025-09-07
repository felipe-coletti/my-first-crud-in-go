package controller

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	FindAllUsers(c *gin.Context)
	FindUserByUsername(c *gin.Context)
	FindMe(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateMe(c *gin.Context)
	DeleteMe(c *gin.Context)
}

type userController struct {
	service service.UserServiceInterface
}

func NewUserController(serviceInterface service.UserServiceInterface) UserControllerInterface {
	return &userController{
		service: serviceInterface,
	}
}