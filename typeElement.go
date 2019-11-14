package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

type Element struct {
	selfElement js.Value
	Document
}

func (el *Element) InitializeById(id string) {
	el.Document = NewDocument()
	el.selfElement = el.selfDocument.Call("getElementById", id)
}

func (el *Element) Get() js.Value {
	return el.selfElement
}
