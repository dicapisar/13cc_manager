package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey, autoIncrement, not null, unique"`
	Name        string    `gorm:"type:text; not null, unique"`
	Password    string    `gorm:"type:text; not null"`
	Phone       string    `gorm:"type:text; not null"`
	Email       string    `gorm:"type:text; not null, unique"`
	Active      bool      `gorm:"not null"`
	CreateDate  time.Time `gorm:"not null"`
	UpdateDate  time.Time `gorm:"not null"`
	CreatedByID uint      `gorm:"not null"`
	UpdatedByID uint      `gorm:"not null"`
}
