package models

import "gorm.io/gorm"

type Administraciones struct {
	gorm.Model
	ID          int  `json:"id" gorm:"primaryKey"`
	Adeudo      bool `json:"adeudo"`
	Estado      int  `json:"estado"`
	Dinero      int  `json:"dinero"`
	AlumnoRefer int  `json:"alumno_id"`
}
