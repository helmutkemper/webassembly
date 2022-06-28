package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeSpotLight
//
// English:
//
// The <feSpotLight> SVG filter primitive defines a light source that can be used to create a spotlight effect.
// It is used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>.
//
// Português:
//
// A primitiva de filtro SVG <feSpotLight> define uma fonte de luz que pode ser usada para criar um efeito de spotlight.
// Ele é usado dentro de uma primitiva de filtro de iluminação: <feDiffeLighting> ou <feSpecularLighting>.
func NewTagSvgFeSpotLight() (ref *html.TagSvgFeSpotLight) {
	ref = &html.TagSvgFeSpotLight{}
	ref.Init()

	return ref
}
