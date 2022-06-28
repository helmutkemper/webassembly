package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeOffset
//
// English:
//
// The <feOffset> SVG filter primitive allows to offset the input image. The input image as a whole is offset by the
// values specified in the dx and dy attributes.
//
// Português:
//
// A primitiva de filtro SVG <feOffset> permite deslocar a imagem de entrada. A imagem de entrada é compensada pelos
// valores especificados nos atributos dx e dy.
func NewTagSvgFeOffset() (ref *html.TagSvgFeOffset) {
	ref = &html.TagSvgFeOffset{}
	ref.Init()

	return ref
}
