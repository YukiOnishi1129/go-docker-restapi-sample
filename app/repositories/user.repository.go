package repositories

import (
	"myapp/db"
	"myapp/models"
)

/*
 emailに紐づくユーザーリストを取得
*/
func GetUserByEmail(user *models.User, email string) error {
	db := db.GetDB()
	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		return err
	}

	return nil
}

/*
 emailに紐づくユーザーリストを取得
*/
func GetAllUserByEmail(users *[]models.User, email string) error {
	db := db.GetDB()
	if err := db.Where("email=?", email).Find(&users).Error; err != nil {
		return err
	}

	return nil
}

/*
  ユーザーデータ新規登録
*/
func CreateUser(createUsers *models.User) error {
	db := db.GetDB()
	if err := db.Create(&createUsers).Error; err != nil {
		return err
	}

	return nil
}