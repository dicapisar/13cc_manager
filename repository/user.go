package repository

import (
	"github.com/dicapisar/13cc_manager/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveNewUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
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

func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := u.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
