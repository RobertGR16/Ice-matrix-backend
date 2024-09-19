package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string  `gorm:"not null"`
	Email          string  `gorm:"not null;unique"`
	Password       string  `gorm:"not null"`
	AvatarURL      string  `gorm:"default:''"`                                                              // Пустая строка по умолчанию
	FavoriteTeamID *string `gorm:"default:null"`                                                            // Указатель на строку, чтобы можно было хранить NULL
	FavoriteTeam   Team    `gorm:"foreignKey:FavoriteTeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Связь с командой
}

type Team struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	Name           string
	LogoURL        string
	DateOfFounding string
	City           string
	Address        string
	Email          string
	Website        string
	Trainer        string
}

type Match struct {
	gorm.Model
	Team1ID string
	Team2ID string
	Team1   Team `gorm:"foreignKey:Team1ID"`
	Team2   Team `gorm:"foreignKey:Team2ID"`
	Date    string

	ScoreTeam1 int
	ScoreTeam2 int

	MajorPucksTeam1       int // Шайбы в большинстве (Team 1)
	MajorPucksTeam2       int // Шайбы в большинстве (Team 2)
	OutnumberedPucksTeam1 int // Шайбы в меньшинстве (Team 1)
	OutnumberedPucksTeam2 int // Шайбы в меньшинстве (Team 2)

	AdvantagesTeam1 int // Численное преимущество (Team 1)
	AdvantagesTeam2 int // Численное преимущество (Team 2)

	FaceoffsWonTeam1 int // Выигранные вбрасывания (Team 1)
	FaceoffsWonTeam2 int // Выигранные вбрасывания (Team 2)

	PenaltyMinutesTeam1 int // Штрафное время (Team 1)
	PenaltyMinutesTeam2 int // Штрафное время (Team 2)

	ShotsOnGoalTeam1 int // Броски по воротам (Team 1)
	ShotsOnGoalTeam2 int // Броски по воротам (Team 2)

	DistanceCoveredTeam1 float64 // Пройденная дистанция в км (Team 1)
	DistanceCoveredTeam2 float64 // Пройденная дистанция в км (Team 2)

	PuckPossessionTimeTeam1 string // Время владения шайбой (Team 1)
	PuckPossessionTimeTeam2 string // Время владения шайбой (Team 2)
}

type Player struct {
	gorm.Model
	Name          string
	PhotoURL      string
	TeamID        string // Внешний ключ на команду
	Team          Team   `gorm:"foreignKey:TeamID"` // Связь с таблицей Team
	Position      string
	Weight        int
	Height        int
	DateOfBirth   string
	Number        int
	Сitizenship   string
	Grip          string
	ContractUntil string
	ContractType  string
}

type News struct {
	gorm.Model
	Title       string
	Content     string
	PublishedAt string
	Source      string
}
