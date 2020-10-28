package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)


func GetUserHandler(c * fiber.Ctx) error {
	fmt.Fprintf(c, "%s", c.Params("search"))
	return nil
}
