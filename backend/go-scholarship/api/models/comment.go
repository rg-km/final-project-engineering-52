package models

type Comment struct {
	ID            int64  `json:"id" form:"id"`
	Content       string `json:"content" form:"content" binding:"required"`
	UserID        int64  `json:"user_id" form:"user_id" binding:"required"`
	ScholarshipID int64  `json:"scholarship_id" form:"scholarship_id" binding:"required"`
	CreatedAt     string `json:"created_at" form:"created_at"`
}

type CommentRepository interface {
	Fetch() ([]Comment, error)
}

type CommentUseCase interface {}
