package handlers

import (
	auth2 "github.com/dicapisar/13cc_manager/auth"
	"github.com/dicapisar/13cc_manager/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Dependencies struct {
	UserService     *services.UserServiceImpl
	LoginService    *services.LoginServiceImpl
	UserRolService  *services.UserRolServiceImpl
	ItemTypeService *services.ItemTypeServiceImpl
	Auth            *auth2.Auth
	SessionStore    *session.Store
}

func AddingHandlers(app *fiber.App, dependencies *Dependencies) error {
	app.Get("/ping", pingHandler).Name("ping")
	setUsersGroupHandler(app, dependencies)
	setHomeHandler(app, dependencies)
	setLoginHandler(app, dependencies)
	setLogoutHandler(app, dependencies)
	setInventoryHandler(app, dependencies)
	return nil
}

func setUsersGroupHandler(app *fiber.App, dependencies *Dependencies) {
	userHandler := UserHandlerImpl{
		userService: dependencies.UserService,
	}
	// TODO: PONER RESTRICCIÓN DE SOLAMENTE ADMIN Y MANAGER
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

func setInventoryHandler(app *fiber.App, dependencies *Dependencies) {
	inventoryHandler := InventoryHandlerImpl{
		SessionStore:    dependencies.SessionStore,
		ItemTypeService: dependencies.ItemTypeService,
	}
	inventoryRouteGroupHandler := app.Group("/inventory").Name("inventory:")
	inventoryRouteGroupHandler.Get("/items_type/new", inventoryHandler.ItemsTypeNewGet).Name("items_type_new")
	inventoryRouteGroupHandler.Post("/items_type/new", inventoryHandler.ItemsTypeNewPost).Name("items_type_new_post")
	inventoryRouteGroupHandler.Get("/items_type/list", inventoryHandler.ItemsTypeListGet).Name("items_type_list")
	inventoryRouteGroupHandler.Get("/items_type/:id", inventoryHandler.ItemsTypeEditByIDGet).Name("items_type_edit_by_id")
	inventoryRouteGroupHandler.Post("/items_type/:id", inventoryHandler.ItemsTypeEditByIDPost).Name("items_type_edit_by_id_post")
}
