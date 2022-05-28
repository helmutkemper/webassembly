package html

type SvgClipPathUnits string

func (e SvgClipPathUnits) String() string {
	return string(e)
}

const (
	// KSvgClipPathUnitsUserSpaceOnUse
	//
	// English:
	//
	//  This value indicates that all coordinates inside the <clipPath> element refer to the user coordinate system as
	//  defined when the clipping path was created.
	//
	// Português:
	//
	//  Este valor indica que todas as coordenadas dentro do elemento <clipPath> referem-se ao sistema de coordenadas do
	//  usuário conforme definido quando o caminho de recorte foi criado.
	KSvgClipPathUnitsUserSpaceOnUse SvgClipPathUnits = "userSpaceOnUse"

	// KSvgClipPathUnitsObjectBoundingBox
	//
	// English:
	//
	//  This value indicates that all coordinates inside the <clipPath> element are relative to the bounding box of the
	//  element the clipping path is applied to. It means that the origin of the coordinate system is the top left corner
	//  of the object bounding box and the width and height of the object bounding box are considered to have a length of
	//  1 unit value.
	//
	// Português:
	//
	//  Esse valor indica que todas as coordenadas dentro do elemento <clipPath> são relativas à caixa delimitadora do
	//  elemento ao qual o caminho de recorte é aplicado. Isso significa que a origem do sistema de coordenadas é o canto
	//  superior esquerdo da caixa delimitadora do objeto e a largura e a altura da caixa delimitadora do objeto são
	//  consideradas como tendo um comprimento de 1 valor unitário.
	KSvgClipPathUnitsObjectBoundingBox SvgClipPathUnits = "objectBoundingBox"
)
