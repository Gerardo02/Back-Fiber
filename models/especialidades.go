package models

import "gorm.io/gorm"

type Especialidades struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"primaryKey"`
	Materia      string `json:"materia"`
	Especialidad string `json:"especialidad"`
}
