package html

import (
	"image/color"
	"strconv"
)

// RGBAToJs
//
// English:
//
//  Convert the color format color.RGBA to javascript format rgba(R, G, B, A)
//
// Português:
//
//  Converte uma cor em formato color.RGBA para o formato javascript rgba(R, G, B, A)
func RGBAToJs(color color.RGBA) string {
	return "rgba( " +
		strconv.Itoa(int(color.R)) +
		", " +
		strconv.Itoa(int(color.G)) +
		", " +
		strconv.Itoa(int(color.B)) +
		", " +
		strconv.FormatFloat(float64(color.A)/255.0, 'g', -1, 64) +
		" )"
}
