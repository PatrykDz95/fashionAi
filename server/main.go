package main

import (
	"log"

	"fasion.ai/server/ai"
	"fasion.ai/server/auth"
	"fasion.ai/server/db"
	"fasion.ai/server/recommendation"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database := db.InitDB()
	services := db.InitServices(database)
	userHandler := auth.NewUserHandler(services.UserService)
	recommendationHandler := recommendation.NewRecommendationHandler(services.RecommendationService)

	r := gin.Default()
	api := r.Group("/api")
	//api.Use(auth.JWTMiddleware())

	api.POST("/styleAdvice", ai.GetStyleAdvice)

	api.GET("/recommendations", recommendationHandler.GetRecommendations)
	api.GET("/recommendations/:id", recommendationHandler.GetRecommendationById)
	api.POST("/recommendations", recommendationHandler.SaveRecommendation)

	r.POST("/login", userHandler.LoginUser)
	r.POST("/register", userHandler.RegisterUser)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
