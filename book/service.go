package book

type Service interface {
	FindAll() ([]Buku, error)
	FindByID(ID int) (Buku, error)
	Create(buku BukuRequest) (Buku, error)
	Update(ID int, bukuRequest BukuRequest) (Buku, error) // Ubah dari 'Int' menjadi 'int'
	Delete(ID int) (Buku, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Buku, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Buku, error) { // Mengembalikan Buku
	return s.repository.FindByID(ID)
}

func (s *service) Create(bukuRequest BukuRequest) (Buku, error) {
	harga, _ := bukuRequest.Harga.Int64()
	buku := Buku{
		Judul:     bukuRequest.Judul,
		Karya:     bukuRequest.Karya,
		Harga:     int(harga),
		Deskripsi: bukuRequest.Deskripsi,
		Rating:    bukuRequest.Rating,
		Diskon:    bukuRequest.Diskon,
	}
	return s.repository.Create(buku)
}

func (s *service) Update(ID int, bukuRequest BukuRequest) (Buku, error) {
	// Temukan buku berdasarkan ID
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Buku{}, err
	}

	// Update atribut buku dengan data dari bukuRequest
	harga, _ := bukuRequest.Harga.Int64()

	book.Judul = bukuRequest.Judul
	book.Karya = bukuRequest.Karya
	book.Harga = int(harga)
	book.Deskripsi = bukuRequest.Deskripsi
	book.Rating = bukuRequest.Rating
	book.Diskon = bukuRequest.Diskon

	// Simpan perubahan ke repository
	updatedBook, err := s.repository.Update(book)
	if err != nil {
		return Buku{}, err
	}

	return updatedBook, nil
}

func (s *service) Delete(ID int) (Buku, error) {
	// Temukan buku berdasarkan ID
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Buku{}, err
	}

	updatedBook, err := s.repository.Delete(book)
	if err != nil {
		return Buku{}, err
	}

	return updatedBook, nil
}
