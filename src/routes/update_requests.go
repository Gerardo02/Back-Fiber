package routes

import (
	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateDocuments(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var documents models.Documentos

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findDocuments(id, &documents); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateDocument struct {
		ActaNacimiento       bool `json:"acta_de_nacimiento"`
		Curp                 bool `json:"curp"`
		ComprobanteDomicilio bool `json:"comprobante_de_domicilio"`
		MayorQuince          bool `json:"mayor_quince"`
		Fotos                bool `json:"fotos"`
	}

	var UpdateData UpdateDocument

	if err := c.BodyParser(&UpdateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	documents.ActaNacimiento = UpdateData.ActaNacimiento
	documents.Curp = UpdateData.Curp
	documents.ComprobanteDomicilio = UpdateData.ComprobanteDomicilio
	documents.MayorQuince = UpdateData.MayorQuince
	documents.Fotos = UpdateData.Fotos

	database.Database.Db.Save(&documents)

	responseDocuments := CreateDocumentosResponse(documents)

	return c.Status(200).JSON(responseDocuments)
}

func UpdateAlumnos(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var alumnos models.Alumnos

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findAlumno(id, &alumnos); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedAlumno struct {
		Nombre          string `json:"nombre"`
		Apellidos       string `json:"apellidos"`
		Matricula       string `json:"matricula"`
		FechaNacimiento string `json:"fecha_nacimiento"`
		Edad            int    `json:"edad"`
		NombreTutor     string `json:"nombre_tutor"`
		CelularTutor    string `json:"celular_tutor"`
		Curp            string `json:"curp"`
		Localidad       string `json:"localidad"`
		CodigoPostal    string `json:"codigo_postal"`
		Direccion       string `json:"direccion"`
		TelefonoFijo    string `json:"telefono_fijo"`
		Celular         string `json:"celular"`
		Correo          string `json:"correo"`
	}

	var UpdatedData UpdatedAlumno

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	alumnos.Nombre = UpdatedData.Nombre
	alumnos.Apellidos = UpdatedData.Apellidos
	alumnos.Matricula = UpdatedData.Matricula
	alumnos.FechaNacimiento = UpdatedData.FechaNacimiento
	alumnos.Edad = UpdatedData.Edad
	alumnos.NombreTutor = UpdatedData.NombreTutor
	alumnos.CelularTutor = UpdatedData.CelularTutor
	alumnos.Curp = UpdatedData.Curp
	alumnos.Localidad = UpdatedData.Localidad
	alumnos.CodigoPostal = UpdatedData.CodigoPostal
	alumnos.Direccion = UpdatedData.Direccion
	alumnos.TelefonoFijo = UpdatedData.TelefonoFijo
	alumnos.Celular = UpdatedData.Celular
	alumnos.Correo = UpdatedData.Correo

	database.Database.Db.Save(&alumnos)

	responseAlumnos := CreateAlumnosResponse(alumnos)

	return c.Status(200).JSON(responseAlumnos)
}

func UpdateAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var adeudo, estado string
	var admin models.Administraciones

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findAdmin(id, &admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedAdmin struct {
		Adeudo bool `json:"adeudo"`
		Estado int  `json:"estado"`
	}

	var UpdatedData UpdatedAdmin

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	admin.Adeudo = UpdatedData.Adeudo
	admin.Estado = UpdatedData.Estado

	if Pendiente == admin.Estado {
		estado = "Pendiente"
	} else if EnProceso == admin.Estado {
		estado = "En proceso"
	} else if Listo == admin.Estado {
		estado = "Listo"
	}

	if admin.Adeudo {
		adeudo = "Debe"
	} else {
		adeudo = "Al corriente"
	}

	database.Database.Db.Save(&admin)

	responseAdmin := CreateAdminResponse(admin, adeudo, estado)

	return c.Status(200).JSON(responseAdmin)
}

func UpdateUserName(c *fiber.Ctx) error {
	user := c.Params("user")
	var userModel models.Usuarios
	var usuario models.Usuarios

	if !validString(user) {
		return c.SendString("Invalid user parameter")
	}

	if err := findUser(user, &userModel); err != nil {
		return c.Status(200).JSON(err.Error())
	}

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if usuario.Password != userModel.Password {
		return c.Status(200).JSON("Wrong password")
	}

	type UpdatedUserName struct {
		Usuario  string `json:"usuario"`
		Password string `json:"password"`
	}

	var UpdatedData UpdatedUserName

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userModel.Usuario = UpdatedData.Usuario
	userModel.Password = UpdatedData.Password

	database.Database.Db.Save(&userModel)

	return c.Status(200).JSON("Usuario actualizado succesfully")
}

func UpdateUserPassWord(c *fiber.Ctx) error {

	var userModel models.Usuarios
	var usuario models.Usuarios

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findUser(usuario.Usuario, &userModel); err != nil {
		return c.Status(200).JSON(err.Error())
	}

	type UpdatedUserPassWord struct {
		Usuario  string `json:"usuario"`
		Password string `json:"password"`
	}

	var UpdatedData UpdatedUserPassWord

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userModel.Usuario = UpdatedData.Usuario
	userModel.Password = UpdatedData.Password

	database.Database.Db.Save(&userModel)

	return c.Status(200).JSON("Password actualizado succesfully")
}

func UpdatePermisoUser(c *fiber.Ctx) error {

	var userModel models.Usuarios
	var usuario models.Usuarios

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findUser(usuario.Usuario, &userModel); err != nil {
		return c.Status(200).JSON(err.Error())
	}

	type UpdatedUserPermiso struct {
		Usuario string `json:"usuario"`
		Permiso int    `json:"permisos_id"`
	}

	var UpdatedData UpdatedUserPermiso

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userModel.Usuario = UpdatedData.Usuario
	userModel.PermisosRefer = UpdatedData.Permiso

	database.Database.Db.Save(&userModel)

	return c.Status(200).JSON("Permiso updated succesfully")
}

func UpdateRelacionAlumnoEspecialidad(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	var relacionAumno models.RelacionAlumnoGrupo

	if err != nil {
		return c.SendString("id is not an int")
	}

	type UpdatedRelacionAlumnoGrupo struct {
		EspecialidadesRefer int `json:"especialidad_id"`
	}

	var UpdatedData UpdatedRelacionAlumnoGrupo

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findRelacionEspecialidad(id, UpdatedData.EspecialidadesRefer, &relacionAumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	relacionAumno.EspecialidadRefer = 0

	database.Database.Db.Save(&relacionAumno)

	result := database.Database.Db.Exec("DELETE FROM relacion_alumno_grupos WHERE grupos_activos_refer = 0 AND grupos_aprobados_refer = 0 AND especialidad_refer = 0")

	if result.Error != nil {
		// Handle the error
		// Example: log the error, return an error response, etc.
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Check the number of rows affected
	if result.RowsAffected == 0 {
		// Rows were not deleted
		return c.Status(404).JSON(fiber.Map{"message": "No matching rows found"})
	}

	return c.Status(200).JSON("Relacion updated succesfully")
}

func UpdateGrupoActivo(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	var grupo models.GruposActivos

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findGrupoActivo(id, &grupo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedGrupoActivo struct {
		Nombre            string `json:"nombre"`
		NombreMaestro     string `json:"nombre_maestro"`
		Dia               string `json:"dia"`
		Entrada           string `json:"entrada"`
		Salida            string `json:"salida"`
		CantidadAlumnos   int    `json:"cantidad_de_alumnos"`
		Trimestre         int    `json:"trimestre"`
		EspecialidadRefer int    `json:"especialidad_id"`
	}

	var UpdatedData UpdatedGrupoActivo

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	grupo.Nombre = UpdatedData.Nombre
	grupo.NombreMaestro = UpdatedData.NombreMaestro
	grupo.Dia = UpdatedData.Dia
	grupo.Entrada = UpdatedData.Entrada
	grupo.Salida = UpdatedData.Salida
	grupo.CantidadAlumnos = UpdatedData.CantidadAlumnos
	grupo.Trimestre = UpdatedData.Trimestre
	grupo.EspecialidadRefer = UpdatedData.EspecialidadRefer

	database.Database.Db.Save(&grupo)

	return c.Status(200).JSON("Grupo updated succesfully")
}

func UpdateRelacionAlumnoGrupo(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	var relacion models.RelacionAlumnoGrupo
	var grupo models.GruposActivos

	if err != nil {
		return c.SendString("id is not an int")
	}

	type UpdatedRelacionAlumnoGrupo struct {
		AlumnoRefer int `json:"alumno_id"`
	}

	var UpdatedData UpdatedRelacionAlumnoGrupo

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findRelacionAlumnoGrupo(id, UpdatedData.AlumnoRefer, &relacion); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findGrupoActivo(id, &grupo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	relacion.GruposActivosRefer = 0
	grupo.CantidadAlumnos -= 1

	database.Database.Db.Save(&relacion)
	database.Database.Db.Save(&grupo)

	return c.Status(200).JSON("Grupo updated succesfully")
}

func UpdateCicloEscolar(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	var ciclo models.CicloEscolar

	if err != nil {
		return c.SendString("id is not an int")
	}

	if err := findCicloEscolar(id, &ciclo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedCicloEscolar struct {
		Nombre    string `json:"nombre"`
		Year      string `json:"year"`
		Trimestre int    `json:"trimestre"`
		Activo    bool   `json:"activo"`
	}

	var UpdatedData UpdatedCicloEscolar

	if err := c.BodyParser(&UpdatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	ciclo.Nombre = UpdatedData.Nombre
	ciclo.Year = UpdatedData.Year
	ciclo.Trimestre = UpdatedData.Trimestre
	ciclo.Activo = UpdatedData.Activo

	database.Database.Db.Save(&ciclo)

	return c.Status(200).JSON("Ciclo updated succesfully")
}

/*** Methods unrelated from routes and updates ***/

func validString(s string) bool {
	return s != ""
}
