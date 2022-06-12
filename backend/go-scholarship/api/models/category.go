package models

import "context"

type Category struct {
	ID           int    `json:"id" from:"id"`
	CategoryName string `json:"category_name" from:"category_name" binding:"required"`
	CreatedAt    string `json:"created_at" from:"created_at"`
}

type CategoryRepository interface {
	Fetch(ctx context.Context) ([]Category, error)
	FetchById(ctx context.Context, id int64) (Category, error)
	Create(ctx context.Context, c *Category) (Category, error)
}
