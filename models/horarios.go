package models

type Horarios struct {
	ID            int           `json:"id" gorm:"primaryKey"`
	Dia           string        `json:"dia"`
	Entrada       string        `json:"entrada"`
	Salida        string        `json:"salida"`
	DiaData       int           `json:"diaData"`
	GrupoRefer    int           `json:"grupo_id"`
	GruposActivos GruposActivos `gorm:"foreignKey:GrupoRefer"`
}
