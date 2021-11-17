package repositories

import (
	"myapp/db"
	"myapp/models"
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
func GetTodoById(todo *[]models.Todo, id string, userId int) error {
	db := db.GetDB()
	if err := db.Joins("User").Where("user_id=?", userId).First(&todo, id).Error; err != nil {
		return err
	}

	return nil
}

/*
 Todo新規登録
*/
func InsertTodo(todo *models.Todo) error {
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
	if err := db.Where("id=? AND user_id=?", id, userId).Delete(&models.Todo{}).Error; err != nil {
		return err
	}

	return nil
}

/*
 Todo更新処理
*/
func UpdateTodo(todo *models.Todo, id string) error {
	db := db.GetDB()
	if err := db.Model(&todo).Where("id=? AND user_id=?", id, todo.UserId).Updates(
        map[string]interface{}{
            "title":     todo.Title,
            "comment":    todo.Comment,
			"user_id": todo.UserId,
        }).Error; err != nil {
			return err
		}
	return nil
}