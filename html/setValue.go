package html

import "log"

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
