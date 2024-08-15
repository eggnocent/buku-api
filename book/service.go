package book

type Service interface {
	FindAll() ([]Buku, error)
	FindByID(ID int) ([]Buku, error)
	Create(buku BukuRequest) (Buku, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Buku, error) {
	buku, err := s.repository.FindAll()
	return buku, err
	//return s.repository.FindAll() // cara simple
}

func (s *service) FindByID(ID int) (Buku, error) {
	buku, err := s.repository.FindByID(ID)
	return buku, err
}

func (s *service) Create(bukuRequest BukuRequest) (Buku, error) {

	harga, _ := bukuRequest.Harga.Int64()

	buku := Buku{
		Judul: bukuRequest.Judul,
		Harga: int(harga),
	}
	newBuku, err := s.repository.Create(buku)
	return newBuku, err
}
