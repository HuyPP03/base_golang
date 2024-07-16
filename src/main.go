package main

import (
	"log"
	"os"

	"github.com/HuyPP03/learn/src/database"
	"github.com/HuyPP03/learn/src/loaders"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	database.ConnectDatabase()
	loaders.ConfigLoader(r)

	r.Run(":" + port)
}
