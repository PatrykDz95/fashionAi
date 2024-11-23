package ai

import (
	"net/http"

	"fasion.ai/server/internal/application/ai"
	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	aiService *ai.AIService
}

func NewAIHandler(aiService *ai.AIService) *AIHandler {
	return &AIHandler{aiService: aiService}
}

func (h *AIHandler) GetStyleAdvice(c *gin.Context) {
	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.aiService.GetStyleAdvice(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"advice": response})
}
