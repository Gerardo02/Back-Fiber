package models

import "gorm.io/gorm"

type Grupos struct {
	gorm.Model
	ID                uint           `json:"id" gorm:"primaryKey"`
	Nombre            string         `json:"nombre"`
	CantidadAlumnos   int            `json:"cantidad_de_alumnos"`
	Finalizado        bool           `json:"finalizado"`
	Trimestre         int            `json:"trimestre"`
	EspecialidadRefer int            `json:"especialidad_id"`
	Especialidad      Especialidades `gorm:"foreignKey:EspecialidadRefer"`
}
