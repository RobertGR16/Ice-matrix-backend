package handlers

import (
	"icematrix/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	newsService *services.NewsService
}

func NewNewsHandler(newsService *services.NewsService) *NewsHandler {
	return &NewsHandler{newsService: newsService}
}

// Получение всех новостей
func (h *NewsHandler) GetAllNews(c *gin.Context) {
	news, err := h.newsService.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve news"})
		return
	}

	c.JSON(http.StatusOK, news)
}

// Получение новости по ID
func (h *NewsHandler) GetNewsByID(c *gin.Context) {
	idParam := c.Param("id")
	newsID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid news ID"})
		return
	}

	news, err := h.newsService.GetNewsByID(uint(newsID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "news not found"})
		return
	}

	c.JSON(http.StatusOK, news)
}
