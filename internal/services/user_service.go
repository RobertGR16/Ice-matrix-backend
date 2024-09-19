package services

import (
	"errors"
	"icematrix/internal/models"
	"icematrix/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Регистрация нового пользователя
func (s *UserService) RegisterUser(username, email, password string) (*models.User, error) {
	// Проверка на существование пользователя с таким же email
	existingUser, _ := s.userRepo.GetAllUsers()
	for _, user := range existingUser {
		if user.Email == email {
			return nil, errors.New("user with this email already exists")
		}
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Создание нового пользователя
	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Логин пользователя
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	// Поиск пользователя по email
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var foundUser *models.User
	for _, user := range users {
		if user.Email == email {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		return nil, errors.New("user not found")
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return foundUser, nil
}

// Получение информации о пользователе по ID
func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

// Удаление аккаунта пользователя
func (s *UserService) DeleteUser(userID uint) error {
	return s.userRepo.DeleteUser(userID)
}

// Обновление информации о пользователе
func (s *UserService) UpdateUser(userID uint, updates *models.User) (*models.User, error) {
	// Получаем пользователя из базы данных
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Обновляем только те поля, которые переданы
	if updates.Username != "" {
		user.Username = updates.Username
	}
	if updates.Email != "" {
		user.Email = updates.Email
	}
	if updates.AvatarURL != "" {
		user.AvatarURL = updates.AvatarURL
	}
	if updates.FavoriteTeamID != nil {
		user.FavoriteTeamID = updates.FavoriteTeamID
	}

	// Сохраняем обновления в базе данных
	if err := s.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
