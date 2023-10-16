package models

type Documentos struct {
	ID                   uint   `json:"id" gorm:"primaryKey"`
	AlumnoRefer          uint   `json:"alumno_id"`
	Nombre               string `gorm:"foreignKey:AlumnoRefer"`
	Apellidos            string `gorm:"foreignKey:AlumnoRefer"`
	Matricula            string `gorm:"foreignKey:AlumnoRefer"`
	ActaNacimiento       bool   `json:"acta_de_nacimiento"`
	Curp                 bool   `json:"curp"`
	ComprobanteDomicilio bool   `json:"comprobante_de_domicilio"`
	MayorQuince          bool   `json:"mayor_quince"`
	Fotos                bool   `json:"fotos"`
}
