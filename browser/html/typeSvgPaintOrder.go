package html

type SvgPaintOrder string

func (e SvgPaintOrder) String() string {
	return string(e)
}

const (
	// KSvgPaintOrderNormal
	//
	// English:
	//
	// This value indicates that the fill will be painted first, then the stroke, and finally the markers.
	//
	// Português:
	//
	// Esse valor indica que o preenchimento será pintado primeiro, depois o traçado e, finalmente, os marcadores.
	KSvgPaintOrderNormal SvgPaintOrder = "normal"

	// KSvgPaintOrderFill
	//
	// English:
	//
	// The order of these three keywords indicates the order in which the painting happens, from left to right. If any of
	// the three painting components is omitted, they will be painted in their default order after the specified
	// components.
	//
	// For example, using stroke is equivalent to stroke fill markers.
	//
	// Português:
	//
	// A ordem dessas três palavras-chave indica a ordem em que a pintura acontece, da esquerda para a direita. Se algum
	// dos três componentes de pintura for omitido, eles serão pintados em sua ordem padrão após os componentes
	// especificados.
	//
	// Por exemplo, usar traço é equivalente a marcadores de preenchimento de traço.
	KSvgPaintOrderFill SvgPaintOrder = "fill"

	// KSvgPaintOrderStroke
	//
	// English:
	//
	// The order of these three keywords indicates the order in which the painting happens, from left to right. If any of
	// the three painting components is omitted, they will be painted in their default order after the specified
	// components.
	//
	// For example, using stroke is equivalent to stroke fill markers.
	//
	// Português:
	//
	// A ordem dessas três palavras-chave indica a ordem em que a pintura acontece, da esquerda para a direita. Se algum
	// dos três componentes de pintura for omitido, eles serão pintados em sua ordem padrão após os componentes
	// especificados.
	//
	// Por exemplo, usar traço é equivalente a marcadores de preenchimento de traço.
	KSvgPaintOrderStroke SvgPaintOrder = "stroke"

	// KSvgPaintOrderMarkers
	//
	// English:
	//
	// The order of these three keywords indicates the order in which the painting happens, from left to right. If any of
	// the three painting components is omitted, they will be painted in their default order after the specified
	// components.
	//
	// For example, using stroke is equivalent to stroke fill markers.
	//
	// Português:
	//
	// A ordem dessas três palavras-chave indica a ordem em que a pintura acontece, da esquerda para a direita. Se algum
	// dos três componentes de pintura for omitido, eles serão pintados em sua ordem padrão após os componentes
	// especificados.
	//
	// Por exemplo, usar traço é equivalente a marcadores de preenchimento de traço.
	KSvgPaintOrderMarkers SvgPaintOrder = "markers"
)
