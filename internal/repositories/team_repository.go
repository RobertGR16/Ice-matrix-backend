package repositories

import (
	"icematrix/internal/models"

	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// Получение всех команд
func (r *TeamRepository) GetAllTeams() ([]models.Team, error) {
	var teams []models.Team
	if err := r.db.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

// Получение команды по ID
func (r *TeamRepository) GetTeamByID(teamID string) (*models.Team, error) {
	var team models.Team
	if err := r.db.First(&team, "id = ?", teamID).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
