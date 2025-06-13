package rulesConnection

import (
	"errors"
	"fmt"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
)

// golangManagerTypeConversion Convert types and colors for Golang
type golangManagerTypeConversion struct {
	err error
}

// GetError Returns the list of errors during type conversion
func (e *golangManagerTypeConversion) GetError() (err error) {
	return e.err
}

// Verify Check the syntax looking for errors
func (e *golangManagerTypeConversion) Verify(dataType string) {
	switch dataType {
	case "bool", "int", "int64", "uint", "uint64", "float64",
		"slice", "string", "struct":
	default:
		e.err = errors.Join(e.err, fmt.Errorf("unknown data type `%s` for `%v`", dataType, TypeOfDataCurrentlyInEffect))
	}
}

// TypeToColor For each type of data, numerical, string, boolean ... a visual identity color is attributed
func (e *golangManagerTypeConversion) TypeToColor(dataType string) (color color.RGBA) {
	e.Verify(dataType)

	switch dataType {
	case "bool":
		color = factoryColor.NewGreen()
	case "int":
		color = factoryColor.NewBlue()
	case "int64":
		color = factoryColor.NewBlue()
	case "uint":
		color = factoryColor.NewBlueViolet()
	case "uint64":
		color = factoryColor.NewBlueViolet()
	case "float64":
		color = factoryColor.NewYellowGreen()
	case "slice":
		color = factoryColor.NewDarkMagenta()
	case "string":
		color = factoryColor.NewMediumTurquoise()
	case "struct":
		color = factoryColor.NewGainsboro()
	default:
		color = factoryColor.NewRed()
	}

	return
}
