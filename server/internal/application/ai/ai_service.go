package ai

import (
	"time"

	appAuth "fasion.ai/server/internal/application/auth"
	appRec "fasion.ai/server/internal/application/recommendation"
	domainAI "fasion.ai/server/internal/domain/ai"
	"fasion.ai/server/internal/domain/recommendation"
	"fasion.ai/server/internal/infrastructure/ai"
)

type AIService struct {
	aiClient              ai.Client
	recommendationService *appRec.RecommendationService
	authService           *appAuth.AuthService
}

func NewAIService(aiClient ai.Client,
	recommendationService *appRec.RecommendationService,
	authService *appAuth.AuthService) *AIService {
	return &AIService{
		aiClient:              aiClient,
		recommendationService: recommendationService,
		authService:           authService}
}

func (s *AIService) GetStyleAdvice(input, username, season, occasion string) ([]recommendation.Item, error) {
	user, err := s.authService.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	prompt := domainAI.ReadPrompt(input, season, occasion)
	response, err := s.aiClient.GetChatGPTResponse(prompt)
	if err != nil {
		return nil, err
	}

	outfit := &recommendation.Outfit{
		UserID:           user.ID,
		Season:           recommendation.Season(season),
		Occasion:         occasion,
		RecommendedItems: response,
		DateCreated:      time.Now(),
	}
	if err := s.recommendationService.SaveRecommendation(outfit); err != nil {
		return nil, err
	}

	return response, nil
}
