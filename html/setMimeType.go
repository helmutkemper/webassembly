package html

import "log"

// SetMimeType
//
// English:
//
//  Hints at the linked URL's format with a MIME type. No built-in functionality.
//
// Português:
//
//  Dicas no formato do URL vinculado com um tipo MIME. Nenhuma funcionalidade embutida.
func (e *GlobalAttributes) SetMimeType(typeProperty Mime) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support type property")
	}

	e.selfElement.Set("type", typeProperty)
	return e
}
