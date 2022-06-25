package html

type SvgSide string

func (e SvgSide) String() string {
	return string(e)
}

const (

	// KSvgSideLeft
	//
	// English:
	//
	// This value places the text on the left side of the path (relative to the path direction).
	//
	// Português:
	//
	// Esse valor coloca o texto no lado esquerdo do caminho (em relação à direção do caminho).
	KSvgSideLeft SvgSide = "left"

	// KSvgSideRight
	//
	// English:
	//
	// This value places the text on the right side of the path (relative to the path direction).
	// This effectively reverses the path direction.
	//
	// Português:
	//
	// Esse valor coloca o texto no lado direito do caminho (em relação à direção do caminho).
	// Isso efetivamente inverte a direção do caminho.
	KSvgSideRight SvgSide = "right"
)
