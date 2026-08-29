package http

import (
	"awesomeProject/internal/service"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Shorten(c *gin.Context) {
	url := c.PostForm("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url parameter is required"})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}

	host := c.Request.Host
	hash := service.Hash(context.Background(), url)

	shortenUrl := fmt.Sprintf("%s://%s/%s", scheme, host, hash)

	c.JSON(http.StatusOK, gin.H{
		"short_url": shortenUrl,
	})
}
