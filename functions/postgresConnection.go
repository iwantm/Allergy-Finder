package functions

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"os"
)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), 5432, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected" +
		"")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func AddProductToDatabase(product *Product) {
	db := Connect()
	sqlStatement := `
INSERT INTO products (barcode, brand, product_name, ingredients, allergens, allergens_tags, traces)
VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(sqlStatement, product.Barcode, product.Brand, product.ProductName, pq.Array(product.Ingredients), product.Allergens, pq.Array(product.AllergensTags), pq.Array(product.Traces))
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func SearchDatabase(barcode string) (*Product, error) {
	db := Connect()
	sqlStatement := `SELECT * FROM products WHERE barcode=$1;`
	var product Product
	row := db.QueryRow(sqlStatement, barcode)
	err := row.Scan(&product.Barcode, &product.Brand, &product.ProductName, pq.Array(&product.Ingredients), &product.Allergens, pq.Array(&product.AllergensTags), pq.Array(&product.Traces))
	switch err {
	case sql.ErrNoRows:
		product, err := SearchApi(barcode)
		if err == nil {
			AddProductToDatabase(product)
			return product, nil
		}
		return nil, err
	case nil:
		return &product, nil
	default:
		return nil, err
	}
}
