package html

type SvgTextDecorationStyle string

func (e SvgTextDecorationStyle) String() string {
	return string(e)
}

const (
	// KSvgTextDecorationStyleSolid
	//
	// English:
	//
	// Draws a single line.
	//
	// Português:
	//
	// Desenha uma única linha.
	KSvgTextDecorationStyleSolid SvgTextDecorationStyle = "solid"

	// KSvgTextDecorationStyleDouble
	//
	// English:
	//
	// Draws a double line.
	//
	// Português:
	//
	// Desenha uma linha dupla.
	KSvgTextDecorationStyleDouble SvgTextDecorationStyle = "double"

	// KSvgTextDecorationStyleDotted
	//
	// English:
	//
	// Draws a dotted line.
	//
	// Português:
	//
	// Desenha uma linha pontilhada.
	KSvgTextDecorationStyleDotted SvgTextDecorationStyle = "dotted"

	// KSvgTextDecorationStyleDashed
	//
	// English:
	//
	// Draws a dashed line.
	//
	// Português:
	//
	// Desenha uma linha tracejada.
	KSvgTextDecorationStyleDashed SvgTextDecorationStyle = "dashed"

	// KSvgTextDecorationStyleWavy
	//
	// English:
	//
	// Draws a wavy line.
	//
	// Português:
	//
	// Desenha uma linha ondulada.
	KSvgTextDecorationStyleWavy SvgTextDecorationStyle = "wavy"
)
