package models

import "time"

type Administracion struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	AlumnoRefer int    `json:"alumno_id"`
	Alumno      Alumno `gorm:"foreignKey:AlumnoRefer"`
	Debe        bool   `json:"debe"`
}
