package rulesConnection

import "image/color"

type RulesDataType interface {

	// GetError Returns the list of errors during type conversion
	GetError() (err error)

	// Verify Check the syntax looking for errors
	Verify(dataType string) (err error)

	// TypeToColor For each type of data, numerical, string, boolean ... a visual identity color is attributed
	//
	//   Rule:
	//     * Leave the red color reserved for errors
	TypeToColor(dataType string) (color color.RGBA)
}
