package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"nama": "M Iksan",
			"bio":  "pemain sepakbola asal haiti",
		})
	})
	router.Run()
}
