package recommendation

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
