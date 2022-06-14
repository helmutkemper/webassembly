package html

type SvgTextAnchor string

func (e SvgTextAnchor) String() string {
	return string(e)
}

const (
	// KSvgTextAnchorStart
	//
	// English:
	//
	// The rendered characters are aligned such that the start of the text string is at the initial current text position.
	// For an element with a direction property value of ltr (typical for most European languages), the left side of the
	// text is rendered at the initial text position. For an element with a direction property value of rtl (typical for
	// Arabic and Hebrew), the right side of the text is rendered at the initial text position. For an element with a
	// vertical primary text direction (often typical for Asian text), the top side of the text is rendered at the initial
	// text position.
	//
	// Português:
	//
	// Os caracteres renderizados são alinhados de forma que o início da string de texto esteja na posição inicial do
	// texto atual. Para um elemento com um valor de propriedade de direção de ltr (típico para a maioria dos idiomas
	// europeus), o lado esquerdo do texto é renderizado na posição inicial do texto. Para um elemento com um valor de
	// propriedade de direção de rtl (típico para árabe e hebraico), o lado direito do texto é renderizado na posição
	// inicial do texto. Para um elemento com uma direção de texto primária vertical (geralmente típica para texto
	// asiático), o lado superior do texto é renderizado na posição inicial do texto.
	KSvgTextAnchorStart SvgTextAnchor = "start"

	// KSvgTextAnchorMiddle
	//
	// English:
	//
	// The rendered characters are aligned such that the middle of the text string is at the current text position.
	// (For text on a path, conceptually the text string is first laid out in a straight line. The midpoint between the
	// start of the text string and the end of the text string is determined. Then, the text string is mapped onto the
	// path with this midpoint placed at the current text position.)
	//
	// Português:
	//
	// Os caracteres renderizados são alinhados de forma que o meio da string de texto esteja na posição atual do texto.
	// (Para texto em um caminho, conceitualmente a string de texto é primeiro disposta em uma linha reta.
	// O ponto médio entre o início da string de texto e o final da string de texto é determinado. Em seguida, a string
	// de texto é mapeada no caminho com este ponto médio colocado na posição atual do texto.)
	KSvgTextAnchorMiddle SvgTextAnchor = "middle"

	// KSvgTextAnchorEnd
	//
	// English:
	//
	// The rendered characters are shifted such that the end of the resulting rendered text (final current text position
	// before applying the text-anchor property) is at the initial current text position. For an element with a direction
	// property value of ltr (typical for most European languages), the right side of the text is rendered at the initial
	// text position. For an element with a direction property value of rtl (typical for Arabic and Hebrew), the left
	// side of the text is rendered at the initial text position. For an element with a vertical primary text direction
	// (often typical for Asian text), the bottom of the text is rendered at the initial text position.
	//
	// Português:
	//
	// Os caracteres renderizados são deslocados de forma que o final do texto renderizado resultante (posição final do
	// texto atual antes de aplicar a propriedade text-anchor) fique na posição inicial do texto atual. Para um elemento
	// com um valor de propriedade de direção de ltr (típico para a maioria dos idiomas europeus), o lado direito do texto
	// é renderizado na posição inicial do texto. Para um elemento com um valor de propriedade de direção de rtl (típico
	// para árabe e hebraico), o lado esquerdo do texto é renderizado na posição inicial do texto. Para um elemento com
	// uma direção de texto primária vertical (geralmente típica de texto asiático), a parte inferior do texto é
	// renderizada na posição inicial do texto.
	KSvgTextAnchorEnd SvgTextAnchor = "end"
)
