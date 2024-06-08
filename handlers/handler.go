package handlers

import (
	auth2 "github.com/dicapisar/13cc_manager/auth"
	"github.com/dicapisar/13cc_manager/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Dependencies struct {
	UserService    *services.UserServiceImpl
	LoginService   *services.LoginServiceImpl
	UserRolService *services.UserRolServiceImpl
	Auth           *auth2.Auth
	SessionStore   *session.Store
}

func AddingHandlers(app *fiber.App, dependencies *Dependencies) error {
	app.Get("/ping", pingHandler).Name("ping")
	setUsersGroupHandler(app, dependencies)
	setHomeHandler(app, dependencies)
	setLoginHandler(app, dependencies)
	setLogoutHandler(app, dependencies)
	return nil
}

func setUsersGroupHandler(app *fiber.App, dependencies *Dependencies) {
	userHandler := UserHandlerImpl{
		userService: dependencies.UserService,
	}
	usersRouteGroupHandler := app.Group("/users").Name("users")
	usersRouteGroupHandler.Post("/", userHandler.createNewUserHandler).Name("create_user")
}

func setHomeHandler(app *fiber.App, dependencies *Dependencies) {
	homeHandler := HomeHandlerImpl{
		SessionStore: dependencies.SessionStore,
		Auth:         dependencies.Auth,
	}
	app.Get("/", homeHandler.main).Name("home")
	app.Get("/home", homeHandler.home).Name("main")
}

func setLoginHandler(app *fiber.App, dependencies *Dependencies) {
	loginHandler := LoginHandlerImpl{
		LoginService:   dependencies.LoginService,
		SessionStore:   dependencies.SessionStore,
		UserRolService: dependencies.UserRolService,
	}

	loginRouteGroupHandler := app.Group("/login").Name("login:")
	loginRouteGroupHandler.Get("/", loginHandler.loginGet).Name("get")
	loginRouteGroupHandler.Post("/", loginHandler.loginPost).Name("post")
}

func setLogoutHandler(app *fiber.App, dependencies *Dependencies) {
	logoutHandler := LogoutHandlerImpl{
		SessionStore: dependencies.SessionStore,
	}
	app.Get("/logout", logoutHandler.logoutGet).Name("logout")
}
