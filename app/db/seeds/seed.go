package main

import (
	"fmt"
	"myapp/db"
	"myapp/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func userSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := models.User {
			Name: "ユーザー"+strconv.Itoa(i+1),
			Email: "sample"+strconv.Itoa(i+1)+"@gmail.com",
			Password: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}


func todoSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		var userId int
		if i < 5 {
			userId = 1
		} else {
			userId = 2
		}
		todo := models.Todo {
			Title: "タイトル"+strconv.Itoa(i+1),
			Comment: "コメント"+strconv.Itoa(i+1),
			UserId: userId,
		}

		if err := db.Create(&todo).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}


func main() {
	db.Init()
	dbCon := db.GetDB()
	// dBを閉じる
	defer db.CloseDB()

	if err := userSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}
	
	if err := todoSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}
