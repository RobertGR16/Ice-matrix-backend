package services

import (
	"icematrix/internal/models"
	"icematrix/internal/repositories"
)

type PlayerService struct {
	playerRepo *repositories.PlayerRepository
}

func NewPlayerService(playerRepo *repositories.PlayerRepository) *PlayerService {
	return &PlayerService{playerRepo: playerRepo}
}

// Получение всех игроков команды
func (s *PlayerService) GetPlayersByTeamID(teamID string) ([]models.Player, error) {
	players, err := s.playerRepo.GetPlayersByTeamID(teamID)
	if err != nil {
		return nil, err
	}
	return players, nil
}

// Получение игрока по ID
func (s *PlayerService) GetPlayerByID(playerID uint) (*models.Player, error) {
	player, err := s.playerRepo.GetPlayerByID(playerID)
	if err != nil {
		return nil, err
	}
	return player, nil
}
