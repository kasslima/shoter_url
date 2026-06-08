// cmd/api/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kasslima/url-shortener/internal/http/handler"
  "github.com/kasslima/url-shortener/internal/http"
)

func main() {
	router := gin.Default()

	h := &handler.Handler{}

	http.RegisterRoutes(router, h)

	router.Run(":8080")
}