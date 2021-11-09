package main

import (
	"fmt"
	"myapp/db"
	"myapp/models"

	"gorm.io/gorm"
)

func seeds(dbConn *gorm.DB) error {
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
		if err := dbConn.Create(&item).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}


func main() {
	dbConn := db.OpenConnection()
	// dBを閉じる
	DB, _ := dbConn.DB()
	defer DB.Close()

	if err := seeds(dbConn); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}
