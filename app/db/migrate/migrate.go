package main

import (
	"myapp/db"
	"myapp/models"

	"gorm.io/gorm"
)

func migrate(dbCon *gorm.DB) {
	// Migration実行
	dbCon.AutoMigrate(&models.User{}, &models.Todo{})
}

func main() {
	db.Init()
	dbCon := db.GetDB()
	// dBを閉じる
	defer db.CloseDB()

	// migration実行
	migrate(dbCon)
}