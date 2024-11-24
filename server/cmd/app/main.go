package main

import (
	"log"

	"fasion.ai/server/internal/application/ai"
	"fasion.ai/server/internal/application/auth"
	"fasion.ai/server/internal/application/recommendation"
	infraAI "fasion.ai/server/internal/infrastructure/ai"
	infraAuth "fasion.ai/server/internal/infrastructure/auth"
	"fasion.ai/server/internal/infrastructure/db"
	infraRec "fasion.ai/server/internal/infrastructure/recommendation"
	apiAI "fasion.ai/server/internal/interfaces/api/ai"
	apiAuth "fasion.ai/server/internal/interfaces/api/auth"
	apiRec "fasion.ai/server/internal/interfaces/api/recommendation"
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

	authRepo := infraAuth.NewRepository(database)
	authService := auth.NewAuthService(authRepo)
	authHandler := apiAuth.NewAuthHandler(authService)

	recommendationRepo := infraRec.NewRepository(database)
	recommendationService := recommendation.NewRecommendationService(recommendationRepo)
	recommendationHandler := apiRec.NewRecommendationHandler(recommendationService)

	aiClient := infraAI.NewClient()
	aiService := ai.NewAIService(aiClient, recommendationService, authService)
	aiHandler := apiAI.NewAIHandler(aiService)

	r := gin.Default()

	api := r.Group("/api")
	api.Use(apiAuth.JWTMiddleware())
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
