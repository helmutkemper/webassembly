package html

import "log"

// SetForm
//
// English:
//
//  The <form> element to associate the button with (its form owner). The value of this attribute must be the id of a <form> in the same document. (If this attribute is not set, the <button> is associated with its ancestor <form> element, if any.)
//
// This attribute lets you associate <button> elements to <form>s anywhere in the document, not just inside a <form>. It can also override an ancestor <form> element.
//
// Português:
//
//  O elemento <form> ao qual associar o botão (seu proprietário do formulário). O valor deste atributo deve ser o id de um <form> no mesmo documento. (Se esse atributo não for definido, o <button> será associado ao elemento <form> ancestral, se houver.)
//
// Este atributo permite associar elementos <button> a <form>s em qualquer lugar do documento, não apenas dentro de um <form>. Ele também pode substituir um elemento <form> ancestral.
func (e *GlobalAttributes) SetForm(form string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support form property")
	}

	e.selfElement.Set("form", form)
	return e
}

// SetName
//
// English:
//
//  The name of the button, submitted as a pair with the button's value as part of the form data,
//  when that button is used to submit the form.
//
// Português:
//
//  The name of the button, submitted as a pair with the button's value as part of the form data, when that button is used to submit the form.
func (e *GlobalAttributes) SetName(name string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support name property")
	}

	e.selfElement.Set("name", name)
	return e
}

// SetValue
//
// English:
//
//  Defines the value associated with the button's name when it's submitted with the form data. This value is passed to the server in params when the form is submitted using this button.
//
// Português:
//
//  Defines the value associated with the button's name when it's submitted with the form data. This value is passed to the server in params when the form is submitted using this button.
func (e *GlobalAttributes) SetValue(value string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support value property")
	}

	e.selfElement.Set("value", value)
	return e
}
