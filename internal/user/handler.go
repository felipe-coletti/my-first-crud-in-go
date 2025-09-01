package user

import (
	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {}

func FindUserByUsername(c *gin.Context) {}

func FindMe(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var request UserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}
}

func UpdateMe(c *gin.Context) {}

func DeleteMe(c *gin.Context) {}
