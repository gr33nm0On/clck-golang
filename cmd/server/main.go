package main

import (
	"awesomeProject/internal/delivery/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", http.Shorten)

	r.Run("0.0.0.0:8081")
}
