package db

import (
	"fmt"
	"myapp/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate() {
	// .envを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	// MySQLへの接続情報を定義
	dsn := os.Getenv(("MYSQL_USER")) +":"+os.Getenv(("MYSQL_PASSWORD")) +"@tcp("+ os.Getenv(("MYSQL_HOST")) +":" +os.Getenv(("MYSQL_PORT"))+ ")/"+ os.Getenv(("MYSQL_DATABASE")) +"?charset=utf8mb4&parseTime=True&loc=Local"
	// DBインスタンスを生成
	dbCon, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	// dBを閉じる
	DB, err := dbCon.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer DB.Close()

	// Migration実行
	dbCon.AutoMigrate(&models.Item{})
}