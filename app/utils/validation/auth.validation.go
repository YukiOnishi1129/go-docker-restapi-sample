package validation

import (
	"myapp/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

/*
 ログインパラメータのバリデーション
*/
func SignInValidate(signInRequest models.SingInRequest) error {
	return validation.ValidateStruct(&signInRequest,
		validation.Field(
			&signInRequest.Email,
			validation.Required.Error("メールアドレスは必須入力です。"),
			validation.RuneLength(5, 40).Error("メールアドレスは 5～40 文字です"),
			is.Email.Error("メールアドレスの形式が間違っています。"),
		),
		validation.Field(
			&signInRequest.Password,
			validation.Required.Error("パスワードは必須入力です。"),
			validation.Length(6, 20).Error("パスワードは6文字以上、20字以内で入力してください。"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
		),
	)
}