package ai

import (
	"fmt"
	"log"
	"os"
)

func ReadPrompt(userInput string) string {
	content, err := os.ReadFile("prompts/recommendation")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	promptTemplate := string(content)
	return fmt.Sprintf(promptTemplate, userInput)
}
