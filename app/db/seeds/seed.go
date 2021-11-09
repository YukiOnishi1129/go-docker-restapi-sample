package main

import (
	"fmt"
	"myapp/db"
	"myapp/models"

	"gorm.io/gorm"
)

func seeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		item := models.Item{
			JanCode: "111",
			ItemName: "タイトル",
			Price: 111,
			CategoryId: 1,
			SeriesId: 1,
			Stock: 1,
			Discontinued: false,
		}
		if err := db.Create(&item).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}


func main() {
	db.Init()
	dbCon := db.GetDB()
	// dBを閉じる
	defer db.CloseDB()

	if err := seeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}
