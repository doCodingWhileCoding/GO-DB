package main

import (
	"log"

	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceheader"
	"github.com/doCodingWhileCoding/GO-DB/pkg/product"
	"github.com/doCodingWhileCoding/GO-DB/storage"
)

func main() {
	storage.NewMysqlDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
	storageInvoiceHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}
}
