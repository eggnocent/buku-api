package main

import (
	"buku-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/fact", handler.FactHandler)
	v1.GET("/buku/:id/:title", handler.BukuHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/buku", handler.CreateBukuHandler)

	router.Run()
}
