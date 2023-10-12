package models

import "gorm.io/gorm"

type Alumnos struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Matricula string `json:"matricula"`
	Edad      uint   `json:"edad"`
}
