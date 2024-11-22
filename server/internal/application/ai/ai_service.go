package ai

import "fasion.ai/server/internal/domain/ai"

type AIService struct {
	aiHandler *ai.AIHandler
}

func NewAIService(aiHandler *ai.AIHandler) *AIService {
	return &AIService{aiHandler: aiHandler}
}

func (s *AIService) GetStyleAdvice(prompt string) (string, error) {
	return s.aiHandler.GetStyleAdvice(prompt)
}
