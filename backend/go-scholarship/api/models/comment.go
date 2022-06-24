package models

import "context"

type CommentRequest struct {
	ID            int64  `json:"id"`
	Content       string `json:"content" binding:"required"`
	UserID        int64  `json:"user_id" binding:"required"`
	ScholarshipID int64  `json:"scholarship_id" binding:"required"`
	CreatedAt     string `json:"created_at"`
}

type CommentResponse struct {
	ID            int64        `json:"id"`
	Content       string       `json:"content"`
	User          UserResponse `json:"user"`
	ScholarshipID int64        `json:"scholarship_id"`
	CreatedAt     string       `json:"created_at"`
}

type CommentRepository interface {
	Fetch(ctx context.Context) ([]CommentResponse, error)
	FetchById(ctx context.Context, id int64) (CommentResponse, error)
	Create(ctx context.Context, commentReq *CommentRequest) (CommentResponse, error)
	Update(ctx context.Context, id int64, commentReq *CommentRequest) (CommentResponse, error)
	Delete(ctx context.Context, id int64) error
}

type CommentUseCase interface {
	Fetch(ctx context.Context) ([]CommentResponse, error)
	FetchById(ctx context.Context, id int64) (CommentResponse, error)
	Create(ctx context.Context, commentReq *CommentRequest) (CommentResponse, error)
	Update(ctx context.Context, id int64, commentReq *CommentRequest) (CommentResponse, error)
	Delete(ctx context.Context, id int64) error
}
