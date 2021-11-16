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
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type SingInRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func singIn(w http.ResponseWriter, r *http.Request) {
	// RequestのBodyデータを取得
	reqBody, _ := ioutil.ReadAll(r.Body)
	var signInRequestParam SingInRequest
	// Unmarshal: jsonを構造体に変換
	if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
		log.Fatal(err)
	}
	// TODO: バリデーション
	// ユーザー認証
	var user models.User
	db := db.GetDB()
	db.Where("email=?", signInRequestParam.Email).First(&user)
	// requestPassword, _ := bcrypt.GenerateFromPassword([]byte(signInRequestParam.Password), bcrypt.DefaultCost)
	// パスワード照合
	// CompareHashAndPassword
	// 第一引数: hash化したパスワード
	// 第二引数: 文字列のパスワード
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
		fmt.Printf("パスワード照合失敗")
		log.Fatal(err)
	}

	// TODO: jwtトークンを返す
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

func SetAuthRouting(router *mux.Router) {
	router.HandleFunc("/singin", singIn).Methods("POST")
}