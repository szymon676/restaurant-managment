package database

import (
	"time"

	"github.com/fatih/color"
	"github.com/szymon676/ratings-service/internal/models"
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

func SaveRating(stars float64, user, comment string, date time.Time) error {
	query := "INSERT INTO ratings (stars, username, comment, date) VALUES($1, $2, $3, $4)"

	_, err := DB.Query(query, stars, user, comment, date)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRating(id string) error {
	query := "DELETE FROM ratings WHERE id = $1"
	_, err := DB.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
}
