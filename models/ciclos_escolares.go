package models

import "gorm.io/gorm"

type CicloEscolar struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey"`
	Nombre    string `json:"nombre"`
	Year      string `json:"year"`
	Trimestre int    `json:"trimestre"`
	Activo    bool   `json:"activo"`
}
