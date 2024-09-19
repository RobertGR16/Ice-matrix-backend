package handlers

import (
	"icematrix/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	playerService *services.PlayerService
}

func NewPlayerHandler(playerService *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}

// Получение всех игроков команды
func (h *PlayerHandler) GetPlayersByTeam(c *gin.Context) {
	teamID := c.Param("id")
	players, err := h.playerService.GetPlayersByTeamID(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve players"})
		return
	}

	c.JSON(http.StatusOK, players)
}

// Получение игрока по ID
func (h *PlayerHandler) GetPlayerByID(c *gin.Context) {
	idParam := c.Param("id")
	playerID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	player, err := h.playerService.GetPlayerByID(uint(playerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	c.JSON(http.StatusOK, player)
}
