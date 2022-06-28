package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgAnimateTransform
//
// English:
//
// The animateTransform element animates a transformation attribute on its target element, thereby allowing animations
// to control translation, scaling, rotation, and/or skewing.
//
// Português:
//
// O elemento animateTransform anima um atributo de transformação em seu elemento de destino, permitindo assim que as
// animações controlem a tradução, dimensionamento, rotação e ou inclinação.
func NewTagSvgAnimateTransform() (ref *html.TagSvgAnimateTransform) {
	ref = &html.TagSvgAnimateTransform{}
	ref.Init()

	return ref
}
