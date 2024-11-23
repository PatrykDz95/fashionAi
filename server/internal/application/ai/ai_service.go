package ai

import (
	"fasion.ai/server/internal/domain/ai"
	"fasion.ai/server/internal/domain/recommendation"
)

type AIService struct {
	aiHandler *ai.AIHandler
}

func NewAIService(aiHandler *ai.AIHandler) *AIService {
	return &AIService{aiHandler: aiHandler}
}

func (s *AIService) GetStyleAdvice(prompt string) ([]recommendation.Item, error) {
	return s.aiHandler.GetStyleAdvice(prompt)
}
