package logic

import "golang.org/x/crypto/bcrypt"

type UserLogic interface {
	ChangeHashPassword(password string) []byte
}

type userLogic struct {
}

func NewUserLogic() UserLogic {
	return &userLogic{}
}

// ChangeHashPassword パスワードのハッシュ化
func (ul *userLogic) ChangeHashPassword(password string) []byte {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashPassword
}
