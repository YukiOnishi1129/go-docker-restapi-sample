package repositories

import (
	"myapp/db"
	"myapp/models"

	"github.com/pkg/errors"
)

/*
  Todoリストを取得
*/
func GetAllTodos(todo *[]models.Todo, userId int) error {
	db := db.GetDB()
	if err := db.Joins("User").Where("user_id=?", userId).Find(&todo).Error; err != nil {
		return err
	}

	return nil
}

/*
  Idに紐づくTodoデータを取得
*/
func GetTodoById(todo *models.Todo, id string, userId int) error {
	db := db.GetDB()
	if err := db.Joins("User").Where("user_id=?", userId).First(&todo, id).Error; err != nil {
		return err
	}

	return nil
}

/*
 新規登録したTodoデータを取得
*/
func GetTodoLastByUserId(todo *models.Todo, userId int) error {
	db := db.GetDB()
	if err := db.Joins("User").Where("user_id=?", userId).Last(&todo).Error; err != nil {
		return err
	}

	return nil
}

/*
 Todo新規登録
*/
func CreateTodo(todo *models.Todo) error {
	db := db.GetDB()
	if err := db.Create(&todo).Error; err != nil {
		return err
	}

	return nil
}

/*
 Todo削除処理
*/
func DeleteTodo(id string, userId int) error {
	db := db.GetDB()
	db.Where("id=? AND user_id=?", id, userId).Delete(&models.Todo{})
	// https://stackoverflow.com/questions/67154864/how-to-handle-gorm-error-at-delete-function
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		return errors.Errorf("id=%w のTodoデータが存在しません。", id)
	}

	return nil
}

/*
 Todo更新処理
*/
func UpdateTodo(todo *models.Todo, id string, userId int) error {
	db := db.GetDB()
	if err := db.Model(&todo).Where("id=? AND user_id=?", id, userId).Updates(
        map[string]interface{}{
            "title":     todo.Title,
            "comment":    todo.Comment,
        }).Error; err != nil {
			return err
		}
	return nil
}