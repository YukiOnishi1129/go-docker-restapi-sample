package services

import (
	"myapp/db"
	"myapp/models"
)

func GetAllTodos(todo *[]models.Todo) {
	db := db.GetDB()
	db.Find(&todo)
}

func GetTodoById(todo *models.Todo, id string) {
	db := db.GetDB()
	db.First(&todo, id)
}

func InsertTodo(todo *models.Todo) {
	db := db.GetDB()
	db.Create(&todo)
}

func DeleteTodo(id string) {
	db := db.GetDB()
	db.Where("id=?", id).Delete(&models.Todo{})
}

func UpdateTodo(todo *models.Todo, id string) {
	db := db.GetDB()
	db.Model(&todo).Where("id=?", id).Updates(
        map[string]interface{}{
            "title":     todo.Title,
            "comment":    todo.Comment,
        })
}