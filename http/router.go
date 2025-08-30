package http

import (
	"github.com/felipe-coletti/my-first-crud-in-go/internal/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", user.FindAll)
	r.GET("/users/:username", user.FindUserByUsername)
	r.GET("/users/me", user.FindMe)
	r.POST("/users", user.CreateUser)
	r.PUT("/users/me", user.UpdateMe)
	r.DELETE("/users/me", user.DeleteMe)
}
