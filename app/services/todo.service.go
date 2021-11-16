package services

import (
	"myapp/db"
	"myapp/models"
)

func GetAllTodos(todo *[]models.Todo, userId int) {
	db := db.GetDB()
	db.Where("user_id=?", userId).Find(&todo)
}

func GetTodoById(todo *models.Todo, id string, userId int) {
	db := db.GetDB()
	db.Where("user_id=?", userId).First(&todo, id)
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