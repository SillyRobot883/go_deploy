package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

//IsPublic  bool      `json:"is_public"`
//Followers []int     `json:"followers"`
//Following []int     `json:"following"`
