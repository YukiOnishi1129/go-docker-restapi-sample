package main

import (
	"myapp/databases"
	"myapp/models"

	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	// Migration実行
	db.AutoMigrate(&models.Item{})
}

func main() {
	db := databases.OpenConnection()
	// dBを閉じる
	DB, _ := db.DB()
	defer DB.Close()

	// migration実行
	migrate(db)
}