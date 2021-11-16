package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myapp/db"
	"myapp/models"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

/*
 ログイン処理
*/
func singIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// RequestのBodyデータを取得
		reqBody, _ := ioutil.ReadAll(r.Body)
		var signInRequestParam models.SingInRequest
		// Unmarshal: jsonを構造体に変換
		if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
			log.Fatal(err)
		}

		// バリデーション
		if err := validation.SignInValidate(signInRequestParam); err != nil {
			response := map[string]interface{}{
				"error": err,
			}
			responseBody, _ := json.Marshal(response)
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest) // ステータスコード
			w.Write(responseBody)
			return
		}

		// ユーザー認証
		var user models.User
		db := db.GetDB()
		if err := db.Where("email=?", signInRequestParam.Email).First(&user).Error; err != nil {
			response := map[string]interface{}{
				"error": "メールアドレスに該当するユーザーが存在しません。",
			}
			responseBody, _ := json.Marshal(response)
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized) // ステータスコード
			w.Write(responseBody)
			return
		}

		// パスワード照合
		// CompareHashAndPassword
		// 第一引数: hash化したパスワード
		// 第二引数: 文字列のパスワード
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
			response := map[string]interface{}{
				"error": "パスワードが間違っています。",
			}
			responseBody, _ := json.Marshal(response)
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized) // ステータスコード
			w.Write(responseBody)
			return
		}

		// jwtトークンを返す
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
}

func SetAuthRouting(router *mux.Router) {
	router.HandleFunc("/singin", singIn).Methods("POST")
}