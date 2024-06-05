package auth

import (
	"github.com/dicapisar/13cc_manager/config"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Config *config.AuthConfig
}

func (a *Auth) HashPassword(password *string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), int(a.Config.Cost))
	return string(bytes), err
}

func (a *Auth) CheckPasswordHash(password *string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(*password))
	return err == nil
}

func (a *Auth) IsLoggedIn(session *session.Session) bool {
	loggedIn, ok := session.Get("logged_in").(bool)
	return ok && loggedIn
}
