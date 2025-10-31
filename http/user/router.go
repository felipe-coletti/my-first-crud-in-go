package user

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userHandler UserHandlerInterface) {
	r.GET("/users", userHandler.FindAllUsers)
	r.GET("/users/:username", userHandler.FindUserByUsername)
	r.GET("/users/me", userHandler.FindMe)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/me", userHandler.UpdateMe)
	r.DELETE("/users/me", userHandler.DeleteMe)
}
