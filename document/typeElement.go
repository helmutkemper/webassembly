package document

import (
	"syscall/js"
)

type Element struct {
	SelfElement js.Value
	Document    Document
}

func (el *Element) NewCanvas(id string) js.Value {
	return el.Create("canvas", id)
}

func (el *Element) Create(name, id string) js.Value {
	el.Document.Initialize()
	el.SelfElement = el.Document.SelfDocument.Call("createElement", name)
	el.SelfElement.Set("id", id)

	return el.SelfElement
}

func (el *Element) InitializeExistentElementById(id string) {
	el.Document = NewDocument()
	el.SelfElement = el.Document.SelfDocument.Call("getElementById", id)
}

func (el *Element) InitializeDocument() {
	el.Document = NewDocument()
}

func (el *Element) Get() js.Value {
	return el.SelfElement
}

func (el *Element) AppendElementToDocumentBody() {
	el.Document.AppendChildToDocumentBody(el.SelfElement)
}
