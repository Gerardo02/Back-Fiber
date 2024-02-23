package routes

import (
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateUser(c *fiber.Ctx) error {
	var usuario models.Usuarios
	var usersModel models.Usuarios
	var permiso models.Permisos

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findLoginUser(usuario.Usuario, &usersModel); err != nil {
		return c.Status(201).JSON(err.Error())
	}

	if usuario.Password != usersModel.Password {
		return c.Status(202).JSON("Wrong password")
	}

	if err := findPermisos(usersModel.PermisosRefer, &permiso); err != nil {
		return c.Status(404).JSON("Permit not found")
	}

	responsePermiso := CreatePermisosResponse(permiso)

	return c.Status(200).JSON(responsePermiso)
}
