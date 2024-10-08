package recommendation

import (
	"net/http"

	"fasion.ai/server/ai"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewRecommendationHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetRecommendations(c *gin.Context) {
	recommendations, err := h.Service.GetRecommendations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}

func (h *Handler) GetRecommendationById(c *gin.Context) {
	recommendation, err := h.Service.GetRecommendationByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendation)
}

func (h *Handler) Recommend(c *gin.Context) {
	var userInput struct {
		Input string `json:"input"`
	}
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	aiHandler := ai.AIHandler{}
	aiHandler.GetStyleAdvice(userInput.Input, c)
}

func (h *Handler) SaveRecommendation(c *gin.Context) {
	var recommendation Outfit
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	if err := h.Service.SaveRecommendation(&recommendation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, recommendation)
}
