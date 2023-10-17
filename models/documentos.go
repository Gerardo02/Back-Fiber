package models

type Documentos struct {
	ID                   int  `json:"id" gorm:"primaryKey"`
	AlumnoRefer          int  `json:"alumno_id"`
	ActaNacimiento       bool `json:"acta_de_nacimiento"`
	Curp                 bool `json:"curp"`
	ComprobanteDomicilio bool `json:"comprobante_de_domicilio"`
	MayorQuince          bool `json:"mayor_quince"`
	Fotos                bool `json:"fotos"`
}
