package models

import "gorm.io/gorm"

type GruposActivos struct {
	gorm.Model
	ID                uint           `json:"id" gorm:"primaryKey"`
	Nombre            string         `json:"nombre"`
	CantidadAlumnos   int            `json:"cantidad_de_alumnos"`
	Trimestre         int            `json:"trimestre"`
	ListaAsistencia   string         `json:"lista_asistencia"` //arreglo
	EspecialidadRefer int            `json:"especialidad_id"`
	Especialidad      Especialidades `gorm:"foreignKey:EspecialidadRefer"`
}
