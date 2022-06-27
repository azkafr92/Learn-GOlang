package domain

import (
	"time"

	"gorm.io/gorm"
)

type Common struct {
	ID        int             `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}