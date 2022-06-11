package html

type SvgMaskUnits string

func (e SvgMaskUnits) String() string {
	return string(e)
}

const (
	// KSvgMaskUnitsUserSpaceOnUse
	//
	// English:
	//
	// This value indicates that all coordinates inside the <mask> element refer to the user coordinate system as defined
	// when the mask was created.
	//
	// Português:
	//
	// Este valor indica que todas as coordenadas dentro do elemento <mask> referem-se ao sistema de coordenadas do
	// usuário conforme definido quando a máscara foi criada.
	KSvgMaskUnitsUserSpaceOnUse SvgMaskUnits = "userSpaceOnUse"

	// KSvgMaskUnitsObjectBoundingBox
	//
	// English:
	//
	// This value indicates that all coordinates inside the <mask> element are relative to the bounding box of the element
	// the mask is applied to. A bounding box could be considered the same as if the content of the <mask> were bound to a
	// "0 0 1 1" viewbox.
	//
	// Português:
	//
	// Esse valor indica que todas as coordenadas dentro do elemento <mask> são relativas à caixa delimitadora do elemento
	// ao qual a máscara é aplicada. Uma caixa delimitadora pode ser considerada como se o conteúdo da <mask>
	// estivesse vinculado a uma caixa de visualização "0 0 1 1".
	KSvgMaskUnitsObjectBoundingBox SvgMaskUnits = "objectBoundingBox"
)
