package routes

import (
	"errors"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Permisos struct {
	Permiso string `json:"permiso"`
}

func CreatePermisosResponse(permisosModel models.Permisos) Permisos {
	return Permisos{Permiso: permisosModel.Permiso}
}

func findPermisos(id int, permisos *models.Permisos) error {
	database.Database.Db.Find(&permisos, "id = ?", id)

	if permisos.ID == 0 {
		return errors.New("id no es un int")
	}

	return nil
}

func CreatePermiso(c *fiber.Ctx) error {
	var permiso models.Permisos

	if err := c.BodyParser(&permiso); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&permiso)
	responsePermiso := CreatePermisosResponse(permiso)

	return c.Status(200).JSON(responsePermiso)
}
