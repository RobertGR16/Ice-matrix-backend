package repositories

import (
	"icematrix/internal/models"

	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

// Получение всех игроков команды
func (r *PlayerRepository) GetPlayersByTeamID(teamID string) ([]models.Player, error) {
	var players []models.Player
	if err := r.db.Preload("Team").Where("team_id = ?", teamID).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

// Метод для получения игрока по ID
func (r *PlayerRepository) GetPlayerByID(playerID uint) (*models.Player, error) {
	var player models.Player
	if err := r.db.Preload("Team").First(&player, playerID).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
