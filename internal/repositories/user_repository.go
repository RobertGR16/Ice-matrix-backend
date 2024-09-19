package repositories

import (
	"icematrix/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Создание пользователя
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// Получение пользователя по ID
func (r *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Получение всех пользователей
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Обновление пользователя
func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

// Удаление пользователя
func (r *UserRepository) DeleteUser(userID uint) error {
	return r.db.Delete(&models.User{}, userID).Error
}
