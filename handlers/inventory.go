package handlers

import (
	"github.com/dicapisar/13cc_manager/commos/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type InventoryHandler interface {
}

type InventoryHandlerImpl struct {
	SessionStore *session.Store
}

func (i *InventoryHandlerImpl) ItemsTypeNewGet(c *fiber.Ctx) error {

	userData, err := utils.GetUserDataFromSessionStorage(i.SessionStore, c)

	if err != nil {
		// TODO: CREAR HANDLER DE ERRORES 500 Y 400
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userData["title"] = "New Item Type"

	return c.Render("inventory/items_type/new", userData, "layouts/main")
}

func (i *InventoryHandlerImpl) ItemsTypeListGet(c *fiber.Ctx) error {
	userData, err := utils.GetUserDataFromSessionStorage(i.SessionStore, c)

	if err != nil {
		// TODO: CREAR HANDLER DE ERRORES 500 Y 400
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userData["title"] = "Items Type List"

	return c.Render("inventory/items_type/list", userData, "layouts/main")
}
