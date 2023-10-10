package models

import (
	"gorm.io/gorm"
)

type Administracion struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	AlumnoRefer int    `json:"alumno_id"`
	Alumno      Alumno `gorm:"foreignKey:AlumnoRefer"`
	Debe        bool   `json:"debe"`
}
