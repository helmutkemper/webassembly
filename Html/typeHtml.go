package Html

import (
	"syscall/js"
)

type Html struct{}

func (el Html) NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad bool) Image {
	img := Image{
		SetProperty: propertiesList,
		WaitLoad:    waitLoad,
	}
	img.SetParent(parent)
	img.Create()

	return img
}

func (el Html) Append(document, element interface{}) {
	document.(js.Value).Get("body").Call("appendChild", element)
}

func (el Html) Remove(document, element interface{}) {
	document.(js.Value).Get("body").Call("removeChild", element)
}

func (el Html) GetDocumentWidth(document interface{}) int {
	return document.(js.Value).Get("body").Get("clientWidth").Int()
}

func (el Html) GetDocumentHeight(document interface{}) int {
	return document.(js.Value).Get("body").Get("clientHeight").Int()
}
