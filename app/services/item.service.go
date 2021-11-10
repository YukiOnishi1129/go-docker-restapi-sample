package service_item

import (
	"myapp/db"
	"myapp/models"
)

func GetAllItem(item *[]models.Item) {
	db := db.GetDB()
	db.Find(&item)
}

func GetSIngleItem(item *models.Item, key string) {
	db := db.GetDB()
	db.First(&item, key)
}

func InsertItem(item *models.Item) {
	db := db.GetDB()
	db.Create(&item)
}

func DeleteItem(key string) {
	db := db.GetDB()
	db.Where("id=?", key).Delete(&models.Item{})
}

func UpdateItem(item *models.Item, key string) {
	db := db.GetDB()
	db.Model(&item).Where("id=?", key).Updates(
        map[string]interface{}{
            "jan_code":     item.JanCode,
            "item_name":    item.ItemName,
            "price":        item.Price,
            "category_id":  item.CategoryId,
            "series_id":    item.SeriesId,
            "stock":        item.Stock,
            "discontinued": item.Discontinued,
            "release_date": item.ReleaseDate,
        })
}