package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"go-with-redis/controller/GetUserHandler.go"
)

func main() {
	fmt.Println("MAIN")

	app := fiber.New()

	const url = "https://api.github.com/users/"

	app.Get("/api/:search", controller.GetUserHandler())

	fmt.Println("Listening on port: 1234")
	log.Fatal(app.Listen(":1234"))
}
