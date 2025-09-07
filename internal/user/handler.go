package user

import (
	"net/http"

	"github.com/felipe-coletti/my-first-crud-in-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandlerInterface interface {
	FindAllUsers(c *gin.Context)
	FindUserByUsername(c *gin.Context)
	FindMe(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateMe(c *gin.Context)
	DeleteMe(c *gin.Context)
}

type userHandler struct {
	service UserServiceInterface
}

func NewUserHandler(serviceInterface UserServiceInterface) UserHandlerInterface {
	return &userHandler{
		service: serviceInterface,
	}
}

func (userHandler *userHandler) FindAllUsers(c *gin.Context) {}

func (userHandler *userHandler) FindUserByUsername(c *gin.Context) {}

func (userHandler *userHandler) FindMe(c *gin.Context) {}

func (userHandler *userHandler) CreateUser(c *gin.Context) {
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

	if err := userHandler.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, ToUserResponse(domain))
}

func (userHandler *userHandler) UpdateMe(c *gin.Context) {}

func (userHandler *userHandler) DeleteMe(c *gin.Context) {}
