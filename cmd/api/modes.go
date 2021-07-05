package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jwambugu/images-transformer/pkg/primitive"
)

func (app *application) getModesHandler(c *fiber.Ctx) error {
	// Get all the modes we have
	modes := primitive.Modes()

	return app.successResponse(c, fiber.StatusOK, map[string]interface{}{
		"modes": modes,
	})
}
