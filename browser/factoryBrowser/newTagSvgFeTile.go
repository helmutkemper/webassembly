package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeTile
//
// English:
//
// The <feTile> SVG filter primitive allows to fill a target rectangle with a repeated, tiled pattern of an input image.
// The effect is similar to the one of a <pattern>.
//
// Português:
//
// A primitiva de filtro SVG <feTile> permite preencher um retângulo de destino com um padrão repetido e lado a lado de
// uma imagem de entrada.
// O efeito é semelhante ao de um <pattern>.
func NewTagSvgFeTile() (ref *html.TagSvgFeTile) {
	ref = &html.TagSvgFeTile{}
	ref.Init()

	return ref
}
