package repository

import (
	"github.com/dicapisar/13cc_manager/models"
	"gorm.io/gorm"
)

type UserRolRepository interface {
	GetUserRolByUserId(userId uint) (*models.UserRol, error)
}

type UserRolRepositoryImpl struct {
	DB *gorm.DB
}

func (u *UserRolRepositoryImpl) GetUserRolByUserId(userId uint) (*models.UserRol, error) {

	var userRol models.UserRol

	result := u.DB.Preload("Rol").Preload("User").Where("user_id = ?", userId).First(&userRol)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userRol, nil

}
