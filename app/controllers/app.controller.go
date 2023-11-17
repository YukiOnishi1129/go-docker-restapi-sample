package controllers

import (
	"fmt"
	"net/http"
)

type AppController interface {
	RootPage(w http.ResponseWriter, r *http.Request)
}

type appController struct{}

func NewAppController() AppController {
	return &appController{}
}

// RootPage ルートAPI
func (apc *appController) RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}
