package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeDisplacementMap
//
// English:
//
// The <feDisplacementMap> SVG filter primitive uses the pixel values from the image from in2 to spatially displace the
// image from in.
//
// The formula for the transformation looks like this:
//
// P'(x,y) ← P( x + scale * (XC(x,y) - 0.5), y + scale * (YC(x,y) - 0.5))
//
// where P(x,y) is the input image, in, and P'(x,y) is the destination. XC(x,y) and YC(x,y) are the component values of
// the channel designated by xChannelSelector and yChannelSelector.
//
// Português:
//
// A primitiva de filtro SVG <feDisplacementMap> usa os valores de pixel da imagem de in2 para deslocar espacialmente a
// imagem de in.
//
// A fórmula da transformação fica assim:
//
// P'(x,y) ← P( x + scale * (XC(x,y) - 0.5), y + scale * (YC(x,y) - 0.5))
//
// onde P(x,y) é a imagem de entrada, in, e P'(x,y) é o destino. XC(x,y) e YC(x,y) são os valores componentes do canal
// designado por xChannelSelector e yChannelSelector.
func NewTagSvgFeDisplacementMap() (ref *html.TagSvgFeDisplacementMap) {
	ref = &html.TagSvgFeDisplacementMap{}
	ref.Init()

	return ref
}
