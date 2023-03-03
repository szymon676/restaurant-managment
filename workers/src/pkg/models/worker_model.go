package models

type BindWorker struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Number string `json:"number" binding:"required"`
}

type Worker struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Number string `json:"number"`
}
