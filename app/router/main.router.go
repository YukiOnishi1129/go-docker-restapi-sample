package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MainRouter interface {
	setupRouting() *mux.Router
	StartWebServer() error
}

type mainRouter struct {
	appR  AppRouter
	authR AuthRouter
	todoR TodoRouter
}

func NewMainRouter(appR AppRouter, authR AuthRouter, todoR TodoRouter) MainRouter {
	return &mainRouter{appR, authR, todoR}
}

const PORT = 4000

// routing definition
func (mainRouter *mainRouter) setupRouting() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	mainRouter.appR.SetAppRouting(router)
	mainRouter.authR.SetAuthRouting(router)
	mainRouter.todoR.SetTodoRouting(router)

	return router
}

// StartWebServer server startup
func (mainRouter *mainRouter) StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	// // ルーティング設定
	// setupRouting()

	return http.ListenAndServe(fmt.Sprintf(":%d", PORT), mainRouter.setupRouting())
}
