package models

type RelacionGrupoLista struct {
	ID                 int           `json:"id" gorm:"primaryKey"`
	ListaAsistencia    string        `json:"lista_asistencia"`
	GruposActivosRefer int           `json:"grupo_id"`
	GruposActivos      GruposActivos `gorm:"foreignKey:GruposActivosRefer"`
}
