package ornament

import "github.com/helmutkemper/webassembly/browser/html"

// Draw Draw the visual elements of the device
type Draw interface {

	// Init Initializes the instance
	Init() (err error)

	// Update Draw the element design
	Update(width, height int) (err error)

	// GetSvg Returns the SVG tag with the element design
	GetSvg() (svg *html.TagSvg)

	// SetWarning sets the visibility of the warning mark
	SetWarning(warning bool)
}
