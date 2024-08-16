package book

type BukuResponse struct {
	ID        int    `json:"id"`
	Judul     string `json:"judul"`
	Karya     string `json:"karya"`
	Deskripsi string `json:"deskripsi"`
	Harga     int    `json:"harga"`
	Diskon    int    `json:"diskon"`
	Rating    int    `json:"rating"`
}
