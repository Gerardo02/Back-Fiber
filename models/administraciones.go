package models

import "gorm.io/gorm"

type Administraciones struct {
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey"`
	Adeudo bool `json:"adeudo"`
}
