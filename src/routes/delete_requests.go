package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteAlumno(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var alumno models.Alumnos
	var admin models.Administraciones
	var docs models.Documentos

	if err != nil {
		return c.SendString("id is not a number")
	}

	if err := findAlumno(id, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findAdmin(id, &admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findDocuments(id, &docs); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&alumno).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&admin).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&docs).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("succesfully deleted alumno")
}

func DeleteSingleGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var grupoActivo models.GruposActivos
	gruposRefer := []models.RelacionAlumnoGrupo{}
	gruposReferListas := []models.RelacionGrupoLista{}

	if err != nil {
		return c.SendString("id is not a number")
	}

	if err := findGrupoActivo(id, &grupoActivo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Where("grupos_activos_refer = ?", id).Find(&gruposRefer)
	database.Database.Db.Where("grupos_activos_refer = ?", id).Find(&gruposReferListas)

	for _, grupoRefer := range gruposRefer {
		if grupoRefer.GruposActivosRefer == 0 {
			break
		}
		grupoRefer.GruposActivosRefer = 0
		database.Database.Db.Save(&grupoRefer)
	}

	for _, grupoReferLista := range gruposReferListas {
		if grupoReferLista.GruposActivosRefer == 0 {
			break
		}
		grupoReferLista.GruposActivosRefer = 0
		database.Database.Db.Save(&grupoReferLista)
	}

	if err := database.Database.Db.Delete(&grupoActivo).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("succesfully deleted grupo")
}

func DeleteAllGruposActivos(c *fiber.Ctx) error {
	gruposRefer := []models.RelacionAlumnoGrupo{}
	gruposReferListas := []models.RelacionGrupoLista{}

	database.Database.Db.Where("grupos_activos_refer != 0").Find(&gruposRefer)
	database.Database.Db.Where("grupos_activos_refer != 0").Find(&gruposReferListas)

	for _, grupoRefer := range gruposRefer {
		if grupoRefer.GruposActivosRefer == 0 {
			break
		}
		grupoRefer.GruposActivosRefer = 0
		database.Database.Db.Save(&grupoRefer)
	}

	for _, grupoReferLista := range gruposReferListas {
		if grupoReferLista.GruposActivosRefer == 0 {
			break
		}
		grupoReferLista.GruposActivosRefer = 0
		database.Database.Db.Save(&grupoReferLista)
	}

	database.Database.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.GruposActivos{})
	return c.Status(200).SendString("succesfully softly deleted all grupos")
}

func DropSoftDeletesGruposActivos(c *fiber.Ctx) error {
	database.Database.Db.Exec("DELETE FROM grupos_activos WHERE deleted_at IS NOT NULL")
	return c.Status(200).SendString("succesfully deleted all grupos fr fr no way back now")
}

func DeleteUser(c *fiber.Ctx) error {
	var usuario models.Usuarios
	var usersModel models.Usuarios

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findLoginUser(usuario.Usuario, &usersModel); err != nil {
		return c.Status(200).JSON(err.Error())
	}

	if usuario.Password != usersModel.Password {
		return c.Status(200).JSON("Wrong password")
	}

	if err := database.Database.Db.Delete(&usersModel).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Succesfully deleted user")
}
