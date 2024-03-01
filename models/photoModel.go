package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Caption  string
	PhotoUrl string `gorm:"not null"`
	UserId   uint   `gorm:"not null"`
}
