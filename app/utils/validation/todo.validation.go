package validation

import (
	"myapp/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TodoValidation interface {
	MutationTodoValidate(mutationTodoRequest models.MutationTodoRequest) error
}

type todoValidation struct{}

func NewTodoValidation() TodoValidation {
	return &todoValidation{}
}

// MutationTodoValidate 更新時のリクエストパラメータのバリデーション
func (tv *todoValidation) MutationTodoValidate(mutationTodoRequest models.MutationTodoRequest) error {
	return validation.ValidateStruct(&mutationTodoRequest,
		validation.Field(
			&mutationTodoRequest.Title,
			validation.Required.Error("タイトルは必須入力です。"),
			validation.RuneLength(1, 10).Error("タイトルは 1～10 文字です"),
		),
		validation.Field(
			&mutationTodoRequest.Comment,
			validation.Required.Error("コメントは必須入力です。"),
		),
	)
}
