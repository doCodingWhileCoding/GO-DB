package product

import (
	"errors"
	"fmt"
)

var (
	//ErrIDNotFound es un error personalizado
	ErrIDNotFound = errors.New("El producto no contiene un ID")
)

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
}

//Esto es como el ToString() de Java
func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d\n",
		m.ID, m.Name, m.Observations, m.Price)
}

//Models slice of Model
type Models []*Model

//Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
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

//GetAll is used for get all the products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

//GetByID is used for get a product
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

//Update is used for update a product
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	return s.storage.Update(m)
}

//Delete is used for delete a product
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
