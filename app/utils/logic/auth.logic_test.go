package logic_test

import (
	"fmt"
	"github.com/joho/godotenv"
	"myapp/utils/logic"
	"net/http"
	"testing"
)

//func TestGetUserIdFromContextSuccess(t *testing.T) {
//	// env読み込み
//	err := godotenv.Load("../../.env.sample")
//	if err != nil {
//		t.Errorf(".envファイル読み込みエラー")
//		return
//	}
//	expectedUserId := 1
//	invalidToken := os.Getenv("TEST_JWT_TOKEN")
//	// リクエストの生成
//	req, _ := http.NewRequest(http.MethodGet, "/api/v1/todo", nil)
//
//	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", invalidToken))
//
//	// テスト実行
//	authLogic := logic.NewAuthLogic()
//	actual, err := authLogic.GetUserIdFromContext(req)
//
//	if err != nil || actual != expectedUserId {
//		t.Errorf("actual %v\nwant %v", actual, expectedUserId)
//	}
//}

func TestGetUserIdFromContextNotAuthenticationTokenError(t *testing.T) {
	// env読み込み
	err := godotenv.Load("../../.env.sample")
	if err != nil {
		t.Errorf(".envファイル読み込みエラー")
		return
	}
	expectedUserId := 1
	invalidToken := ""

	// リクエストの生成
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/todo", nil)

	req.Header.Add("Authorization", invalidToken)

	// テスト実行
	authLogic := logic.NewAuthLogic()
	actual, err := authLogic.GetUserIdFromContext(req)

	expectedError := "not token"

	if err.Error() != expectedError || actual != 0 {
		t.Errorf("actual %v\nwant %v", actual, expectedUserId)
	}
}

func TestGetUserIdFromContextEmptyTokenError(t *testing.T) {
	// env読み込み
	err := godotenv.Load("../../.env.sample")
	if err != nil {
		t.Errorf(".envファイル読み込みエラー")
		return
	}
	expectedUserId := 1
	invalidToken := ""

	// リクエストの生成
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/todo", nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", invalidToken))

	// テスト実行
	authLogic := logic.NewAuthLogic()
	actual, err := authLogic.GetUserIdFromContext(req)

	expectedError := "トークンが空文字です。"

	if err.Error() != expectedError || actual != 0 {
		fmt.Print(err)
		t.Errorf("actual %v\nwant %v", actual, expectedUserId)
	}
}
