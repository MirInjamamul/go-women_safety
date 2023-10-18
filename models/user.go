package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Password  string    `json:"password"`
	Token     string    `json:"token" gorm:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
