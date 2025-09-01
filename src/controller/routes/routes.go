package routes

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", controller.FindAllUsers)
	r.GET("/users/:username", controller.FindUserByUsername)
	r.GET("/users/me", controller.FindMe)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/me", controller.UpdateMe)
	r.DELETE("/users/me", controller.DeleteMe)
}
