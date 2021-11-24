package logic_test

import (
	"myapp/utils/logic"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestChangeHashPasswordSuccess(t *testing.T) {
	expectedPassword := "password"
	userLogic := logic.NewUserLogic()
	// テスト実行
	actual := userLogic.ChangeHashPassword(expectedPassword)

	if err := bcrypt.CompareHashAndPassword(actual, []byte(expectedPassword)); err != nil {
		t.Errorf("actual %v\nwant %v", actual, []byte(expectedPassword))
	}
}