package models

type Permisos struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Permiso string `json:"permiso"`
}
