package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgPattern
//
// English:
//
// The <pattern> element defines a graphics object which can be redrawn at repeated x- and y-coordinate intervals
// ("tiled") to cover an area.
//
// The <pattern> is referenced by the fill and/or stroke attributes on other graphics elements to fill or stroke those
// elements with the referenced pattern.
//
// Português:
//
// O elemento <pattern> define um objeto gráfico que pode ser redesenhado em intervalos repetidos de coordenadas x e y
// ("lado a lado") para cobrir uma área.
//
// O <pattern> é referenciado pelos atributos fill andor stroke em outros elementos gráficos para preencher ou traçar
// esses elementos com o padrão referenciado.
func NewTagSvgPattern() (ref *html.TagSvgPattern) {
	ref = &html.TagSvgPattern{}
	ref.Init()

	return ref
}
