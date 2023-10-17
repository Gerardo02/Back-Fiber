package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateDocuments(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var documents models.Documentos

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findDocuments(id, &documents); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateDocument struct {
		ActaNacimiento       bool `json:"acta_de_nacimiento"`
		Curp                 bool `json:"curp"`
		ComprobanteDomicilio bool `json:"comprobante_de_domicilio"`
		MayorQuince          bool `json:"mayor_quince"`
		Fotos                bool `json:"fotos"`
	}

	var UpdateData UpdateDocument

	if err := c.BodyParser(&UpdateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	documents.ActaNacimiento = UpdateData.ActaNacimiento
	documents.Curp = UpdateData.Curp
	documents.ComprobanteDomicilio = UpdateData.ComprobanteDomicilio
	documents.MayorQuince = UpdateData.MayorQuince
	documents.Fotos = UpdateData.Fotos

	database.Database.Db.Save(&documents)

	responseDocuments := CreateDocumentosResponse(documents)

	return c.Status(200).JSON(responseDocuments)
}

func UpdateAlumnos(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var alumnos models.Alumnos

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findAlumno(id, &alumnos); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedAlumno struct {
		Nombre          string `json:"nombre"`
		Apellidos       string `json:"apellidos"`
		Matricula       string `json:"matricula"`
		FechaNacimiento string `json:"fecha_nacimiento"`
		Edad            int    `json:"edad"`
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

	var UpdatedData UpdatedAlumno

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	alumnos.Nombre = UpdatedData.Nombre
	alumnos.Apellidos = UpdatedData.Apellidos
	alumnos.Matricula = UpdatedData.Matricula
	alumnos.FechaNacimiento = UpdatedData.FechaNacimiento
	alumnos.Edad = UpdatedData.Edad
	alumnos.NombreTutor = UpdatedData.NombreTutor
	alumnos.CelularTutor = UpdatedData.CelularTutor
	alumnos.Curp = UpdatedData.Curp
	alumnos.Localidad = UpdatedData.Localidad
	alumnos.CodigoPostal = UpdatedData.CodigoPostal
	alumnos.Direccion = UpdatedData.Direccion
	alumnos.TelefonoFijo = UpdatedData.TelefonoFijo
	alumnos.Celular = UpdatedData.Celular
	alumnos.Correo = UpdatedData.Correo

	database.Database.Db.Save(&alumnos)

	responseAlumnos := CreateAlumnosResponse(alumnos)

	return c.Status(200).JSON(responseAlumnos)
}

func UpdateAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var admin models.Administraciones

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findAdmin(id, &admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedAdmin struct {
		Adeudo bool `json:"adeudo"`
		Estado int  `json:"estado"`
	}

	var UpdatedData UpdatedAdmin

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	admin.Adeudo = UpdatedData.Adeudo
	admin.Estado = UpdatedData.Estado

	database.Database.Db.Save(&admin)

	responseAdmin := CreateAdminResponse(admin)

	return c.Status(200).JSON(responseAdmin)
}
