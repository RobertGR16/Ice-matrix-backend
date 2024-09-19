package repositories

import (
	"icematrix/internal/models"

	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

// Получение матча по ID
func (r *MatchRepository) GetMatchByID(matchID uint) (*models.Match, error) {
	var match models.Match
	if err := r.db.Preload("Team1").Preload("Team2").First(&match, matchID).Error; err != nil {
		return nil, err
	}
	return &match, nil
}

// Получение всех матчей, отсортированных по дате
func (r *MatchRepository) GetAllMatchesSortedByDate() ([]models.Match, error) {
	var matches []models.Match
	if err := r.db.Preload("Team1").Preload("Team2").Order("date desc").Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

// Получение предстоящих матчей, отсортированных по дате (по возрастанию)
func (r *MatchRepository) GetUpcomingMatches() ([]models.Match, error) {
	var matches []models.Match
	if err := r.db.Preload("Team1").Preload("Team2").
		Where("CAST(date AS TIMESTAMP) > NOW()").
		Order("CAST(date AS TIMESTAMP) asc").Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

// Получение прошедших матчей, отсортированных по дате (по убыванию)
func (r *MatchRepository) GetFinishedMatches() ([]models.Match, error) {
	var matches []models.Match
	if err := r.db.Preload("Team1").Preload("Team2").
		Where("CAST(date AS TIMESTAMP) <= NOW()").
		Order("CAST(date AS TIMESTAMP) desc").Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

// Получение всех матчей конкретной команды, сначала один ближайший предстоящий матч, затем прошедшие
func (r *MatchRepository) GetMatchesByTeamID(teamID string) ([]models.Match, error) {
	var upcomingMatch models.Match // Для одного ближайшего предстоящего матча
	var finishedMatches []models.Match

	// Получаем один ближайший предстоящий матч команды
	if err := r.db.Preload("Team1").Preload("Team2").
		Where("(team1_id = ? OR team2_id = ?) AND CAST(date AS TIMESTAMP) > NOW()", teamID, teamID).
		Order("CAST(date AS TIMESTAMP) asc").Limit(1).Find(&upcomingMatch).Error; err != nil {
		return nil, err
	}

	// Получаем прошедшие матчи команды, отсортированные по дате по убыванию
	if err := r.db.Preload("Team1").Preload("Team2").
		Where("(team1_id = ? OR team2_id = ?) AND CAST(date AS TIMESTAMP) <= NOW()", teamID, teamID).
		Order("CAST(date AS TIMESTAMP) desc").Find(&finishedMatches).Error; err != nil {
		return nil, err
	}

	// Создаем массив для всех матчей
	allMatches := []models.Match{}

	// Добавляем ближайший предстоящий матч, если он существует
	if upcomingMatch.ID != 0 {
		allMatches = append(allMatches, upcomingMatch)
	}

	// Добавляем все прошедшие матчи
	allMatches = append(allMatches, finishedMatches...)

	return allMatches, nil
}
