package main

import (
	"fasion.ai/server/ai"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//db.InitDB()

	r := gin.Default()
	//api := r.Group("/api")
	//api.Use(common.JWTMiddleware())

	r.POST("/styleAdvice", ai.GetStyleAdvice)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
