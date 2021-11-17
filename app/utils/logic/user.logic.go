package logic

import "golang.org/x/crypto/bcrypt"

/*
 パスワードのハッシュ化
*/
func ChangeHashPassword(password string) []byte {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashPassword
}