package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (app *application) routes() *fiber.App {
	fiberApp := fiber.New()

	// Register the fiber middleware
	registerFiberMiddleware(fiberApp)

	// API V1 route group
	v1 := fiberApp.Group("/v1")

	// Images route group
	images := v1.Group("/images")

	// Mode route group
	modes := v1.Group("/modes")

	// Images routes
	images.Post("/", app.uploadImagesHandler)

	// Modes routes
	modes.Get("/", app.getModesHandler)
	modes.Get("/no-of-shapes", app.getNumberOfShapesHandler)

	// File server
	v1.Static(PublicImagePrefix, fmt.Sprintf("%s/files", storagePath))

	// 404 Handler
	fiberApp.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Route not found 🤭",
		})
	})

	return fiberApp
}
