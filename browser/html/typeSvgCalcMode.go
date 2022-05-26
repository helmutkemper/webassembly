package html

type SvgCalcMode string

func (e SvgCalcMode) String() string {
	return string(e)
}

const (
	// KSvgCalcModeDiscrete
	//
	// English:
	//
	//  This specifies that the animation function will jump from one value to the next without any interpolation.
	//
	// Português:
	//
	//  Isso especifica que a função de animação saltará de um valor para o próximo sem qualquer interpolação.
	KSvgCalcModeDiscrete SvgCalcMode = "discrete"

	// KSvgCalcModeLinear
	//
	// English:
	//
	//  Simple linear interpolation between values is used to calculate the animation function. Except for
	//  <animateMotion>, this is the default value.
	//
	// Português:
	//
	//  A interpolação linear simples entre valores é usada para calcular a função de animação. Exceto para
	//  <animateMotion>, este é o valor padrão.
	KSvgCalcModeLinear SvgCalcMode = "linear"

	// KSvgCalcModePaced
	//
	// English:
	//
	//  Defines interpolation to produce an even pace of change across the animation.
	//
	// This is only supported for values that define a linear numeric range, and for which some notion of "distance"
	// between points can be calculated (e.g. position, width, height, etc.).
	// If paced is specified, any keyTimes or keySplines will be ignored.
	//
	// Português:
	//
	//  Define a interpolação para produzir um ritmo uniforme de mudança na animação.
	//
	// Ele é suportado apenas para valores que definem um intervalo numérico linear e para os quais alguma noção de
	// "distância" entre pontos pode ser calculada (por exemplo, posição, largura, altura etc.).
	// Se paced for especificado, quaisquer keyTimes ou keySplines serão ignorados.
	KSvgCalcModePaced SvgCalcMode = "paced"

	// KSvgCalcModeSpline
	//
	// English:
	//
	//  Interpolates from one value in the values list to the next according to a time function defined by a cubic Bézier
	//  spline. The points of the spline are defined in the keyTimes attribute, and the control points for each interval
	//  are defined in the keySplines attribute.
	//
	// Português:
	//
	//  Interpola de um valor na lista de valores para o próximo de acordo com uma função de tempo definida por uma spline
	//  de Bézier cúbica. Os pontos do spline são definidos no atributo keyTimes e os pontos de controle para cada
	//  intervalo são definidos no atributo keySplines.
	KSvgCalcModeSpline SvgCalcMode = "spline"
)
