package html

import "log"

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
