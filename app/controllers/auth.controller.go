package controllers

import (
	"myapp/services"
	"net/http"
)

type AuthController interface {
	SingIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	as services.AuthService
}

func NewAuthController(as services.AuthService) AuthController {
	return &authController{as}
}

/*
 ログイン処理
*/
func (ac *authController) SingIn(w http.ResponseWriter, r *http.Request) {
	// ログイン処理
	user, err := ac.as.SignIn(w, r);
	if err != nil {
		return
	}

	// ログインAPIのレスポンス送信
	ac.as.SendAuthResponse(w, &user, http.StatusOK)
}

/*
 会員登録処理
*/
func (ac *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	// ログイン処理
	createUser, err := ac.as.SignUp(w, r);
	if err != nil {
		return
	}

	// 会員登録APIのレスポンス送信
	ac.as.SendAuthResponse(w, &createUser, http.StatusCreated)
}

/*
 auth controllerのルーティング設定
*/
// func SetAuthRouting(router *mux.Router) {
// 	router.HandleFunc("/api/v1/signin", singIn).Methods("POST")
// 	router.HandleFunc("/api/v1/signup", signUp).Methods("POST")
// }