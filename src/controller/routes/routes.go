package routes

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", controller.FindUserById)
	r.GET("/user/:email", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user", controller.UpdateUser)
	r.DELETE("/user", controller.DeleteUser)
}
