package controller

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/config/validation"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/request"
	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {}

func FindUserByUsername(c *gin.Context) {}

func FindMe(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}
}

func UpdateMe(c *gin.Context) {}

func DeleteMe(c *gin.Context) {}
