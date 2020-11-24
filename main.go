package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	url := "https://api.github.com/users/"
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	app.Get("/:user", func(c *fiber.Ctx) error {

		user := c.Params("user")

		val, err := rdb.Get(ctx, "vindecodex").Result()
		if err != nil {
			log.Fatalln("Error 30:")
		}

		if len(val) > 0 {
			resp, err := http.Get(url + user)
			if err != nil {
				log.Fatalln("Error 36:")
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln("Error 40:")
			}

			err = rdb.Set(ctx, "user", body, 0).Err()
			if err != nil {
				log.Fatalln("Error 45:")
			}

			return c.Send(body)
		}

		return c.SendString(val)
	})

	app.Listen(":4321")
}
