package main

import (
	"myapp/db"
	"myapp/models"

	"gorm.io/gorm"
)

func migrate(dbConn *gorm.DB) {
	// Migration実行
	dbConn.AutoMigrate(&models.Item{})
}

func main() {
	dbConn := db.OpenConnection()
	// dBを閉じる
	DB, _ := dbConn.DB()
	defer DB.Close()

	// migration実行
	migrate(dbConn)
}