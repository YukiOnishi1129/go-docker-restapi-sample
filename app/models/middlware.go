package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	ID        uint        `gorm:"primary_key" json:"id"`
    CreatedAt *time.Time  `json:"created_at"`
    UpdatedAt *time.Time  `json:"updated_at"`
    DeletedAt *time.Time  `json:"deleted_at"`
}

type Item struct {
    Model
    JanCode      string     `gorm:"size:255" json:"jan_code,omitempty"`
    ItemName     string     `gorm:"size:255" json:"item_name,omitempty"`
    Price        int        `json:"price,omitempty"`
    CategoryId   int        `json:"category_id,omitempty"`
    SeriesId     int        `json:"series_id,omitempty"`
    Stock        int        `json:"stock,omitempty"`
    Discontinued bool       `json:"discontinued"`
    ReleaseDate  *time.Time `json:"release_date,omitempty"`
}