package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

type Element struct {
	selfElement js.Value
	Document
}

func (el *Element) NewCanvas(id string) js.Value {
	return el.Create("canvas", id)
}

func (el *Element) Create(name, id string) js.Value {
	el.Document.Initialize()
	el.selfElement = el.selfDocument.Call("createElement", name)
	el.selfElement.Set("id", id)

	return el.selfElement
}

func (el *Element) InitializeExistentElementById(id string) {
	el.Document = NewDocument()
	el.selfElement = el.selfDocument.Call("getElementById", id)
}

func (el *Element) InitializeDocument() {
	el.Document = NewDocument()
}

func (el *Element) Get() js.Value {
	return el.selfElement
}

func (el *Element) AppendElementToDocumentBody() {
	el.Document.AppendChildToDocumentBody(el.selfElement)
}
