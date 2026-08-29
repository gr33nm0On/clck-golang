package main

import (
	"awesomeProject/internal/delivery/http"
	"awesomeProject/internal/repository"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pool := repository.ConnectDB(context.Background())

	err := repository.InitDB(context.Background(), pool)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/", http.GetShortenHandler(context.Background(), pool))

	r.Run("0.0.0.0:8081")

	defer pool.Close()
}
