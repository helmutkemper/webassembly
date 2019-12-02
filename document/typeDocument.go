package document

import (
	"syscall/js"
)

type Document struct {
	hasInitialized bool
	SelfDocument   js.Value
}

func (el *Document) Initialize() {
	el.hasInitialized = true
	el.SelfDocument = js.Global().Get("document")
}

func (el *Document) Get() js.Value {

	if el.hasInitialized == false {
		el.Initialize()
	}

	return el.SelfDocument
}

func (el *Document) AppendChildToDocumentBody(value interface{}) {

	if el.hasInitialized == false {
		el.Initialize()
	}

	el.SelfDocument.Get("body").Call("appendChild", value)
}

func (el *Document) AppendChild(element string, value interface{}) {

	if el.hasInitialized == false {
		el.Initialize()
	}

	el.SelfDocument.Get(element).Call("appendChild", value)
}
