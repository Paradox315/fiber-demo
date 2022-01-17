package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type helloReq struct {
	name string `json:"name"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		var req helloReq
		c.BodyParser(&req)
		return c.JSON(fiber.Map{"message": "hello" + req.name})
	})

	log.Fatal(app.Listen(":30000"))
}
