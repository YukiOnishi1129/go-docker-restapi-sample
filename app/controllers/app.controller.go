package controllers

import (
	"fmt"
	"net/http"
)

type AppController interface {
	RootPage(w http.ResponseWriter, r *http.Request)
}

type appController struct {}


func NewAppController() AppController {
	return &appController{}
}

/*
* ルートAPI
 */
func (apc *appController) RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
    fmt.Println("Root endpoint is hooked!")
}

// func SetAppRouting(router *mux.Router) {
// 	router.HandleFunc("/api/v1", rootPage).Methods("GET")
// }