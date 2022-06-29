package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeDiffuseLighting
//
// English:
//
// The <feDiffuseLighting> SVG filter primitive lights an image using the alpha channel as a bump map. The resulting
// image, which is an RGBA opaque image, depends on the light color, light position and surface geometry of the input
// bump map.
//
// The light map produced by this filter primitive can be combined with a texture image using the multiply term of the
// arithmetic operator of the <feComposite> filter primitive. Multiple light sources can be simulated by adding several
// of these light maps together before applying it to the texture image.
//
// Português:
//
// A primitiva de filtro SVG <feDiffuseLighting> ilumina uma imagem usando o canal alfa como um mapa de relevo. A imagem
// resultante, que é uma imagem opaca RGBA, depende da cor da luz, posição da luz e geometria da superfície do mapa de
// relevo de entrada.
//
// O mapa de luz produzido por esta primitiva de filtro pode ser combinado com uma imagem de textura usando o termo
// multiplicar do operador aritmético da primitiva de filtro <feComposite>. Várias fontes de luz podem ser simuladas
// adicionando vários desses mapas de luz antes de aplicá-los à imagem de textura.
func NewTagSvgFeDiffuseLighting() (ref *html.TagSvgFeDiffuseLighting) {
	ref = &html.TagSvgFeDiffuseLighting{}
	ref.Init()

	return ref
}
