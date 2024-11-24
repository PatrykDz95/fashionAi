package recommendation

import (
	"net/http"
	"strconv"

	"fasion.ai/server/internal/application/recommendation"
	domainRec "fasion.ai/server/internal/domain/recommendation"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *recommendation.RecommendationService
}

func NewRecommendationHandler(service *recommendation.RecommendationService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetRecommendations(c *gin.Context) {
	recommendations, err := h.service.GetRecommendations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}

func (h *Handler) GetRecommendationById(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid ID"})
		return
	}
	recommendation, err := h.service.GetRecommendationByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendation)
}

func (h *Handler) SaveRecommendation(c *gin.Context) {
	var recommendation domainRec.Outfit
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	if err := h.service.SaveRecommendation(&recommendation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, recommendation)
}
