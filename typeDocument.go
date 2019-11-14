package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

type Document struct {
	selfDocument js.Value
}

func (el *Document) Initialize() {
	el.selfDocument = js.Global().Get("document")
}

func (el *Document) Get() js.Value {
	return el.selfDocument
}
