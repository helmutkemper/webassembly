package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgStyle
//
// English:
//
// The SVG <style> element allows style sheets to be embedded directly within SVG content.
//
//	Notes:
//	  * SVG's style element has the same attributes as the corresponding element in HTML
//	    (see HTML's <style> element).
//
// Português:
//
// O elemento SVG <style> permite que as folhas de estilo sejam incorporadas diretamente no conteúdo SVG.
//
//	Notas:
//	  * O elemento de estilo SVG tem os mesmos atributos que o elemento correspondente em HTML
//	    (definir elemento HTML <style>).
func NewTagSvgStyle() (ref *html.TagSvgStyle) {
	ref = &html.TagSvgStyle{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
