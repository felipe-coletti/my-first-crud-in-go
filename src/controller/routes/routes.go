package routes

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", controller.FindAll)
	r.GET("/users/:id", controller.FindUserById)
	r.GET("/users/:username", controller.FindUserByUsername)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:username", controller.UpdateUser)
	r.DELETE("/users/:username", controller.DeleteUser)
}
