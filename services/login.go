package services

import (
	"errors"
	auth2 "github.com/dicapisar/13cc_manager/auth"
	"github.com/dicapisar/13cc_manager/models"
	"github.com/dicapisar/13cc_manager/repository"
)

type LoginService interface {
	Login(email, password string) (*models.User, error)
}

type LoginServiceImpl struct {
	UserRepository *repository.UserRepositoryImpl
	Auth           *auth2.Auth
}

func (l *LoginServiceImpl) Login(email, password string) (*models.User, error) {
	user, err := l.UserRepository.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	isMatched := l.Auth.CheckPasswordHash(&password, user.Password)

	if !isMatched {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
