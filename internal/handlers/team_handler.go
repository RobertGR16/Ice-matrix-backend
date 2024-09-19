package handlers

import (
	"icematrix/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService *services.TeamService
}

func NewTeamHandler(teamService *services.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

// Получение всех команд
func (h *TeamHandler) GetAllTeams(c *gin.Context) {
	teams, err := h.teamService.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve teams"})
		return
	}

	c.JSON(http.StatusOK, teams)
}

// Получение команды по ID
func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	idParam := c.Param("id")

	team, err := h.teamService.GetTeamByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}
