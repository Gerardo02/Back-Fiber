package routes

import (
	"errors"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Alumno struct {
	ID                  uint   `json:"id" gorm:"primaryKey"`
	ControlEscolarRefer int    `json:"control_escolar_id"`
	Nombre              string `json:"nombre"`
	Matricula           string `json:"matricula"`
}

func CreateResponseAlumno(alumnoModel models.Alumno) Alumno {
	return Alumno{
		ID:                  alumnoModel.ID,
		ControlEscolarRefer: alumnoModel.ControlEscolarRefer,
		Nombre:              alumnoModel.Nombre,
		Matricula:           alumnoModel.Matricula,
	}
}

func findAlumno(id int, alumno *models.Alumno) error {
	database.Database.Db.Find(&alumno, "id = ?", id)

	if alumno.ID == 0 {
		return errors.New("Alumno does not exist")
	}

	return nil
}

func CreateAlumno(c *fiber.Ctx) error {
	var alumno models.Alumno

	if err := c.BodyParser(&alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&alumno)
	responseAlumno := CreateResponseAlumno(alumno)

	return c.Status(200).JSON(responseAlumno)
}

func GetAlumnos(c *fiber.Ctx) error {
	alumnos := []models.Alumno{}

	database.Database.Db.Find(&alumnos)
	responseAlumnos := []Alumno{}

	for _, alumno := range alumnos {
		responseAlumno := CreateResponseAlumno(alumno)
		responseAlumnos = append(responseAlumnos, responseAlumno)
	}

	return c.Status(200).JSON(responseAlumnos)

}

func GetAlumno(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var alumno models.Alumno

	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}

	if err := findAlumno(id, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseAlumno := CreateResponseAlumno(alumno)
	return c.Status(200).JSON(responseAlumno)

}

func UpdateAlumno(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var alumno models.Alumno

	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}

	if err := findAlumno(id, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedAlumno struct {
		Nombre    string `json:"nombre"`
		Matricula string `json:"matricula"`
	}

	var updateData UpdatedAlumno

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	alumno.Nombre = updateData.Nombre
	alumno.Matricula = updateData.Matricula

	database.Database.Db.Save(&alumno)

	responseAlumno := CreateResponseAlumno(alumno)

	return c.Status(200).JSON(responseAlumno)

}

func DeleteAlumno(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var alumno models.Alumno

	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}

	if err := findAlumno(id, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&alumno).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Succesfully deleted Alumno")
}
