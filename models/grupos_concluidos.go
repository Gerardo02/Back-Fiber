package models

import "gorm.io/gorm"

type GruposConcluidos struct {
	gorm.Model
	ID                int            `json:"id" gorm:"primaryKey"`
	Nombre            string         `json:"nombre"`
	CantidadAlumnos   int            `json:"cantidad_de_alumnos"`
	EspecialidadRefer int            `json:"especialidad_id"`
	Especialidad      Especialidades `gorm:"foreignKey:EspecialidadRefer"`
	CicloEscolarRefer int            `json:"ciclo_escolar_id"`
	CicloEscolar      CicloEscolar   `gorm:"foreignKey:CicloEscolarRefer"`
}
