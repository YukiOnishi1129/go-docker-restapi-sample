package logic

import (
	"myapp/models"
	"reflect"
	"testing"
)

/*
 【正常系】CreateAllTodoResponseが正常に処理されること
*/
func TestCreateAllTodoResponseSuccess(t *testing.T) {
	// テスト対象の引数
	var argTodo models.Todo
	argTodo.BaseModel.ID = 1
	argTodo.Title = "テスト1"
	argTodo.Comment = "テスト1"
	argTodoList := []models.Todo{argTodo, argTodo}

	// 予測値
	var expectedBaseTodoResponse models.BaseTodoResponse
	expectedBaseTodoResponse.BaseModel.ID = 1
	expectedBaseTodoResponse.Title = "テスト1"
	expectedBaseTodoResponse.Comment = "テスト1"
	expectedBaseTodoResponseList := [] models.BaseTodoResponse{expectedBaseTodoResponse,expectedBaseTodoResponse}

	tr := NewTodoLogic()
	// テスト対象の処理を実行
	actual := tr.CreateAllTodoResponse(&argTodoList)

	// テスト実行
	if !reflect.DeepEqual(actual, expectedBaseTodoResponseList) {
		t.Errorf("actual %v\nwant %v", actual, expectedBaseTodoResponseList)
	}
}

/*
 【正常系】CreateAllTodoResponseに空配列の引数を渡した際に、空配列が返却されること
*/
func TestCreateAllTodoResponseNotEmptySuccess(t *testing.T) {
	// テスト対象の引数 (空配列)
	argTodoList := []models.Todo{}

	// 予測値 (空配列)
	expectedBaseTodoResponseList := [] models.BaseTodoResponse{}

	tr := NewTodoLogic()
	// テスト対象の処理を実行
	actual := tr.CreateAllTodoResponse(&argTodoList)

	// テスト実行
	if len(actual) != len(expectedBaseTodoResponseList) {
		t.Errorf("actual %v\nwant %v", actual, expectedBaseTodoResponseList)
	}
}

/*
 【正常系】CreateTodoResponseが正常に処理されること
*/
func TestCreateTodoResponseSuccess(t *testing.T) {
	// テスト対象の引数
	var argTodo models.Todo
	argTodo.BaseModel.ID = 1
	argTodo.Title = "テスト1"
	argTodo.Comment = "テスト1"

	// 予測値
	var expectedBaseTodoResponse models.BaseTodoResponse
	expectedBaseTodoResponse.BaseModel.ID = 1
	expectedBaseTodoResponse.Title = "テスト1"
	expectedBaseTodoResponse.Comment = "テスト1"

	tr := NewTodoLogic()
	// テスト対象の処理を実行
	actual := tr.CreateTodoResponse(&argTodo)

	// テスト実行
	if !reflect.DeepEqual(actual, expectedBaseTodoResponse) {
		t.Errorf("actual %v\nwant %v", actual, expectedBaseTodoResponse)
	}
}