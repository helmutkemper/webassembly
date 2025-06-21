package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgScript
//
// English:
//
// The SVG script element allows to add scripts to an SVG document.
//
//	Notes:
//	  * While SVG's script element is equivalent to the HTML <script> element, it has some discrepancies, like it uses
//	    the href attribute instead of src and it doesn't support ECMAScript modules so far (See browser compatibility
//	    below for details)
//
// Português:
//
// O elemento script SVG permite adicionar scripts a um documento SVG.
//
//	Notas:
//	  * Embora o elemento script do SVG seja equivalente ao elemento HTML <script>, ele tem algumas discrepâncias, como
//	    usar o atributo href em vez de src e não suportar módulos ECMAScript até agora (consulte a compatibilidade do
//	    navegador abaixo para obter detalhes)
func NewTagSvgScript() (ref *html.TagSvgScript) {
	ref = &html.TagSvgScript{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
