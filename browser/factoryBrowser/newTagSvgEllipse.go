package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewTagSvgEllipse
//
// English:
//
// The <ellipse> element is an SVG basic shape, used to create ellipses based on a center coordinate, and both their x
// and y radius.
//
//	Notes:
//	  * Ellipses are unable to specify the exact orientation of the ellipse (if, for example, you wanted to draw an
//	    ellipse tilted at a 45 degree angle), but it can be rotated by using the transform attribute.
//
// Português:
//
// O elemento <ellipse> é uma forma básica SVG, usada para criar elipses com base em uma coordenada central e em seus
// raios x e y.
//
//	Note:
//	  * As elipses não podem especificar a orientação exata da elipse (se, por exemplo, você quiser desenhar uma
//	    elipse inclinada em um ângulo de 45 graus), mas ela pode ser girada usando o atributo transform.
func NewTagSvgEllipse() (ref *html.TagSvgEllipse) {
	ref = &html.TagSvgEllipse{}
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}
