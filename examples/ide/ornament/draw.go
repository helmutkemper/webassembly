package ornament

import "github.com/helmutkemper/webassembly/browser/html"

type Draw interface {
	Init() (err error)
	Update(width, height int) (err error)
	GetSvg() (svg *html.TagSvg)
}
