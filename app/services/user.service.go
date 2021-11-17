package services

import (
	"myapp/db"
	"myapp/models"
)

/*
 emailに紐づくユーザーを取得
*/
func GetUserByEmail(users *[]models.User, email string) {
	db := db.GetDB()
	db.Where("email=?", email).Find(&users)
}