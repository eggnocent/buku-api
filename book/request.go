package book

import "encoding/json"

type BukuRequest struct {
	Judul string      `json:"judul" binding:"required"`
	Harga json.Number `json:"harga" binding:"required,number"`
}
