package ai

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var req struct {
	Prompt string `json:"prompt"`
}

func GetStyleAdvice(c *gin.Context) {

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := GetChatGPTResponse(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": response})
}

func readPrompt(userInput string) string {
	content, err := os.ReadFile("server/prompts/recommendation")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	promptTemplate := string(content)

	return fmt.Sprintf(promptTemplate, userInput)
}
