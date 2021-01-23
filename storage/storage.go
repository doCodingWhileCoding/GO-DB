package storage

import (
	"database/sql"
	"log"
	"sync"

	//hjhj h
	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

//NewMysqlDB es una función que limita el comienzo de la db a una sola vez, de manera que si se vuelve a llamr a la función esta no se volverá a ejecutar una vez ya ejecutada
func NewMysqlDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/pruebago")
		if err != nil {
			log.Fatalf("can´t open db: %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("can´t do ping: %v", err)
		}
		//fmt.Println("Prueba")
	})

}

//Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}
