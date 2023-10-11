package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Administraciones struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Adeudo          bool   `json:"adeudo"`
	Estado          int    `json:"estado"`
	AlumnoNombre    string `json:"alumno_nombre"`
	AlumnoMatricula string `json:"alumno_matricula"`
}

func CreateResponse(adminModel models.Administraciones, alumnoNombre string, alumnoMatricula string) Administraciones {
	return Administraciones{
		ID:              adminModel.ID,
		Adeudo:          adminModel.Adeudo,
		Estado:          adminModel.Estado,
		AlumnoNombre:    alumnoNombre,
		AlumnoMatricula: alumnoMatricula,
	}
}

func CreateCuentaAdmin(c *fiber.Ctx) error {
	var admin models.Administraciones

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&admin)

	var alumno models.Alumnos

	if err := findAlumno(admin.AlumnoRefer, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseAdmin := CreateResponse(admin, alumno.Nombre, alumno.Matricula)

	return c.Status(200).JSON(responseAdmin)

}
