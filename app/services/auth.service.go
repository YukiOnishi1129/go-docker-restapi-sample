package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"

	"github.com/pkg/errors"
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
 会員登録処理
*/
func SignUp(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// RequestのBodyデータを取得
	reqBody, _ := ioutil.ReadAll(r.Body)
	var signUpRequestParam models.SignUpRequest
	if err := json.Unmarshal(reqBody, &signUpRequestParam); err != nil {
		log.Fatal(err)
		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
		return models.User{}, err
	}

	// バリデーション
	if err := validation.SignUpValidate(signUpRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		logic.SendResponse(w, logic.CreateErrorResponse(err), http.StatusBadRequest)
		return models.User{}, err
	}
	// 同じメールアドレスのユーザーがいないか検証
	var users []models.User
	// emailに紐づくユーザーをチェック
	if err := repositories.GetAllUserByEmail(&users, signUpRequestParam.Email); err != nil {
		logic.SendResponse(w, logic.CreateErrorStringResponse("DBエラー"), http.StatusInternalServerError)
		return models.User{}, err
	}

	if len(users) != 0 {
		logic.SendResponse(w, logic.CreateErrorStringResponse("入力されたメールアドレスは既に登録されています。"),http.StatusUnauthorized)
		return models.User{}, errors.Errorf("「%w」 のユーザーは既に登録されています。", signUpRequestParam.Email)
	}

	var createUser models.User
	hashPassword := logic.ChangeHashPassword(signUpRequestParam.Password)
	// 登録データを作成
	createUser.Name = signUpRequestParam.Name
	createUser.Email = signUpRequestParam.Email
	createUser.Password = string(hashPassword)
	// ユーザー登録処理
	if err := repositories.CreateUser(&createUser); err != nil {
		logic.SendResponse(w, logic.CreateErrorStringResponse("ユーザー登録処理に失敗"), http.StatusInternalServerError)
		return models.User{}, err
	}

	// jwtトークンを作成
	logic.CreateJwtToken(&createUser)

	return createUser, nil
}

/*
 ログインAPI・会員登録APIのレスポンス送信処理
*/
func SendAuthResponse(w http.ResponseWriter, user *models.User, code int) {
	var response models.SignUpResponse
	response.Token = logic.GetJwtToken()
	response.User.BaseModel.ID = user.ID
	response.User.BaseModel.CreatedAt = user.CreatedAt
	response.User.BaseModel.UpdatedAt = user.UpdatedAt
	response.User.BaseModel.DeletedAt = user.DeletedAt
	response.User.Name = user.Name
	response.User.Email = user.Email
	// レスポンスデータ作成
	responseBody, _ := json.Marshal(response)

	// レスポンス送信
	logic.SendResponse(w, responseBody, code)
}