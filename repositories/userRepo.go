package repositories

import (
	"template-go/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(limit int) ([]models.Users, error)
	GetUserByID(id uint) (models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
	DeleteUser(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetAllUsers(limit int) ([]models.Users, error) {
	var users []models.Users
	if err := r.DB.Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id uint) (models.Users, error) {
	var user models.Users
	if err := r.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user models.Users) (models.Users, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user models.Users) (models.Users, error) {
	if err := r.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id uint) error {
	if err := r.DB.Delete(&models.Users{}, id).Error; err != nil {
		return err
	}
	return nil
}
