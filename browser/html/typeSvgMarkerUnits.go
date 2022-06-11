package html

type SvgMarkerUnits string

func (e SvgMarkerUnits) String() string {
	return string(e)
}

const (
	// KSvgMarkerUnitsUserSpaceOnUse
	//
	// English:
	//
	// This value specifies that the markerWidth and markerHeight attributes and the contents of the <marker> element
	// represent values in the current user coordinate system in place for the graphic object referencing the marker
	// (i.e., the user coordinate system for the element referencing the <marker> element via a marker, marker-start,
	// marker-mid, or marker-end property).
	//
	// Português:
	//
	// Este valor especifica que os atributos markerWidth e markerHeight e o conteúdo do elemento <marker> representam
	// valores no sistema de coordenadas do usuário atual em vigor para o objeto gráfico que faz referência ao marcador
	// (ou seja, o sistema de coordenadas do usuário para o elemento que faz referência ao <marker> elemento por meio de
	// um marcador, marcador inicial, marcador intermediário ou marcador final).
	KSvgMarkerUnitsUserSpaceOnUse SvgMarkerUnits = "userSpaceOnUse"

	// KSvgMarkerUnitsStrokeWidth
	//
	// English:
	//
	// This value specifies that the markerWidth and markerHeight attributes and the contents of the <marker> element
	// represent values in a coordinate system which has a single unit equal the size in user units of the current stroke
	// width (see the stroke-width attribute) in place for the graphic object referencing the marker.
	//
	// Português:
	//
	// Este valor especifica que os atributos markerWidth e markerHeight e o conteúdo do elemento <marker> representam
	// valores em um sistema de coordenadas que tem uma única unidade igual ao tamanho em unidades de usuário da largura
	// do traço atual (consulte o atributo stroke-width) no lugar para o objeto gráfico que faz referência ao marcador.
	KSvgMarkerUnitsStrokeWidth SvgMarkerUnits = "strokeWidth"
)
