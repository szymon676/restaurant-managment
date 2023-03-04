package database

func CreateOrder(product string, table int) error {
	query := "INSERT INTO orders (product,tableid) VALUES($1,$2);"

	_, err := DB.Exec(query, product, table)
	if err != nil {
		return err
	}

	return nil
}
