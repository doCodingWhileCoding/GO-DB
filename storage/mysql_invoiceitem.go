package storage

import (
	"database/sql"
	"fmt"

	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceitem"
)

const (
	mySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS Invoice_items (
			id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
			invoice_header_id INT NOT NULL,
			product_id INT NOT NULL,
			CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY 
			(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE 
			RESTRICT ON DELETE RESTRICT,
			CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) 
			REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
		)`
	mySQLCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES (?, ?)`
)

//MySQLInvoiceItem used for work with mysql-InvoiceItem
type MySQLInvoiceItem struct {
	db *sql.DB
}

//NewMySQLInvoiceItem return a new pointer of MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

//Migrate implements the interface InvoiceItem.storage
func (p *MySQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migración del InvoiceItemo ejecutada")
	return nil
}

//CreateTx implements the interface InvoiceItem.storage
func (p *MySQLInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(headerID, item.ProductID)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		item.ID = uint(id)

	}

	return nil
}
