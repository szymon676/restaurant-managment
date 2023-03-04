package models

type Order struct {
	ID      int64  `json:"id"`
	Product string `json:"product"`
	Table   int    `json:"table"`
}

type BindOrder struct {
	Product string `json:"product" binding:"required"`
	Table   int    `json:"table" binding:"required"`
}
