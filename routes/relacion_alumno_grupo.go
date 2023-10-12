package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type RelacionAlumnoGrupo struct {
	ID                   uint `json:"id"`
	AlumnoRefer          int  `json:"alumno_id"`
	GruposActivosRefer   int  `json:"grupo_activo_id"`
	GruposAprobadosRefer int  `json:"grupo_aprobado_id"`
}

func CreateRelacionResponse(relacionModel models.RelacionAlumnoGrupo) RelacionAlumnoGrupo {
	return RelacionAlumnoGrupo{
		ID:                   relacionModel.ID,
		AlumnoRefer:          relacionModel.AlumnoRefer,
		GruposActivosRefer:   relacionModel.GruposActivosRefer,
		GruposAprobadosRefer: relacionModel.GruposAprobadosRefer,
	}
}

/*
func findAlumnoRelacion(id uint, alumno *models.RelacionAlumnoGrupo) error {
	database.Database.Db.Find(&alumno, "alumno_refer = ?", id)

	if alumno.AlumnoRefer == 0 {
		return errors.New("alumno does not exist")
	}

	return nil
}*/

func GetAllRelacionAlumnosGrupos(c *fiber.Ctx) error {
	relaciones := []models.RelacionAlumnoGrupo{}

	database.Database.Db.Find(&relaciones)

	responseRelaciones := []RelacionAlumnoGrupo{}

	for _, relacion := range relaciones {
		responseRelacion := CreateRelacionResponse(relacion)
		responseRelaciones = append(responseRelaciones, responseRelacion)
	}

	return c.Status(200).JSON(responseRelaciones)
}

func CreateRelacionAlumnosGrupos(c *fiber.Ctx) error {
	var relacion models.RelacionAlumnoGrupo

	if err := c.BodyParser(&relacion); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&relacion)

	relacionResponse := CreateRelacionResponse(relacion)
	return c.Status(200).JSON(relacionResponse)
}
