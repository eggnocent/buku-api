package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Buku, error)
	FindByID(id int) (Buku, error)
	Create(buku Buku) (Buku, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Buku, error) {
	var bukus []Buku
	err := r.db.Find(&bukus).Error
	return bukus, err
}

func (r *repository) FindByID(ID int) (Buku, error) { // Mengembalikan Buku
	var buku Buku
	err := r.db.First(&buku, ID).Error
	return buku, err
}

func (r *repository) Create(buku Buku) (Buku, error) {

	err := r.db.Create(&buku).Error
	return buku, err
}
