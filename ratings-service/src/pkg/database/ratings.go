package database

import (
	"github.com/fatih/color"
	"github.com/szymon676/ratings-service/src/pkg/models"
)

func GetRatings() ([]models.Rating, error) {
	query := "SELECT * FROM ratings;"
	rows, err := DB.Query(query)
	if err != nil {
		color.Red("%c", err)
		return nil, err
	}

	defer rows.Close()

	var ratings []models.Rating

	for rows.Next() {
		var rating models.Rating

		err := rows.Scan(&rating.ID, &rating.Stars, &rating.User, &rating.Comment, &rating.Date)
		if err != nil {
			color.Red("%c", err)
			return nil, err
		}

		ratings = append(ratings, rating)
	}

	err = rows.Err()
	if err != nil {
		color.Red("%c", err)
		return nil, err
	}

	return ratings, nil
}
