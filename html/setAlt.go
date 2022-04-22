package html

import "log"

// SetAlt
//
// English:
//
//  The alt attribute provides alternative text for the image, displaying the value of the attribute
//  if the image src is missing or otherwise fails to load.
//
// Português:
//
//  O atributo alt fornece texto alternativo para a imagem, exibindo o valor do atributo se o src da
//  imagem estiver ausente ou falhar ao carregar.
func (e *GlobalAttributes) SetAlt(alt string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support alt property")
	}

	e.selfElement.Set("alt", alt)
	return e
}
