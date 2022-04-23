package html

import "log"

// SetWarp
//
// English:
//
//  Indicates how the control wraps text.
//
// Português:
//
//  Indica como o controle quebra o texto.
func (e *GlobalAttributes) SetWarp(warp Warp) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagTextarea:
	default:
		log.Printf("tag " + e.tag.String() + " does not support warp property")
	}

	e.selfElement.Set("warp", warp.String())
	return e
}
