package _global

import "log"

// SetInputType
//
// English:
//
//  How an <input> works varies considerably depending on the value of its type attribute, hence the
//  different types are covered in their own separate reference pages.
//
// If this attribute is not specified, the default type adopted is text.
//
// Português:
//
//  Como um <input> funciona varia consideravelmente dependendo do valor de seu atributo type,
//  portanto, os diferentes tipos são abordados em suas próprias páginas de referência separadas.
//
// Se este atributo não for especificado, o tipo padrão adotado é texto.
func (e *GlobalAttributes) SetInputType(inputType InputType) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support type property")
	}

	e.selfElement.Set("type", inputType.String())
	return e
}
