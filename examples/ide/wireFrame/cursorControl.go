package wireFrame

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"image/color"
)

type CursorControl interface {
	// Init Initializes the instance
	Init()
	SetColor(color color.RGBA)
	GetSvg() (svg *html.TagSvg)
	GetVisible() (visible string)
	SetFatherId(fatherId string)
	Hide()
	ShowVertical()
	ShowHorizontal()
	Move(x, y, width, height int)
	Update(width, height int) (err error)
	ChangesDirection()
}
