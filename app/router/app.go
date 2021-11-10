package router

import (
	"fmt"
	controller_item "myapp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

/*
 ルーティング定義
*/
func setupRouting() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller_item.RootPage)
    router.HandleFunc("/items", controller_item.FetchAllItems).Methods("GET")
    router.HandleFunc("/item/{id}", controller_item.FetchSingleItem).Methods("GET")

    router.HandleFunc("/item", controller_item.CreateItem).Methods("POST")
    router.HandleFunc("/item/{id}", controller_item.DeleteItem).Methods("DELETE")
    router.HandleFunc("/item/{id}", controller_item.UpdateItem).Methods("PUT")

	Router = router
}

/*
 サーバー起動
*/
func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	// ルーティング設定
	setupRouting()

	return http.ListenAndServe(fmt.Sprintf(":%d", 3000), Router)
}

