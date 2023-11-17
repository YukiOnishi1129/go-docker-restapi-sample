package logic

import (
	"net/http"
	"os"
	"strings"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/pkg/errors"
)

type AuthLogic interface {
	GetUserIdFromContext(r *http.Request) (int, error)
}

type authLogic struct {
}

func NewAuthLogic() AuthLogic {
	return &authLogic{}
}

// GetUserIdFromContext トークン情報よりuserIdを取得
func (al *authLogic) GetUserIdFromContext(r *http.Request) (int, error) {
	// トークンからuserIdを取得
	// 昔のやり方
	// user := r.Context().Value("user")
	// claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	// userId, ok := claims["id"].(float64)

	// https://ichi.pro/golang-de-no-jwt-no-jissen-154586293449920
	// https://stackoverflow.com/questions/56415581/how-to-test-authenticate-jwt-routes-in-go

	clientToken := r.Header.Get("Authorization")
	if clientToken == "" {
		return 0, errors.New("not token")
	}

	extractToken := strings.Split(clientToken, "Bearer ")
	secretKey := os.Getenv("JWT_KEY")

	if extractToken[1] == "" {
		return 0, errors.Errorf("トークンが空文字です。")
	}

	token, err := jwt.Parse(extractToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("トークンをjwtにparseできません。")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, claimOk := token.Claims.(jwt.MapClaims)
	if !claimOk || !token.Valid {
		return 0, errors.New("id type not match")
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("id type not match")
	}

	return int(userId), nil
}
