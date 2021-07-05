package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func registerFiberMiddleware(app *fiber.App) {
	app.Use(
		// Enable the cors middleware.
		cors.New(),
		// Enable the recover middleware.
		recover.New(recover.Config{
			EnableStackTrace: true,
		}),
		// Enable the compress middleware.
		compress.New(),
		// RequestID middleware for Fiber that adds an identifier to the response.
		requestid.New(),
		// Enable the logger middleware.
		logger.New(logger.Config{
			//Output:     logFile,
			Format:     "PID: ${pid} Status: ${status} RequestID: ${locals:requestid} - ${method} ${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "Africa/Nairobi",
		}),
	)
}
