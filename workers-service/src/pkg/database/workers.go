package database

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/szymon676/workers-service/src/pkg/models"
)

func GetWorkers() ([]models.Worker, error) {
	query := "SELECT * FROM workers;"
	rows, err := DB.Query(query)
	if err != nil {
		color.Red("%c", err)
		return nil, err
	}

	defer rows.Close()

	var workers []models.Worker

	for rows.Next() {
		var worker models.Worker

		err := rows.Scan(&worker.ID, &worker.Name, &worker.Email, &worker.Number)
		if err != nil {
			color.Red("%c", err)
			return nil, err
		}

		workers = append(workers, worker)
	}

	err = rows.Err()
	if err != nil {
		color.Red("%c", err)
		return nil, err
	}

	return workers, nil
}

func SaveWorker(name, email, number string) error {
	query := "INSERT INTO workers (name, email, number) VALUES($1, $2, $3)"

	_, err := DB.Query(query, name, email, number)
	if err != nil {
		return err
	}

	return nil
}

func UpdateWorker(id, name, email, number string) error {
	var count int

	if err := DB.QueryRow("SELECT COUNT(*) FROM workers WHERE id = $1", id).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("worker with id %s does not exist", id)
	}

	query := "UPDATE workers SET name = $1, email = $2, number = $3 WHERE id = $4;"
	_, err := DB.Exec(query, name, email, number, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	var count int

	if err := DB.QueryRow("SELECT COUNT(*) FROM workers WHERE id = $1", id).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("worker with id %s does not exist", id)
	}

	query := "DELETE FROM workers WHERE id = $1"

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
