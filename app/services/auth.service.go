package services

import (
	"encoding/json"
	"fmt"
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"
)

/*
 ログインバリデーション処理
*/
func ValidateSignIn(w http.ResponseWriter, signInRequestParam models.SignInRequest) error {
	// バリデーション
	if err := validation.SignInValidate(signInRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		logic.SendResponse(w, logic.CreateErrorResponse(err), http.StatusBadRequest)
		return err
	}

	return nil
}

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
 メールアドレスに紐づくユーザーを取得
*/
func FindUserByEmail(w http.ResponseWriter, user *models.User, email string) error {
	// emailに紐づくユーザーをチェック
	if err := repositories.GetUserByEmail(user, email); err != nil {
		errMessage := "メールアドレスに該当するユーザーが存在しません。"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return err
	}

	return nil
}

/*
 同じメールアドレスのユーザーがないか検証
*/
func CheckSameEmailUser(w http.ResponseWriter, users *[]models.User, email string) error {
	// emailに紐づくユーザーをチェック
	if err := repositories.GetAllUserByEmail(users, email); err != nil {
		logic.SendResponse(w, logic.CreateErrorStringResponse("DBエラー"), http.StatusInternalServerError)
		return err
	}

	if len(*users) != 0 {
		logic.SendResponse(w, logic.CreateErrorStringResponse("入力されたメールアドレスは既に登録されています。"),http.StatusUnauthorized)
		return fmt.Errorf("%d is minus value",10)
	}

	return  nil
}

/*
 会員登録APIのレスポンス送信処理
*/
func SendSignUpResponse(w http.ResponseWriter, createUser *models.User) {
	var response models.SignUpResponse
	response.Token = logic.GetJwtToken()
	response.User.BaseModel.ID = createUser.ID
	response.User.BaseModel.CreatedAt = createUser.CreatedAt
	response.User.BaseModel.UpdatedAt = createUser.UpdatedAt
	response.User.BaseModel.DeletedAt = createUser.DeletedAt
	response.User.Name = createUser.Name
	response.User.Email = createUser.Email
	// レスポンスデータ作成
	responseBody, _ := json.Marshal(response)

	// レスポンス送信
	logic.SendResponse(w, responseBody, http.StatusCreated)
}