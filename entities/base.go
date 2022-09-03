package entities

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
