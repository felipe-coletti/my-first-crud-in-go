package routes

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/users", userController.FindAllUsers)
	r.GET("/users/:username", userController.FindUserByUsername)
	r.GET("/users/me", userController.FindMe)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/me", userController.UpdateMe)
	r.DELETE("/users/me", userController.DeleteMe)
}
