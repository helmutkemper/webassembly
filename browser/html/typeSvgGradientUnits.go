package html

type SvgGradientUnits string

func (e SvgGradientUnits) String() string {
	return string(e)
}

const (

	// KSvgGradientUnitsUserSpaceOnUse
	//
	// English:
	//
	//  This value indicates that the attributes represent values in the coordinate system that results from taking the
	//  current user coordinate system in place at the time when the gradient element is referenced (i.e., the user
	//  coordinate system for the element referencing the gradient element via a fill or stroke property) and then
	//  applying the transform specified by attribute gradientTransform. Percentages represent values relative to the
	//  current SVG viewport.
	//
	// Portuguese
	//
	//  Este valor indica que os atributos representam valores no sistema de coordenadas que resulta da tomada do sistema
	//  de coordenadas do usuário atual no momento em que o elemento gradiente é referenciado (ou seja, o sistema de
	//  coordenadas do usuário para o elemento referenciando o elemento gradiente por meio de um preenchimento ou stroke
	//  property) e, em seguida, aplicar a transformação especificada pelo atributo gradientTransform.
	//  As porcentagens representam valores relativos à janela de visualização SVG atual.
	KSvgGradientUnitsUserSpaceOnUse SvgGradientUnits = "userSpaceOnUse"

	// KSvgGradientUnitsObjectBoundingBox
	//
	// English:
	//
	//  This value indicates that the user coordinate system for the attributes is established using the bounding box of
	//  the element to which the gradient is applied and then applying the transform specified by attribute
	//  gradientTransform.
	//
	// Percentages represent values relative to the bounding box for the object.
	//
	// With this value and gradientTransform being the identity matrix, the normal of the linear gradient is
	// perpendicular to the gradient vector in object bounding box space (i.e., the abstract coordinate system where
	// (0,0) is at the top/left of the object bounding box and (1,1) is at the bottom/right of the object bounding box).
	// When the object's bounding box is not square, the gradient normal which is initially perpendicular to the gradient
	// vector within object bounding box space may render non-perpendicular relative to the gradient vector in user space.
	// If the gradient vector is parallel to one of the axes of the bounding box, the gradient normal will remain
	// perpendicular.
	// This transformation is due to application of the non-uniform scaling transformation from bounding box space to
	// user space.
	//
	// Portuguese
	//
	//  Esse valor indica que o sistema de coordenadas do usuário para os atributos é estabelecido usando a caixa
	//  delimitadora do elemento ao qual o gradiente é aplicado e, em seguida, aplicando a transformação especificada
	//  pelo atributo gradientTransform.
	//
	// As porcentagens representam valores relativos à caixa delimitadora do objeto.
	//
	// Com este valor e gradientTransform sendo a matriz identidade, a normal do gradiente linear é perpendicular ao
	// vetor gradiente no espaço da caixa delimitadora do objeto (ou seja, o sistema de coordenadas abstrato onde (0,0)
	// está no canto superior esquerdo da caixa delimitadora do objeto e (1,1) está na parte inferior direita da caixa
	// delimitadora do objeto). Quando a caixa delimitadora do objeto não é quadrada, o gradiente normal que é
	// inicialmente perpendicular ao vetor gradiente dentro do espaço da caixa delimitadora do objeto pode tornar-se não
	// perpendicular em relação ao vetor gradiente no espaço do usuário. Se o vetor gradiente for paralelo a um dos eixos
	// da caixa delimitadora, a normal do gradiente permanecerá perpendicular. Essa transformação se deve à aplicação da
	// transformação de dimensionamento não uniforme do espaço da caixa delimitadora para o espaço do usuário.
	KSvgGradientUnitsObjectBoundingBox SvgGradientUnits = "objectBoundingBox"
)
