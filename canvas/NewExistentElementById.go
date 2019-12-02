package canvas

import "github.com/helmutkemper/iotmaker.platform.webbrowser/document"

func NewExistentElementById(id string) document.Element {
	el := document.Element{}
	el.InitializeExistentElementById(id)

	return el
}
