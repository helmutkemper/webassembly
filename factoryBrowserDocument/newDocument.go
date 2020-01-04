package factoryBrowserDocument

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"

func NewDocument() document.Document {
	el := document.Document{}
	el.Initialize()

	return el
}
