package html

import "log"

// SetFor
//
// English:
//
//  The value of the for attribute must be a single id for a labelable form-related element in the
//  same document as the <label> element. So, any given label element can be associated with only
//  one form control.
//
//   Note:
//     * To programmatically set the for attribute, use htmlFor.
//
// The first element in the document with an id attribute matching the value of the for attribute is
// the labeled control for this label element — if the element with that id is actually a labelable
// element. If it is not a labelable element, then the for attribute has no effect. If there are
// other elements that also match the id value, later in the document, they are not considered.
//
// Multiple label elements can be given the same value for their for attribute; doing so causes the
// associated form control (the form control that for value references) to have multiple labels.
//
//   Note:
//     * A <label> element can have both a for attribute and a contained control element, as long as
//       the for attribute points to the contained control element.
//
// Português:
//
//  O valor do atributo for deve ser um único id para um elemento rotulável relacionado ao formulário
//  no mesmo documento que o elemento <label>. Portanto, qualquer elemento de rótulo pode ser
//  associado a apenas um controle de formulário.
//
//   Nota:
//     * Programaticamente definir o atributo for, use htmlFor.
//
// O primeiro elemento no documento com um atributo id correspondente ao valor do atributo é o
// controle rotulado para este elemento label - se o elemento com esse ID é realmente um elemento
// labelable. Se não é um elemento labelable, em seguida, o atributo for tem nenhum efeito.
// Se existem outros elementos que também correspondem ao valor id, mais adiante no documento,
// eles não são considerados.
//
// Vários elementos de rótulo podem receber o mesmo valor para seu atributo for; isso faz com que o
// controle de formulário associado (o controle de formulário para referências de valor) tenha
// vários rótulos.
//
//   Nota:
//     * Um elemento <label> pode ter um atributo for e um elemento de controle contido, desde que
//       o atributo for aponte para o elemento de controle contido.
func (e *GlobalAttributes) SetFor(value string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagLabel:
	default:
		log.Printf("tag " + e.tag.String() + " does not support for property")
	}

	e.selfElement.Set("for", value)
	return e
}
