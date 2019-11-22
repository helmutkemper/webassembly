package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

type DrawImage struct {
	// en: Specifies the image, canvas, or video element to use
	Image js.Value

	// en: The x coordinate where to place the image on the canvas
	X iotmaker_types.Pixel

	// en: The y coordinate where to place the image on the canvas
	Y iotmaker_types.Pixel

	// en: Optional. The x coordinate where to start clipping
	SX iotmaker_types.Pixel

	// en: Optional. The y coordinate where to start clipping
	SY iotmaker_types.Pixel

	// en: Optional. The width of the clipped image
	SWidth iotmaker_types.Pixel

	// en: Optional. The height of the clipped image
	SHeight iotmaker_types.Pixel

	// en: Optional. The width of the image to use (stretch or reduce the image)
	Width iotmaker_types.Pixel

	// en: Optional. The height of the image to use (stretch or reduce the image)
	Height iotmaker_types.Pixel
}
