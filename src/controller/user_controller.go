package controller

import (
	"net/http"

	"github.com/felipe-coletti/my-first-crud-in-go/src/config/logger"
	"github.com/felipe-coletti/my-first-crud-in-go/src/config/validation"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/domain"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/request"
	"github.com/felipe-coletti/my-first-crud-in-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface domain.UserDomainInterface
)

func (userController *userController) FindAllUsers(c *gin.Context) {}

func (userController *userController) FindUserByUsername(c *gin.Context) {}

func (userController *userController) FindMe(c *gin.Context) {}

func (userController *userController) CreateUser(c *gin.Context) {
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

	domain := domain.NewUserDomain(
		userRequest.DisplayName,
		userRequest.Username,
		userRequest.Email,
		userRequest.Password,
	)

	if err := userController.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, view.ToUserResponse(domain))
}

func (userController *userController) UpdateMe(c *gin.Context) {}

func (userController *userController) DeleteMe(c *gin.Context) {}
