package controllers

import (
	"myapp/services"
	"net/http"
)

type TodoController interface {
	FetchAllTodos(w http.ResponseWriter, r *http.Request)
	FetchTodoById(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	ts services.TodoService
	as services.AuthService
}

func NewTodoController(ts services.TodoService, as services.AuthService) TodoController {
	return &todoController{ts, as}
}

// FetchAllTodos Todoリスト取得
func (tc *todoController) FetchAllTodos(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// todoリスト取得処理
	alltodo, err := tc.ts.GetAllTodos(w, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.ts.SendAllTodoResponse(w, &alltodo)
}

// FetchTodoById idに紐づくTodoを取得
func (tc *todoController) FetchTodoById(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}
	// todoデータ取得処理
	responseTodo, err := tc.ts.GetTodoById(w, r, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.ts.SendTodoResponse(w, &responseTodo)
}

// CreateTodo Todo新規登録
func (tc *todoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// トークンからuserIdを取得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// todoデータ取得処理
	responseTodo, err := tc.ts.CreateTodo(w, r, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.ts.SendCreateTodoResponse(w, &responseTodo)
}

// DeleteTodo 削除処理
func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// トークンからuserIdを取得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// データ削除処理
	if err := tc.ts.DeleteTodo(w, r, userId); err != nil {
		return
	}

	// レスポンス送信処理
	tc.ts.SendDeleteTodoResponse(w)
}

// UpdateTodo Todo更新処理
func (tc *todoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// todo更新処理
	responseTodo, err := tc.ts.UpdateTodo(w, r, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.ts.SendCreateTodoResponse(w, &responseTodo)
}
