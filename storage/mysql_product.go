package storage

import (
	"database/sql"
	"fmt"

	"github.com/doCodingWhileCoding/GO-DB/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	mySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
			id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
			name VARCHAR(25) NOT NULL,
			observations VARCHAR(100),
			price INT NOT NULL
		)`
	mySQLCreateProduct  = `INSERT INTO products(name, observations, price) VALUES(?,?,?)`
	mySQLGetAllProduct  = `SELECT id, name, observations, price FROM products`
	mySQLGetProductByID = mySQLGetAllProduct + " WHERE id = ?"
	mySQLUpdateProduct  = `UPDATE products SET name = ?, observations = ? , price = ? WHERE id = ?`
	mySQLDeleteProduct  = `DELETE FROM products WHERE id = ?`
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
		stringToNull(m.Observations),
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

	fmt.Printf("Se creo el producto correctamente con ID: %d\n", m.ID)
	return nil
}

//GetAll implements the interface product.storage
func (p *MySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mySQLGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

//GetByID implements the interface product.storage
func (p *MySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mySQLGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

//Update implements the interface product.storage
func (p *MySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con el id: %d", m.ID)
	}

	fmt.Println("se actualizo el producto correctamente")
	return nil
}

//Delete implements the interface product.storage
func (p *MySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mySQLDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con el id: %d", id)
	}

	fmt.Println("se eliminó el producto correctamente")
	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
	)
	if err != nil {
		return &product.Model{}, err
	}
	m.Observations = observationNull.String
	return m, nil
}
