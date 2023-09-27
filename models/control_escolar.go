package models

import "time"

type ControlEscolar struct {
	ID              uint `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time
	NombreSalon     string `json:"nombre_salon"`
	Especialidad    string `json:"especialidad"`
	CantidadAlumnos int    `json:"cantidad_alumnos"`
}
