//go:generate stringer -type=Mode
package primitive

import (
	"fmt"
	"os/exec"
	"strings"
	"unicode"
)

// Mode defines the shapes used when transforming the images
type Mode int

// ModeData represent the data of a particular Mode
type ModeData struct {
	Mode Mode   `json:"mode"`
	Name string `json:"name"`
}

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

	for i := 0; i < 8; i++ {
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
