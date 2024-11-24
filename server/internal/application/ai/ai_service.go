package ai

import (
	"time"

	app_rec "fasion.ai/server/internal/application/recommendation"
	domain_ai "fasion.ai/server/internal/domain/ai"
	"fasion.ai/server/internal/domain/recommendation"
	"fasion.ai/server/internal/infrastructure/ai"
)

type AIService struct {
	aiClient              ai.Client
	recommendationService *app_rec.RecommendationService
}

func NewAIService(aiClient ai.Client, recommendationService *app_rec.RecommendationService) *AIService {
	return &AIService{aiClient: aiClient, recommendationService: recommendationService}
}

func (s *AIService) GetStyleAdvice(input string) ([]recommendation.Item, error) {
	prompt := domain_ai.ReadPrompt(input)
	response, err := s.aiClient.GetChatGPTResponse(prompt)
	if err != nil {
		return nil, err
	}

	// Save the recommendation
	outfit := &recommendation.Outfit{
		// Populate the outfit fields based on the response
		RecommendedItems: response,
		DateCreated:      time.Now(),
	}
	if err := s.recommendationService.SaveRecommendation(outfit); err != nil {
		return nil, err
	}

	return response, nil
}
