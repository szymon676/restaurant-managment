package models

type Rating struct {
	ID      int     `json:"id"`
	Stars   float64 `json:"stars"`
	User    string  `json:"user"`
	Comment string  `json:"comment"`
	Date    string  `json:"date"`
}

type BindRating struct {
	Stars   float64 `json:"stars" binding:"required"`
	User    string  `json:"user" binding:"required"`
	Comment string  `json:"comment" binding:"required"`
	Date    string  `json:"date" binding:"required"`
}
