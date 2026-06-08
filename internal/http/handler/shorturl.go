package handler

import (
	gin "github.com/gin-gonic/gin"
	shorturldto "github.com/kasslima/url-shortener/internal/http/dto/shorturl"
)

func (h *Handler) CreateShortURL(c *gin.Context) {
	var req shorturldto.CreateShortURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"shorturl": req.LongURL})
}

func (h *Handler) ListShortURL(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ListShortURL"})
}
