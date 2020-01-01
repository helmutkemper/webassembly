package Html

import (
	"sync"
	"syscall/js"
)

// en: The Image{} struct creates a new HTMLImageElement instance. It is equivalent
// to document.createElement('Img').
type Image struct {
	// en: html parent element
	HtmlParent js.Value
	element    js.Value

	// en: list of html properties
	SetProperty map[string]interface{}

	// en: wait onLoad event to release create() method
	WaitLoad  bool
	waitGroup sync.WaitGroup
}

func (el *Image) SetParent(parent interface{}) {
	el.HtmlParent = parent.(js.Value)
}

func (el *Image) Get() js.Value {
	return el.element
}

func (el *Image) Create() {
	if el.element != js.Undefined() {
		el.element.Call("delete")
	}

	el.element = el.HtmlParent.Call("createElement", "Img")

	for property, value := range el.SetProperty {
		el.element.Set(property, value)
	}

	if el.WaitLoad == true {
		el.waitGroup.Add(1)
		el.element.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			el.waitGroup.Done()
			return nil
		}))

		el.waitGroup.Wait()
	}
}
