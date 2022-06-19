package html

type SvgWritingMode string

func (e SvgWritingMode) String() string {
	return string(e)
}

const (
	// KSvgWritingModeHorizontalTb
	//
	// English:
	//
	// This value defines a top-to-bottom block flow direction. Both the writing mode and the typographic mode are
	// horizontal.
	//
	// Português:
	//
	// Este valor define uma direção de fluxo de bloco de cima para baixo. Tanto o modo de escrita quanto o modo
	// tipográfico são horizontais.
	KSvgWritingModeHorizontalTb SvgWritingMode = "horizontal-tb"

	// KSvgWritingModeVerticalRl
	//
	// English:
	//
	// This value defines a right-to-left block flow direction. Both the writing mode and the typographic mode are
	// vertical.
	//
	// Português:
	//
	// Este valor define uma direção de fluxo de bloco da direita para a esquerda. Tanto o modo de escrita quanto o modo
	// tipográfico são verticais.
	KSvgWritingModeVerticalRl SvgWritingMode = "vertical-rl"

	// KSvgWritingModeVerticalLr
	//
	// English:
	//
	// This value defines a left-to-right block flow direction. Both the writing mode and the typographic mode are
	// vertical.
	//
	// Português:
	//
	// Este valor define uma direção de fluxo de bloco da esquerda para a direita. Tanto o modo de escrita quanto o modo
	// tipográfico são verticais.
	KSvgWritingModeVerticalLr SvgWritingMode = "vertical-lr"
)
