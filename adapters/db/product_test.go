package db_test

import (
	"database/sql"
	"github.com/fabioods/fc_hexagonal_arch/adapters/db"
	"github.com/fabioods/fc_hexagonal_arch/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
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
	insert := `INSERT INTO products (id, name, price, status) VALUES ("123", "Product test", 0.0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalf("Error preparing insert: %s", err.Error())
	}
	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setup()
	defer Db.Close()
	productDB := db.NewProductDB(Db)
	product, err := productDB.Get("123")
	require.Nil(t, err)
	require.Equal(t, "123", product.GetID())
	require.Equal(t, "Product test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setup()
	defer Db.Close()
	productDB := db.NewProductDB(Db)
	product := application.NewProduct()
	product.Name = "Product test"
	product.Price = 10.0

	productResult, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Name = "Product test 2"
	product.Price = 20.0
	productUpdated, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productUpdated.GetName())
	require.Equal(t, product.Price, productUpdated.GetPrice())
	require.Equal(t, product.Status, productUpdated.GetStatus())
}
