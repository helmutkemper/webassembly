package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

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
func NewTagSvgFeDistantLight(id string) (ref *html.TagSvgFeDistantLight) {
	ref = &html.TagSvgFeDistantLight{}
	ref.Init(id)

	return ref
}
