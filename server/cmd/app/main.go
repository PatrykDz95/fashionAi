package main

import (
	"log"

	"fasion.ai/server/internal/application/ai"
	"fasion.ai/server/internal/application/auth"
	"fasion.ai/server/internal/application/recommendation"
	domain_ai "fasion.ai/server/internal/domain/ai"
	infra_auth "fasion.ai/server/internal/infrastructure/auth"
	"fasion.ai/server/internal/infrastructure/db"
	infra_rec "fasion.ai/server/internal/infrastructure/recommendation"
	api_ai "fasion.ai/server/internal/interfaces/api/ai"
	api_auth "fasion.ai/server/internal/interfaces/api/auth"
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
	recommendationRepo := infra_rec.NewRepository(database)
	recommendationService := recommendation.NewRecommendationService(recommendationRepo)
	recommendationHandler := api_rec.NewRecommendationHandler(recommendationService)

	aiHandler := api_ai.NewAIHandler(ai.NewAIService(&domain_ai.AIHandler{}))

	authRepo := infra_auth.NewRepository(database)
	authService := auth.NewAuthService(authRepo)
	authHandler := api_auth.NewAuthHandler(authService)

	r := gin.Default()

	api := r.Group("/api")
	api.Use(api_auth.JWTMiddleware())
	{
		api.GET("/recommendations", recommendationHandler.GetRecommendations)
		api.GET("/recommendations/:id", recommendationHandler.GetRecommendationById)
		api.POST("/recommendations", recommendationHandler.SaveRecommendation)
	}

	r.POST("/styleAdvice", aiHandler.GetStyleAdvice)
	// r.GET("/recommendations", recommendationHandler.GetRecommendations)
	// r.GET("/recommendations/:id", recommendationHandler.GetRecommendationById)
	// r.POST("/recommendations", recommendationHandler.SaveRecommendation)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.Run(":8080")
}
