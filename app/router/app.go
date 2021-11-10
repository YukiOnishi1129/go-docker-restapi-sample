package router

import (
	"fmt"
	"myapp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

/*
 ルーティング定義
*/
func setupRouting() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.RootPage)
    router.HandleFunc("/items", controllers.FetchAllItems).Methods("GET")
    router.HandleFunc("/item/{id}", controllers.FetchSingleItem).Methods("GET")

    router.HandleFunc("/item", controllers.CreateItem).Methods("POST")
    router.HandleFunc("/item/{id}", controllers.DeleteItem).Methods("DELETE")
    router.HandleFunc("/item/{id}", controllers.UpdateItem).Methods("PUT")

	Router = router
}

/*
 サーバー起動
*/
func  StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	setupRouting()

	return http.ListenAndServe(fmt.Sprintf(":%d", 3000), Router)
}

