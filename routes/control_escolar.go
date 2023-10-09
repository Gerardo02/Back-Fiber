package routes

import (
	"errors"
	"time"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type ControlEscolar struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time `json:"fecha_de_alta"`
	NombreSalon     string    `json:"nombre_salon"`
	Especialidad    string    `json:"especialidad"`
	CantidadAlumnos int       `json:"cantidad_alumnos"`
}

func findControl(id int, control *models.ControlEscolar) error {
	database.Database.Db.Find(&control, "id = ?", id)

	if control.ID == 0 {
		return errors.New("control does not exist")
	}

	return nil
}

func CreateResponseControl(controlEscolarModel models.ControlEscolar) ControlEscolar {
	return ControlEscolar{
		ID:              controlEscolarModel.ID,
		CreatedAt:       controlEscolarModel.CreatedAt,
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

func UpdateControlEscolar(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var control models.ControlEscolar

	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}

	if err := findControl(id, &control); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedData struct {
		NombreSalon     string `json:"nombre_salon"`
		Especialidad    string `json:"especialidad"`
		CantidadAlumnos int    `json:"cantidad_alumnos"`
	}

	var updatedControl UpdatedData

	if err := c.BodyParser(&updatedControl); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	control.NombreSalon = updatedControl.NombreSalon
	control.Especialidad = updatedControl.Especialidad
	control.CantidadAlumnos = updatedControl.CantidadAlumnos

	database.Database.Db.Save(&control)

	responseControl := CreateResponseControl(control)

	return c.Status(200).JSON(responseControl)
}
