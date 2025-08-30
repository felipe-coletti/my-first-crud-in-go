package main

import (
	"log"
	"os"

	"github.com/felipe-coletti/my-first-crud-in-go/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	http.InitRoutes(&router.RouterGroup)

	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
