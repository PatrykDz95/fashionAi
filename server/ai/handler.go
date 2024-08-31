package ai

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
