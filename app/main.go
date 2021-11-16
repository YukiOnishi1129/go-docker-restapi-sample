package main

import (
	"myapp/db"
	"myapp/router"
)

func main() {
	// DB接続
	db.Init()

	// API起動
	router.StartWebServer()
}