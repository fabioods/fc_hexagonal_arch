package main

import (
	"database/sql"
	db2 "github.com/fabioods/fc_hexagonal_arch/adapters/db"
	"github.com/fabioods/fc_hexagonal_arch/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDB(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product test", 10.0)
	product.UpdatePrice(0)
	productService.Disable(product)

}
