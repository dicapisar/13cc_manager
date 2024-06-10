package models

import "time"

type ItemType struct {
	ID          uint      `gorm:"primaryKey, autoIncrement, not null, unique"`
	Name        string    `gorm:"not null, unique"`
	Description string    `gorm:"not null"`
	Status      bool      `gorm:"not null"`
	CreateDate  time.Time `gorm:"not null"`
	UpdateDate  time.Time `gorm:"not null"`
	CreatedByID uint      `gorm:"not null"`
	UpdatedByID uint      `gorm:"not null"`
}

func (i *ItemType) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          i.ID,
		"name":        i.Name,
		"description": i.Description,
		"status":      i.Status,
		"create_date": i.CreateDate,
		"update_date": i.UpdateDate,
		"created_by":  i.CreatedByID,
		"updated_by":  i.UpdatedByID,
	}
}
