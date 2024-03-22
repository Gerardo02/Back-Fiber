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

func CreateCicloEscolar(c *fiber.Ctx) error {
	var ciclo models.CicloEscolar

	if err := c.BodyParser(&ciclo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&ciclo)

	return c.Status(200).JSON("Ciclo escolar created succesfully")
}

func CreateHorario(c *fiber.Ctx) error {
	var horario models.Horarios
	var grupo models.GruposActivos

	if err := c.BodyParser(&horario); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Retrieve the last grupo
	result := database.Database.Db.Order("id DESC").First(&grupo)
	if result.Error != nil && result.RowsAffected == 0 {
		// If there are no rows, set grupo.ID to the last created ID
		lastID := 0
		database.Database.Db.Raw("SELECT last_value FROM grupos_activos").Row().Scan(&lastID)
		grupo.ID = lastID
	} else if result.Error != nil {
		// Handle error if any
		return c.Status(500).JSON(result.Error.Error())
	}

	horario.GrupoRefer = grupo.ID

	database.Database.Db.Create(&horario)

	return c.Status(200).JSON("Horario created succesfully")
}

func CreateRelacionAlumnoEspecialidad(c *fiber.Ctx) error {

	var relacion models.RelacionAlumnoGrupo

	if err := c.BodyParser(&relacion); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&relacion)

	return c.Status(200).JSON("Relacion created succesfully")
}

func CreateAlumno(c *fiber.Ctx) error {
	var alumnoArray Alumnos
	var alumno models.Alumnos
	var docs models.Documentos
	var admin models.Administraciones

	if err := c.BodyParser(&alumnoArray); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	//alumno.ID = alumnoArray.ID
	alumno.Nombre = alumnoArray.Nombre
	alumno.Apellidos = alumnoArray.Apellidos
	alumno.Matricula = alumnoArray.Matricula
	alumno.FechaNacimiento = alumnoArray.FechaNacimiento
	alumno.Edad = alumnoArray.Edad
	alumno.NombreTutor = alumnoArray.NombreTutor
	alumno.CelularTutor = alumnoArray.CelularTutor
	alumno.Curp = alumnoArray.Curp
	alumno.Localidad = alumnoArray.Localidad
	alumno.CodigoPostal = alumnoArray.CodigoPostal
	alumno.Direccion = alumnoArray.Direccion
	alumno.TelefonoFijo = alumnoArray.TelefonoFijo
	alumno.Celular = alumnoArray.Celular
	alumno.Correo = alumnoArray.Correo

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

	type ResponseJson struct {
		Nombre            string `json:"nombre"`
		CantidadAlumnos   int    `json:"cantidad_de_alumnos"`
		EspecialidadRefer int    `json:"especialidad_id"`
		CicloEscolarRefer int    `json:"ciclo_escolar_id"`
		TemporaryId       int    `json:"id_grupo_place_holder"`
	}

	var grabData ResponseJson
	var grupoConcluido models.GruposConcluidos
	var relacion models.RelacionAlumnoGrupo
	var relacionModel models.RelacionAlumnoGrupo

	if err := c.BodyParser(&grabData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	grupoConcluido.CantidadAlumnos = grabData.CantidadAlumnos
	grupoConcluido.Nombre = grabData.Nombre
	grupoConcluido.EspecialidadRefer = grabData.EspecialidadRefer
	grupoConcluido.CicloEscolarRefer = grabData.CicloEscolarRefer

	database.Database.Db.Create(&grupoConcluido)

	if err := findRelacionGrupo(grabData.TemporaryId, &relacion); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	relacionModel.AlumnoRefer = relacion.AlumnoRefer
	relacionModel.Estado = 1
	relacionModel.GruposAprobadosRefer = grupoConcluido.ID

	database.Database.Db.Create(&relacionModel)

	return c.Status(200).JSON("Grupo concluido creado successfully")
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

	return c.Status(200).JSON("relacion created succesfully")
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

func CreateAdminHistorial(c *fiber.Ctx) error {
	var historial models.HistorialLogs

	if err := c.BodyParser(&historial); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&historial)

	return c.Status(200).JSON("Historial creado")

}
