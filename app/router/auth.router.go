package router

import (
	"myapp/controllers"

	"github.com/gorilla/mux"
)

type AuthRouter interface {
	SetAuthRouting(router *mux.Router)
}

type authRouter struct {
	ac controllers.AuthController
}

func NewAuthRouter(ac controllers.AuthController) AuthRouter {
	return &authRouter{ac}
}


func (ar *authRouter) SetAuthRouting(router *mux.Router) {
	router.HandleFunc("/api/v1/signin", ar.ac.SingIn).Methods("POST")
	router.HandleFunc("/api/v1/signup", ar.ac.SignUp).Methods("POST")
}