package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Usuarios struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Usuario  string   `json:"usuario"`
	Password string   `json:"password"`
	Permisos Permisos `json:"permisos"`
}

func CreateUsuariosResponse(usuariosModel models.Usuarios, permiso Permisos) Usuarios {
	return Usuarios{
		ID:       usuariosModel.ID,
		Usuario:  usuariosModel.Usuario,
		Password: usuariosModel.Password,
		Permisos: permiso,
	}
}

func CreateUsuarios(c *fiber.Ctx) error {
	var usuario models.Usuarios

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&usuario)

	var permisos models.Permisos

	if err := findPermisos(usuario.PermisosRefer, &permisos); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responsePermisos := CreatePermisosResponse(permisos)
	responseUsuario := CreateUsuariosResponse(usuario, responsePermisos)
	return c.Status(200).JSON(responseUsuario)
}
