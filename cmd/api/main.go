// cmd/api/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kasslima/url-shortener/internal/feature/shorturl"
	"github.com/kasslima/url-shortener/internal/http"
)

func main() {
	router := gin.Default()

	shortURLService := shorturl.NewService()
	shortURLHandler := shorturl.NewHandler(shortURLService)

	http.RegisterRoutes(router, shortURLHandler)

	router.Run(":8080")
}
