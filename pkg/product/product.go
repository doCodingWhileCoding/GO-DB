package product

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
}

//Models slice of Model
type Models []*Model

//Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GetById(uint) (*Model, error)
	//Delete(uint) error
}

//Service of product
type Service struct {
	storage Storage
}

//NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate is used for Migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

//Create is used for Create a product
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
