package main

import (
	"myapp/controllers"
	"myapp/db"
)

// func handler(writer http.ResponseWriter, _ *http.Request) {
// 	fmt.Fprint(writer, "Hello World")
// }

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":3000", nil)
	// godotenv.Load(".env")
	db.Migrate()
	controllers.StartWebServer()
}