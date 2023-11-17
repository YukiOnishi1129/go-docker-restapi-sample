package logic

import (
	"fmt"
	"myapp/models"
	"os"
	"strconv"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/joho/godotenv"
)

type JWTLogic interface {
	CreateJwtToken(user *models.User) (string, error)
}

type jwtLogic struct{}

func NewJWTLogic() JWTLogic {
	return &jwtLogic{}
}

// CreateJwtToken jwtトークンの新規作成
func (jl *jwtLogic) CreateJwtToken(user *models.User) (string, error) {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)
	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = strconv.Itoa(int(user.ID)) + user.Email + user.Name
	claims["id"] = user.ID
	claims["name"] = user.Name
	// latを取り除かないとミドルウェアで「Token used before issued」エラーになる
	// https://github.com/dgrijalva/jwt-go/issues/314#issuecomment-812775567
	// claims["iat"] = time.Now() // jwtの発行時間
	// 経過時間
	// 経過時間を過ぎたjetは処理しないようになる
	// ここでは24時間の経過時間をリミットにしている
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// .envを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return tokenString, nil
}

// JwtMiddleware jwt認証のミドルウェア
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
