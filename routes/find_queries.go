package routes

import (
	"errors"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
)

func findPermisos(id int, permisos *models.Permisos) error {
	database.Database.Db.Find(&permisos, "id = ?", id)

	if permisos.ID == 0 {
		return errors.New("id no es un int")
	}

	return nil
}

func findAlumno(id int, alumno *models.Alumnos) error {
	database.Database.Db.Find(&alumno, "id = ?", id)

	if alumno.ID == 0 {
		return errors.New("alumno does not exist")
	}

	return nil
}
