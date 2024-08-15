package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/fact", factHandler)
	router.GET("/buku/:id/:title", bukuHandler)
	router.GET("/query", queryHandler)
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
