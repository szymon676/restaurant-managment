package models

import "time"

type Rating struct {
	ID      int       `json:"id"`
	Stars   float64   `json:"stars"`
	User    string    `json:"user"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}

type BindRating struct {
	Stars   float64   `json:"stars" binding:"required"`
	User    string    `json:"user" binding:"required"`
	Comment string    `json:"comment" binding:"required"`
	Date    time.Time `json:"date" binding:"required"`
}
