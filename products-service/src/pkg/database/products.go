package database

import (
	"fmt"

	"github.com/szymon676/products-service/src/pkg/models"
)

func SaveProduct(title string, price float64, desc string) error {
	query := "INSERT INTO products (title, price, description) VALUES($1, $2, $3)"

	_, err := DB.Query(query, title, price, desc)
	if err != nil {
		return err
	}

	return nil
}

func GetProducts() (
	[]models.Product, error) {
	query := "SELECT * FROM products"

	rows, err := DB.Query(query)
	if err != nil {
		fmt.Printf("%c", err)
		return nil, err
	}

	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Title, &product.Price, &product.Description)
		if err != nil {
			fmt.Printf("%c", err)
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func UpdateProduct(id, title string, price float64, desc string) error {
	var count int

	if err := DB.QueryRow("SELECT COUNT(*) FROM products WHERE id = $1", id).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("product with id %s does not exist", id)
	}

	query := "UPDATE products SET title = $1, price = $2, description = $3 WHERE id = $4;"
	_, err := DB.Exec(query, title, price, desc, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id string) error {
	var count int

	if err := DB.QueryRow("SELECT COUNT(*) FROM products WHERE id = $1", id).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("product with id %s does not exist", id)
	}

	query := "DELETE FROM products WHERE id = $1"

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
