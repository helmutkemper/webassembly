package html

type SvgVectorEffect string

func (e SvgVectorEffect) String() string {
	return string(e)
}

const (
	// KSvgVectorEffectNone
	//
	// English:
	//
	// This value specifies that no vector effect shall be applied, i.e. the default rendering behavior is used which is
	// to first fill the geometry of a shape with a specified paint, then stroke the outline with a specified paint.
	//
	// Português:
	//
	// Este valor especifica que nenhum efeito vetorial deve ser aplicado, ou seja, o comportamento de renderização padrão
	// é usado, que é primeiro preencher a geometria de uma forma com uma tinta especificada e, em seguida, traçar o
	// contorno com uma tinta especificada.
	KSvgVectorEffectNone SvgVectorEffect = "none"

	// KSvgVectorEffectNonScalingStroke
	//
	// English:
	//
	// This value modifies the way an object is stroked. Normally stroking involves calculating stroke outline of the
	// shape's path in current user coordinate system and filling that outline with the stroke paint (color or gradient).
	// The resulting visual effect of this value is that the stroke width is not dependent on the transformations of the
	// element (including non-uniform scaling and shear transformations) and zoom level.
	//
	// Português:
	//
	// Esse valor modifica a maneira como um objeto é traçado. Normalmente, o traçado envolve o cálculo do contorno do
	// traçado do caminho da forma no sistema de coordenadas do usuário atual e o preenchimento desse contorno com a
	// pintura do traçado (cor ou gradiente).
	// O efeito visual resultante desse valor é que a largura do traço não depende das transformações do elemento
	// (incluindo escala não uniforme e transformações de cisalhamento) e do nível de zoom.
	KSvgVectorEffectNonScalingStroke SvgVectorEffect = "non-scaling-stroke"

	// KSvgVectorEffectNonScalingSize
	//
	// English:
	//
	// This value specifies a special user coordinate system used by the element and its descendants.
	// The scale of that user coordinate system does not change in spite of any transformation changes from a host
	// coordinate space.
	// However, it does not specify the suppression of rotation and skew. Also, it does not specify the origin of the user
	// coordinate system. Since this value suppresses scaling of the user coordinate system, it also has the
	// characteristics of non-scaling-stroke.
	//
	// Português:
	//
	// Este valor especifica um sistema de coordenadas do usuário especial usado pelo elemento e seus descendentes.
	// A escala desse sistema de coordenadas do usuário não muda apesar de qualquer mudança de transformação de um espaço
	// de coordenadas do hospedeiro. No entanto, não especifica a supressão de rotação e inclinação.
	// Além disso, não especifica a origem do sistema de coordenadas do usuário. Como este valor suprime a escala do
	// sistema de coordenadas do usuário, ele também possui as características de curso sem escala.
	KSvgVectorEffectNonScalingSize SvgVectorEffect = "non-scaling-size"

	// KSvgVectorEffectNonRotation
	//
	// English:
	//
	// This value specifies a special user coordinate system used by the element and its descendants.
	// The rotation and skew of that user coordinate system is suppressed in spite of any transformation changes from a
	// host coordinate space. However, it does not specify the suppression of scaling. Also, it does not specify the
	// origin of user coordinate system.
	//
	// Português:
	//
	// Este valor especifica um sistema de coordenadas do usuário especial usado pelo elemento e seus descendentes.
	// A rotação e a inclinação desse sistema de coordenadas do usuário são suprimidas apesar de quaisquer alterações de
	// transformação de um espaço de coordenadas do host. No entanto, ele não especifica a supressão de dimensionamento.
	// Além disso, não especifica a origem do sistema de coordenadas do usuário.
	KSvgVectorEffectNonRotation SvgVectorEffect = "non-rotation"

	// KSvgVectorEffectFixedPosition
	//
	// English:
	//
	// This value specifies a special user coordinate system used by the element and its descendants. The position of user
	// coordinate system is fixed in spite of any transformation changes from a host coordinate space.
	// However, it does not specify the suppression of rotation, skew and scaling. When this vector effect and the
	// transform property are defined at the same time, that property is consumed for this effect.
	//
	// Português:
	//
	// Este valor especifica um sistema de coordenadas do usuário especial usado pelo elemento e seus descendentes.
	// A posição do sistema de coordenadas do usuário é fixa apesar de quaisquer mudanças de transformação de um espaço de
	// coordenadas do hospedeiro. No entanto, não especifica a supressão de rotação, inclinação e dimensionamento.
	// Quando esse efeito vetorial e a propriedade de transformação são definidos ao mesmo tempo, essa propriedade é
	// consumida para esse efeito.
	KSvgVectorEffectFixedPosition SvgVectorEffect = "fixed-position"
)
