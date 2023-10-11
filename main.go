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

	app.Post("/api/alumnos", routes.CreateAlumno)

	app.Post("/api/administracion", routes.CreateCuentaAdmin)

	app.Post("/api/permisos", routes.CreatePermiso)

	app.Post("/api/usuarios", routes.CreateUsuarios)

	app.Post("/api/grupos", routes.CreateGrupoActivo)

	app.Post("/api/grupos/aprobados", routes.CreateGrupoConcluido)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3030"))
}
