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

func findRelacionEspecialidad(alumnoRefer, especialidadRefer int, relacion *models.RelacionAlumnoGrupo) error {
	database.Database.Db.Where("alumno_refer = ? AND especialidad_refer = ?", alumnoRefer, especialidadRefer).First(&relacion)

	if relacion.ID == 0 {
		return errors.New("relacion does not exist")
	}

	return nil
}

func findDocuments(id int, documents *models.Documentos) error {
	database.Database.Db.Find(&documents, "alumno_refer = ?", id)

	if documents.AlumnoRefer == 0 {
		return errors.New("archivadero de documentos del alumno no existe")
	}

	return nil
}

func findAdmin(id int, admin *models.Administraciones) error {
	database.Database.Db.Find(&admin, "alumno_refer = ?", id)

	if admin.AlumnoRefer == 0 {
		return errors.New("alumno does not exist")
	}

	return nil
}

func findGrupoActivo(id int, grupo *models.GruposActivos) error {
	database.Database.Db.Find(&grupo, "id = ?", id)

	if grupo.ID == 0 {
		return errors.New("grupo does not exist")
	}

	return nil
}

func findUser(user string, usuario *models.Usuarios) error {
	database.Database.Db.Find(&usuario, "usuario = ?", user)

	if usuario.Usuario == "" {
		return errors.New("usuario no existe")
	}

	return nil
}

// func findGrupoActivoRefer(id int, grupo *models.RelacionAlumnoGrupo) error {
// 	database.Database.Db.Find(&grupo, "grupos_activos_refer = ?", id)

// 	if grupo.GruposActivosRefer == 0 {
// 		return errors.New("grupo does not exist")
// 	}

// 	return nil
// }
