package model

import "time"

// Session this is the representation of session
type Session struct {
	ID          int       `gorm:"primary_key" json:"id"`
	IDWorkspace int       `validate:"required" json:"id_workspace"`
	Name        string    `validate:"gte=10,lte=50" json:"name"`
	Description string    `validate:"gte=10,lte=255" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBY   string    `json:"updated_by"`
}
