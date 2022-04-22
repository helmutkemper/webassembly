package html

import "log"

// SetFormNovalidate
//
// English:
//
//  This Boolean attribute indicates that the form shouldn't be validated when submitted.
//
// If this attribute is not set (and therefore the form is validated), it can be overridden by a
// formnovalidate attribute on a <button>, <input type="submit">, or <input type="image"> element
// belonging to the form.
//
// Português:
//
//  Este atributo booleano indica que o formulário não deve ser validado quando enviado.
//
// Se este atributo não estiver definido (e, portanto, o formulário for validado), ele poderá ser
// substituído por um atributo formnovalidate em um elemento <button>, <input type="submit"> ou
// <input type="image"> pertencente a a forma.
func (e *GlobalAttributes) SetFormNovalidate(novalidate string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support novalidate property")
	}

	e.selfElement.Set("novalidate", novalidate)
	return e
}
