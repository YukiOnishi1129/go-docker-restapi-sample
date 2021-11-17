package services

import (
	"encoding/json"
	"errors"
	"myapp/db"
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

/*
 Todoリストを取得
*/
func GetAllTodos(w http.ResponseWriter, userId int) ([]models.BaseTodoResponse, error) {
	var todos []models.Todo
	// todoリストデータ取得
	if err := repositories.GetAllTodos(&todos, userId); err != nil {
		logic.SendResponse(w, logic.CreateErrorStringResponse("データ取得に失敗"), http.StatusInternalServerError)
		return nil, err
	}
	// レスポンス用の構造体に変換
	responseTodos := logic.CreateAllTodoResponse(&todos)

	return responseTodos, nil
}

/*
 IDに紐づくTodoを取得
*/
func GetTodoById(w http.ResponseWriter,r *http.Request, userId int) (models.BaseTodoResponse, error) {
	// getパラメータからIDを取得
	vars := mux.Vars(r)
    id := vars["id"]
	var todo models.Todo
	// todoデータ取得処理
	if err := repositories.GetTodoById(&todo, id, userId); err != nil {
		var errMessage string
		var statusCode int
		// https://gorm.io/ja_JP/docs/error_handling.html
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "データ取得に失敗しました。"
		}
		logic.SendResponse(w, logic.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseTodoResponse{}, err
	}

	// レスポンス用の構造体に変換
	responseTodos := logic.CreateTodoResponse(&todo)

	return responseTodos, nil
}

func InsertTodo(todo *models.Todo) {
	db := db.GetDB()
	db.Create(&todo)
}

func DeleteTodo(id string, userId int) {
	db := db.GetDB()
	db.Where("id=? AND user_id=?", id, userId).Delete(&models.Todo{})
}

func UpdateTodo(todo *models.Todo, id string) {
	db := db.GetDB()
	db.Model(&todo).Where("id=? AND user_id=?", id, todo.UserId).Updates(
        map[string]interface{}{
            "title":     todo.Title,
            "comment":    todo.Comment,
			"user_id": todo.UserId,
        })
}


/*
 Todoリスト取得APIのレスポンス送信処理
*/
func SendAllTodoResponse(w http.ResponseWriter, todos *[]models.BaseTodoResponse) {
	var response models.AllTodoResponse
	response.Todos = *todos
	// レスポンスデータ作成
	responseBody, _ := json.Marshal(response)

	// レスポンス送信
	logic.SendResponse(w, responseBody, http.StatusOK)
}

/*
 Todoデータのレスポンス送信処理
*/
func SendTodoResponse(w http.ResponseWriter, todo *models.BaseTodoResponse) {
	var response models.TodoResponse
	response.Todo = *todo
	// レスポンスデータ作成
	responseBody, _ := json.Marshal(response)
	// レスポンス送信
	logic.SendResponse(w, responseBody, http.StatusOK)
}