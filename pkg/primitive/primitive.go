//go:generate stringer -type=Mode
package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

// Mode defines the shapes used when transforming the images
type Mode int

type (
	// ModeData represent the data of a particular Mode
	ModeData struct {
		Mode Mode   `json:"mode"`
		Name string `json:"name"`
	}
)

// List of all modes supported by the primitive CLI
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRectangle
	ModeEllipse
	ModeCircle
	ModeRotatedRectangle
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

// splitCamelCase splits a camelcase string and returns a new string separated by a white space
func splitCamelCase(s string) string {
	var parts []string
	start := 0

	for i, r := range s {
		if i != 0 && unicode.IsUpper(r) {
			parts = append(parts, s[start:i])
			start = i
		}
	}

	if start != len(s) {
		parts = append(parts, s[start:])
	}

	return strings.Join(parts, " ")
}

// Modes returns a slice of ModeData
func Modes() []*ModeData {
	var modes []*ModeData

	for i := 0; i <= 8; i++ {
		m := Mode(i).String()[4:]
		m = splitCamelCase(m)

		data := &ModeData{
			Mode: Mode(i),
			Name: m,
		}

		modes = append(modes, data)
	}

	return modes
}

func primitive(inputFile, outputFile string, numberOfShapes int, args ...string) ([]byte, error) {
	argsStr := fmt.Sprintf("-i %s -o %s -n %d", inputFile, outputFile, numberOfShapes)
	args = append(strings.Fields(argsStr), args...)

	cmd := exec.Command("primitive", args...)

	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("primitive.RunCommand:: %v", err)
	}

	return output, nil
}

// createTempFile creates temp file in the default temp file dir
func createTempFile(prefix, extension string) (*os.File, error) {
	tempFile, err := ioutil.TempFile("", prefix)

	if err != nil {
		return nil, fmt.Errorf("primitive.createTempFile.TempFile:: %v", err)
	}

	defer func(name string) {
		_ = os.Remove(name)
	}(tempFile.Name())

	fileToCreate := fmt.Sprintf("%s.%s", tempFile.Name(), extension)

	file, err := os.Create(fileToCreate)

	if err != nil {
		return nil, fmt.Errorf("primitive.createTempFile.OSCreate:: %v", err)
	}

	return file, nil
}

// WithMode returns the Mode to use to transform the image
// Default Mode is ModeTriangle
func WithMode(m Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", m)}
	}
}

// Transform takes the provided image and applies primitive transformation to it then returns a reader
// to the modified image
func Transform(image io.Reader, extension string, numberOfShapes int, opts ...func() []string) (io.Reader, error) {
	var args []string

	for _, opt := range opts {
		args = append(args, opt()...)
	}

	inputTempFile, err := createTempFile("input_", extension)

	if err != nil {
		return nil, fmt.Errorf("primitive: failed to create temp input file:: %v", err)
	}

	outputTempFile, err := createTempFile("output_", extension)

	if err != nil {
		return nil, fmt.Errorf("primitive: failed to create temp output file:: %v", err)
	}

	// Read the image
	_, err = io.Copy(inputTempFile, image)

	if err != nil {
		return nil, fmt.Errorf("primitive.Transform.CopyTempImage:: %v", err)
	}

	_, err = primitive(inputTempFile.Name(), outputTempFile.Name(), numberOfShapes, args...)

	if err != nil {
		return nil, err
	}

	// Read out into a reader, return reader and delete out
	b := bytes.NewBuffer(nil)

	_, err = io.Copy(b, outputTempFile)

	if err != nil {
		return nil, fmt.Errorf("primitive.Transform.CopyTempOutputImage:: %v", err)
	}

	return b, nil
}
