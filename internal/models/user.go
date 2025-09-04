package models

import "time"

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      *string    `json:"role,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}
