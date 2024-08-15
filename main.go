package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/fact", factHandler)
	router.GET("/buku/:id/:title", bukuHandler)
	router.GET("/query", queryHandler)
	router.POST("/buku", createBukuHandler)
	router.Run()
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"nama": "M Iksan",
		"bio":  "pemain sepakbola asal haiti",
	})
}

func factHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"funfact":         "Bisa melakukan dribbling dengan mata kaki",
		"makanan favorit": "tengkleng pak bayu",
	})
}

func bukuHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(ctx *gin.Context) {
	judul := ctx.Query("judul")
	harga := ctx.Query("harga")
	ctx.JSON(http.StatusOK, gin.H{
		"judul": judul,
		"harga": harga,
	})
}

type BukuInput struct {
	Judul string `json:"judul" binding:"required"`
	Harga int    `json:"harga" binding:"required,number"`
}

func createBukuHandler(ctx *gin.Context) {
	var BukuInput BukuInput

	err := ctx.ShouldBindJSON(&BukuInput)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on fiel %s, condition: %s", e.Field(), e.ActualTag())
			ctx.JSON(http.StatusBadRequest, errorMessage)
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{
		"judul": BukuInput.Judul,
		"harga": BukuInput.Harga,
	})
}
