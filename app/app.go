package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/avila-r/env"
	"github.com/avila-r/gor"
	"github.com/avila-r/tasker"
	"github.com/avila-r/tasker/router"
)

var (
	url = env.Get("SERVER_URL")
)

func main() {
	if err := env.Load(tasker.RootPath); err != nil {
		log.Fatalf(err.Error())
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: gor.ErrHandler,
	})

	router.Run(app)

	app.Listen(url)
}
