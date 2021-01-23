package storage

import (
	"database/sql"
	"fmt"

	"github.com/doCodingWhileCoding/GO-DB/pkg/invoice"
	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceheader"
	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceitem"
)

//MySQLInvoice used for work with mysql-Invoice
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

//NewMySQLInvoice return a new pointer of MySQLInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

//Create implements the interface Invoice.storage
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return nil
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("items creados: %d \n", len(m.Items))
	return tx.Commit()

}
