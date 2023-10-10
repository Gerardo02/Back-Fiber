package models

type Derechos struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Derecho string `json:"derecho"`
}
