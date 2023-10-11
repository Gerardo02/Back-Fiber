package models

import "gorm.io/gorm"

type Alumnos struct {
	gorm.Model
	ID                   uint             `json:"id" gorm:"primaryKey"`
	Nombre               string           `json:"nombre"`
	Apellidos            string           `json:"apellidos"`
	Matricula            string           `json:"matricula"`
	Edad                 uint             `json:"edad"`
	GrupoActivoRefer     int              `json:"grupo_id"`
	GrupoActivo          GruposActivos    `gorm:"foreignKey:GrupoActivoRefer"`
	GruposAprobadosRefer int              `json:"grupo_aprobado_id"`
	GruposAprobados      GruposConcluidos `gorm:"foreignKey:GruposAprobadosRefer"`
}
