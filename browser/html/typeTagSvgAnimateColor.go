package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"image/color"
	"strconv"
	"sync"
	"syscall/js"
	"time"
)

// TagSvgAnimateColor
//
// English:
//
// The SVG <animate> element provides a way to animate an attribute of an element over time.
//
// Português:
//
// O elemento SVG <animate> fornece uma maneira de animar um atributo de um elemento ao longo do tempo.
type TagSvgAnimateColor struct {

	// id
	//
	// English:
	//
	//  Unique id, standard html id property.
	//
	// Português:
	//
	//  Id único, propriedade padrão id do html.
	id string

	// selfElement
	//
	// English:
	//
	//  Reference to self element as js.Value.
	//
	// Português:
	//
	//  Referencia ao próprio elemento na forma de js.Value.
	selfElement js.Value

	cssClass *css.Class

	x int
	y int

	// listener
	//
	// English:
	//
	//  The javascript function removeEventListener needs to receive the function passed in addEventListener
	//
	// Português:
	//
	//  A função javascript removeEventListener necessitam receber a função passada em addEventListener
	listener *sync.Map

	// drag

	// stage
	//
	// English:
	//
	//  Browser main document reference captured at startup.
	//
	// Português:
	//
	//  Referencia do documento principal do navegador capturado na inicialização.
	stage js.Value

	// isDragging
	//
	// English:
	//
	//  Indicates the process of dragging the element.
	//
	// Português:
	//
	//  Indica o processo de arrasto do elemento.
	isDragging bool

	// dragDifX
	//
	// English:
	//
	//  Used in calculating element drag.
	//
	// Português:
	//
	//  Usado no cálculo do arrasto de elemento.
	dragDifX int

	// dragDifX
	//
	// English:
	//
	//  Used in calculating element drag.
	//
	// Português:
	//
	//  Usado no cálculo do arrasto de elemento.
	dragDifY int

	// deltaMovieX
	//
	// English:
	//
	//  Additional value added in the SetX() function: (x = x + deltaMovieX) and subtracted in the
	//  GetX() function: (x = x - deltaMovieX).
	//
	// Português:
	//
	//  Valor adicional adicionado na função SetX(): (x = x + deltaMovieX)  e subtraído na função
	//  GetX(): (x = x - deltaMovieX).
	deltaMovieX int

	// deltaMovieY
	//
	// English:
	//
	//  Additional value added in the SetY() function: (y = y + deltaMovieY) and subtracted in the
	//  GetY() function: (y = y - deltaMovieY).
	//
	// Português:
	//
	//  Valor adicional adicionado na função SetY(): (y = y + deltaMovieY)  e subtraído na função
	//  GetY(): (y = y - deltaMovieY).
	deltaMovieY int

	// tween
	//
	// English:
	//
	//  Easing tween.
	//
	// Receives an identifier and a pointer of the tween object to be used in case of multiple
	// functions.
	//
	// Português:
	//
	//  Facilitador de interpolação.
	//
	// Recebe um identificador e um ponteiro do objeto tween para ser usado em caso de múltiplas
	// funções.
	tween map[string]interfaces.TweenInterface

	points    *[]algorithm.Point
	pointsLen int

	rotateDelta float64
}

// Accumulate
//
// English:
//
//  The accumulate attribute controls whether or not an animation is cumulative.
//
//   Input:
//     KSvgAccumulateSum: Specifies that each repeat iteration after the first builds upon the last value of the
//       previous iteration;
//     KSvgAccumulateNone: Specifies that repeat iterations are not cumulative.
//
// It is frequently useful for repeated animations to build upon the previous results, accumulating with each iteration.
// This attribute said to the animation if the value is added to the previous animated attribute's value on each
// iteration.
//
//   Notes:
//     * This attribute is ignored if the target attribute value does not support addition, or if the animation element
//       does not repeat;
//     * This attribute will be ignored if the animation function is specified with only the to attribute.
//
// Português
//
//  O atributo acumular controla se uma animação é cumulativa ou não.
//
//   Input:
//     KSvgAccumulateSum: Especifica que cada iteração repetida após a primeira se baseia no último valor da iteração
//       anterior;
//     KSvgAccumulateNone: Especifica que as iterações repetidas não são cumulativas.
//
// Frequentemente, é útil que as animações repetidas se baseiem nos resultados anteriores, acumulando a cada iteração.
// Este atributo é dito à animação se o valor for adicionado ao valor do atributo animado anterior em cada iteração.
//
//   Notas:
//     * Esse atributo será ignorado se o valor do atributo de destino não suportar adição ou se o elemento de animação
//       não se repetir;
//     * Este atributo será ignorado se a função de animação for especificada apenas com o atributo to.
func (e *TagSvgAnimateColor) Accumulate(accumulate SvgAccumulate) (ref *TagSvgAnimateColor) {
	e.selfElement.Call("setAttribute", "accumulate", accumulate.String())
	return e
}

// Additive
//
// English:
//
//  The additive attribute controls whether or not an animation is additive.
//
//   Input:
//     KSvgAdditiveSum: Specifies that the animation will add to the underlying value of the attribute and other
//       lower priority animations.
//     KSvgAdditiveReplace: (default) Specifies that the animation will override the underlying value of the attribute
//       and other lower priority animations.
//
// It is frequently useful to define animation as an offset or delta to an attribute's value, rather than as
// absolute values.
//
// Português
//
//  O atributo aditivo controla se uma animação é ou não aditiva.
//
//   Input:
//     KSvgAdditiveSum: Especifica que a animação será adicionada ao valor subjacente do atributo e outras animações de
//       prioridade mais baixa.
//     KSvgAdditiveReplace: (default) Especifica que a animação substituirá o valor subjacente do atributo e outras
//       animações de prioridade mais baixa.
//
// É frequentemente útil definir a animação como um deslocamento ou delta para o valor de um atributo, em vez de
// valores absolutos.
func (e *TagSvgAnimateColor) Additive(additive SvgAdditive) (ref *TagSvgAnimateColor) {
	e.selfElement.Call("setAttribute", "additive", additive.String())
	return e
}

// AttributeName
//
// English:
//
//  The attributeName attribute indicates the name of the CSS property or attribute of the target element that is going
//  to be changed during an animation.
//
// Português
//
//  O atributo attributeName indica o nome da propriedade CSS ou atributo do elemento de destino que será alterado
//  durante uma animação.
func (e *TagSvgAnimateColor) AttributeName(attributeName string) (ref *TagSvgAnimateColor) {
	e.selfElement.Call("setAttribute", "attributeName", attributeName)
	return e
}

// Begin
//
// English:
//
//  The begin attribute defines when an animation should begin or when an element should be discarded.
//
//   Input:
//     begin: defines when an animation should begin or when an element should be discarded.
//       offset-value: This value defines a clock-value that represents a point in time relative to the beginning of the
//         SVG document (usually the load or DOMContentLoaded event). Negative values are valid.
//         (e.g. time.Second*5 or "5s")
//       syncbase-value: This value defines a syncbase and an optional offset from that syncbase. The element's
//         animation start time is defined relative to the begin or active end of another animation.
//         A valid syncbase-value consists of an ID reference to another animation element followed by a dot and either
//         begin or end to identify whether to synchronize with the beginning or active end of the referenced animation
//         element. An optional offset value as defined in <offset-value> can be appended.
//         (e.g. "0s;third.end", "first.end" or "second.end")
//       event-value: This value defines an event and an optional offset that determines the time at which the element's
//         animation should begin. The animation start time is defined relative to the time that the specified event is
//         fired.
//         A valid event-value consists of an element ID followed by a dot and one of the supported events for that
//         element. All valid events (not necessarily supported by all elements) are defined by the DOM and HTML
//         specifications. Those are: 'focus', 'blur', 'focusin', 'focusout', 'activate', 'auxclick', 'click',
//         'dblclick', 'mousedown', 'mouseenter', 'mouseleave', 'mousemove', 'mouseout', 'mouseover', 'mouseup',
//         'wheel', 'beforeinput', 'input', 'keydown', 'keyup', 'compositionstart', 'compositionupdate',
//         'compositionend', 'load', 'unload', 'abort', 'error', 'select', 'resize', 'scroll', 'beginEvent', 'endEvent',
//         and 'repeatEvent'. An optional offset value as defined in <offset-value> can be appended.
//         (e.g. "startButton.click")
//       repeat-value: This value defines a qualified repeat event. The element animation start time is defined relative
//         to the time that the repeat event is raised with the specified iteration value.
//         A valid repeat value consists of an element ID followed by a dot and the function repeat() with an integer
//         value specifying the number of repetitions as parameter. An optional offset value as defined in
//         <offset-value> can be appended.
//         (e.g. "0s;myLoop.end", "myLoop.begin", "myLoop.repeat(1)" or "myLoop.repeat(2)")
//       accessKey-value: This value defines an access key that should trigger the animation. The element animation will
//         begin when the user presses the specified key.
//         A valid accessKey-value consists of the function accessKey() with the character to be input as parameter. An
//         optional offset value as defined in <offset-value> can be appended.
//         (e.g. "accessKey(s)")
//       wallclock-sync-value: This value defines the animation start time as a real-world clock time.
//         A valid wallclock-sync-value consists of the function wallclock() with a time value as parameter. The time
//         syntax is based upon the syntax defined in ISO 8601.
//         (e.g. time.Now() )
//       indefinite: The begin of the animation will be determined by a beginElement() method call or a hyperlink
//         targeted to the element.
//         (e.g. "infinite")
//
// The attribute value is a semicolon separated list of values. The interpretation of a list of start times is detailed
// in the SMIL specification in "Evaluation of begin and end time lists". Each individual value can be one of the
// following: <offset-value>, <syncbase-value>, <event-value>, <repeat-value>, <accessKey-value>, <wallclock-sync-value>
// or the keyword 'indefinite'.
//
// Português:
//
//  O atributo begin define quando uma animação deve começar ou quando um elemento deve ser descartado.
//
//   Entrada:
//     begin: define quando uma animação deve começar ou quando um elemento deve ser descartado.
//       offset-value: Esse valor define um valor de relógio que representa um ponto no tempo relativo ao início do
//         documento SVG (geralmente o evento load ou DOMContentLoaded). Valores negativos são válidos.
//         (e.g. time.Second*5 or "5s")
//       syncbase-value: Esse valor define uma base de sincronização e um deslocamento opcional dessa base de
//         sincronização. A hora de início da animação do elemento é definida em relação ao início ou fim ativo de outra
//         animação.
//         Um valor syncbase válido consiste em uma referência de ID para outro elemento de animação seguido por um
//         ponto e um início ou fim para identificar se deve ser sincronizado com o início ou o final ativo do elemento
//         de animação referenciado. Um valor de deslocamento opcional conforme definido em <offset-value> pode ser
//         anexado.
//         (e.g. "0s;third.end", "first.end" or "second.end")
//       event-value: Esse valor define um evento e um deslocamento opcional que determina a hora em que a animação do
//         elemento deve começar. A hora de início da animação é definida em relação à hora em que o evento especificado
//         é acionado.
//         Um valor de evento válido consiste em um ID de elemento seguido por um ponto e um dos eventos com suporte
//         para esse elemento. Todos os eventos válidos (não necessariamente suportados por todos os elementos) são
//         definidos pelas especificações DOM e HTML. Esses valores são: 'focus', 'blur', 'focusin', 'focusout',
//         'activate', 'auxclick', 'click', 'dblclick', 'mousedown', 'mouseenter', 'mouseleave', 'mousemove',
//         'mouseout', 'mouseover', 'mouseup', 'wheel', 'beforeinput', 'input', 'keydown', 'keyup', 'compositionstart',
//         'compositionupdate', 'compositionend', 'load', 'unload', 'abort', 'error', 'select', 'resize', 'scroll',
//         'beginEvent', 'endEvent', e 'repeatEvent'. Um valor de deslocamento opcional conforme definido em
//         <offset-value> pode ser anexado.
//         (e.g. "startButton.click")
//       repeat-value: Esse valor define um evento de repetição qualificado. A hora de início da animação do elemento é
//         definida em relação à hora em que o evento de repetição é gerado com o valor de iteração especificado.
//         Um valor de repetição válido consiste em um ID de elemento seguido por um ponto e a função repeat() com um
//         valor inteiro especificando o número de repetições como parâmetro. Um valor de deslocamento opcional conforme
//         definido em <offset-value> pode ser anexado.
//         (e.g. "0s;myLoop.end", "myLoop.begin", "myLoop.repeat(1)" or "myLoop.repeat(2)")
//       accessKey-value: Este valor define uma chave de acesso que deve acionar a animação. A animação do elemento
//         começará quando o usuário pressionar a tecla especificada.
//         Um valor válido de accessKey consiste na função accessKey() com o caractere a ser inserido como parâmetro.
//         Um valor de deslocamento opcional conforme definido em <valor de deslocamento> pode ser anexado.
//         (e.g. "accessKey(s)")
//       wallclock-sync-value: Esse valor define a hora de início da animação como uma hora do relógio do mundo real.
//         Um valor wallclock-sync válido consiste na função wallclock() com um valor de tempo como parâmetro. A sintaxe
//         de tempo é baseada na sintaxe definida na ISO 8601.
//         (e.g. time.Now() )
//       indefinite: O início da animação será determinado por uma chamada de método beginElement() ou um hiperlink
//         direcionado ao elemento.
//         (e.g. "infinite")
//
// O valor do atributo é uma lista de valores separados por ponto e vírgula. A interpretação de uma lista de horários de
// início é detalhada na especificação SMIL em "Avaliação de listas de horários de início e término". Cada valor
// individual pode ser um dos seguintes: <offset-value>, <syncbase-value>, <event-value>, <repeat-value>,
// <accessKey-value>, <wallclock-sync-value> ou a palavra-chave 'indefinite'.
func (e *TagSvgAnimateColor) Begin(begin interface{}) (ref *TagSvgAnimateColor) {
	if converted, ok := begin.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "begin", converted.String())
		return e
	}

	if converted, ok := begin.(time.Time); ok {
		e.selfElement.Call("setAttribute", "begin", converted.String())
		return e
	}

	if converted, ok := begin.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "baseline-shift", p)
		return e
	}

	e.selfElement.Call("setAttribute", "begin", begin)
	return e
}

// By
//
// English:
//
//  The by attribute specifies a relative offset value for an attribute that will be modified during an animation.
//
//   Input:
//     by: specifies a relative offset value for an attribute
//
// The starting value for the attribute is either indicated by specifying it as value for the attribute given in the
// attributeName or the from attribute.
//
// Português:
//
//  O atributo by especifica um valor de deslocamento relativo para um atributo que será modificado durante uma
//  animação.
//
//   Entrada:
//     by: especifica um valor de deslocamento relativo para um atributo
//
// O valor inicial para o atributo é indicado especificando-o como valor para o atributo fornecido no attributeName ou
// no atributo from.
func (e *TagSvgAnimateColor) By(by float64) (ref *TagSvgAnimateColor) {
	e.selfElement.Call("setAttribute", "by", by)
	return e
}

// CalcMode
//
// English:
//
//  The calcMode attribute specifies the interpolation mode for the animation.
//
//   Input:
//     KSvgCalcModeDiscrete: This specifies that the animation function will jump from one value to the next without
//       any interpolation.
//     KSvgCalcModeLinear: Simple linear interpolation between values is used to calculate the animation function.
//       Except for <animateMotion>, this is the default value.
//     KSvgCalcModePaced: Defines interpolation to produce an even pace of change across the animation.
//     KSvgCalcModeSpline: Interpolates from one value in the values list to the next according to a time function
//       defined by a cubic Bézier spline. The points of the spline are defined in the keyTimes attribute, and the
//       control points for each interval are defined in the keySplines attribute.
//
// The default mode is linear, however if the attribute does not support linear interpolation (e.g. for strings), the
// calcMode attribute is ignored and discrete interpolation is used.
//
//   Notes:
//     Default value: KSvgCalcModePaced
//
// Português
//
//  O atributo calcMode especifica o modo de interpolação para a animação.
//
//   Entrada:
//     KSvgCalcModeDiscrete: Isso especifica que a função de animação saltará de um valor para o próximo sem qualquer
//       interpolação.
//     KSvgCalcModeLinear: A interpolação linear simples entre valores é usada para calcular a função de animação.
//       Exceto para <animateMotion>, este é o valor padrão.
//     KSvgCalcModePaced: Define a interpolação para produzir um ritmo uniforme de mudança na animação.
//     KSvgCalcModeSpline: Interpola de um valor na lista de valores para o próximo de acordo com uma função de tempo
//       definida por uma spline de Bézier cúbica. Os pontos do spline são definidos no atributo keyTimes e os pontos
//       de controle para cada intervalo são definidos no atributo keySplines.
//
// O modo padrão é linear, no entanto, se o atributo não suportar interpolação linear (por exemplo, para strings), o
// atributo calcMode será ignorado e a interpolação discreta será usada.
//
//   Notas:
//     * Valor padrão: KSvgCalcModePaced
func (e *TagSvgAnimateColor) CalcMode(calcMode SvgCalcMode) (ref *TagSvgAnimateColor) {
	e.selfElement.Call("setAttribute", "calcMode", calcMode.String())
	return e
}

// ColorInterpolation
//
// English:
//
//  The color-interpolation attribute specifies the color space for gradient interpolations, color animations, and alpha
//  compositing.
//
// The color-interpolation property chooses between color operations occurring in the sRGB color space or in a (light
// energy linear) linearized RGB color space. Having chosen the appropriate color space, component-wise linear
// interpolation is used.
//
// When a child element is blended into a background, the value of the color-interpolation property on the child
// determines the type of blending, not the value of the color-interpolation on the parent.
// For gradients which make use of the href or the deprecated xlink:href attribute to reference another gradient, the
// gradient uses the property's value from the gradient element which is directly referenced by the fill or stroke
// property. When animating colors, color interpolation is performed according to the value of the color-interpolation
// property on the element being animated.
//
//   Notes:
//     * For filter effects, the color-interpolation-filters property controls which color space is used.
//     * As a presentation attribute, color-interpolation can be used as a CSS property.
//
// Português:
//
//  O atributo color-interpolation especifica o espaço de cores para interpolações de gradiente, animações de cores e
//  composição alfa.
//
// A propriedade de interpolação de cores escolhe entre operações de cores que ocorrem no espaço de cores sRGB ou em um
// espaço de cores RGB linearizado (energia de luz linear). Tendo escolhido o espaço de cor apropriado, a interpolação
// linear de componentes é usada.
//
// Quando um elemento filho é mesclado em um plano de fundo, o valor da propriedade color-interpolation no filho
// determina o tipo de mesclagem, não o valor da interpolação de cores no pai.
// Para gradientes que usam o href ou o atributo obsoleto xlink:href para referenciar outro gradiente, o gradiente usa
// o valor da propriedade do elemento gradiente que é diretamente referenciado pela propriedade fill ou stroke.
// Ao animar cores, à interpolação de cores é executada de acordo com o valor da propriedade color-interpolation no
// elemento que está sendo animado.
//
//   Notas:
//     * Para efeitos de filtro, a propriedade color-interpolation-filters controla qual espaço de cor é usado.
//     * Como atributo de apresentação, a interpolação de cores pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateColor) ColorInterpolation(value interface{}) (ref *TagSvgAnimateColor) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color-interpolation", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color-interpolation", value)
	return e
}
