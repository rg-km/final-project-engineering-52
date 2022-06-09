package models

import (
	"time"
)

type User struct {
	ID           int64     `json:"id" form:"id"`
	Name         string    `json:"name" form:"name" binding:"required"`
	Image        string    `json:"image" form:"image"`
	Email        string    `json:"email" form:"email" binding:"required,email"`
	Password     string    `json:"password" form:"password" binding:"required,min=4"`
	Role         string    `json:"role" form:"role"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
}

type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=4"`
}

type UserRepository interface {
	Fetch() ([]User, error)
	FetchById(id int64) (User, error)
	Create(u *User) (User, error)
	Update(id int64, u *User) (User, error)
	Delete(id int64) error
	Login(l *Login) (User, error)
	Register(u *User) (User, error)
}

type UserUseCase interface{}
