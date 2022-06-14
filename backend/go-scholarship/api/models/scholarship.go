package models

import "context"

type Scholarship struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Category    Category `json:"category"`
	User        User     `json:"user"`
	CreatedAt   string   `json:"created_at"`
}

type ScholarshipRepository interface {
	Fetch(ctx context.Context) ([]Scholarship, error)
	FetchById(ctx context.Context, id int64) (Scholarship, error)
	Create(ctx context.Context, scholarship *Scholarship) (Scholarship, error)
	Update(ctx context.Context, id int64, scholarship *Scholarship) (Scholarship, error)
	Delete(ctx context.Context, id int64) error
}

type ScholarshipUseCase interface {
	Fetch(ctx context.Context) ([]Scholarship, error)
	FetchById(ctx context.Context, id int64) (Scholarship, error)
	Create(ctx context.Context, scholarship *Scholarship) (Scholarship, error)
	Update(ctx context.Context, id int64, scholarship *Scholarship) (Scholarship, error)
	Delete(ctx context.Context, id int64) error
}
