package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewUserHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) RegisterUser(c *gin.Context) {
	user := &User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.Service.Create(c, user)
}

func (h *Handler) LoginUser(c *gin.Context) {
	user := &User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.Service.Login(c, user)
}
