package main

import (
	"github.com/gofiber/fiber"
	"github.com/vindecodex/go-with-redis/controller"
)

func main() {
	app := fiber.New()

	app.Get("/", controller.GetUserHandler)
	app.Listen(":4321")
}
