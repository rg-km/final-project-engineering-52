package models

import "context"

type ScholarRequest struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	CategoryID  int64  `json:"category_id"`
	UserID      int64  `json:"user_id"`
	CreatedAt   string `json:"created_at"`
}
type ScholarResponse struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Category    Category `json:"category"`
	User        UserResponse     `json:"user"`
	CreatedAt   string   `json:"created_at"`
}

type ScholarshipRepository interface {
	Fetch(ctx context.Context) ([]ScholarResponse, error)
	FetchById(ctx context.Context, id int64) (ScholarResponse, error)
	Create(ctx context.Context, scholarReq *ScholarRequest) (ScholarResponse, error)
	Update(ctx context.Context, id int64, scholarReq *ScholarRequest) (ScholarResponse, error)
	Delete(ctx context.Context, id int64) error
}

type ScholarshipUseCase interface {
	Fetch(ctx context.Context) ([]ScholarResponse, error)
	FetchById(ctx context.Context, id int64) (ScholarResponse, error)
	Create(ctx context.Context, scholarReq *ScholarRequest) (ScholarResponse, error)
	Update(ctx context.Context, id int64, scholarReq *ScholarRequest) (ScholarResponse, error)
	Delete(ctx context.Context, id int64) error
}
