package html

import (
	"log"
	"syscall/js"
)

// CreateElement
//
// English:
//
//
//
// Português:
//
//  Em um documento HTML, o método Document.createElement() cria o elemento HTML especificado ou um
//  HTMLUnknownElement se o nome do elemento dado não for conhecido.
//
// Em um documento XUL, o elemento XUL especificado é criado.
//
// Em outros documentos, ele cria um elemento com um namespace URI null.
//
// nomeDaTag é uma string que especifica o tipo do elemento a ser criado. O nodeName (en-US) do elemento criado é inicializado com o valor da nomeDaTag. Não use nomes qualificados (como "html:a") com este método.
//
//
// fixme: terminar de documentar
//
//
func (e *GlobalAttributes) CreateElement(tag Tag) (ref *GlobalAttributes) {
	e.selfElement = js.Global().Get("document").Call("createElement", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}
	e.tag = tag

	return e
}
