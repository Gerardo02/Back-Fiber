package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetCiclosEscolares(c *fiber.Ctx) error {
	ciclos := []models.CicloEscolar{}

	database.Database.Db.Find(&ciclos)

	responseCiclos := []CicloEscolar{}

	for _, ciclo := range ciclos {

		responseCiclo := CreateGetCicloEscolarResponse(ciclo)
		responseCiclos = append(responseCiclos, responseCiclo)
	}

	return c.Status(200).JSON(responseCiclos)
}

func GetGruposConcluidos(c *fiber.Ctx) error {
	grupos := []models.GruposConcluidos{}

	database.Database.Db.Find(&grupos)

	responseGrupos := []GruposConcluidos{}

	for _, grupo := range grupos {

		var ciclo models.CicloEscolar
		var especialidad models.Especialidades

		if err := findCicloEscolar(grupo.CicloEscolarRefer, &ciclo); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		if err := findEspecialidad(grupo.EspecialidadRefer, &especialidad); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		responseEspecialidad := CreateEspecialidadResponse(especialidad)
		responseCiclo := CreateGetCicloEscolarResponse(ciclo)
		responseGrupo := CreateGruposConcluidosResponse(grupo, responseEspecialidad, responseCiclo)
		responseGrupos = append(responseGrupos, responseGrupo)
	}

	return c.Status(200).JSON(responseGrupos)
}

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
				responseGrupoConcluido := CreateAlumnosGruposConcluidosResponse(grupoConcluido, responseEspecialidad)
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
		var cicloEscolar models.CicloEscolar
		database.Database.Db.Where("grupos_activos_refer = ?", grupo.ID).Find(&relaciones)

		for _, relacion := range relaciones {
			if relacion.ListaAsistencia != "" {

				listas = append(listas, relacion.ListaAsistencia)

			}
		}

		// if err := findCicloEscolar(grupo.CicloRefer, &cicloEscolar); err != nil {
		// 	return c.Status(400).JSON(err.Error())
		// }

		database.Database.Db.First(&especialidad, grupo.EspecialidadRefer)
		database.Database.Db.First(&cicloEscolar, grupo.CicloRefer)

		responseEspecialidad := CreateEspecialidadResponse(especialidad)
		responseCicloEscolar := CreateGetCicloEscolarResponse(cicloEscolar)
		responseGrupo := CreateGruposActivosResponse(grupo, responseEspecialidad, listas, responseCicloEscolar)
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

func GetAllEspecialidades(c *fiber.Ctx) error {
	especialidades := []models.Especialidades{}

	database.Database.Db.Find(&especialidades)

	responseEspecialidades := []Especialidades{}

	for _, especialidad := range especialidades {
		responseEspecialidad := CreateEspecialidadResponse(especialidad)
		responseEspecialidades = append(responseEspecialidades, responseEspecialidad)
	}

	return c.Status(200).JSON(responseEspecialidades)
}

func GetAdministraciones(c *fiber.Ctx) error {
	administraciones := []models.Administraciones{}

	database.Database.Db.Find(&administraciones)

	responseAdministraciones := []Administraciones{}

	for _, administracion := range administraciones {
		var alumno models.Alumnos
		var adeudo, estado string

		if err := findAlumno(administracion.AlumnoRefer, &alumno); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		if Pendiente == administracion.Estado {
			estado = "Pendiente"
		} else if EnProceso == administracion.Estado {
			estado = "En proceso"
		} else if Listo == administracion.Estado {
			estado = "Listo"
		}

		if administracion.Adeudo {
			adeudo = "Al corriente"
		} else {
			adeudo = "Debe"
		}

		responseAdministracion := CreateGetAdminResponse(administracion, alumno.Nombre, alumno.Apellidos, alumno.Matricula, adeudo, estado)
		responseAdministraciones = append(responseAdministraciones, responseAdministracion)
	}

	return c.Status(200).JSON(responseAdministraciones)
}

func GetDocumentosEntregados(c *fiber.Ctx) error {
	documentos := []models.Documentos{}

	database.Database.Db.Find(&documentos)

	responseDocumentos := []Documentos{}

	for _, documento := range documentos {
		var alumno models.Alumnos

		if err := findAlumno(documento.AlumnoRefer, &alumno); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		responseDocumento := CreateGetDocumentosResponse(documento, alumno.Nombre, alumno.Apellidos, alumno.Matricula)
		responseDocumentos = append(responseDocumentos, responseDocumento)
	}

	return c.Status(200).JSON(responseDocumentos)
}

func GetUsuarios(c *fiber.Ctx) error {
	usuarios := []models.Usuarios{}

	database.Database.Db.Find(&usuarios)

	responseUsuarios := []Usuarios{}

	for _, usuario := range usuarios {
		var permiso models.Permisos

		if err := findPermisos(usuario.PermisosRefer, &permiso); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responsePermiso := CreatePermisosResponse(permiso)
		responseUsuario := CreateUsuariosResponse(usuario, responsePermiso)
		responseUsuarios = append(responseUsuarios, responseUsuario)
	}

	return c.Status(200).JSON(responseUsuarios)
}

func GetPermisos(c *fiber.Ctx) error {
	permisos := []models.Permisos{}

	database.Database.Db.Find(&permisos)

	responsePermisos := []Permisos{}

	for _, permiso := range permisos {

		responsePermiso := CreatePermisosResponse(permiso)
		responsePermisos = append(responsePermisos, responsePermiso)
	}

	return c.Status(200).JSON(responsePermisos)
}

func GetAlumnosNombres(c *fiber.Ctx) error {
	alumnos := []models.Alumnos{}

	database.Database.Db.Find(&alumnos)

	responseAlumnos := []AlumnosNombres{}

	for _, alumno := range alumnos {
		relaciones := []models.RelacionAlumnoGrupo{}
		database.Database.Db.Where("alumno_refer = ?", alumno.ID).Find(&relaciones)

		responseEspecialidades := []Especialidades{}

		for _, relacion := range relaciones {
			if relacion.EspecialidadRefer != 0 {
				var especialidad models.Especialidades

				database.Database.Db.First(&especialidad, relacion.EspecialidadRefer)

				responseEspecialidad := CreateEspecialidadResponse(especialidad)
				responseEspecialidades = append(responseEspecialidades, responseEspecialidad)
			}

		}

		responseAlumno := CreateAlumnoNombreResponse(alumno, responseEspecialidades)
		responseAlumnos = append(responseAlumnos, responseAlumno)
	}

	return c.Status(200).JSON(responseAlumnos)
}

func GetHistorialAdimn(c *fiber.Ctx) error {
	historiales := []models.HistorialLogs{}

	database.Database.Db.Find(&historiales)

	responseHistoriales := []HistorialLogs{}

	for _, historial := range historiales {
		var alumno models.Alumnos

		if err := findAlumno(historial.AlumnoRefer, &alumno); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		responseAlumno := CreateAlumnoNombreNombreResponse(alumno)
		responseHistorial := CreateHistorialResponse(historial, responseAlumno)
		responseHistoriales = append(responseHistoriales, responseHistorial)
	}

	return c.Status(200).JSON(responseHistoriales)
}
