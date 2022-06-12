package html

type SvgUnits string

func (e SvgUnits) String() string {
	return string(e)
}

const (
	// KSvgUnitsUserSpaceOnUse
	//
	// English:
	//
	// This value indicates that all coordinates for the geometry properties refer to the user coordinate system as
	// defined when the pattern was applied.
	//
	// Português:
	//
	// Este valor indica que todas as coordenadas para as propriedades de geometria referem-se ao sistema de coordenadas
	// do usuário conforme definido quando o padrão foi aplicado.
	KSvgUnitsUserSpaceOnUse SvgUnits = "userSpaceOnUse"

	// KSvgUnitsObjectBoundingBox
	//
	// English:
	//
	// This value indicates that all coordinates for the geometry properties represent fractions or percentages of the
	// bounding box of the element to which the pattern is applied. A bounding box could be considered the same as if the
	// content of the <pattern> were bound to a "0 0 1 1" viewbox.
	//
	// Português:
	//
	// Este valor indica que todas as coordenadas das propriedades geométricas representam frações ou porcentagens da
	// caixa delimitadora do elemento ao qual o padrão é aplicado. Uma caixa delimitadora pode ser considerada como se o
	// conteúdo do <pattern> estivesse vinculado a uma caixa de visualização "0 0 1 1".
	KSvgUnitsObjectBoundingBox SvgUnits = "objectBoundingBox"
)
