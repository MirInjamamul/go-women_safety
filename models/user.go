package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Complain struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Thana        string    `json:"thana"`
	ComplainType string    `json:"complain_type"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
