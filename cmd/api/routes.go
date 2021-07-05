package main

import "github.com/gofiber/fiber/v2"

func (app *application) routes() *fiber.App {
	fiberApp := fiber.New()

	// Register the fiber middleware
	registerFiberMiddleware(fiberApp)

	// API V1 route group
	v1 := fiberApp.Group("/v1")

	// Images route group
	images := v1.Group("/images")

	// Images routes
	images.Get("/", func(ctx *fiber.Ctx) error {
		return nil
	})

	// 404 Handler
	fiberApp.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Route not found 🤭",
		})
	})

	return fiberApp
}
