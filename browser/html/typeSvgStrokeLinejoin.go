package html

type SvgStrokeLinejoin string

func (e SvgStrokeLinejoin) String() string {
	return string(e)
}

const (
	// KSvgStrokeLinejoinArcs
	//
	// English:
	//
	// The arcs value indicates that an arcs corner is to be used to join path segments. The arcs shape is formed by
	// extending the outer edges of the stroke at the join point with arcs that have the same curvature as the outer edges
	// at the join point.
	//
	// Português:
	//
	// O valor de arcos indica que um canto de arco deve ser usado para unir segmentos de caminho. A forma dos arcos é
	// formada estendendo as bordas externas do traço no ponto de junção com arcos que têm a mesma curvatura das bordas
	// externas no ponto de junção.
	//
	// todo: exemplo: https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linejoin#arcs
	KSvgStrokeLinejoinArcs SvgStrokeLinejoin = "arcs"

	// KSvgStrokeLinejoinBevel
	//
	// English:
	//
	// The bevel value indicates that a bevelled corner is to be used to join path segments.
	//
	// Português:
	//
	// O valor de chanfro indica que um canto chanfrado deve ser usado para unir segmentos de caminho.
	//
	// todo: exemplo: https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linejoin#bevel
	KSvgStrokeLinejoinBevel SvgStrokeLinejoin = "bevel"

	// KSvgStrokeLinejoinMiter
	//
	// English:
	//
	// The miter value indicates that a sharp corner is to be used to join path segments. The corner is formed by
	// extending the outer edges of the stroke at the tangents of the path segments until they intersect.
	//
	// Português:
	//
	// O valor da mitra indica que um canto agudo deve ser usado para unir segmentos de caminho. O canto é formado
	// estendendo as bordas externas do traçado nas tangentes dos segmentos de caminho até que eles se cruzem.
	//
	// todo: exemplo: https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linejoin#miter
	KSvgStrokeLinejoinMiter SvgStrokeLinejoin = "miter"

	// KSvgStrokeLinejoinMiterClip
	//
	// English:
	//
	// The miter-clip value indicates that a sharp corner is to be used to join path segments. The corner is formed by
	// extending the outer edges of the stroke at the tangents of the path segments until they intersect.
	//
	// If the stroke-miterlimit is exceeded, the miter is clipped at a distance equal to half the stroke-miterlimit value
	// multiplied by the stroke width from the intersection of the path segments. This provides a better rendering than
	// miter on very sharp join or in case of an animation.
	//
	// Português:
	//
	// O valor de clipe de mitra indica que um canto agudo deve ser usado para unir segmentos de caminho. O canto é
	// formado estendendo as bordas externas do traçado nas tangentes dos segmentos de caminho até que eles se cruzem.
	//
	// Se o stroke-miterlimit for excedido, a mitra é cortada a uma distância igual à metade do valor stroke-miterlimit
	// multiplicado pela largura do traço da interseção dos segmentos do caminho. Isso fornece uma renderização melhor do
	// que a mitra em uma junção muito nítida ou no caso de uma animação.
	//
	// todo: exemplo: https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linejoin#miter-clip
	KSvgStrokeLinejoinMiterClip SvgStrokeLinejoin = "miter-clip"

	// KSvgStrokeLinejoinRound
	//
	// English:
	//
	// The round value indicates that a round corner is to be used to join path segments.
	//
	// Português:
	//
	// O valor arredondado indica que um canto arredondado deve ser usado para unir segmentos de caminho.
	//
	// todo: exemplo: https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linejoin#round
	KSvgStrokeLinejoinRound SvgStrokeLinejoin = "round"
)
