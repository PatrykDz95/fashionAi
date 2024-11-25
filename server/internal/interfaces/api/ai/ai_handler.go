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
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in token"})
		return
	}

	var req struct {
		Prompt   string `json:"prompt"`
		Season   string `json:"season"`
		Occasion string `json:"occasion"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.aiService.GetStyleAdvice(req.Prompt, username.(string), req.Season, req.Occasion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"advice": response})
}
