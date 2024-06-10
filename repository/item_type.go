package repository

import (
	"github.com/dicapisar/13cc_manager/dtos/request"
	"github.com/dicapisar/13cc_manager/models"
	"gorm.io/gorm"
	"time"
)

type ItemTypeRepository interface {
}

type ItemTypeRepositoryImpl struct {
	DB *gorm.DB
}

func (i *ItemTypeRepositoryImpl) SaveNewItemType(itemType request.NewItemType, userID uint) (*models.ItemType, error) {
	newItemType := models.ItemType{
		Name:        itemType.Name,
		Description: itemType.Description,
		Status:      true,
		CreateDate:  time.Now(),
		UpdateDate:  time.Now(),
		CreatedByID: userID,
		UpdatedByID: userID,
	}

	result := i.DB.Create(&newItemType)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newItemType, nil
}

func (i *ItemTypeRepositoryImpl) FindAllItemTypes() ([]models.ItemType, error) {
	var itemTypes []models.ItemType

	result := i.DB.Order("id").Find(&itemTypes)

	if result.Error != nil {
		return nil, result.Error
	}

	return itemTypes, nil
}

func (i *ItemTypeRepositoryImpl) FindItemTypeByID(id uint) (*models.ItemType, error) {
	var itemType models.ItemType

	result := i.DB.First(&itemType, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &itemType, nil
}

func (i *ItemTypeRepositoryImpl) UpdateItemType(itemType *models.ItemType) (*models.ItemType, error) {
	result := i.DB.Save(itemType)

	if result.Error != nil {
		return nil, result.Error
	}

	return itemType, nil
}
