package services

import (
	"github.com/RedContritio/E-commerce-system/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) CreateUser(user models.User) error {
	result := s.DB.Create(&user)
	return result.Error
}