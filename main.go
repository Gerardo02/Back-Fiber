package main

import (
	//"github.com/Gerardo02/Back-Fiber/database"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Welcome to the jungle")
}

func main() {
	//database.ConnectDb()
	app := fiber.New()

	app.Get("/api", helloWorld)

	app.Listen(":3030")
}
