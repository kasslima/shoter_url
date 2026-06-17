package shorturl

import (
	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateShortURL(c *gin.Context) {
	var req CreateShortURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"shorturl": req.LongURL})
}

func (h *Handler) ListShortURL(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ListShortURL"})
}
