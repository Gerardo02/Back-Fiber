package models

import "gorm.io/gorm"

type Administraciones struct {
	gorm.Model
	ID              uint   `json:"id" gorm:"primaryKey"`
	Adeudo          bool   `json:"adeudo"`
	Estado          int    `json:"estado"`
	AlumnoRefer     int    `json:"alumno_id"`
	AlumnoNombre    string `gorm:"foreignKey:AlumnoRefer"`
	AlumnoMatricula string `gorm:"foreignKey:AlumnoRefer"`
}
