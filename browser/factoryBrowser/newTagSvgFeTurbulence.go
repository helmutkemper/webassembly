package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeTurbulence
//
// English:
//
// The <feTurbulence> SVG filter primitive creates an image using the Perlin turbulence function. It allows the
// synthesis of artificial textures like clouds or marble. The resulting image will fill the entire filter primitive
// subregion.
//
// Português:
//
// A primitiva de filtro SVG <feTurbulence> cria uma imagem usando a função de turbulência Perlin. Permite a síntese
// de texturas artificiais como nuvens ou mármore. A imagem resultante preencherá toda a sub-região primitiva do filtro.
func NewTagSvgFeTurbulence(id string) (ref *html.TagSvgFeTurbulence) {
	ref = &html.TagSvgFeTurbulence{}
	ref.Init(id)

	return ref
}
