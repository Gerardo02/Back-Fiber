package models

type Permisos struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Permiso  string `json:"permiso"`
	Password string `json:"password"`
}
