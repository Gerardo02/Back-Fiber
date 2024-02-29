package main

import (
	"log"
	"os"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/src/routes"
	"github.com/Gerardo02/Back-Fiber/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {

	app.Post("/login", utils.GenerateToken)

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(secretKey)},
	// }))

	app.Get("/restricted", utils.Restricted, func(c *fiber.Ctx) error {
		return c.SendString("asdasd")
	})

	// alumnos
	app.Post("/api/alumnos", routes.CreateEspecialidadRelacionMiddleware, routes.CreateAlumno)
	app.Get("/api/alumnos", routes.GetAllAlumnos)
	app.Get("/api/alumnos/nombres", routes.GetAlumnosNombres)
	app.Put("/api/alumnos/:id", routes.UpdateAlumnos)
	app.Delete("/api/alumnos/:id", routes.DeleteAlumno)

	// grupos
	app.Get("/api/grupos", routes.GetAllGruposActivos)
	app.Post("/api/grupos", routes.CreateGrupoActivo)
	app.Put("/api/grupos/:id", routes.UpdateGrupoActivo)
	app.Delete("/api/grupos", routes.DeleteAllGruposActivos)
	app.Delete("/api/grupos/purge", routes.DropSoftDeletesGruposActivos)
	app.Delete("/api/grupos/:id", routes.DeleteSingleGroup)
	app.Post("/api/grupos/aprobados", routes.CreateGrupoConcluido)

	// relacion alumnos - grupos
	app.Get("/api/alumnos/grupos", routes.GetAllRelacionAlumnosGrupos)
	app.Post("/api/alumnos/grupos", routes.CreateRelacionAlumnosGrupos)
	app.Put("/api/alumnos/grupos/especialidad/:id", routes.UpdateRelacionAlumnoEspecialidad)
	app.Put("/api/alumnos/grupos/:id", routes.UpdateRelacionAlumnoGrupo)

	// relacion grupos - listas
	app.Post("/api/grupos/listas", routes.CreateRelacionGrupoListas)

	// administracion
	app.Put("/api/administracion/:id", routes.UpdateAdmin)
	app.Get("/api/administracion", routes.GetAdministraciones)

	// permisos
	app.Get("/api/usuarios", routes.GetUsuarios)
	app.Get("/api/permisos", routes.GetPermisos)
	app.Put("/api/permisos", routes.UpdatePermisoUser)
	app.Delete("/api/permisos/:id", routes.DeletePermiso)

	app.Delete("/api/usuarios", routes.DeleteUser)
	app.Post("/api/permisos", routes.CreatePermiso)
	app.Post("/api/usuarios", routes.CreateUsuarios)
	app.Post("/api/login", routes.AuthenticateUser)
	app.Put("/api/usuarios/:user", routes.UpdateUserName)
	app.Put("/api/password", routes.UpdateUserPassWord)

	// especialidad
	app.Post("/api/especialidad", routes.CreateEspecialidad)
	app.Get("/api/especialidad", routes.GetAllEspecialidades)

	// documentos entregados
	app.Get("/api/documentos", routes.GetDocumentosEntregados)
	app.Put("/api/documentos/:id", routes.UpdateDocuments)

}

func main() {

	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))

	// secretKey := utils.GoDotEnvVariable("SECRET_KEY")
	setupRoutes(app)
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}
