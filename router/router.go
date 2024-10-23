package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/avila-r/tasker/domain/tasks"
)

var (
	TaskDomainHandler = tasks.DefaultHandler
)

// Run function initializes the application routes and middleware.
//
// Note that this method use only default handlers by default.
// Change '${domain}DomainHandler' vars to use custom handlers
func Run(app *fiber.App) {
	// Health check route to verify API connection
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Group for API routes
	api := app.Group("api")
	{
		// Version 1 of the API
		v1 := api.Group("/v1")
		{
			tasks := v1.Group("/tasks")

			TaskDomainHandler.Route(tasks)
		}
	}
}
