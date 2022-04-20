package html

import "log"

// SetLabel
//
// English:
//
//  This attribute is text for the label indicating the meaning of the option. If the label attribute
//  isn't defined, its value is that of the element text content.
//
// Português:
//
//  Este atributo é um texto para o rótulo que indica o significado da opção. Se o atributo label não
//  estiver definido, seu valor será o do conteúdo do texto do elemento.
func (e *GlobalAttributes) SetLabel(label string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagOption:
	default:
		log.Printf("tag " + e.tag.String() + " does not support label property")
	}

	e.selfElement.Set("label", label)
	return e
}
