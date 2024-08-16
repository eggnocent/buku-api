package book

type Service interface {
	FindAll() ([]Buku, error)
	FindByID(ID int) (Buku, error) // Mengembalikan Buku, bukan []Buku
	Create(buku BukuRequest) (Buku, error)
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
		Judul: bukuRequest.Judul,
		Harga: int(harga),
	}
	return s.repository.Create(buku)
}
