package services

import (
	"myapp/models"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"
)

/*
 会員登録バリデーション処理
*/
func SignUpValidation(w http.ResponseWriter, signUpRequestParam models.SignUpRequest) error {
	// バリデーション
	if err := validation.SignUpValidate(signUpRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		logic.SendResponse(w, logic.CreateErrorResponse(err), http.StatusBadRequest)
		return err
	}

	return nil
}