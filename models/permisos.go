package models

type Permisos struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Permiso        string `json:"permiso"`
	Administracion bool   `json:"administracion"`
	ControlEscolar bool   `json:"control_escolar"`
	Administrador  bool   `json:"administrador"`
	Inscripcion    bool   `json:"inscripcion"`
}
