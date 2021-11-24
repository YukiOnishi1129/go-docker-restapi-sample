package router

import (
	"myapp/controllers"

	"github.com/gorilla/mux"
)

type AppRouter interface {
	SetAppRouting(router *mux.Router)
}

type appRouter struct {
	apc controllers.AppController
}

func NewAppRouter(apc controllers.AppController) AppRouter {
	return &appRouter{apc}
}

func (apr *appRouter) SetAppRouting(router *mux.Router) {
	router.HandleFunc("/api/v1", apr.apc.RootPage).Methods("GET")
}