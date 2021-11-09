package models

import (
	"time"
)

type Item struct {
    BaseModel
    JanCode      string     `gorm:"size:255" json:"jan_code,omitempty"`
    ItemName     string     `gorm:"size:255" json:"item_name,omitempty"`
    Price        int        `json:"price,omitempty"`
    CategoryId   int        `json:"category_id,omitempty"`
    SeriesId     int        `json:"series_id,omitempty"`
    Stock        int        `json:"stock,omitempty"`
    Discontinued bool       `json:"discontinued"`
    ReleaseDate  *time.Time `json:"release_date,omitempty"`
}