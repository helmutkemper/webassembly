package factoryBrowserElement

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/element"
)

func NewElement() element.Element {
	el := element.Element{}
	el.InitializeDocument()

	return el
}
