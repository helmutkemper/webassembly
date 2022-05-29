package html

type SvgIn string

func (e SvgIn) String() string {
	return string(e)
}

const (

	// KSvgInSourceGraphic
	//
	// English:
	//
	//  This keyword represents the graphics elements that were the original input into the <filter> element.
	//
	// Portuguese
	//
	//  Esta palavra-chave representa os elementos gráficos que foram a entrada original no elemento <filter>.
	KSvgInSourceGraphic SvgIn = "SourceGraphic"

	// KSvgInSourceAlpha
	//
	// English:
	//
	//  This keyword represents the graphics elements that were the original input into the <filter> element.
	//  SourceAlpha has all of the same rules as SourceGraphic except that only the alpha channel is used.
	//
	// Portuguese
	//
	//  Esta palavra-chave representa os elementos gráficos que foram a entrada original no elemento <filter>.
	//  SourceAlpha tem todas as mesmas regras que SourceGraphic, exceto que apenas o canal alfa é usado.
	KSvgInSourceAlpha SvgIn = "SourceAlpha"

	// KSvgInBackgroundImage
	//
	// English:
	//
	//  This keyword represents an image snapshot of the SVG document under the filter region at the time that the
	//  <filter> element was invoked.
	//
	// Portuguese
	//
	//  Essa palavra-chave representa um instantâneo de imagem do documento SVG na região do filtro no momento em que o
	//  elemento <filter> foi invocado.
	KSvgInBackgroundImage SvgIn = "BackgroundImage"

	// KSvgInBackgroundAlpha
	//
	// English:
	//
	//  Same as BackgroundImage except only the alpha channel is used.
	//
	// Portuguese
	//
	//  O mesmo que BackgroundImage, exceto que apenas o canal alfa é usado.
	KSvgInBackgroundAlpha SvgIn = "BackgroundAlpha"

	// KSvgInFillPaint
	//
	// English:
	//
	//  This keyword represents the value of the fill property on the target element for the filter effect.
	//
	// In many cases, the FillPaint is opaque everywhere, but that might not be the case if a shape is painted with a
	// gradient or pattern which itself includes transparent or semi-transparent parts.
	//
	// Portuguese
	//
	//  Essa palavra-chave representa o valor da propriedade fill no elemento de destino para o efeito de filtro.
	//
	// Em muitos casos, o FillPaint é opaco em todos os lugares, mas esse pode não ser o caso se uma forma for pintada
	// com um gradiente ou padrão que inclui partes transparentes ou semitransparentes.
	KSvgInFillPaint SvgIn = "FillPaint"

	// KSvgInStrokePaint
	//
	// English:
	//
	//  This keyword represents the value of the stroke property on the target element for the filter effect.
	//
	// In many cases, the StrokePaint is opaque everywhere, but that might not be the case if a shape is painted with a
	// gradient or pattern which itself includes transparent or semi-transparent parts.
	//
	// Portuguese
	//
	//  Essa palavra-chave representa o valor da propriedade stroke no elemento de destino para o efeito de filtro.
	//
	// Em muitos casos, o StrokePaint é opaco em todos os lugares, mas esse pode não ser o caso se uma forma for pintada
	// com um gradiente ou padrão que inclui partes transparentes ou semitransparentes.
	KSvgInStrokePaint SvgIn = "StrokePaint"
)
