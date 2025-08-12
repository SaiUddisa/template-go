package services

import (
	"template-go/models"
	"template-go/repositories"
)

type UserService interface {
	GetUsers(limit int) ([]models.Users, error)
	GetUserByID(id uint) (models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUsers(limit int) ([]models.Users, error) {
	return s.repo.GetAllUsers(limit)
}

func (s *userService) GetUserByID(id uint) (models.Users, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) CreateUser(user models.Users) (models.Users, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) UpdateUser(user models.Users) (models.Users, error) {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
