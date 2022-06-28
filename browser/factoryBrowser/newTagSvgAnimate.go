package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgAnimate
//
// English:
//
//  The SVG <animate> element provides a way to animate an attribute of an element over time.
//
// Português:
//
//  O elemento SVG <animate> fornece uma maneira de animar um atributo de um elemento ao longo do tempo.
func NewTagSvgAnimate() (ref *html.TagSvgAnimate) {
	ref = &html.TagSvgAnimate{}
	ref.Init()

	return ref
}
