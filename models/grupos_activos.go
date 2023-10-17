package models

import "gorm.io/gorm"

type GruposActivos struct {
	gorm.Model
	ID                int            `json:"id" gorm:"primaryKey"`
	Nombre            string         `json:"nombre"`
	CantidadAlumnos   int            `json:"cantidad_de_alumnos"`
	Trimestre         int            `json:"trimestre"`
	EspecialidadRefer int            `json:"especialidad_id"`
	Especialidad      Especialidades `gorm:"foreignKey:EspecialidadRefer"`
}
