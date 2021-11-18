package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/services"
	"myapp/utils/logic"
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
	services.SendAuthResponse(w, &user)
}

/*
 会員登録処理
*/
func signUp(w http.ResponseWriter, r *http.Request) {
	// RequestのBodyデータを取得
	reqBody, _ := ioutil.ReadAll(r.Body)
	var signUpRequestParam models.SignUpRequest
	if err := json.Unmarshal(reqBody, &signUpRequestParam); err != nil {
		log.Fatal(err)
		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
	}

	// バリデーション
	if err := services.ValidateSignUp(w, signUpRequestParam); err != nil {
		return
	}

	// 同じメールアドレスのユーザーがいないか検証
	var users []models.User
	if err := services.CheckSameEmailUser(w, &users, signUpRequestParam.Email); err != nil {
		return
	}

	var createUser models.User

	// ユーザー登録処理
	if err := services.CreateUser(w, &createUser, signUpRequestParam); err != nil {
		return
	}

	// jwtトークンを作成
	logic.CreateJwtToken(&createUser)

	// 会員登録APIのレスポンス送信
	services.SendAuthResponse(w, &createUser)
}

/*
 auth controllerのルーティング設定
*/
func SetAuthRouting(router *mux.Router) {
	router.HandleFunc("/signin", singIn).Methods("POST")
	router.HandleFunc("/signup", signUp).Methods("POST")
}