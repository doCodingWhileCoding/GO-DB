package invoiceitem

//Model of invoice
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
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

//Service of invoiceitem
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
