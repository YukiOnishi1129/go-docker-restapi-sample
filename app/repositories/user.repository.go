package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	GetAllUserByEmail(users *[]models.User, email string) error
	CreateUser(createUsers *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// GetUserByEmail emailに紐づくユーザーリストを取得
func (ur *userRepository) GetUserByEmail(user *models.User, email string) error {
	// db := db.GetDB()
	if err := ur.db.Where("email=?", email).First(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetAllUserByEmail emailに紐づくユーザーリストを取得
func (ur *userRepository) GetAllUserByEmail(users *[]models.User, email string) error {
	// db := db.GetDB()
	if err := ur.db.Where("email=?", email).Find(&users).Error; err != nil {
		return err
	}

	return nil
}

// CreateUser ユーザーデータ新規登録
func (ur *userRepository) CreateUser(createUsers *models.User) error {
	// db := db.GetDB()
	if err := ur.db.Create(&createUsers).Error; err != nil {
		return err
	}

	return nil
}
