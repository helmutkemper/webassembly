package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgText
//
// English:
//
// The SVG <text> element draws a graphics element consisting of text. It's possible to apply a gradient, pattern,
// clipping path, mask, or filter to <text>, like any other SVG graphics element.
//
// If text is included in SVG not inside of a <text> element, it is not rendered. This is different than being hidden
// by default, as setting the display property won't show the text.
//
// Português:
//
// O elemento SVG <text> desenha um elemento gráfico que consiste em texto. É possível aplicar um gradiente, padrão,
// caminho de recorte, máscara ou filtro a <text>, como qualquer outro elemento gráfico SVG.
//
// Se o texto for incluído no SVG fora de um elemento <text>, ele não será renderizado. Isso é diferente de estar
// oculto por padrão, pois definir a propriedade de exibição não mostrará o texto.
func NewTagSvgText() (ref *html.TagSvgText) {
	ref = &html.TagSvgText{}
	ref.Init()

	return ref
}
