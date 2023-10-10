package models

import "gorm.io/gorm"

type Usuarios struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
	Derechos string `json:"derechos"`
}
