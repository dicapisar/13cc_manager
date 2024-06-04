package services

import (
	"github.com/dicapisar/13cc_manager/dtos/request"
	"github.com/dicapisar/13cc_manager/models"
	"github.com/dicapisar/13cc_manager/repository"
	"time"
)

type UserService interface {
	CreateNewUser(newUserRequest *request.CreateNewUser) (*models.User, error)
}

type UserServiceImpl struct {
	UserRepository *repository.UserRepositoryImpl
}

func (u *UserServiceImpl) CreateNewUser(newUserRequest *request.CreateNewUser) (*models.User, error) {

	newUser := models.User{
		Name:        newUserRequest.Name,
		Password:    newUserRequest.Password,
		Phone:       newUserRequest.Phone,
		Email:       newUserRequest.Email,
		Active:      true,
		CreateDate:  time.Now(),
		UpdateDate:  time.Now(),
		CreatedByID: 1,
		UpdatedByID: 1,
	}

	newUserCreated, err := u.UserRepository.SaveNewUser(&newUser)

	if err != nil {
		return nil, err
	}

	return newUserCreated, nil
}
