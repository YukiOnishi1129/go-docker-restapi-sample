package logic

import (
	"errors"
	"net/http"

	jwt "github.com/form3tech-oss/jwt-go"
)

type AuthLogic interface {
	GetUserIdFromContext(r *http.Request) (int, error)
}

type authLogic struct {
}

func NewAuthLogic() AuthLogic {
	return &authLogic{}
}

/*
 トークン情報よりuserIdを取得
*/
func (al *authLogic) GetUserIdFromContext(r *http.Request) (int, error) {
	// トークンからuserIdを取得
	user := r.Context().Value("user")
    claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	userId, ok := claims["id"].(float64)

	if !ok {
        return 0, errors.New("id type not match")
    }

	return int(userId), nil
}