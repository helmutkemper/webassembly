package _global

import "log"

// SetFormAction
//
// English:
//
//  A string indicating the URL to which to submit the data. This takes precedence over the action
//  attribute on the <form> element that owns the <input>.
//
// This attribute is also available on <input type="image"> and <button> elements.
//
// Português:
//
//  Uma string indicando o URL para o qual enviar os dados. Isso tem precedência sobre o atributo
//  action no elemento <form> que possui o <input>.
//
// Este atributo também está disponível nos elementos <input type="image"> e <button>.
func (e *GlobalAttributes) SetFormAction(action string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagLabel:
	default:
		log.Printf("tag " + e.tag.String() + " does not support for property")
	}

	e.selfElement.Set("formaction", action)
	return e
}
