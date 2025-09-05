package controller

import (
	"net/http"

	"github.com/felipe-coletti/my-first-crud-in-go/src/config/logger"
	"github.com/felipe-coletti/my-first-crud-in-go/src/config/validation"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/request"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func FindAllUsers(c *gin.Context) {}

func FindUserByUsername(c *gin.Context) {}

func FindMe(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser in user handler",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserDomain(
		userRequest.DisplayName,
		userRequest.Username,
		userRequest.Email,
		userRequest.Password,
	)

	service := service.NewUserService()

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
