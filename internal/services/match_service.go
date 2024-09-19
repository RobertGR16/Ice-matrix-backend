package services

import (
	"icematrix/internal/models"
	"icematrix/internal/repositories"
)

type MatchService struct {
	matchRepo *repositories.MatchRepository
}

func NewMatchService(matchRepo *repositories.MatchRepository) *MatchService {
	return &MatchService{matchRepo: matchRepo}
}

// Получение матча по ID
func (s *MatchService) GetMatchByID(matchID uint) (*models.Match, error) {
	match, err := s.matchRepo.GetMatchByID(matchID)
	if err != nil {
		return nil, err
	}
	return match, nil
}

// Получение всех матчей, отсортированных по дате
func (s *MatchService) GetAllMatchesSortedByDate() ([]models.Match, error) {
	matches, err := s.matchRepo.GetAllMatchesSortedByDate()
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// Получение предстоящих матчей
func (s *MatchService) GetUpcomingMatches() ([]models.Match, error) {
	matches, err := s.matchRepo.GetUpcomingMatches()
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// Получение прошедших матчей
func (s *MatchService) GetFinishedMatches() ([]models.Match, error) {
	matches, err := s.matchRepo.GetFinishedMatches()
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// Получение всех матчей для конкретной команды
func (s *MatchService) GetMatchesByTeamID(teamID string) ([]models.Match, error) {
	matches, err := s.matchRepo.GetMatchesByTeamID(teamID)
	if err != nil {
		return nil, err
	}
	return matches, nil
}
