package models

import "gorm.io/gorm"

type HistorialLogs struct {
	gorm.Model
	ID         int    `json:"id" gorm:"primaryKey"`
	Nombre     string `json:"nombre"`
	Movimiento string `json:"movimiento"`
	Monto      int    `json:"monto"`
}
