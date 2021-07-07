package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"path/filepath"
	"strings"
	"time"
)

func (app *application) uploadImagesHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("photos")

	if err != nil {
		log.Fatalln(err)
	}

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	filename = strings.ToLower(filename)

	fileExtension := filepath.Ext(filename)[1:]

	path := fmt.Sprintf("%s/files/%s", storagePath, filename)
	fmt.Println(path, fileExtension)

	if err := c.SaveFile(file, path); err != nil {
		return app.errorResponse(c, fiber.StatusInternalServerError, err)
	}

	return app.successResponse(c, fiber.StatusOK, map[string]interface{}{
		"filename": filename,
	})
}
