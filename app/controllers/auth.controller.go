package controllers

import (
	"encoding/json"
	"fmt"
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
	// RequestのBodyデータを取得
	reqBody, _ := ioutil.ReadAll(r.Body)
	var signInRequestParam models.SignInRequest
	// Unmarshal: jsonを構造体に変換
	if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
		log.Fatal(err)
		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
	}

	// バリデーション
	if err := services.ValidateSignIn(w, signInRequestParam); err != nil {
		return
	}

	// ユーザー認証
	var user models.User
	if err := services.FindUserByEmail(w, &user, signInRequestParam.Email); err != nil {
		return
	}

	if err := services.VerificationPassword(w, user.Password, signInRequestParam.Password); err != nil {
		return
	}

	// jwtトークンを作成
	logic.CreateJwtToken(&user)

	// レスポンスの構造体を作る
	response := map[string]interface{}{
		"token": logic.GetJwtToken(),
		"user": user,
	}

	// レスポンスデータ作成
	responseBody, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("レスポンスデータ失敗")
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(responseBody)
	
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
	services.SendSignUpResponse(w, &createUser)
}

func SetAuthRouting(router *mux.Router) {
	router.HandleFunc("/signin", singIn).Methods("POST")
	router.HandleFunc("/signup", signUp).Methods("POST")
}