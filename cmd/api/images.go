package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jwambugu/images-transformer/pkg/primitive"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type generateImageOptions struct {
	NumberOfShapes int
	Mode           primitive.Mode
}

func createTempFile(prefix, extension string) (*os.File, error) {
	tempFile, err := ioutil.TempFile(imageStorageDir, prefix)

	if err != nil {
		return nil, fmt.Errorf("primitive.createTempFile.TempFile:: %v", err)
	}

	defer func(name string) {
		_ = os.Remove(name)
	}(tempFile.Name())

	fileToCreate := fmt.Sprintf("%s.%s", tempFile.Name(), extension)

	return os.Create(fileToCreate)
}

func (app *application) generateImage(filename string, opt generateImageOptions) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	fileExtension := filepath.Ext(file.Name())[1:]

	output, err := primitive.Transform(file, fileExtension, opt.NumberOfShapes, primitive.WithMode(opt.Mode))

	if err != nil {
		return "", err
	}

	outputFile, err := createTempFile("", fileExtension)

	if err != nil {
		return "", err
	}

	_, err = io.Copy(outputFile, output)

	if err != nil {
		return "", err
	}

	defer func(outputFile *os.File) {
		_ = outputFile.Close()
	}(outputFile)

	return outputFile.Name(), nil
}

func (app *application) uploadImagesHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("photos")

	if err != nil {
		return app.errorResponse(c, fiber.StatusBadRequest, err)
	}

	selectedMode := c.FormValue("mode")
	selectedNumberOfShapes := c.FormValue("shapes")

	numberOfShapes, err := strconv.Atoi(selectedNumberOfShapes)

	if err != nil {
		return app.errorResponse(c, fiber.StatusBadRequest, err)
	}

	mode, err := strconv.Atoi(selectedMode)

	if err != nil {
		return app.errorResponse(c, fiber.StatusBadRequest, err)
	}

	imageOptions := generateImageOptions{
		NumberOfShapes: numberOfShapes,
		Mode:           primitive.Mode(mode),
	}

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	filename = strings.ToLower(filename)
	path := fmt.Sprintf("%s/files/%s", storagePath, filename)

	// Save the file to the disk
	if err := c.SaveFile(file, path); err != nil {
		return app.errorResponse(c, fiber.StatusInternalServerError, err)
	}

	generatedImage, err := app.generateImage(path, imageOptions)
	generatedImage = filepath.Base(generatedImage)

	if err != nil {
		return app.errorResponse(c, fiber.StatusInternalServerError, err)
	}

	publicURL := fmt.Sprintf("%s/v1%s/%s", app.config.AppURL, PublicImagePrefix, generatedImage)

	return app.successResponse(c, fiber.StatusOK, map[string]interface{}{
		"filename": generatedImage,
		"url":      publicURL,
	})
}
