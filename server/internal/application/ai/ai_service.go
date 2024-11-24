package ai

import (
	"time"

	app_auth "fasion.ai/server/internal/application/auth"
	app_rec "fasion.ai/server/internal/application/recommendation"
	domain_ai "fasion.ai/server/internal/domain/ai"
	"fasion.ai/server/internal/domain/recommendation"
	"fasion.ai/server/internal/infrastructure/ai"
)

type AIService struct {
	aiClient              ai.Client
	recommendationService *app_rec.RecommendationService
	authService           *app_auth.AuthService
}

func NewAIService(aiClient ai.Client,
	recommendationService *app_rec.RecommendationService,
	authService *app_auth.AuthService) *AIService {
	return &AIService{
		aiClient:              aiClient,
		recommendationService: recommendationService,
		authService:           authService}
}

func (s *AIService) GetStyleAdvice(input, username string) ([]recommendation.Item, error) {
	prompt := domain_ai.ReadPrompt(input)
	response, err := s.aiClient.GetChatGPTResponse(prompt)
	if err != nil {
		return nil, err
	}

	user, err := s.authService.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	outfit := &recommendation.Outfit{
		UserID:           user.ID,
		RecommendedItems: response,
		DateCreated:      time.Now(),
	}
	if err := s.recommendationService.SaveRecommendation(outfit); err != nil {
		return nil, err
	}

	return response, nil
}
