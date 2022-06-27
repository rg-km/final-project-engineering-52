package models

import (
	"context"
	"time"
)

type User struct {
	ID         int64     `json:"id" form:"id"`
	Name       string    `json:"name" form:"name" binding:"required"`
	Pendidikan string    `json:"pendidikan" form:"pendidikan"`
	Image      string    `json:"image" form:"image"`
	Email      string    `json:"email" form:"email" binding:"required,email"`
	Password   string    `json:"password" form:"password" binding:"required,min=4"`
	Role       string    `json:"role" form:"role"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
}

type UserResponse struct {
	ID         int64     `json:"id" form:"id"`
	Name       string    `json:"name" form:"name"`
	Pendidikan string    `json:"pendidikan" form:"pendidikan"`
	Image      string    `json:"image" form:"image"`
	Email      string    `json:"email" form:"email"`
	Role       string    `json:"role" form:"role"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
}

type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=4"`
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]UserResponse, error)
	FetchById(ctx context.Context, id int64) (UserResponse, error)
	Create(ctx context.Context, u *User) (UserResponse, error)
	Update(ctx context.Context, id int64, u *User) (UserResponse, error)
	Delete(ctx context.Context, id int64) error
	Login(ctx context.Context, l *Login) (UserResponse, error)
	Register(ctx context.Context, u *User) (UserResponse, error)
}
