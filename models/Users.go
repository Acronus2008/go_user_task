package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id       int64
	Name     string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Task     []Task
}
