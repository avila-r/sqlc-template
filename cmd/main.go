package main

import (
	"log"

	"github.com/avila-r/env"
	"github.com/avila-r/tasker"

	"github.com/gofiber/fiber/v2"
)

var (
	url = env.Get("SERVER_URL")
)

func main() {
	if err := env.Load(tasker.RootPath); err != nil {
		log.Fatalf(err.Error())
	}

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Listen(url)
}
