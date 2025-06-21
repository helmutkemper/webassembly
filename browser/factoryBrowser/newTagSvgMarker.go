package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgMarker
//
// English:
//
// The <marker> element defines the graphic that is to be used for drawing arrowheads or polymarkers on a given <path>,
// <line>, <polyline> or <polygon> element.
//
// Markers are attached to shapes using the marker-start, marker-mid, and marker-end properties.
//
// Português:
//
// O elemento <marker> define o gráfico que deve ser usado para desenhar pontas de seta ou polimarcadores em um
// determinado elemento <path>, <line>, <polyline> ou <polygon>.
//
// Os marcadores são anexados às formas usando as propriedades de início do marcador, meio do marcador e final do
// marcador.
func NewTagSvgMarker() (ref *html.TagSvgMarker) {
	ref = &html.TagSvgMarker{}
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
