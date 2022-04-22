package html

import "log"

// SetRel
//
// English:
//
//  The relationship of the linked URL as space-separated link types.
//
// Português:
//
//  O relacionamento da URL vinculada como tipos de link separados por espaço.
func (e *GlobalAttributes) SetRel(rel string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support rel property")
	}

	e.selfElement.Set("rel", rel)
	return e
}
