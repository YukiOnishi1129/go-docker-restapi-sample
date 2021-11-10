package service_item

import (
	"myapp/db"
	"myapp/models"
)

func GetAllItem(item *[]models.Item) {
	db := db.GetDB()
	db.Find(&item)
}

