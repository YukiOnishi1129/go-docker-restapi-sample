package logic

import "myapp/models"

type TodoLogic interface {
	CreateAllTodoResponse(todos *[]models.Todo) []models.BaseTodoResponse
	CreateTodoResponse(todo *models.Todo) models.BaseTodoResponse
}

type todoLogic struct {}

func NewTodoLogic() TodoLogic {
	return &todoLogic{}
}

/*
 レスポンス用のTodoリストの構造体を作成
*/
func (tl *todoLogic) CreateAllTodoResponse(todos *[]models.Todo) []models.BaseTodoResponse {
	var responseTodos []models.BaseTodoResponse
	for _, todo := range *todos {
		var newTodo models.BaseTodoResponse
		newTodo.BaseModel.ID = todo.BaseModel.ID
		newTodo.BaseModel.CreatedAt = todo.BaseModel.CreatedAt
		newTodo.BaseModel.UpdatedAt = todo.BaseModel.UpdatedAt
		newTodo.BaseModel.DeletedAt = todo.BaseModel.DeletedAt
		newTodo.Title = todo.Title
		newTodo.Comment = todo.Comment
		responseTodos = append(responseTodos, newTodo)
	}

	return responseTodos
}

/*
 レスポンス用のTodoの構造体を作成
*/
func (tl *todoLogic) CreateTodoResponse(todo *models.Todo) models.BaseTodoResponse {
	var responseTodo models.BaseTodoResponse
	responseTodo.BaseModel.ID = todo.BaseModel.ID
	responseTodo.BaseModel.CreatedAt = todo.BaseModel.CreatedAt
	responseTodo.BaseModel.UpdatedAt = todo.BaseModel.UpdatedAt
	responseTodo.BaseModel.DeletedAt = todo.BaseModel.DeletedAt
	responseTodo.Title = todo.Title
	responseTodo.Comment = todo.Comment

	return responseTodo
}