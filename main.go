package main

import (
	"log"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("welcome to the jungle")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", helloWorld)

	// alumnos
	app.Post("/api/alumnos", routes.CreateAlumno)
	app.Get("/api/alumnos", routes.GetAllAlumnos)

	// grupos
	app.Post("/api/grupos", routes.CreateGrupoActivo)
	app.Post("/api/grupos/aprobados", routes.CreateGrupoConcluido)

	// relacion alumnos - grupos
	app.Get("/api/alumnos/grupos", routes.GetAllRelacionAlumnosGrupos)
	app.Post("/api/alumnos/grupos", routes.CreateRelacionAlumnosGrupos)

	// administracion
	app.Post("/api/administracion", routes.CreateCuentaAdmin)

	// permisos
	app.Post("/api/permisos", routes.CreatePermiso)
	app.Post("/api/usuarios", routes.CreateUsuarios)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3030"))
}
