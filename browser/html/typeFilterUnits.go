package html

type SvgFilterUnits string

func (e SvgFilterUnits) String() string {
	return string(e)
}

const (

	// KSvgFilterUnitsUserSpaceOnUse
	//
	// English:
	//
	//  x, y, width and height represent values in the current coordinate system that results from taking the current user
	//  coordinate system in place at the time when the <filter> element is referenced (i.e., the user coordinate system
	//  for the element referencing the <filter> element via a filter attribute).
	//
	// Portuguese
	//
	//  x, y, largura e altura representam valores no sistema de coordenadas atual que resulta da tomada do sistema de
	//  coordenadas do usuário atual no momento em que o elemento <filter> é referenciado (ou seja, o sistema de
	//  coordenadas do usuário para o elemento que faz referência ao <filter > elemento através de um atributo de filtro).
	KSvgFilterUnitsUserSpaceOnUse SvgFilterUnits = "userSpaceOnUse"

	// KSvgFilterUnitsObjectBoundingBox
	//
	// English:
	//
	//  In that case, x, y, width and height represent fractions or percentages of the bounding box on the referencing
	//  element.
	//
	// Portuguese
	//
	//  Nesse caso, x, y, largura e altura representam frações ou porcentagens da caixa delimitadora no elemento de
	//  referência.
	KSvgFilterUnitsObjectBoundingBox SvgFilterUnits = "objectBoundingBox"
)
