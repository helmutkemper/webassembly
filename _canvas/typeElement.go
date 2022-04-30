package canvas

import (
	"github.com/helmutkemper/iotmaker.webassembly"
	"syscall/js"
)

type Element struct {
	SelfElement js.Value
	Document
}

func (el *Element) NewCanvas(id string) js.Value {
	return el.Create("canvas", id)
}

func (el *Element) Create(name, id string) js.Value {
	el.Document.Initialize()
	el.SelfElement = el.selfDocument.Call("createElement", name)
	el.SelfElement.Set("id", id)

	return el.SelfElement
}

func (el *Element) InitializeExistentElementById(id string) {
	el.Document = iotmaker_platform_webbrowser.NewDocument()
	el.SelfElement = el.selfDocument.Call("getElementById", id)
}

func (el *Element) InitializeDocument() {
	el.Document = iotmaker_platform_webbrowser.NewDocument()
}

func (el *Element) Get() js.Value {
	return el.SelfElement
}

func (el *Element) AppendElementToDocumentBody() {
	el.Document.AppendChildToDocumentBody(el.SelfElement)
}
