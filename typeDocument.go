package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

type Document struct {
	hasInitialized bool
	selfDocument   js.Value
}

func (el *Document) Initialize() {
	el.hasInitialized = true
	el.selfDocument = js.Global().Get("document")
}

func (el *Document) Get() js.Value {

	if el.hasInitialized == false {
		el.Initialize()
	}

	return el.selfDocument
}

func (el *Document) AppendChildToDocumentBody(value interface{}) {

	if el.hasInitialized == false {
		el.Initialize()
	}

	el.selfDocument.Get("body").Call("appendChild", value)
}

func (el *Document) AppendChild(element string, value interface{}) {

	if el.hasInitialized == false {
		el.Initialize()
	}

	el.selfDocument.Get(element).Call("appendChild", value)
}
