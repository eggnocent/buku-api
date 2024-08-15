package book

import "time"

type Buku struct {
	ID        int
	Judul     string
	Karya     string
	Deskripsi string
	Harga     int
	Rating    int
	CreateAt  time.Time
	UpdateAt  time.Time
}
