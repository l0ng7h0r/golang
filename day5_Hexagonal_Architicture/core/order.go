package core

type Order struct {
	ID uint `gorm:"primary_key; autoIncrement"`
	Total float64 `json:"total"`
}