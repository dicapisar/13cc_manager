package repository

import (
	"github.com/dicapisar/13cc_manager/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveNewUser(user *models.User) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (u *UserRepositoryImpl) SaveNewUser(user *models.User) (*models.User, error) {

	result := u.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
