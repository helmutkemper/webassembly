package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFePointLight
//
// English:
//
// The <fePointLight> filter primitive defines a light source which allows to create a point light effect.
// It that can be used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>.
//
// Português:
//
// A primitiva de filtro <fePointLight> define uma fonte de luz que permite criar um efeito de luz pontual.
// Ele que pode ser usado dentro de uma primitiva de filtro de iluminação: <feDiffuseLighting> ou <feSpecularLighting>.
func NewTagSvgFePointLight(id string) (ref *html.TagSvgFePointLight) {
	ref = &html.TagSvgFePointLight{}
	ref.Init(id)

	return ref
}
