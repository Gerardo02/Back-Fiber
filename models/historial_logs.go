package models

import "gorm.io/gorm"

type HistorialLogs struct {
	gorm.Model
	ID          int     `json:"id" gorm:"primaryKey"`
	Movimiento  string  `json:"movimiento"`
	Monto       int     `json:"monto"`
	Hora        string  `json:"hora"`
	Fecha       string  `json:"fecha"`
	AlumnoRefer int     `json:"alumno_id"`
	Alumno      Alumnos `gorm:"foreignKey:AlumnoRefer"`
}
