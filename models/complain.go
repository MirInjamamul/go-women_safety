package models

import "time"

type Complain struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Thana        string    `json:"thana"`
	ComplainType string    `json:"complain_type"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
