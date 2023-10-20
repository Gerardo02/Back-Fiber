package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllAlumnos(c *fiber.Ctx) error {
	alumnos := []models.Alumnos{}
	database.Database.Db.Find(&alumnos)

	responseAlumnos := []Alumnos{}

	for _, alumno := range alumnos {
		relaciones := []models.RelacionAlumnoGrupo{}
		database.Database.Db.Where("alumno_refer = ?", alumno.ID).Find(&relaciones)

		responseGruposActivos := []GruposActivos{}
		responseGruposConcluidos := []GruposConcluidos{}
		responseEspecialidades := []Especialidades{}

		for _, relacion := range relaciones {
			if relacion.GruposActivosRefer != 0 {
				var grupoActivo models.GruposActivos
				var especialidad models.Especialidades

				database.Database.Db.First(&grupoActivo, relacion.GruposActivosRefer)
				database.Database.Db.First(&especialidad, grupoActivo.EspecialidadRefer)

				responseEspecialidad := CreateEspecialidadResponse(especialidad)
				responseGrupoActivo := CreateGruposActivosAlumnosResponse(grupoActivo, responseEspecialidad)
				responseGruposActivos = append(responseGruposActivos, responseGrupoActivo)
			}

			if relacion.GruposAprobadosRefer != 0 {
				var grupoConcluido models.GruposConcluidos
				var especialidad models.Especialidades

				database.Database.Db.First(&grupoConcluido, relacion.GruposAprobadosRefer)
				database.Database.Db.First(&especialidad, grupoConcluido.EspecialidadRefer)

				responseEspecialidad := CreateEspecialidadResponse(especialidad)
				responseGrupoConcluido := CreateGruposConcluidosResponse(grupoConcluido, responseEspecialidad)
				responseGruposConcluidos = append(responseGruposConcluidos, responseGrupoConcluido)
			}

			if relacion.EspecialidadRefer != 0 {
				var especialidad models.Especialidades

				database.Database.Db.First(&especialidad, relacion.EspecialidadRefer)

				responseEspecialidad := CreateEspecialidadResponse(especialidad)
				responseEspecialidades = append(responseEspecialidades, responseEspecialidad)
			}
		}

		responseAlumno := CreateGetAllAlumnosResponse(alumno, responseGruposActivos, responseGruposConcluidos, responseEspecialidades)
		responseAlumnos = append(responseAlumnos, responseAlumno)
	}

	return c.Status(200).JSON(responseAlumnos)
}

func GetAllGruposActivos(c *fiber.Ctx) error {
	grupos := []models.GruposActivos{}
	database.Database.Db.Find(&grupos)

	responseGrupos := []GruposActivos{}

	for _, grupo := range grupos {
		relaciones := []models.RelacionGrupoLista{}
		listas := []string{}

		var especialidad models.Especialidades
		database.Database.Db.Where("grupos_activos_refer = ?", grupo.ID).Find(&relaciones)

		for _, relacion := range relaciones {
			if relacion.ListaAsistencia != "" {

				listas = append(listas, relacion.ListaAsistencia)

			}
		}

		database.Database.Db.First(&especialidad, grupo.EspecialidadRefer)

		responseEspecialidad := CreateEspecialidadResponse(especialidad)
		responseGrupo := CreateGruposActivosResponse(grupo, responseEspecialidad, listas)
		responseGrupos = append(responseGrupos, responseGrupo)

	}
	return c.Status(200).JSON(responseGrupos)
}

func GetAllRelacionAlumnosGrupos(c *fiber.Ctx) error {
	relaciones := []models.RelacionAlumnoGrupo{}

	database.Database.Db.Find(&relaciones)

	responseRelaciones := []RelacionAlumnoGrupo{}

	for _, relacion := range relaciones {
		responseRelacion := CreateRelacionAlumnoGrupoResponse(relacion)
		responseRelaciones = append(responseRelaciones, responseRelacion)
	}

	return c.Status(200).JSON(responseRelaciones)
}
