package models

type Scholarship struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CategoryID  int    `json:"category_id"`
	CreatedAt   string `json:"created_at"`
}

type ScholarshipRepository interface {
	Fetch() ([]Scholarship, error)
}

type ScholarshipUseCase interface{}
