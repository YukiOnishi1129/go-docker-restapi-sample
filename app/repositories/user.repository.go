package repositories

import (
	"fmt"
	"myapp/db"
	"myapp/models"
)

/*
 emailに紐づくユーザーを取得
*/
func GetUserByEmail(users *[]models.User, email string) error {
	db := db.GetDB()
	if err := db.Where("email=?", email).Find(&users).Error; err != nil {
		fmt.Print(users)
		return err
	}

	return nil
}