package models

import "gorm.io/gorm"

type Alumnos struct {
	gorm.Model
	ID              uint   `json:"id" gorm:"primaryKey"`
	Nombre          string `json:"nombre"`
	Apellidos       string `json:"apellidos"`
	Matricula       string `json:"matricula"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Edad            uint   `json:"edad"`
	NombreTutor     string `json:"nombre_tutor"`
	CelularTutor    string `json:"celular_tutor"`
	Curp            string `json:"curp"`
	Localidad       string `json:"localidad"`
	CodigoPostal    string `json:"codigo_postal"`
	Direccion       string `json:"direccion"`
	TelefonoFijo    string `json:"telefono_fijo"`
	Celular         string `json:"celular"`
	Correo          string `json:"correo"`
}
