package logic

import (
	"fmt"
	"myapp/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var JwtToken []byte

/*
 jwtトークンの新規作成
*/
func CreateJwtToken(user *models.User)  {
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
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	JwtToken = []byte(tokenString)
}

/*
 jwtトークンを取得
*/
func GetJwtToken() []byte {
	return JwtToken
}