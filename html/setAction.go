package html

import "log"

// SetAction
//
// English:
//
//  The URL that processes the form submission. This value can be overridden by a formaction
//  attribute on a <button>, <input type="submit">, or <input type="image"> element.
//
// This attribute is ignored when method="dialog" is set.
//
// Português:
//
//  A URL que processa o envio do formulário. Esse valor pode ser substituído por um atributo
//  formaction em um elemento <button>, <input type="submit"> ou <input type="image">.
//
// Este atributo é ignorado quando method="dialog" é definido.
func (e *GlobalAttributes) SetAction(action string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support action property")
	}

	e.selfElement.Set("action", action)
	return e
}
