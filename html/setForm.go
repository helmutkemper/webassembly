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
