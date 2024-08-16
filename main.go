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
		log.Fatal("DB connection error:", err)
	}
	db.AutoMigrate(&book.Buku{})
	fmt.Println("Database berhasil terkoneksi")

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// FIND ALL
	// bukus, err := bookRepository.FindAll()
	// if err != nil {
	// 	fmt.Println("Error retrieving books:", err)
	// 	return
	// }

	// for _, b := range bukus {
	// 	fmt.Println("Title:", b.Judul)
	// }

	//FIND BY ID
	// bukus, err := bookRepository.FindByID(3)

	// fmt.Println("Judul: ", bukus.Judul)

	// buku := book.Buku{
	// 	Judul:     "Dengarkan Tanah",
	// 	Deskripsi: "dari dongker",
	// 	Harga:     120000,
	// 	Rating:    4,
	// 	Diskon:    10,
	// }

	//bookRepository.Create(buku)
	// bookRequest := book.BukuRequest{
	// 	Judul: "Merusak Kebahagiaan",
	// 	Harga: "200000",
	// }
	// bookService.Create(bookRequest)

	// CREATE
	// buku := book.Buku{}
	// buku.Judul = "kucing langit"
	// buku.Harga = 280000
	// buku.Diskon = 11
	// buku.Rating = 55
	// buku.Deskripsi = "baguy bangay"

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
	// 	fmt.Println("gagal menambahkan data")
	// 	fmt.Println("========================")
	// }

	// fmt.Println("Judul: ", buku.Judul)
	// fmt.Printf("objek buku %v", buku)

	// UPDATE
	// var buku book.Buku
	// err = db.Debug().Where("id = ?", 1).First(&buku).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("gagal mengupdate data")
	// 	fmt.Println("========================")
	// }

	// buku.Judul = "resep warung cah bagus"
	// err = db.Save(&buku).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("gagal mengupdate data")
	// 	fmt.Println("========================")
	// }

	// DELETE
	// var buku book.Buku
	// err = db.Debug().Where("id = ?", 1).First(&buku).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("gagal menemukan data")
	// 	fmt.Println("========================")
	// }

	// err = db.Delete(&buku).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("gagal menghapus data")
	// 	fmt.Println("========================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/fact", bookHandler.FactHandler)
	v1.GET("/buku/:id/:title", bookHandler.BukuHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/bukus", bookHandler.GetBuku)
	v1.GET("/bukus/:id", bookHandler.GetBukus)
	v1.POST("/buku", bookHandler.CreateBukuHandler)
	v1.PUT("/bukus/:id", bookHandler.UpdateBukuHandler)
	router.Run()
}
