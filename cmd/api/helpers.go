package main

import "github.com/gofiber/fiber/v2"

// successResponse will return all the successful responses
// All the responses will sets the content header to application/json
func (app *application) successResponse(c *fiber.Ctx, status int, data map[string]interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"data":   data,
	})
}

// errorResponse will return all the error responses (4xx and 5xx)
// All the responses will sets the content header to application/json
func (app *application) errorResponse(c *fiber.Ctx, status int, error interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"data":   error,
	})
}
