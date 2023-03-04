package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `
		CREATE TABLE products (
		    "id" string,
		    "name" string,
		    "price" float,
		    "status" string
		);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatalf("Error preparing table: %s", err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ("123", "Product test", 12.0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalf("Error preparing insert: %s", err.Error())
	}
	stmt.Exec()

}
