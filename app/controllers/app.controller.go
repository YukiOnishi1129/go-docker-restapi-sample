package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
* ルートAPI
 */
func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
    fmt.Println("Root endpoint is hooked!")
}

func SetAppRouting(router *mux.Router) {
	router.HandleFunc("/", rootPage).Methods("GET")
}