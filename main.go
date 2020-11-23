package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	url := "https://api.github.com/users/"
	var data = nil
	app.Get("/:user", func(c *fiber.Ctx) error {
		resp, err := http.Get(url + c.Params("user"))
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal([]byte(resp), &data)
	})
	app.Listen(":4321")
}
