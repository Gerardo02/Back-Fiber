package models

import "time"

type Alumno struct {
	ID                  uint `json:"id" gorm:"primaryKey"`
	CreatedAt           time.Time
	ControlEscolarRefer int            `json:"control_escolar_id"`
	ControlEscolar      ControlEscolar `gorm:"foreignKey:ControlEscolarRefer"`
	Nombre              string         `json:"nombre"`
	Matricula           string         `json:"matricula"`
}
