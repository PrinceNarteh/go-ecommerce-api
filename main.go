package main

import (
	"log"

	"github.com/PrinceNarteh/go-ecommerce-api/database"
	"github.com/PrinceNarteh/go-ecommerce-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome E-Commerce API")
}

func routesSetup(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)

	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetAllUsers)
	app.Get("/api/users/:userId", routes.GetUser)
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	routesSetup(app)

	log.Fatal(app.Listen(":3000"))
}
