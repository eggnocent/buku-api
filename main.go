package main

import (
	"buku-api/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "egiwira:12345@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}

	fmt.Println("Database berhasil terkoneksi")
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/fact", handler.FactHandler)
	v1.GET("/buku/:id/:title", handler.BukuHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/buku", handler.CreateBukuHandler)

	router.Run()
}
