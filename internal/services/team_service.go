package services

import (
	"icematrix/internal/models"
	"icematrix/internal/repositories"
)

type TeamService struct {
	teamRepo *repositories.TeamRepository
}

func NewTeamService(teamRepo *repositories.TeamRepository) *TeamService {
	return &TeamService{teamRepo: teamRepo}
}

// Получение всех команд
func (s *TeamService) GetAllTeams() ([]models.Team, error) {
	teams, err := s.teamRepo.GetAllTeams()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

// Получение команды по ID
func (s *TeamService) GetTeamByID(teamID string) (*models.Team, error) {
	team, err := s.teamRepo.GetTeamByID(teamID)
	if err != nil {
		return nil, err
	}
	return team, nil
}
