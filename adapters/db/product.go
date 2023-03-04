package db

import (
	"database/sql"
	"github.com/fabioods/fc_hexagonal_arch/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil

}
