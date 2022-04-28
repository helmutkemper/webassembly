package _global

import "log"

// SetRequired
//
// English:
//
//  A Boolean attribute indicating that an option with a non-empty string value must be selected.
//
// Português:
//
//  Um atributo booleano que indica que uma opção com um valor de string não vazio deve ser
//  selecionada.
func (e *GlobalAttributes) SetRequired(required bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagSelect:
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support required property")
	}

	e.selfElement.Set("required", required)
	return e
}
