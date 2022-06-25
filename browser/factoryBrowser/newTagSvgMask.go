package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgMask
//
// English:
//
// The <mask> element defines an alpha mask for compositing the current object into the background. A mask is
// used/referenced using the mask property.
//
// Português:
//
// O elemento <mask> define uma máscara alfa para compor o objeto atual em segundo plano. Uma máscara é used/referenced
// usando a propriedade mask.
func NewTagSvgMask(id string) (ref *html.TagSvgMask) {
	ref = &html.TagSvgMask{}
	ref.Init(id)

	return ref
}
