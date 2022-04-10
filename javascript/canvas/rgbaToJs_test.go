//go:build js
// +build js

package canvas

import (
	"fmt"
	"image/color"
)

func ExampleRGBAToJs() {
	color := color.RGBA{
		R: 10,
		G: 20,
		B: 30,
		A: 100,
	}
	fmt.Printf("Color: %v\n", RGBAToJs(color))

	// Output:
	// Color: rgba( 10, 20, 30, 100 )
}
