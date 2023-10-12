package routes

import (
	"errors"
	"log"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type GruposActivos struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Nombre          string         `json:"nombre"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Trimestre       int            `json:"trimestre"`
	ListaAsistencia string         `json:"lista_asistencia"`
	Especialidad    Especialidades `json:"especialidad"`
}

func CreateGruposActivosResponse(gruposActivosModel models.GruposActivos) GruposActivos {
	return GruposActivos{
		ID:              gruposActivosModel.ID,
		Nombre:          gruposActivosModel.Nombre,
		CantidadAlumnos: gruposActivosModel.CantidadAlumnos,
		Trimestre:       gruposActivosModel.Trimestre,
		ListaAsistencia: gruposActivosModel.ListaAsistencia,
	}
}

func findGrupoActivo(id int, grupoActivo *models.GruposActivos) error {

	log.Print(id)
	if id == 0 {
		return nil
	}

	database.Database.Db.Find(&grupoActivo, "id = ?", id)

	if grupoActivo.ID == 0 {
		return errors.New("grupo activo does not exist")
	}

	return nil
}

func CreateGrupoActivo(c *fiber.Ctx) error {
	var grupoActivo models.GruposActivos

	if err := c.BodyParser(&grupoActivo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&grupoActivo)

	responseGrupoActivo := CreateGruposActivosResponse(grupoActivo)
	return c.Status(200).JSON(responseGrupoActivo)
}
