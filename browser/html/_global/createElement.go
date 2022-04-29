package _global

import (
	"log"
	"syscall/js"
)

// CreateElement
//
// English:
//
//  In an HTML document, the Document.createElement() method creates the specified HTML element or an
//  HTMLUnknownElement if the given element name is not known.
//
// Português:
//
//  Em um documento HTML, o método Document.createElement() cria o elemento HTML especificado ou um
//  HTMLUnknownElement se o nome do elemento dado não for conhecido.
func (e *GlobalAttributes) CreateElement(tag Tag) (ref *GlobalAttributes) {
	e.selfElement = js.Global().Get("document").Call("createElement", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}
	e.tag = tag

	return e
}
