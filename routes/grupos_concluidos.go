package routes

import (
	"errors"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type GruposConcluidos struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Nombre          string         `json:"nombre"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Especialidad    Especialidades `json:"especialidad"`
}

func CreateGruposConcluidosResponse(gruposConcluidosModel models.GruposConcluidos) GruposConcluidos {
	return GruposConcluidos{
		ID:              gruposConcluidosModel.ID,
		Nombre:          gruposConcluidosModel.Nombre,
		CantidadAlumnos: gruposConcluidosModel.CantidadAlumnos,
	}
}

func findGrupoConcluido(id int, grupoConcluido *models.GruposConcluidos) error {

	if id == 0 {
		return nil
	}

	database.Database.Db.Find(&grupoConcluido, "id = ?", id)

	if grupoConcluido.ID == 0 {
		return errors.New("grupo concluido does not exist")
	}

	return nil
}

func CreateGrupoConcluido(c *fiber.Ctx) error {
	var grupoConcluido models.GruposConcluidos

	if err := c.BodyParser(&grupoConcluido); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&grupoConcluido)

	responseGrupoConcluido := CreateGruposConcluidosResponse(grupoConcluido)
	return c.Status(200).JSON(responseGrupoConcluido)
}
