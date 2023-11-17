package services

import (
	"encoding/json"
	"io"
	"log"
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GetUserIdFromToken(w http.ResponseWriter, r *http.Request) (int, error)
	SignIn(w http.ResponseWriter, r *http.Request) (models.User, error)
	SignUp(w http.ResponseWriter, r *http.Request) (models.User, error)
	SendAuthResponse(w http.ResponseWriter, user *models.User, code int)
}

type authService struct {
	ur repositories.UserRepository
	al logic.AuthLogic
	ul logic.UserLogic
	rl logic.ResponseLogic
	jl logic.JWTLogic
	av validation.AuthValidation
}

func NewAuthService(ur repositories.UserRepository, al logic.AuthLogic, ul logic.UserLogic, rl logic.ResponseLogic, jl logic.JWTLogic, av validation.AuthValidation) AuthService {
	return &authService{ur, al, ul, rl, jl, av}
}

// GetUserIdFromToken tokenよりuserIdを取得
func (as *authService) GetUserIdFromToken(w http.ResponseWriter, r *http.Request) (int, error) {
	// トークンからuserIdを取得
	userId, err := as.al.GetUserIdFromContext(r)
	if err != nil {
		errMessage := "認証エラー"
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return 0, err
	}

	return userId, nil
}

// SignIn ログイン処理
func (as *authService) SignIn(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// RequestのBodyデータを取得
	reqBody, _ := io.ReadAll(r.Body)
	var signInRequestParam models.SignInRequest
	// Unmarshal: jsonを構造体に変換
	if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
		log.Fatal(err)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("リクエストパラメータを構造体へ変換処理でエラー発生"), http.StatusInternalServerError)
		return models.User{}, err
	}

	// バリデーション
	if err := as.av.SignInValidate(signInRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		as.rl.SendResponse(w, as.rl.CreateErrorResponse(err), http.StatusBadRequest)
		return models.User{}, err
	}

	// ユーザー認証
	var user models.User
	// emailに紐づくユーザーをチェック
	if err := as.ur.GetUserByEmail(&user, signInRequestParam.Email); err != nil {
		errMessage := "メールアドレスに該当するユーザーが存在しません。"
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return models.User{}, err
	}
	// パスワード照合
	// CompareHashAndPassword
	// 第一引数: hash化したパスワード
	// 第二引数: 文字列のパスワード
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
		errMessage := "パスワードが間違っています。"
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), http.StatusUnauthorized)
		return models.User{}, err
	}

	return user, nil
}

// SignUp 会員登録処理
func (as *authService) SignUp(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// RequestのBodyデータを取得
	reqBody, _ := io.ReadAll(r.Body)
	var signUpRequestParam models.SignUpRequest
	if err := json.Unmarshal(reqBody, &signUpRequestParam); err != nil {
		log.Fatal(err)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("リクエストパラメータを構造体へ変換処理でエラー発生"), http.StatusInternalServerError)
		return models.User{}, err
	}

	// バリデーション
	if err := as.av.SignUpValidate(signUpRequestParam); err != nil {
		// バリデーションエラーのレスポンスを送信
		as.rl.SendResponse(w, as.rl.CreateErrorResponse(err), http.StatusBadRequest)
		return models.User{}, err
	}
	// 同じメールアドレスのユーザーがいないか検証
	var users []models.User
	// emailに紐づくユーザーをチェック
	if err := as.ur.GetAllUserByEmail(&users, signUpRequestParam.Email); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("DBエラー"), http.StatusInternalServerError)
		return models.User{}, err
	}

	if len(users) != 0 {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("入力されたメールアドレスは既に登録されています。"), http.StatusUnauthorized)
		return models.User{}, errors.Errorf("「%w」 のユーザーは既に登録されています。", signUpRequestParam.Email)
	}

	var createUser models.User
	hashPassword := as.ul.ChangeHashPassword(signUpRequestParam.Password)
	// 登録データを作成
	createUser.Name = signUpRequestParam.Name
	createUser.Email = signUpRequestParam.Email
	createUser.Password = string(hashPassword)
	// ユーザー登録処理
	if err := as.ur.CreateUser(&createUser); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("ユーザー登録処理に失敗"), http.StatusInternalServerError)
		return models.User{}, err
	}

	return createUser, nil
}

// SendAuthResponse ログインAPI・会員登録APIのレスポンス送信処理
func (as *authService) SendAuthResponse(w http.ResponseWriter, user *models.User, code int) {
	// jwtトークンを作成
	token, err := as.jl.CreateJwtToken(user)
	if err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("トークン作成に失敗"), http.StatusInternalServerError)
		return
	}
	// レスポンスデータ作成
	var response models.AuthResponse
	response.Token = token
	response.User.BaseModel.ID = user.ID
	response.User.BaseModel.CreatedAt = user.CreatedAt
	response.User.BaseModel.UpdatedAt = user.UpdatedAt
	response.User.BaseModel.DeletedAt = user.DeletedAt
	response.User.Name = user.Name
	response.User.Email = user.Email

	// レスポンスデータをjson形式に変換
	responseBody, _ := json.Marshal(response)

	// レスポンス送信
	as.rl.SendResponse(w, responseBody, code)
}
