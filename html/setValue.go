package html

import "log"

// SetValue
//
// English:
//
//  Defines the value associated with the element.
//
// Português:
//
//  Define o valor associado ao elemento.
func (e *GlobalAttributes) SetValue(value string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	case KTagOption:
	default:
		log.Printf("tag " + e.tag.String() + " does not support buttonValue property")
	}

	e.selfElement.Set("value", value)
	return e
}
