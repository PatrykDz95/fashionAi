package ai

import (
	"fmt"
	"log"
	"os"

	"fasion.ai/server/internal/domain/recommendation"
	"fasion.ai/server/internal/infrastructure/ai"
)

type AIHandler struct{}

type AIService struct {
	aiHandler AIHandler
}

// TODO remove/move to interface
func (h *AIHandler) GetStyleAdvice(input string) ([]recommendation.Item, error) {
	prompt := readPrompt(input)
	print(prompt)
	response, err := ai.GetChatGPTResponse(prompt)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func readPrompt(userInput string) string {
	content, err := os.ReadFile("prompts/recommendation")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	promptTemplate := string(content)

	return fmt.Sprintf(promptTemplate, userInput)
}
