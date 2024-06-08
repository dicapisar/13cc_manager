package services

import (
	"github.com/dicapisar/13cc_manager/models"
	"github.com/dicapisar/13cc_manager/repository"
)

type UserRolService interface {
}

type UserRolServiceImpl struct {
	UserRolRepository *repository.UserRolRepositoryImpl
}

func (u *UserRolServiceImpl) GetUserRolByID(id int) (*models.UserRol, error) {
	userRol, err := u.UserRolRepository.GetUserRolByUserId(uint(id))

	if err != nil {
		return nil, err
	}

	return userRol, nil
}
