package main

import (
	"log"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/src/routes"
	"github.com/Gerardo02/Back-Fiber/src/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("welcome to the jungle")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", helloWorld)

	app.Post("/login", utils.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("SECRET_KEY")},
	}))

	app.Get("/restricted", utils.Restricted)

	// alumnos
	app.Post("/api/alumnos", routes.CreateAlumno)
	app.Get("/api/alumnos", routes.GetAllAlumnos)
	app.Put("/api/alumnos/:id", routes.UpdateAlumnos)

	// grupos
	app.Get("/api/grupos", routes.GetAllGruposActivos)
	app.Post("/api/grupos", routes.CreateGrupoActivo)
	app.Post("/api/grupos/aprobados", routes.CreateGrupoConcluido)

	// relacion alumnos - grupos
	app.Get("/api/alumnos/grupos", routes.GetAllRelacionAlumnosGrupos)
	app.Post("/api/alumnos/grupos", routes.CreateRelacionAlumnosGrupos)

	// relacion grupos - listas
	app.Post("/api/grupos/listas", routes.CreateRelacionGrupoListas)

	// administracion
	app.Put("/api/administracion/:id", routes.UpdateAdmin)

	// permisos
	app.Post("/api/permisos", routes.CreatePermiso)
	app.Post("/api/usuarios", routes.CreateUsuarios)

	// especialidad
	app.Post("/api/especialidad", routes.CreateEspecialidad)

	// documentos entregados
	app.Put("api/documentos/:id", routes.UpdateDocuments)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3030"))
}
