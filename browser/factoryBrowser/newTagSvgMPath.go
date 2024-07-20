package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgMPath
//
// English:
//
// The <mpath> sub-element for the <animateMotion> element provides the ability to reference an external <path> element
// as the definition of a motion path.
//
// Português:
//
// O subelemento <mpath> para o elemento <animateMotion> fornece a capacidade de referenciar um elemento <path> externo
// como a definição de um caminho de movimento.
func NewTagSvgMPath() (ref *html.TagSvgMPath) {
	ref = &html.TagSvgMPath{}
	ref.Init()

	return ref
}
