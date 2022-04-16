package html

import "log"

// SetAutofocus
//
// English:
//
//  This Boolean attribute specifies that the button should have input focus when the page loads.
//  Only one element in a document can have this attribute.
//
// Português:
//
//  Este atributo booleano especifica que o botão deve ter foco de entrada quando a página for
//  carregada. Apenas um elemento em um documento pode ter esse atributo.
func (e *GlobalAttributes) SetAutofocus(autofocus bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support autofocus property")
	}

	e.selfElement.Set("autofocus", autofocus)
	return e
}
