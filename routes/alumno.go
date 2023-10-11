package routes

import (
	"errors"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Alumnos struct {
	ID              uint             `json:"id" gorm:"primaryKey"`
	Nombre          string           `json:"nombre"`
	Apellidos       string           `json:"apellidos"`
	Matricula       string           `json:"matricula"`
	Edad            uint             `json:"edad"`
	GrupoActivo     GruposActivos    `json:"grupos_activos"`   //arreglo
	GruposAprobados GruposConcluidos `json:"grupos_aprobados"` //arreglo
}

func CreateResponseAlumnos(alumnoModel models.Alumnos) Alumnos {
	return Alumnos{
		ID:        alumnoModel.ID,
		Nombre:    alumnoModel.Nombre,
		Apellidos: alumnoModel.Apellidos,
		Matricula: alumnoModel.Matricula,
		Edad:      alumnoModel.Edad,
		//GrupoActivo:     grupoActivo,
		//GruposAprobados: gruposAprobados,
	}
}

func findAlumno(id int, alumno *models.Alumnos) error {
	database.Database.Db.Find(&alumno, "id = ?", id)

	if alumno.ID == 0 {
		return errors.New("alumno does not exist")
	}

	return nil
}

func CreateAlumno(c *fiber.Ctx) error {
	var alumno models.Alumnos

	if err := c.BodyParser(&alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&alumno)

	responseAlumno := CreateResponseAlumnos(alumno)

	return c.Status(200).JSON(responseAlumno)
}

/*
func GetAllAlumnos(c *fiber.Ctx) error {
	alumnos := []models.Alumnos{}

	database.Database.Db.Find(&alumnos)

	responseAlumnos := []Alumnos{}

	for _, alumno := range alumnos {
		var grupoActivo models.GruposActivos
		var grupoAprobado models.GruposConcluidos


	}

	return c.Status(200).JSON(responseAlumnos)
}
*/
