//go:build js
// +build js

package html

import (
	"fmt"
	"image/color"
)

func ExampleRGBAToJs() {
	colorRGBA := color.RGBA{
		R: 10,
		G: 20,
		B: 30,
		A: 100,
	}
	fmt.Printf("Color: %v\n", RGBAToJs(colorRGBA))

	// Output:
	// Color: rgba( 10, 20, 30, 100 )
}
