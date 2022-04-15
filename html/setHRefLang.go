package html

import "log"

// SetHRefLang
//
// English:
//
//  Hints at the human language of the linked URL. No built-in functionality. Allowed values are the
//  same as the global lang attribute.
//
// Português:
//
//  Dicas para a linguagem humana da URL vinculada. Nenhuma funcionalidade embutida. Os valores
//  permitidos são os mesmos do atributo lang global.
func (e *GlobalAttributes) SetHRefLang(hreflang string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support hreflang property")
	}

	e.selfElement.Set("hreflang", hreflang)
	return e
}
