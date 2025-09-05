package user

import (
	"net/http"

	"github.com/felipe-coletti/my-first-crud-in-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FindAllUsers(c *gin.Context) {}

func FindUserByUsername(c *gin.Context) {}

func FindMe(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser in user handler",
		zap.String("journey", "createUser"),
	)
	var request UserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)

		restErr := ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := NewUserDomain(
		request.DisplayName,
		request.Username,
		request.Email,
		request.Password,
	)

	service := NewUserService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusOK, "")
}

func UpdateMe(c *gin.Context) {}

func DeleteMe(c *gin.Context) {}
