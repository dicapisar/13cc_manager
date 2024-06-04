package models

type Rol struct {
	ID   uint   `gorm:"primaryKey, autoIncrement, not null, unique"`
	Name string `gorm:"not null, unique"`
}
