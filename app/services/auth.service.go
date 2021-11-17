package services

import (
	"fmt"
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"
)

/*
 会員登録バリデーション処理
*/
func ValidateSignUp(w http.ResponseWriter, signUpRequestParam models.SignUpRequest) error {
	// バリデーション
	if err := validation.SignUpValidate(signUpRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		logic.SendResponse(w, logic.CreateErrorResponse(err), http.StatusBadRequest)
		return err
	}

	return nil
}

/*
 同じメールアドレスのユーザーがないか検証
*/
func CheckSameEmailUser(w http.ResponseWriter, users *[]models.User, email string) error {
	// emailに紐づくユーザーをチェック
	if err := repositories.GetUserByEmail(users, email); err != nil {
		logic.SendResponse(w, logic.CreateErrorStringResponse("DBエラー"), http.StatusInternalServerError)
		return err
	}

	if len(*users) != 0 {
		logic.SendResponse(w, logic.CreateErrorStringResponse("入力されたメールアドレスは既に登録されています。"),http.StatusUnauthorized)
		return fmt.Errorf("%d is minus value",10)
	}

	return  nil
}