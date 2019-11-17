package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

type DrawImage struct {
	// en: Specifies the image, canvas, or video element to use
	Image js.Value

	// en: The x coordinate where to place the image on the canvas
	X float64

	// en: The y coordinate where to place the image on the canvas
	Y float64

	// en: Optional. The x coordinate where to start clipping
	SX float64

	// en: Optional. The y coordinate where to start clipping
	SY float64

	// en: Optional. The width of the clipped image
	SWidth float64

	// en: Optional. The height of the clipped image
	SHeight float64

	// en: Optional. The width of the image to use (stretch or reduce the image)
	Width float64

	// en: Optional. The height of the image to use (stretch or reduce the image)
	Height float64
}
