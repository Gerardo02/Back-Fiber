package main

import (
	"log"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Welcome to the jungle")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	//alumnos
	app.Post("/alumnos", routes.CreateAlumno)
	app.Get("/alumnos", routes.GetAlumnos)
	app.Get("/alumnos/:id", routes.GetAlumno)
	app.Delete("/alumnos/:id", routes.DeleteAlumno)
	app.Put("/alumnos/:id", routes.UpdateAlumno)

	// Control escolar
	app.Post("/control", routes.CreateControlEscolar)
	app.Put("/control/:id", routes.UpdateControlEscolar)
	app.Get("/control", routes.GetAllControlEscolar)
	app.Get("/control/:id", routes.GetControlEscolar)
	app.Delete("/control/:id", routes.DeleteControlEscolar)

	// Administracion

	app.Get("/admin", routes.GetAllAdministracion)
	app.Get("/admin/:id", routes.GetAdministracion)
	app.Post("/admin", routes.CreateAdmin)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
