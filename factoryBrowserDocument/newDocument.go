package factoryBrowserDocument

import "github.com/helmutkemper/iotmaker.platform.webbrowser/document"

func NewDocument() document.Document {
	el := document.Document{}
	el.Initialize()

	return el
}
