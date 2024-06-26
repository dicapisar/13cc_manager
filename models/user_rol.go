package models

import "time"

type UserRol struct {
	ID          uint      `gorm:"primaryKey, autoIncrement, not null, unique"`
	UserID      uint      `gorm:"not null"`
	RolID       uint      `gorm:"not null"`
	User        *User     `gorm:"foreignKey:UserID"`
	Rol         *Rol      `gorm:"foreignKey:RolID"`
	Status      bool      `gorm:"not null"`
	CreateDate  time.Time `gorm:"not null"`
	UpdateDate  time.Time `gorm:"not null"`
	CreatedByID uint      `gorm:"not null"`
	UpdatedByID uint      `gorm:"not null"`
}
