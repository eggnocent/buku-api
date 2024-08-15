package book

import "encoding/json"

type BukuRequest struct {
	Judul     string      `json:"judul" binding:"required"`
	Harga     json.Number `json:"harga" binding:"required,number"`
	Deskripsi string      `json:"deskripsi" binding:"required"`
	Rating    int         `json:"rating" binding:"required,number"`
	Diskon    int         `json:"diskon" binding:"required,number"`
}
