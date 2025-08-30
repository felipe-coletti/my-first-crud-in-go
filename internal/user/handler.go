package user

import (
	"fmt"

	"github.com/felipe-coletti/my-first-crud-in-go/pkg/errors"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {}

func FindUserByUsername(c *gin.Context) {}

func FindMe(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewRestErr("bad_request", fmt.Sprintf("Há campos incorretos. Erro = %s", err.Error()), nil)

		c.JSON(restErr.Code, restErr)

		return
	}
}

func UpdateMe(c *gin.Context) {}

func DeleteMe(c *gin.Context) {}
