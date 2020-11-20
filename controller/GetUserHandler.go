package controller

import "github.com/gofiber/fiber"

func GetUserHandler(c *fiber.Ctx) error {
	c.SendString("Hello")
	return nil
}
