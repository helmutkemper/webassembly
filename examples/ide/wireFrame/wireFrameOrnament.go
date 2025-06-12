package wireFrame

import "github.com/helmutkemper/webassembly/browser/html"

type Ornament interface {
	BindingType(binding string)
	AddPoint(x, y int, bindingType string)
	GetPointsLength() (length int)
	Init() (err error)
	GetSvg() (svg *html.TagSvg)
	Update(width, height int) (err error)
}
