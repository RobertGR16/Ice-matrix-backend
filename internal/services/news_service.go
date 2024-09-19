package services

import (
	"icematrix/internal/models"
	"icematrix/internal/repositories"
)

type NewsService struct {
	newsRepo *repositories.NewsRepository
}

func NewNewsService(newsRepo *repositories.NewsRepository) *NewsService {
	return &NewsService{newsRepo: newsRepo}
}

// Получение всех новостей, отсортированных по дате от новейших к старейшим
func (s *NewsService) GetAllNews() ([]models.News, error) {
	news, err := s.newsRepo.GetAllNewsSortedByDate()
	if err != nil {
		return nil, err
	}
	return news, nil
}

// Получение новости по ID
func (s *NewsService) GetNewsByID(newsID uint) (*models.News, error) {
	news, err := s.newsRepo.GetNewsByID(newsID)
	if err != nil {
		return nil, err
	}
	return news, nil
}
