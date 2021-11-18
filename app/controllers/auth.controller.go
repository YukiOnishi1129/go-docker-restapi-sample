package controllers

import (
	"myapp/services"
	"net/http"

	"github.com/gorilla/mux"
)

/*
 ログイン処理
*/
func singIn(w http.ResponseWriter, r *http.Request) {
	// ログイン処理
	user, err := services.SignIn(w, r);
	if err != nil {
		return
	}

	// ログインAPIのレスポンス送信
	services.SendAuthResponse(w, &user, http.StatusOK)
}

/*
 会員登録処理
*/
func signUp(w http.ResponseWriter, r *http.Request) {
	// ログイン処理
	createUser, err := services.SignUp(w, r);
	if err != nil {
		return
	}

	// 会員登録APIのレスポンス送信
	services.SendAuthResponse(w, &createUser, http.StatusCreated)
}

/*
 auth controllerのルーティング設定
*/
func SetAuthRouting(router *mux.Router) {
	router.HandleFunc("/signin", singIn).Methods("POST")
	router.HandleFunc("/signup", signUp).Methods("POST")
}