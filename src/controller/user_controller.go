package controller

import (
	"fmt"

	"github.com/felipe-coletti/my-first-crud-in-go/src/config/rest_err"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/request"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {}

func FindUserByUsername(c *gin.Context) {}

func FindMe(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewRestErr("bad_request", fmt.Sprintf("Há campos incorretos. Erro = %s", err.Error()), nil)

		c.JSON(restErr.Code, restErr)

		return
	}
}

func UpdateMe(c *gin.Context) {}

func DeleteMe(c *gin.Context) {}
