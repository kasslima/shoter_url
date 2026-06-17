package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kasslima/url-shortener/internal/feature/shorturl"
)



func RegisterRoutes(router *gin.Engine, h *shorturl.Handler) {
	router.POST("/short-url", h.CreateShortURL)

	router.GET("/short-url", h.ListShortURL)
}
