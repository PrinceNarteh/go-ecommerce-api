package main

import (
	"log"

	"github.com/PrinceNarteh/go-ecommerce-api/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome api")
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
