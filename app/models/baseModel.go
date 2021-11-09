package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	gorm.Model
	ID        uint        `gorm:"primary_key" json:"id"`
    CreatedAt *time.Time  `json:"created_at"`
    UpdatedAt *time.Time  `json:"updated_at"`
    DeletedAt *time.Time  `json:"deleted_at"`
}