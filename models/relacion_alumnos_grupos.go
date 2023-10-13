package models

import "gorm.io/gorm"

type RelacionAlumnoGrupo struct {
	gorm.Model
	ID                   uint             `json:"id" gorm:"primaryKey"`
	AlumnoRefer          int              `json:"alumno_id"`
	GruposActivosRefer   int              `json:"grupo_activo_id"`
	GruposAprobadosRefer int              `json:"grupo_aprobado_id"`
	EspecialidadRefer    int              `json:"especialidad_id"`
	Especialidad         Especialidades   `gorm:"foreignKey:EspecialidadRefer"`
	Alumno               Alumnos          `gorm:"foreignKey:AlumnoRefer"`
	GruposActivos        GruposActivos    `gorm:"foreignKey:GruposActivosRefer;default:null"`
	GruposConcluidos     GruposConcluidos `gorm:"foreignKey:GruposAprobadosRefer;default:null"`
}
