package handler

import (
	"buku-api/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bukuHandler struct {
	bookService book.Service
}

func (h *bukuHandler) GetBuku(ctx *gin.Context) {
	buku, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var bukusResponse []book.BukuResponse

	for _, buku := range buku {
		bukuResponse := convertToBukuResponse(buku)
		bukusResponse = append(bukusResponse, bukuResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bukusResponse,
	})
}

func (h *bukuHandler) GetBukus(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	buku, err := h.bookService.FindByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	bukuResponse := convertToBukuResponse(buku)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bukuResponse,
	})
}

func (h *bukuHandler) CreateBukuHandler(ctx *gin.Context) {
	var BukuRequest book.BukuRequest

	err := ctx.ShouldBindJSON(&BukuRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(BukuRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToBukuResponse(book),
	})
}

func (h *bukuHandler) UpdateBukuHandler(ctx *gin.Context) {
	var BukuRequest book.BukuRequest

	err := ctx.ShouldBindJSON(&BukuRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, BukuRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToBukuResponse(book),
	})
}

func (h *bukuHandler) DeleteBukuHandler(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	buku, err := h.bookService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	bukuResponse := convertToBukuResponse(buku)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bukuResponse,
	})
}

func NewBookHandler(bookService book.Service) *bukuHandler {
	return &bukuHandler{bookService}
}

func (h *bukuHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"nama": "M Iksan",
		"bio":  "pemain sepakbola asal haiti",
	})
}

func (h *bukuHandler) FactHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"funfact":         "Bisa melakukan dribbling dengan mata kaki",
		"makanan favorit": "tengkleng pak bayu",
	})
}

func (h *bukuHandler) BukuHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bukuHandler) QueryHandler(ctx *gin.Context) {
	judul := ctx.Query("judul")
	harga := ctx.Query("harga")
	ctx.JSON(http.StatusOK, gin.H{
		"judul": judul,
		"harga": harga,
	})
}

func convertToBukuResponse(buku book.Buku) book.BukuResponse {
	return book.BukuResponse{
		ID:        buku.ID,
		Judul:     buku.Judul,
		Harga:     buku.Harga,
		Deskripsi: buku.Deskripsi,
		Rating:    buku.Rating,
		Diskon:    buku.Diskon,
	}
}
