package main

import (
	"log"

	"github.com/doCodingWhileCoding/GO-DB/pkg/invoice"
	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceheader"
	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceitem"
	"github.com/doCodingWhileCoding/GO-DB/storage"
)

func main() {
	storage.NewMysqlDB()

	//Creación de las tablas//////////////////////////////////////////

	/*storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
	storageInvoiceHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}
	storageInvoiceItem := storage.NewMySQLInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}*/

	//Inserción de datos (CREATE)///////////////////////////////////////////

	/*storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:         "Curso de Js",
		Price:        65,
		Observations: "clasic",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}
	fmt.Printf("%+v\n", m)*/

	//Obtención de datos (READ)/////////////////////////////////////////////////

	//conjunto de filas//////////

	/*storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println(ms)*/

	//Una unica fila////////////

	/*storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(1)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con ese id")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default:
		fmt.Println(m)
	}*/

	//Alteración de datos (UPDATE) /////////////////////////////////////////////////////
	/*storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:  "Curso de testing con GO",
		Price: 50,
		ID:    1,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}*/

	//Eliminación de datos (DELETE) /////////////////////////////////////////////////
	/*storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(1)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}*/

	//Transacciones /////////////////////////////////////////////////////////////////
	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	storageInvoice := storage.NewMySQLInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Alexys",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 2},
			&invoiceitem.Model{ProductID: 3},
			&invoiceitem.Model{ProductID: 4},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
