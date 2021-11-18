package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

/*
 tokenよりuserIdを取得
*/
func GetUserIdFromToken(w http.ResponseWriter, r *http.Request) (int, error) {
	// トークンからuserIdを取得
	userId, err := logic.GetUserIdFromContext(r)
	if err != nil {
		errMessage := "認証エラー"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return 0, err
	}

	return userId, nil
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
 ログイン処理
*/
func SignIn(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// RequestのBodyデータを取得
	reqBody, _ := ioutil.ReadAll(r.Body)
	var signInRequestParam models.SignInRequest
	// Unmarshal: jsonを構造体に変換
	if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
		log.Fatal(err)
		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
		return models.User{}, err
	}

	// バリデーション
	if err := validation.SignInValidate(signInRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		logic.SendResponse(w, logic.CreateErrorResponse(err), http.StatusBadRequest)
		return models.User{}, err
	}

	// ユーザー認証
	var user models.User
	// emailに紐づくユーザーをチェック
	if err := repositories.GetUserByEmail(&user, signInRequestParam.Email); err != nil {
		errMessage := "メールアドレスに該当するユーザーが存在しません。"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return models.User{}, err
	}
	// パスワード照合
	// CompareHashAndPassword
	// 第一引数: hash化したパスワード
	// 第二引数: 文字列のパスワード
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
		errMessage := "パスワードが間違っています。"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return models.User{}, err
	}

	// jwtトークンを作成
	logic.CreateJwtToken(&user)

	return user, nil
}

/*
 ログインAPI・会員登録APIのレスポンス送信処理
*/
func SendAuthResponse(w http.ResponseWriter, createUser *models.User) {
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