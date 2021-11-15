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

	controllers.SetAppRouting(router)
	controllers.SetTodoRouting(router)

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

