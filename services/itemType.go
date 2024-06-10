package services

import (
	"github.com/dicapisar/13cc_manager/dtos/request"
	"github.com/dicapisar/13cc_manager/models"
	"github.com/dicapisar/13cc_manager/repository"
	"time"
)

type ItemTypeService interface {
	CreateNewItemType(newTypeItem request.NewItemType, userID uint) (*models.ItemType, error)
}

type ItemTypeServiceImpl struct {
	ItemTypeRepository *repository.ItemTypeRepositoryImpl
}

func (i *ItemTypeServiceImpl) CreateNewItemType(newTypeItem request.NewItemType, userID uint) (*models.ItemType, error) {
	return i.ItemTypeRepository.SaveNewItemType(newTypeItem, userID)
}

func (i *ItemTypeServiceImpl) GetAllItemTypes() ([]models.ItemType, error) {
	return i.ItemTypeRepository.FindAllItemTypes()
}

func (i *ItemTypeServiceImpl) GetItemTypeByID(id uint) (*models.ItemType, error) {
	return i.ItemTypeRepository.FindItemTypeByID(id)
}

func (i *ItemTypeServiceImpl) UpdateItemTypeByID(itemTypeId uint, req request.NewItemType, updaterId uint) (*models.ItemType, error) {

	itemType, err := i.ItemTypeRepository.FindItemTypeByID(itemTypeId)

	if err != nil {
		return nil, err
	}

	itemType.Name = req.Name
	itemType.Description = req.Description
	itemType.UpdatedByID = updaterId
	itemType.UpdateDate = time.Now()

	return i.ItemTypeRepository.UpdateItemType(itemType)
}
