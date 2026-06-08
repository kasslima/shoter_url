// internal/http/routes.go
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kasslima/url-shortener/internal/http/handler"
)

func RegisterRoutes(router *gin.Engine, h *handler.Handler) {
	router.POST("/short-url", h.CreateShortURL)

	router.GET("/short-url", h.ListShortURL)
}