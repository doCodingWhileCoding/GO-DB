package invoiceheader

import "time"

//Model of invoice header
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	//Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GetById(uint) (*Model, error)
	//Delete(uint) error
}

//Service of invoiceheader
type Service struct {
	storage Storage
}

//NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate is used for Migrate
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
