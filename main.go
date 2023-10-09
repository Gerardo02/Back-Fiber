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

	app.Post("/alumnos", routes.CreateAlumno)
	app.Get("/alumnos", routes.GetAlumnos)
	app.Get("/alumno/:id", routes.GetAlumno)
	app.Delete("/alumnos/:id", routes.DeleteAlumno)
	app.Put("/alumnos/:id", routes.UpdateAlumno)

	// COntrol escolar
	app.Post("/control", routes.CreateControlEscolar)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
