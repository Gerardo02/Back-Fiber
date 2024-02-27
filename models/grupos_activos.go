package models

import "gorm.io/gorm"

type GruposActivos struct {
	gorm.Model
	ID                int            `json:"id" gorm:"primaryKey"`
	Nombre            string         `json:"nombre"`
	NombreMaestro     string         `json:"nombre_maestro"`
	Dia               string         `json:"dia"`
	Entrada           string         `json:"entrada"`
	Salida            string         `json:"salida"`
	CantidadAlumnos   int            `json:"cantidad_de_alumnos"`
	Trimestre         int            `json:"trimestre"`
	EspecialidadRefer int            `json:"especialidad_id"`
	Especialidad      Especialidades `gorm:"foreignKey:EspecialidadRefer"`
}
