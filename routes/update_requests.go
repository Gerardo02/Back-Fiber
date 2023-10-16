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
