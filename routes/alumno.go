package routes

import (
	"errors"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Alumnos struct {
	ID              uint               `json:"id" gorm:"primaryKey"`
	Nombre          string             `json:"nombre"`
	Apellidos       string             `json:"apellidos"`
	Matricula       string             `json:"matricula"`
	Edad            uint               `json:"edad"`
	GrupoActivo     []GruposActivos    `json:"grupos_activos"`   //arreglo
	GruposAprobados []GruposConcluidos `json:"grupos_aprobados"` //arreglo
}

func CreateGetResponseAlumnos(alumnoModel models.Alumnos, grupoActivo []GruposActivos, gruposAprobados []GruposConcluidos) Alumnos {
	return Alumnos{
		ID:              alumnoModel.ID,
		Nombre:          alumnoModel.Nombre,
		Apellidos:       alumnoModel.Apellidos,
		Matricula:       alumnoModel.Matricula,
		Edad:            alumnoModel.Edad,
		GrupoActivo:     grupoActivo,
		GruposAprobados: gruposAprobados,
	}
}

func CreateAlumnosResponse(alumnoModel models.Alumnos) Alumnos {
	return Alumnos{
		ID:        alumnoModel.ID,
		Nombre:    alumnoModel.Nombre,
		Apellidos: alumnoModel.Apellidos,
		Matricula: alumnoModel.Matricula,
		Edad:      alumnoModel.Edad,
	}
}

func findAlumno(id int, alumno *models.Alumnos) error {
	database.Database.Db.Find(&alumno, "id = ?", id)

	if alumno.ID == 0 {
		return errors.New("alumno does not exist")
	}

	return nil
}

// func GetAllAlumnos(c *fiber.Ctx) error {
// 	alumnos := []models.Alumnos{}
// 	relaciones := []models.RelacionAlumnoGrupo{}

// 	var grupoActivo models.GruposActivos
// 	var grupoAprobado models.GruposConcluidos

// 	//var relacion models.RelacionAlumnoGrupo

// 	responseAlumnos := []Alumnos{}
// 	responseGruposActivos := []GruposActivos{}
// 	responseGruposAprobados := []GruposConcluidos{}

// 	database.Database.Db.Find(&alumnos)

// 	for _, alumno := range alumnos {
// 		/*
// 			if err := findAlumnoRelacion(alumno.ID, &relacion); err != nil {
// 				return c.Status(400).JSON(err.Error())
// 			}*/
// 		database.Database.Db.Find(&relaciones, models.RelacionAlumnoGrupo{AlumnoRefer: 1})
// 		for _, relacion := range relaciones {

// 			/*
// 				if err := findGrupoActivo(relacion.GruposActivosRefer, &grupoActivo); err != nil {
// 					return c.Status(400).JSON(err.Error())
// 				}*/

// 			database.Database.Db.Find(&grupoActivo, "id = ?", relacion.GruposActivosRefer)
// 			database.Database.Db.Find(&grupoAprobado, "id = ?", relacion.GruposAprobadosRefer)

// 			//log.Print(relacion.GruposActivosRefer)

// 			responseGrupoActivo := CreateGruposActivosResponse(grupoActivo)
// 			responseGruposActivos = append(responseGruposActivos, responseGrupoActivo)

// 			responseGrupoAprobado := CreateGruposConcluidosResponse(grupoAprobado)
// 			responseGruposAprobados = append(responseGruposAprobados, responseGrupoAprobado)
// 			//log.Print(responseGruposActivos)
// 		}
// 		// database.Database.Db.Find(&gruposActivos)
// 		// database.Database.Db.Find(&gruposAprobados)
// 		responseAlumno := CreateGetResponseAlumnos(alumno, responseGruposActivos, responseGruposAprobados)
// 		responseAlumnos = append(responseAlumnos, responseAlumno)

// 	}
// 	return c.Status(200).JSON(responseAlumnos)

// }

func GetAllAlumnos(c *fiber.Ctx) error {
	alumnos := []models.Alumnos{}
	database.Database.Db.Find(&alumnos)

	responseAlumnos := []Alumnos{}

	for _, alumno := range alumnos {
		var relaciones []models.RelacionAlumnoGrupo
		database.Database.Db.Where("alumno_refer = ?", alumno.ID).Find(&relaciones)

		responseGruposActivos := []GruposActivos{}
		responseGruposConcluidos := []GruposConcluidos{}

		for _, relacion := range relaciones {
			if relacion.GruposActivosRefer != 0 {
				var grupoActivo models.GruposActivos
				database.Database.Db.First(&grupoActivo, relacion.GruposActivosRefer)
				responseGrupoActivo := CreateGruposActivosResponse(grupoActivo)
				responseGruposActivos = append(responseGruposActivos, responseGrupoActivo)
			}

			if relacion.GruposAprobadosRefer != 0 {
				var grupoConcluido models.GruposConcluidos
				database.Database.Db.First(&grupoConcluido, relacion.GruposAprobadosRefer)
				responseGrupoConcluido := CreateGruposConcluidosResponse(grupoConcluido)
				responseGruposConcluidos = append(responseGruposConcluidos, responseGrupoConcluido)
			}
		}

		responseAlumno := CreateGetResponseAlumnos(alumno, responseGruposActivos, responseGruposConcluidos)
		responseAlumnos = append(responseAlumnos, responseAlumno)
	}

	return c.Status(200).JSON(responseAlumnos)
}

func CreateAlumno(c *fiber.Ctx) error {
	var alumno models.Alumnos

	if err := c.BodyParser(&alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&alumno)

	responseAlumno := CreateAlumnosResponse(alumno)

	return c.Status(200).JSON(responseAlumno)
}
