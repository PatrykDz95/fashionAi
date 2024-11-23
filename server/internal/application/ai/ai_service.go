package ai

import (
	domain_ai "fasion.ai/server/internal/domain/ai"
	"fasion.ai/server/internal/domain/recommendation"
	"fasion.ai/server/internal/infrastructure/ai"
)

type AIService struct {
	aiClient ai.Client
}

func NewAIService(aiClient ai.Client) *AIService {
	return &AIService{aiClient: aiClient}
}

func (s *AIService) GetStyleAdvice(input string) ([]recommendation.Item, error) {
	prompt := domain_ai.ReadPrompt(input)
	response, err := s.aiClient.GetChatGPTResponse(prompt)
	if err != nil {
		return nil, err
	}
	return response, nil
}
