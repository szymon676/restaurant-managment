package models

type BindProduct struct {
	Title       string  `json:"title" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type Product struct {
	ID          int     `json:"id" `
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
