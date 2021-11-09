package main

import (
	"fmt"
	"myapp/databases"
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
	db := databases.OpenConnection()
	// dBを閉じる
	DB, _ := db.DB()
	defer DB.Close()

	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}
