package handlers

import (
	"icematrix/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	matchService *services.MatchService
}

func NewMatchHandler(matchService *services.MatchService) *MatchHandler {
	return &MatchHandler{matchService: matchService}
}

// Получение матча по ID
func (h *MatchHandler) GetMatchByID(c *gin.Context) {
	idParam := c.Param("id")
	matchID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid match ID"})
		return
	}

	match, err := h.matchService.GetMatchByID(uint(matchID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "match not found"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// Получение всех матчей, отсортированных по дате
func (h *MatchHandler) GetAllMatches(c *gin.Context) {
	matches, err := h.matchService.GetAllMatchesSortedByDate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve matches"})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// Получение предстоящих матчей
func (h *MatchHandler) GetUpcomingMatches(c *gin.Context) {
	matches, err := h.matchService.GetUpcomingMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve upcoming matches"})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// Получение завершенных матчей
func (h *MatchHandler) GetFinishedMatches(c *gin.Context) {
	matches, err := h.matchService.GetFinishedMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve finished matches"})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// Получение всех матчей команды (сначала предстоящие, затем прошедшие)
func (h *MatchHandler) GetMatchesByTeamID(c *gin.Context) {
	teamID := c.Param("teamID")

	matches, err := h.matchService.GetMatchesByTeamID(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve matches for the team"})
		return
	}

	c.JSON(http.StatusOK, matches)
}
