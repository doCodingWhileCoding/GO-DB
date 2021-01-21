package storage

import (
	"database/sql"
	"fmt"
)

const (
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
			id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
			client VARCHAR(100) NOT NULL,
		)`
)

//MySQLInvoiceHeader used for work with mysql-InvoiceHeader
type MySQLInvoiceHeader struct {
	db *sql.DB
}

//NewMySQLInvoiceHeader return a new pointer of MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

//Migrate implements the interface InvoiceHeader.storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migración del invoiceHeader ejecutada")
	return nil
}
