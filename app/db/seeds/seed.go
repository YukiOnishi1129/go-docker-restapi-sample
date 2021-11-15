package main

import (
	"fmt"
	"myapp/db"
	"myapp/models"
	"strconv"

	"gorm.io/gorm"
)

func todoSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		todo := models.Todo {
			Title: "タイトル"+strconv.Itoa(i+1),
			Comment: "コメント"+strconv.Itoa(i+1),
		}

		if err := db.Create(&todo).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

// func itemSeeds(db *gorm.DB) error {
// 	for i := 0; i < 10; i++ {
// 		item := models.Item{
// 			JanCode: "00"+strconv.Itoa(i),
// 			ItemName: "item_"+strconv.Itoa(i),
// 			Price: 111 * i,
// 			CategoryId: 1,
// 			SeriesId: 1,
// 			Stock: 1,
// 			Discontinued: false,
// 		}
// 		if err := db.Create(&item).Error; err != nil {
// 			fmt.Printf("%+v", err)
// 		}
// 	}
// 	return nil
// }


func main() {
	db.Init()
	dbCon := db.GetDB()
	// dBを閉じる
	defer db.CloseDB()

	if err := todoSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}
