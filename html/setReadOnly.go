package html

import "log"

// SetReadOnly
//
// English:
//
//  A Boolean attribute which, if present, indicates that the user should not be able to edit the
//  value of the input.
//
// The readonly attribute is supported by the text, search, url, tel, email, date, month, week, time,
// datetime-local, number, and password input types.
//
// Português:
//
//  A Boolean attribute which, if present, indicates that the user should not be able to edit the value of the input. The readonly attribute is supported by the text, search, url, tel, email, date, month, week, time, datetime-local, number, and password input types.
func (e *GlobalAttributes) SetReadOnly(readonly bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support readonly property")
	}

	e.selfElement.Set("readonly", readonly)
	return e
}
