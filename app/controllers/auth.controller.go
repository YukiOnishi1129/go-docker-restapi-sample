package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myapp/db"
	"myapp/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type SingInRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

/*
 ログインパラメータのバリデーション
*/
func SignInValidate(signInRequest SingInRequest) error {
	return validation.ValidateStruct(&signInRequest,
		validation.Field(
			&signInRequest.Email,
			validation.Required.Error("メールアドレスは必須入力です。"),
			validation.RuneLength(5, 40).Error("メールアドレスは 5～40 文字です"),
			is.Email.Error("メールアドレスを入力して下さい"),
		),
		validation.Field(
			&signInRequest.Password,
			validation.Required.Error("パスワードは必須入力です。"),
			validation.Length(6, 20).Error("パスワードは6文字以上、20字以内で入力してください。"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
		),
	)
}

/*
 ログイン処理
*/
func singIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// RequestのBodyデータを取得
		reqBody, _ := ioutil.ReadAll(r.Body)
		var signInRequestParam SingInRequest
		// Unmarshal: jsonを構造体に変換
		if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
			log.Fatal(err)
		}

		// バリデーション
		if err := SignInValidate(signInRequestParam); err != nil {
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
		// headerのセット
		token := jwt.New(jwt.SigningMethodHS256)
		// claimsのセット
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["email"] = user.Email
		claims["name"] = user.Name
		claims["iat"] = time.Now() // jwtの発行時間
		// 経過時間
		// 経過時間を過ぎたjetは処理しないようになる
		// ここでは24時間の経過時間をリミットにしている
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// .envを読み込む
		err := godotenv.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 電子署名
		tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

		// レスポンスの構造体を作る
		response := map[string]interface{}{
			"token": []byte(tokenString),
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