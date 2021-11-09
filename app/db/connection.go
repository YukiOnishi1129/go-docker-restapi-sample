package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
* データベース接続処理
 */
func OpenConnection() *gorm.DB {
	// .envを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// MySQLへの接続情報を定義
	dsn := os.Getenv(("MYSQL_USER")) +":"+os.Getenv(("MYSQL_PASSWORD")) +"@tcp("+ os.Getenv(("MYSQL_HOST")) +":" +os.Getenv(("MYSQL_PORT"))+ ")/"+ os.Getenv(("MYSQL_DATABASE")) +"?charset=utf8mb4&parseTime=True&loc=Local"
	// DBインスタンスを生成
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("データベース接続に失敗しました。", err)
		return  nil
	}

	return db
}