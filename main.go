package main

import (
	"encoding/json"
	"fmt"
	"icematrix/internal/handlers"
	"icematrix/internal/middleware"
	"icematrix/internal/models"
	"icematrix/internal/repositories"
	"icematrix/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Загрузка переменных из .env файла
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Подключение к базе данных
func ConnectDatabase() *gorm.DB {
	LoadEnv()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

// Функция для загрузки данных из JSON файла и сохранения их в базу данных
func LoadDataFromJSON(db *gorm.DB, filepath string, model interface{}) error {
	byteValue, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %v", err)
	}

	err = json.Unmarshal(byteValue, model)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	if err := db.Create(model).Error; err != nil {
		return fmt.Errorf("failed to save data: %v", err)
	}

	return nil
}

// Хендлер для загрузки данных всех моделей
func LoadAllDataHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Загрузка данных для команд
		var teams []models.Team
		if err := LoadDataFromJSON(db, "preload_data/teams.json", &teams); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading teams: " + err.Error()})
			return
		}

		// Загрузка данных для игроков
		var players []models.Player
		if err := LoadDataFromJSON(db, "preload_data/players.json", &players); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading players: " + err.Error()})
			return
		}

		// Загрузка данных для матчей
		var matches []models.Match
		if err := LoadDataFromJSON(db, "preload_data/games.json", &matches); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading matches: " + err.Error()})
			return
		}

		// Загрузка данных для новостей
		var newsList []models.News
		if err := LoadDataFromJSON(db, "preload_data/news.json", &newsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading news: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "All data loaded successfully!"})
	}
}

func main() {
	// Подключаемся к базе данных
	db := ConnectDatabase()

	// Выполняем миграции для всех моделей
	db.AutoMigrate(&models.Team{}, &models.Player{}, &models.Match{}, &models.News{}, &models.User{})

	fmt.Println("Подключение к базе успешно")
	// Инициализация репозиториев и сервисов
	playerRepo := repositories.NewPlayerRepository(db)
	userRepo := repositories.NewUserRepository(db)
	teamRepo := repositories.NewTeamRepository(db)
	matchRepo := repositories.NewMatchRepository(db)
	newsRepo := repositories.NewNewsRepository(db)

	playerService := services.NewPlayerService(playerRepo)
	userService := services.NewUserService(userRepo)
	teamService := services.NewTeamService(teamRepo)
	matchService := services.NewMatchService(matchRepo)
	newsService := services.NewNewsService(newsRepo)

	// Инициализация хендлеров
	jwtSecret := "your_jwt_secret"
	userHandler := handlers.NewUserHandler(userService, jwtSecret)
	playerHandler := handlers.NewPlayerHandler(playerService)
	teamHandler := handlers.NewTeamHandler(teamService)
	matchHandler := handlers.NewMatchHandler(matchService)
	newsHandler := handlers.NewNewsHandler(newsService)

	// Настройка роутера Gin
	router := gin.Default()

	// Публичные маршруты для пользователей
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	// Защищенные маршруты для пользователей с использованием JWT
	authRoutes := router.Group("/")
	authRoutes.Use(middleware.JWTAuthMiddleware(jwtSecret))
	{
		authRoutes.GET("/users/me", userHandler.GetUser)       // Получение информации о текущем пользователе
		authRoutes.DELETE("/users/me", userHandler.DeleteUser) // Удаление аккаунта
		authRoutes.PATCH("/users/me", userHandler.UpdateUser)  // Обновление аккаунта
	}

	// Маршруты для команд
	router.GET("/teams", teamHandler.GetAllTeams)     // Получение всех команд
	router.GET("/teams/:id", teamHandler.GetTeamByID) // Получение команды по ID

	// Маршруты для игроков
	router.GET("/teams/:id/players", playerHandler.GetPlayersByTeam) // Получение игроков по ID команды
	router.GET("/players/:id", playerHandler.GetPlayerByID)          // Получение информации об игроке по его ID

	// Маршруты для матчей
	router.GET("/matches", matchHandler.GetAllMatches)                   // Получение всех матчей
	router.GET("/matches/upcoming", matchHandler.GetUpcomingMatches)     // Получение предстоящих матчей
	router.GET("/matches/finished", matchHandler.GetFinishedMatches)     // Получение завершенных матчей
	router.GET("/matches/:id", matchHandler.GetMatchByID)                // Получение матча по ID
	router.GET("/matches/team/:teamID", matchHandler.GetMatchesByTeamID) // Получение всех матчей команды (сначала предстоящие, затем прошедшие)

	// Маршруты для новостей
	router.GET("/news", newsHandler.GetAllNews)      // Получение всех новостей
	router.GET("/news/:id", newsHandler.GetNewsByID) // Получение новости по ID

	// Маршрут для загрузки всех данных в базу
	router.POST("/load-all-data", LoadAllDataHandler(db)) // Загрузка всех данных (команды, игроки, матчи, новости)

	// Запуск сервера
	router.Run()
}
