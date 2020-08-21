package model

import (
	"time"
)

// WorkspaceModel struct
type WorkspaceModel struct {
	ID            int       `json:"id"`
	IDFacilitator int       `validate:"required" json:"id_facilitator"`
	Name          string    `validate:"gte=10,lte=50" json:"name"`
	Description   string    `validate:"gte=30,lte=255" json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBY     string    `json:"updated_by"`
}
