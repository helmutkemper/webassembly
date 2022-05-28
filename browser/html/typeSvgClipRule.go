package html

type SvgClipRule string

func (e SvgClipRule) String() string {
	return string(e)
}

const (
	// KSvgClipRuleNonzero
	//
	// English:
	//
	//  The value nonzero determines the "insideness" of a point in the shape by drawing a ray from that point to infinity
	//  in any direction, and then examining the places where a segment of the shape crosses the ray. Starting with a
	//  count of zero, add one each time a path segment crosses the ray from left to right and subtract one each time a
	//  path segment crosses the ray from right to left. After counting the crossings, if the result is zero then the
	//  point is outside the path. Otherwise, it is inside.
	//
	// Português:
	//
	//  O valor diferente de zero determina a "interioridade" de um ponto na forma desenhando um raio desse ponto até o
	//  infinito em qualquer direção e, em seguida, examinando os locais onde um segmento da forma cruza o raio. Começando
	//  com uma contagem de zero, adicione um cada vez que um segmento de caminho cruza o raio da esquerda para a direita
	//  e subtraia um cada vez que um segmento de caminho cruza o raio da direita para a esquerda. Depois de contar os
	//  cruzamentos, se o resultado for zero, então o ponto está fora do caminho. Caso contrário, está dentro.
	KSvgClipRuleNonzero SvgClipRule = "nonzero"

	// KSvgClipRuleEvenOdd
	//
	// English:
	//
	//  The value evenodd determines the "insideness" of a point in the shape by drawing a ray from that point to
	//  infinity in any direction and counting the number of path segments from the given shape that the ray crosses. If
	//  this number is odd, the point is inside; if even, the point is outside.
	//
	// Português:
	//
	//  O valor evenodd determina a "interioridade" de um ponto na forma desenhando um raio desse ponto até o infinito em
	//  qualquer direção e contando o número de segmentos de caminho da forma dada que o raio cruza. Se este número for
	//  ímpar, o ponto está dentro; se par, o ponto está fora.
	KSvgClipRuleEvenOdd SvgClipRule = "evenodd"
)
