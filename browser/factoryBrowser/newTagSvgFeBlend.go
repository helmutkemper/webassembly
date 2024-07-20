package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgFeBlend
//
// English:
//
// The <feBlend> SVG filter primitive composes two objects together ruled by a certain blending mode.
//
// This is similar to what is known from image editing software when blending two layers. The mode is defined by the
// mode attribute.
//
// Português:
//
// A primitiva de filtro SVG <feBlend> compõe dois objetos juntos governados por um certo modo de mesclagem.
//
// Isso é semelhante ao que é conhecido no software de edição de imagens ao misturar duas camadas. O modo é definido
// pelo atributo mode.
func NewTagSvgFeBlend() (ref *html.TagSvgFeBlend) {
	ref = &html.TagSvgFeBlend{}
	ref.Init()

	return ref
}
