package ai

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AIHandler struct{}

var req struct {
	Prompt string `json:"prompt"`
}

func (h *AIHandler) GetStyleAdvice(input string, c *gin.Context) {
	prompt := readPrompt(input)
	print(prompt)
	response, err := GetChatGPTResponse(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	print(response)
	c.JSON(http.StatusOK, gin.H{"response": response})
}

func readPrompt(userInput string) string {
	content, err := os.ReadFile("prompts/recommendation")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	promptTemplate := string(content)

	return fmt.Sprintf(promptTemplate, userInput)
}
