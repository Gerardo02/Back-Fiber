package main

import (
	"log"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/src/routes"
	"github.com/Gerardo02/Back-Fiber/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App, secretKey string) {

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
	app.Put("/api/alumnos/:id", routes.UpdateAlumnos)
	app.Delete("/api/alumnos/:id", routes.DeleteAlumno)

	// grupos
	app.Get("/api/grupos", routes.GetAllGruposActivos)
	app.Post("/api/grupos", routes.CreateGrupoActivo)
	app.Delete("/api/grupos", routes.DeleteAllGruposActivos)
	app.Delete("/api/grupos/purge", routes.DropSoftDeletesGruposActivos)
	app.Delete("/api/grupos/:id", routes.DeleteSingleGroup)
	app.Post("/api/grupos/aprobados", routes.CreateGrupoConcluido)

	// relacion alumnos - grupos
	app.Get("/api/alumnos/grupos", routes.GetAllRelacionAlumnosGrupos)
	app.Post("/api/alumnos/grupos", routes.CreateRelacionAlumnosGrupos)

	// relacion grupos - listas
	app.Post("/api/grupos/listas", routes.CreateRelacionGrupoListas)

	// administracion
	app.Put("/api/administracion/:id", routes.UpdateAdmin)
	app.Get("/api/administracion", routes.GetAdministraciones)

	// permisos
	app.Post("/api/permisos", routes.CreatePermiso)
	app.Post("/api/usuarios", routes.CreateUsuarios)

	// especialidad
	app.Post("/api/especialidad", routes.CreateEspecialidad)
	app.Get("/api/especialidad", routes.GetAllEspecialidades)

	// documentos entregados
	app.Put("api/documentos/:id", routes.UpdateDocuments)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))

	secretKey := utils.GoDotEnvVariable("SECRET_KEY")
	setupRoutes(app, secretKey)

	log.Fatal(app.Listen(":3030"))
}
