package main

import (
	"log"
	"os"

	userDomain "github.com/felipe-coletti/my-first-crud-in-go/internal/user"
	userHTTP "github.com/felipe-coletti/my-first-crud-in-go/http/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := userDomain.NewUserService()
	userHandler := userHTTP.NewUserHandler(service)

	router := gin.Default()

	userHTTP.InitRoutes(&router.RouterGroup, userHandler)

	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
