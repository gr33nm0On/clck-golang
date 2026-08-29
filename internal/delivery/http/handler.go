package http

import (
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetShortenHandler(ctx context.Context, pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		hash := service.Hash(url)

		shortenUrl := fmt.Sprintf("%s://%s/%s", scheme, host, hash)

		repository.Save(ctx, pool, url, shortenUrl)

		c.JSON(http.StatusOK, gin.H{
			"short_url": shortenUrl,
		})
	}
}
