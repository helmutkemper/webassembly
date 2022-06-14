package html

// SvgTextDecorationLine
//
// English:
//
// The text-decoration-line CSS property sets the kind of decoration that is used on text in an element, such as an
// underline or overline.
//
// Português:
//
// A propriedade CSS text-decoration-line define o tipo de decoração usada no texto em um elemento, como sublinhado ou
// overline.
type SvgTextDecorationLine string

func (e SvgTextDecorationLine) String() string {
	return string(e)
}

const (
	// KSvgTextDecorationLineNone
	//
	// English:
	//
	// Produces no text decoration.
	//
	// Português:
	//
	// Não produz decoração de texto.
	KSvgTextDecorationLineNone SvgTextDecorationLine = "none"

	// KSvgTextDecorationLineUnderline
	//
	// English:
	//
	// Each line of text has a decorative line beneath it.
	//
	// Português:
	//
	// Cada linha de texto tem uma linha decorativa abaixo dela.
	KSvgTextDecorationLineUnderline SvgTextDecorationLine = "underline"

	// KSvgTextDecorationLineOverline
	//
	// English:
	//
	// Each line of text has a decorative line above it.
	//
	// Português:
	//
	// Cada linha de texto tem uma linha decorativa acima dela.
	KSvgTextDecorationLineOverline SvgTextDecorationLine = "overline"

	// KSvgTextDecorationLineLineThrough
	//
	// English:
	//
	// Each line of text has a decorative line going through its middle.
	//
	// Português:
	//
	// Cada linha de texto tem uma linha decorativa passando pelo meio.
	KSvgTextDecorationLineLineThrough SvgTextDecorationLine = "line-through"
)
