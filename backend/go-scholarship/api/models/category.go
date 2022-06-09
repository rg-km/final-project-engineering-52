package models

type Category struct {
	ID           int    `json:"id" from:"id"`
	CategoryName string `json:"category_name" from:"category_name" binding:"required"`
	CreatedAt    string `json:"created_at" from:"created_at"`
}

type CategoryRepository interface {
	Fetch() ([]Category, error)
}

type CategoryUseCase interface {}