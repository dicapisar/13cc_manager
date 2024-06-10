package handlers

import (
	"github.com/dicapisar/13cc_manager/commos/utils"
	"github.com/dicapisar/13cc_manager/dtos/request"
	"github.com/dicapisar/13cc_manager/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"strconv"
)

type InventoryHandler interface {
}

type InventoryHandlerImpl struct {
	SessionStore    *session.Store
	ItemTypeService *services.ItemTypeServiceImpl
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

func (i *InventoryHandlerImpl) ItemsTypeNewPost(c *fiber.Ctx) error {

	var req request.NewItemType

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userData, err := utils.GetUserDataFromSessionStorage(i.SessionStore, c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err = i.ItemTypeService.CreateNewItemType(req, userData["userId"].(uint))

	return c.Redirect("/inventory/items_type/list")
}

func (i *InventoryHandlerImpl) ItemsTypeListGet(c *fiber.Ctx) error {
	userData, err := utils.GetUserDataFromSessionStorage(i.SessionStore, c)

	if err != nil {
		// TODO: CREAR HANDLER DE ERRORES 500 Y 400
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	itemTypes, err := i.ItemTypeService.GetAllItemTypes()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	mapItemTypes := make([]map[string]interface{}, len(itemTypes))

	for i, itemType := range itemTypes {
		mapItemTypes[i] = itemType.ToMap()
	}

	userData["title"] = "Items Type List"
	userData["itemsType"] = mapItemTypes

	return c.Render("inventory/items_type/list", userData, "layouts/main")
}

func (i *InventoryHandlerImpl) ItemsTypeEditByIDGet(c *fiber.Ctx) error {

	itemTypeID := c.Params("id")

	parseItemTypeID, err := strconv.ParseUint(itemTypeID, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userData, err := utils.GetUserDataFromSessionStorage(i.SessionStore, c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	itemType, err := i.ItemTypeService.GetItemTypeByID(uint(parseItemTypeID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userData["title"] = "Item Type"
	userData["itemType"] = itemType.ToMap()

	return c.Render("inventory/items_type/edit", userData, "layouts/main")

}

func (i *InventoryHandlerImpl) ItemsTypeEditByIDPost(c *fiber.Ctx) error {
	itemTypeID := c.Params("id")

	parseItemTypeID, err := strconv.ParseUint(itemTypeID, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req request.NewItemType

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userData, err := utils.GetUserDataFromSessionStorage(i.SessionStore, c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err = i.ItemTypeService.UpdateItemTypeByID(uint(parseItemTypeID), req, userData["userId"].(uint))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect("/inventory/items_type/list")

}
