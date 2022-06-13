package html

type SvgStrokeLinecap string

func (e SvgStrokeLinecap) String() string {
	return string(e)
}

const (
	// KSvgStrokeLinecapButt
	//
	// English:
	//
	// The butt value indicates that the stroke for each subpath does not extend beyond its two endpoints.
	// On a zero length subpath, the path will not be rendered at all.
	//
	// Português:
	//
	// O valor de extremidade indica que o traço para cada subcaminho não se estende além de seus dois pontos finais.
	// Em um subcaminho de comprimento zero, o caminho não será renderizado.
	//
	// todo: exemplo https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linecap#butt
	KSvgStrokeLinecapButt SvgStrokeLinecap = "butt"

	// KSvgStrokeLinecapRound
	//
	// English:
	//
	// The round value indicates that at the end of each subpath the stroke will be extended by a half circle with a
	// diameter equal to the stroke width. On a zero length subpath, the stroke consists of a full circle centered at the
	// subpath's point.
	//
	// Português:
	//
	// O valor arredondado indica que no final de cada subcaminho o traço será estendido por um semicírculo com um
	// diâmetro igual à largura do traço. Em um subcaminho de comprimento zero, o traçado consiste em um círculo completo
	// centrado no ponto do subcaminho.
	//
	// todo: exemplo: https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linecap#round
	KSvgStrokeLinecapRound SvgStrokeLinecap = "round"

	// KSvgStrokeLinecapSquare
	//
	// English:
	//
	// The square value indicates that at the end of each subpath the stroke will be extended by a rectangle with a width
	// equal to half the width of the stroke and a height equal to the width of the stroke. On a zero length subpath, the
	// stroke consists of a square with its width equal to the stroke width, centered at the subpath's point.
	//
	// Português:
	//
	// O valor quadrado indica que no final de cada subcaminho o traço será estendido por um retângulo com largura igual
	// à metade da largura do traço e altura igual à largura do traço. Em um subcaminho de comprimento zero, o traçado
	// consiste em um quadrado com sua largura igual à largura do traçado, centralizado no ponto do subcaminho.
	//
	// todo: exemplo: https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linecap#square
	KSvgStrokeLinecapSquare SvgStrokeLinecap = "square"
)
