package database

import (
	"github.com/fatih/color"
	"github.com/szymon676/orders-service/internal/models"
)

func GetOrders() ([]models.Order, error) {
	query := "SELECT * FROM orders;"
	rows, err := DB.Query(query)
	if err != nil {
		color.Red("%c", err)
		return nil, err
	}

	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var order models.Order

		err := rows.Scan(&order.ID, &order.Product, &order.Table)
		if err != nil {
			color.Red("%c", err)
			return nil, err
		}

		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		color.Red("%c", err)
		return nil, err
	}

	return orders, nil
}

func CreateOrder(product string, table int) error {
	query := "INSERT INTO orders (product,tableid) VALUES($1,$2);"

	_, err := DB.Exec(query, product, table)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrder(id string) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
