package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey, autoIncrement, not null, unique"`
	Name        string    `gorm:"not null, unique"`
	Password    string    `gorm:"not null"`
	Phone       string    `gorm:"not null"`
	Email       string    `gorm:"not null, unique"`
	Active      bool      `gorm:"not null"`
	CreateDate  time.Time `gorm:"not null"`
	UpdateDate  time.Time `gorm:"not null"`
	CreatedBy   *User     `gorm:"foreignKey:CreatedByID"`
	CreatedByID uint      `gorm:"not null"`
	UpdatedBy   *User     `gorm:"foreignKey:UpdatedByID"`
	UpdatedByID uint      `gorm:"not null"`
}
