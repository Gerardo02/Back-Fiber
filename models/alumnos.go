package models

import "gorm.io/gorm"

type Alumnos struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey"`
	GrupoRefer int    `json:"grupo_id"`
	Grupo      Grupos `gorm:"foreignKey:GrupoRefer"`
}
