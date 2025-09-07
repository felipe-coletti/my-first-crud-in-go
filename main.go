package main

import (
	"log"
	"os"

	"github.com/felipe-coletti/my-first-crud-in-go/src/controller"
	"github.com/felipe-coletti/my-first-crud-in-go/src/controller/routes"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserService()
	userController := controller.NewUserController(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
