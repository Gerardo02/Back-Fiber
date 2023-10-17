package models

type Permisos struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Permiso string `json:"permiso"`
}
