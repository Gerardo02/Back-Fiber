package models

import "gorm.io/gorm"

type Pagos struct {
	gorm.Model
	ID                  uint             `json:"id" gorm:"primaryKey"`
	Historial           string           `json:"historial"`
	AdministracionRefer int              `json:"administracion_id"`
	Administracion      Administraciones `gorm:"foreignKey:AdministracionRefer"`
}
