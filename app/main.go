package main

import (
	"myapp/db"
	"myapp/router"
)

// func handler(writer http.ResponseWriter, _ *http.Request) {
// 	fmt.Fprint(writer, "Hello World")
// }

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":3000", nil)
	// godotenv.Load(".env")
	// db.Migrate()

	// DB接続
	db.Init()

	// API起動
	// controllers.StartWebServer()
	router.StartWebServer()
}