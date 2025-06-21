package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgFeSpecularLighting
//
// English:
//
// The <feSpecularLighting> SVG filter primitive lights a source graphic using the alpha channel as a bump map. The
// resulting image is an RGBA image based on the light color. The lighting calculation follows the standard specular
// component of the Phong lighting model. The resulting image depends on the light color, light position and surface
// geometry of the input bump map. The result of the lighting calculation is added. The filter primitive assumes that
// the viewer is at infinity in the z direction.
//
// This filter primitive produces an image which contains the specular reflection part of the lighting calculation.
// Such a map is intended to be combined with a texture using the add term of the arithmetic <feComposite> method.
// Multiple light sources can be simulated by adding several of these light maps before applying it to the texture
// image.
//
// Português:
//
// A primitiva de filtro SVG <feSpecularLighting> ilumina um gráfico de origem usando o canal alfa como um mapa de
// relevo.
// A imagem resultante é uma imagem RGBA baseada na cor clara. O cálculo da iluminação segue o componente especular
// padrão do modelo de iluminação Phong. A imagem resultante depende da cor da luz, posição da luz e geometria da
// superfície do mapa de relevo de entrada. O resultado do cálculo de iluminação é adicionado. A primitiva de filtro
// assume que o visualizador está no infinito na direção z.
//
// Esta primitiva de filtro produz uma imagem que contém a parte de reflexão especular do cálculo de iluminação.
// Tal mapa deve ser combinado com uma textura usando o termo add do método aritmético <feComposite>.
// Várias fontes de luz podem ser simuladas adicionando vários desses mapas de luz antes de aplicá-los à imagem de
// textura.
func NewTagSvgFeSpecularLighting() (ref *html.TagSvgFeSpecularLighting) {
	ref = &html.TagSvgFeSpecularLighting{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
