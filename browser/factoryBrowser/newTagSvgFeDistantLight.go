package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgFeDistantLight
//
// English:
//
// The <feDistantLight> filter primitive defines a distant light source that can be used within a lighting filter
// primitive: <feDiffuseLighting> or <feSpecularLighting>.
//
// Português:
//
// A primitiva de filtro <feDistantLight> define uma fonte de luz distante que pode ser usada em uma primitiva de filtro
// de iluminação: <feDiffuseLighting> ou <feSpecularLighting>.
func NewTagSvgFeDistantLight() (ref *html.TagSvgFeDistantLight) {
	ref = &html.TagSvgFeDistantLight{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
