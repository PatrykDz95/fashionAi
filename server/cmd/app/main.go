package main

import (
	"log"

	"fasion.ai/server/internal/application/ai"
	"fasion.ai/server/internal/application/auth"
	"fasion.ai/server/internal/application/recommendation"
	infra_ai "fasion.ai/server/internal/infrastructure/ai"
	infra_auth "fasion.ai/server/internal/infrastructure/auth"
	"fasion.ai/server/internal/infrastructure/db"
	infra_rec "fasion.ai/server/internal/infrastructure/recommendation"
	api_ai "fasion.ai/server/internal/interfaces/api/ai"
	api_auth "fasion.ai/server/internal/interfaces/api/auth"
	interface_auth "fasion.ai/server/internal/interfaces/api/auth"
	api_rec "fasion.ai/server/internal/interfaces/api/recommendation"
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

	authRepo := infra_auth.NewRepository(database)
	authService := auth.NewAuthService(authRepo)
	authHandler := interface_auth.NewAuthHandler(authService)

	recommendationRepo := infra_rec.NewRepository(database)
	recommendationService := recommendation.NewRecommendationService(recommendationRepo)
	recommendationHandler := api_rec.NewRecommendationHandler(recommendationService)

	aiClient := infra_ai.NewClient()
	aiService := ai.NewAIService(aiClient, recommendationService, authService)
	aiHandler := api_ai.NewAIHandler(aiService)

	r := gin.Default()

	api := r.Group("/api")
	api.Use(api_auth.JWTMiddleware())
	{
		api.POST("/styleAdvice", aiHandler.GetStyleAdvice)
		api.GET("/recommendations", recommendationHandler.GetRecommendations)
		api.GET("/recommendations/:id", recommendationHandler.GetRecommendationById)
		api.POST("/recommendations", recommendationHandler.SaveRecommendation)
	}

	// r.POST("/styleAdvice", aiHandler.GetStyleAdvice)
	// r.GET("/recommendations", recommendationHandler.GetRecommendations)
	// r.GET("/recommendations/:id", recommendationHandler.GetRecommendationById)
	// r.POST("/recommendations", recommendationHandler.SaveRecommendation)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.Run(":8080")
}
