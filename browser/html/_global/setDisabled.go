package _global

import "log"

// SetDisabled
//
// English:
//
//  Este atributo booleano impede que o usuário interaja com o elemento.
//
// Português:
//
//  Este atributo booleano impede que o usuário interaja com o elemento.
func (e *GlobalAttributes) SetDisabled(disabled bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	case KTagOption:
	case KTagFieldset:
	default:
		log.Printf("tag " + e.tag.String() + " does not support disabled property")
	}

	e.selfElement.Set("disabled", disabled)
	return e
}
