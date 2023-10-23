package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

/*
func CreateCuentaAdmin(c *fiber.Ctx) error {
	var admin models.Administraciones

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&admin)

	var alumno models.Alumnos

	if err := findAlumno(admin.AlumnoRefer, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseAdmin := CreateResponseAdmin(admin, alumno.Nombre, alumno.Matricula)

	return c.Status(200).JSON(responseAdmin)

}*/

func CreateAlumno(c *fiber.Ctx) error {
	var alumno models.Alumnos
	var docs models.Documentos
	var admin models.Administraciones

	if err := c.BodyParser(&alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&alumno)

	docs.AlumnoRefer = alumno.ID
	admin.AlumnoRefer = alumno.ID

	database.Database.Db.Create(&admin)
	database.Database.Db.Create(&docs)

	responseAlumno := CreateAlumnosResponse(alumno)

	return c.Status(200).JSON(responseAlumno)
}

func CreateEspecialidad(c *fiber.Ctx) error {
	var especialidad models.Especialidades

	if err := c.BodyParser(&especialidad); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&especialidad)

	responseEspecialidad := CreateEspecialidadResponse(especialidad)

	return c.Status(200).JSON(responseEspecialidad)
}

func CreateGrupoActivo(c *fiber.Ctx) error {
	var grupoActivo models.GruposActivos

	if err := c.BodyParser(&grupoActivo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&grupoActivo)

	var especialidad models.Especialidades

	database.Database.Db.Find(&especialidad)

	responseEspecialidad := CreateEspecialidadResponse(especialidad)
	responseGrupoActivo := CreateGruposActivosAlumnosResponse(grupoActivo, responseEspecialidad)
	return c.Status(200).JSON(responseGrupoActivo)
}

func CreateGrupoConcluido(c *fiber.Ctx) error {
	var grupoConcluido models.GruposConcluidos

	if err := c.BodyParser(&grupoConcluido); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&grupoConcluido)

	var especialidad models.Especialidades

	database.Database.Db.Find(&especialidad)

	responseEspecialidad := CreateEspecialidadResponse(especialidad)
	responseGrupoConcluido := CreateGruposConcluidosResponse(grupoConcluido, responseEspecialidad)
	return c.Status(200).JSON(responseGrupoConcluido)
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

func CreateRelacionAlumnosGrupos(c *fiber.Ctx) error {
	var relacion models.RelacionAlumnoGrupo

	if err := c.BodyParser(&relacion); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&relacion)

	relacionResponse := CreateRelacionAlumnoGrupoResponse(relacion)
	return c.Status(200).JSON(relacionResponse)
}

func CreateRelacionGrupoListas(c *fiber.Ctx) error {
	var relacion models.RelacionGrupoLista

	if err := c.BodyParser(&relacion); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&relacion)

	relacionResponse := CreateRelacionGrupoListasResponse(relacion)
	return c.Status(200).JSON(relacionResponse)
}

func CreateEspecialidadRelacionMiddleware(c *fiber.Ctx) error {
	var alumno models.Alumnos

	DummyAlumno := struct {
		Especialidad []string `json:"especialidad"`
	}{}

	if err := c.BodyParser(&DummyAlumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	length := len(DummyAlumno.Especialidad)

	if length <= 0 {
		return c.Next()
	}

	database.Database.Db.Last(&alumno)

	alumnoID := alumno.ID + 1

	for i := 0; i < length; i++ {

		var especialidades []models.Especialidades
		database.Database.Db.Where("materia = ?", DummyAlumno.Especialidad[i]).Find(&especialidades)

		for _, especialidad := range especialidades {
			relacion := models.RelacionAlumnoGrupo{
				AlumnoRefer:          alumnoID,
				GruposActivosRefer:   0,
				GruposAprobadosRefer: 0,
				EspecialidadRefer:    especialidad.ID,
			}

			database.Database.Db.Create(&relacion)
		}
	}

	return c.Next()
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
