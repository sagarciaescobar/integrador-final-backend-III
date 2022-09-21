package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file:" + err.Error())
	}
	r := gin.Default()
	r.Run(":8080")
}
