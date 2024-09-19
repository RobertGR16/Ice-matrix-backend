package repositories

import (
	"icematrix/internal/models"

	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) GetAllNewsSortedByDate() ([]models.News, error) {
	var news []models.News
	if err := r.db.Order("published_at desc").Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

// Получение новости по ID
func (r *NewsRepository) GetNewsByID(newsID uint) (*models.News, error) {
	var news models.News
	if err := r.db.First(&news, newsID).Error; err != nil {
		return nil, err
	}
	return &news, nil
}
