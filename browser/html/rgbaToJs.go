package html

import (
	"image/color"
	"strconv"
)

func RGBAToJs(color color.RGBA) string {
	return "rgba( " +
		strconv.Itoa(int(color.R)) +
		", " +
		strconv.Itoa(int(color.G)) +
		", " +
		strconv.Itoa(int(color.B)) +
		", " +
		strconv.FormatFloat(float64(color.A)/255.0, 'f', 4, 64) +
		" )"
}
