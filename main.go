package main

import (
	"buku-api/book"
	"buku-api/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "egiwira:12345@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}
	db.AutoMigrate(&book.Buku{})
	fmt.Println("Database berhasil terkoneksi")

	// CREATE
	// buku := book.Buku{}
	// buku.Judul = "kipas angin kesedot sampah"
	// buku.Harga = 250000
	// buku.Diskon = 11
	// buku.Rating = 5
	// buku.Deskripsi = "baguy"

	// err = db.Create((&buku)).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("gagal menambahkan data")
	// 	fmt.Println("========================")
	// }

	// READ
	// var buku book.Buku
	// err = db.First(&buku).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("gagal mengambil data")
	// 	fmt.Println("========================")
	// }

	// fmt.Println("Judul: ", buku.Judul)
	// fmt.Printf("objek buku %v", buku)

	// READ LAST
	var buku book.Buku
	err = db.Last(&buku).Error
	if err != nil {
		fmt.Println("========================")
		fmt.Println("gagal mengambil data:", err) // Cetak error detail
		fmt.Println("========================")
		return
	}

	fmt.Println("Judul: ", buku.Judul)
	fmt.Printf("objek buku: %+v\n", buku)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/fact", handler.FactHandler)
	v1.GET("/buku/:id/:title", handler.BukuHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/buku", handler.CreateBukuHandler)

	router.Run()
}
