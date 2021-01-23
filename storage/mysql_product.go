package storage

import (
	"database/sql"
	"fmt"

	"github.com/doCodingWhileCoding/GO-DB/pkg/product"
)

const (
	mySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
			id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
			name VARCHAR(25) NOT NULL,
			observations VARCHAR(100),
			price INT NOT NULL
		)`
	mySQLCreateProduct = `INSERT INTO products(name, observations, price) VALUES(?,?,?)`
)

//MySQLProduct used for work with mysql-product
type MySQLProduct struct {
	db *sql.DB
}

//NewMySQLProduct return a new pointer of MySQLProduct
func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

//Migrate implements the interface product.storage
func (p *MySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migración del producto ejecutada")
	return nil
}

//Create implements the interface product.storage
func (p *MySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		m.Observations,
		m.Price,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Printf("Se creo el producto correctamente con ID: %d", m.ID)
	return nil
}
