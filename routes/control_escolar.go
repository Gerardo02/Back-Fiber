package routes

import (
	"time"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type ControlEscolar struct {
	ID              uint `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time
	NombreSalon     string `json:"nombre_salon"`
	Especialidad    string `json:"especialidad"`
	CantidadAlumnos int    `json:"cantidad_alumnos"`
}

func CreateResponseControl(controlEscolarModel models.ControlEscolar) ControlEscolar {
	return ControlEscolar{
		ID:              controlEscolarModel.ID,
		NombreSalon:     controlEscolarModel.NombreSalon,
		Especialidad:    controlEscolarModel.Especialidad,
		CantidadAlumnos: controlEscolarModel.CantidadAlumnos,
	}
}

func CreateControlEscolar(c *fiber.Ctx) error {
	var control models.ControlEscolar

	if err := c.BodyParser(&control); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&control)

	responseControl := CreateResponseControl(control)
	return c.Status(200).JSON(responseControl)
}
