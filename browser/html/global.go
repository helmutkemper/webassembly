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

type TagSvgGlobal struct {

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
func (e *TagSvgGlobal) Accumulate(accumulate SvgAccumulate) (ref *TagSvgGlobal) {
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
func (e *TagSvgGlobal) Additive(additive SvgAdditive) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "additive", additive.String())
	return e
}

// AlignmentBaseline
//
// English:
//
//  The alignment-baseline attribute specifies how an object is aligned with respect to its parent. This property specifies which baseline of this element is to be aligned with the corresponding baseline of the parent. For example, this allows alphabetic baselines in Roman text to stay aligned across font size changes. It defaults to the baseline with the same name as the computed value of the alignment-baseline property.
//
//   Input:
//     alignmentBaseline: specifies how an object is aligned with respect to its parent.
//       string: url(#myClip)
//       consts KSvgAlignmentBaseline... (e.g. KSvgAlignmentBaselineTextBeforeEdge)
//
//   Notes:
//     * As a presentation attribute alignment-baseline can be used as a CSS property.
//     * SVG 2 introduces some changes to the definition of this property. In particular: the values auto, before-edge, and after-edge have been removed. For backwards compatibility, text-before-edge may be mapped to text-top and text-after-edge to text-bottom. Neither text-before-edge nor text-after-edge should be used with the vertical-align property.
//
// Português:
//
//  O atributo alinhamento-base especifica como um objeto é alinhado em relação ao seu pai. Esta propriedade especifica qual linha de base deste elemento deve ser alinhada com a linha de base correspondente do pai. Por exemplo, isso permite que as linhas de base alfabéticas em texto romano permaneçam alinhadas nas alterações de tamanho da fonte. O padrão é a linha de base com o mesmo nome que o valor calculado da propriedade de linha de base de alinhamento.
//
//   Input:
//     alignmentBaseline: especifica como um objeto é alinhado em relação ao seu pai.
//       string: url(#myClip)
//       consts KSvgAlignmentBaseline...  (ex. KSvgAlignmentBaselineTextBeforeEdge)
//
//   Notas:
//     * Como um atributo de apresentação, a linha de base de alinhamento pode ser usada como uma propriedade CSS.
//     * O SVG 2 introduz algumas mudanças na definição desta propriedade. Em particular: os valores auto, before-edge e after-edge foram removidos. Para compatibilidade com versões anteriores, text-before-edge pode ser mapeado para text-top e text-after-edge para text-bottom. Nem text-before-edge nem text-after-edge devem ser usados com a propriedade vertical-align.
func (e *TagSvgGlobal) AlignmentBaseline(alignmentBaseline interface{}) (ref *TagSvgGlobal) {
	if converted, ok := alignmentBaseline.(SvgAlignmentBaseline); ok {
		e.selfElement.Call("setAttribute", "alignment-baseline", converted.String())
	}

	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// Amplitude
//
// English:
//
//  The amplitude attribute controls the amplitude of the gamma function of a component transfer element when its type
//  attribute is gamma.
//
//   Input:
//     amplitude: controls the amplitude of the gamma function
//
// Português:
//
//  O atributo amplitude controla à amplitude da função gama de um elemento de transferência de componente quando seu
//  atributo de tipo é gama.
//
//   Entrada:
//     amplitude: controla a amplitude da função de gama
func (e *TagSvgGlobal) Amplitude(amplitude float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "amplitude", amplitude)
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
func (e *TagSvgGlobal) AttributeName(attributeName string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "attributeName", attributeName)
	return e
}

// Azimuth
//
// English:
//
//  The azimuth attribute specifies the direction angle for the light source on the XY plane (clockwise), in degrees
//  from the x axis.
//
//   Input:
//     azimuth: specifies the direction angle for the light source on the XY plane
//
// Português:
//
//  O atributo azimute especifica o ângulo de direção da fonte de luz no plano XY (sentido horário), em graus a partir
//  do eixo x.
//
//   Input:
//     azimuth: especifica o ângulo de direção para a fonte de luz no plano XY
func (e *TagSvgGlobal) Azimuth(azimuth float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "azimuth", azimuth)
	return e
}

// BaseFrequency
//
// English:
//
//  The baseFrequency attribute represents the base frequency parameter for the noise function of the <feTurbulence> filter primitive.
//
//   Input:
//     baseFrequency: represents the base frequency parameter for the noise function
//
// Português:
//
//  O atributo baseFrequency representa o parâmetro de frequência base para a função de ruído da primitiva de filtro <feTurbulence>.
//
//   Entrada:
//     baseFrequency: representa o parâmetro de frequência base para a função de ruído
func (e *TagSvgGlobal) BaseFrequency(baseFrequency float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "baseFrequency", baseFrequency)
	return e
}

// BaselineShift
//
// English:
//
//  The baseline-shift attribute allows repositioning of the dominant-baseline relative to the dominant-baseline of the parent text content element. The shifted object might be a sub- or superscript.
//
//   Input:
//     baselineShift: allows repositioning of the dominant-baseline relative to the dominant-baseline of the parent text content element.
//       string: url(#myClip)
//       length-percentage: "5%"
//       consts KSvgBaselineShift... (e.g. KSvgBaselineShiftAuto)
//
//   Notes:
//     * As a presentation attribute baseline-shift can be used as a CSS property.
//     * This property is going to be deprecated and authors are advised to use vertical-align instead.
//
// Português:
//
//  O atributo baseline-shift permite o reposicionamento da linha de base dominante em relação à linha de base dominante do elemento de conteúdo de texto pai. O objeto deslocado pode ser um sub ou sobrescrito.
//
//   Input:
//     baselineShift: permite o reposicionamento da linha de base dominante em relação à linha de base dominante do elemento de conteúdo de texto pai.
//       string: url(#myClip)
//       length-percentage: "5%"
//       consts KSvgBaselineShift... (ex. KSvgBaselineShiftAuto)
//
//   Notas:
//     * Como atributo de apresentação, baseline-shift pode ser usado como propriedade CSS.
//     * Essa propriedade será preterida e os autores são aconselhados a usar alinhamento vertical.
func (e *TagSvgGlobal) BaselineShift(baselineShift interface{}) (ref *TagSvgGlobal) {
	if converted, ok := baselineShift.(SvgBaselineShift); ok {
		e.selfElement.Call("setAttribute", "baseline-shift", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "baseline-shift", baselineShift)
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
func (e *TagSvgGlobal) Begin(begin interface{}) (ref *TagSvgGlobal) {
	if converted, ok := begin.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "begin", converted.String())
		return e
	}

	if converted, ok := begin.(time.Time); ok {
		e.selfElement.Call("setAttribute", "begin", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "begin", begin)
	return e
}

// Bias
//
// English:
//
//  The bias attribute shifts the range of the filter. After applying the kernelMatrix of the <feConvolveMatrix> element
//  to the input image to yield a number and applied the divisor attribute, the bias attribute is added to each
//  component. This allows representation of values that would otherwise be clamped to 0 or 1.
//
//   Input:
//     bias: shifts the range of the filter
//
// Português:
//
//  O atributo bias muda o intervalo do filtro. Depois de aplicar o kernelMatrix do elemento <feConvolveMatrix> à imagem
//  de entrada para gerar um número e aplicar o atributo divisor, o atributo bias é adicionado a cada componente. Isso
//  permite a representação de valores que de outra forma seriam fixados em 0 ou 1.
//
//   Entrada:
//     bias: muda o intervalo do filtro
func (e *TagSvgGlobal) Bias(bias float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "bias", bias)
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
func (e *TagSvgGlobal) By(by float64) (ref *TagSvgGlobal) {
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
func (e *TagSvgGlobal) CalcMode(calcMode SvgCalcMode) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "calcMode", calcMode.String())
	return e
}

// Class
//
// English:
//
// Assigns a class name or set of class names to an element. You may assign the same class name or names to any number
// of elements, however, multiple class names must be separated by whitespace characters.
//
//   Input:
//     class: Assigns a class name or set of class names to an element
//
// An element's class name serves two key roles:
//   * As a style sheet selector, for when an author assigns style information to a set of elements.
//   * For general use by the browser.
//
// Português:
//
// Atribui um nome de classe ou um conjunto de nomes de classe à um elemento. Você pode atribuir o mesmo nome ou nomes
// de classe a qualquer número de elementos, no entanto, vários nomes de classe devem ser separados por caracteres de
// espaço em branco.
//
//   Entrada:
//     class: Atribui um nome de classe ou um conjunto de nomes de classe à um elemento.
//
// O nome de classe de um elemento tem duas funções principais:
//   * Como um seletor de folha de estilo, para quando um autor atribui informações de estilo a um conjunto de
//     elementos.
//   * Para uso geral pelo navegador.
func (e *TagSvgGlobal) Class(class string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "class", class)
	return e
}

// ClipPathUnits
//
// English:
//
//  The clipPathUnits attribute indicates which coordinate system to use for the contents of the <clipPath> element.
//
//   Input:
//     clipPathUnits: indicates which coordinate system to used
//       KSvgClipPathUnits... (e.g. KSvgClipPathUnitsUserSpaceOnUse)
//
// Português:
//
//  O atributo clipPathUnits indica qual sistema de coordenadas deve ser usado para o conteúdo do elemento <clipPath>.
//
//   Input:
//     clipPathUnits: indica qual sistema de coordenadas deve ser usado
//       KSvgClipPathUnits... (ex. KSvgClipPathUnitsUserSpaceOnUse)
func (e *TagSvgGlobal) ClipPathUnits(clipPathUnits SvgClipPathUnits) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "clipPathUnits", clipPathUnits.String())
	return e
}

// ClipPath
//
// English:
//
//  It binds the element it is applied to with a given <clipPath> element.
//
//   Input:
//     clipPath: the element it is applied
//       (e.g. "url(#myClip)", "circle() fill-box", "circle() stroke-box" or "circle() view-box")
//
// Português:
//
//  Ele associa o elemento ao qual é aplicado a um determinado elemento <clipPath>.
//
//   Entrada:
//     clipPath: elemento ao qual é aplicado
//       (ex. "url(#myClip)", "circle() fill-box", "circle() stroke-box" ou "circle() view-box")
func (e *TagSvgGlobal) ClipPath(clipPath string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "clip-path", clipPath)
	return e
}

// ClipRule
//
// English:
//
//  It indicates how to determine what side of a path is inside a shape in order to know how a <clipPath> should clip
//  its target.
//
// Português:
//
//  Ele indica como determinar qual lado de um caminho está dentro de uma forma para saber como um <clipPath> deve
//  recortar seu destino.
func (e *TagSvgGlobal) ClipRule(clipRule SvgClipRule) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "clip-rule", clipRule.String())
	return e
}

// Color
//
// English:
//
//  It provides a potential indirect value (currentcolor) for the fill, stroke, stop-color, flood-color and
//  lighting-color presentation attributes.
//
//   Notes:
//     * As a presentation attribute, color can be used as a CSS property. See CSS color for further information.
//
// Português:
//
//  Ele fornece um valor indireto potencial (currentcolor) para os atributos de apresentação de preenchimento, traçado,
//  cor de parada, cor de inundação e cor de iluminação.
//
//   Notas:
//     * Como atributo de apresentação, a cor pode ser usada como propriedade CSS. Veja cor CSS para mais informações.
func (e *TagSvgGlobal) Color(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color", value)
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
func (e *TagSvgGlobal) ColorInterpolation(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color-interpolation", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color-interpolation", value)
	return e
}

// ColorInterpolationFilters
//
// English:
//
//  The color-interpolation-filters attribute specifies the color space for imaging operations performed via filter
//  effects.
//
//   Notes:
//     * This property just has an affect on filter operations. Therefore, it has no effect on filter primitives like
//       <feOffset>, <feImage>, <feTile> or <feFlood>;
//     * color-interpolation-filters has a different initial value than color-interpolation. color-interpolation-filters
//       has an initial value of linearRGB, whereas color-interpolation has an initial value of sRGB. Thus, in the
//       default case, filter effects operations occur in the linearRGB color space, whereas all other color
//       interpolations occur by default in the sRGB color space;
//     * It has no affect on filter functions, which operate in the sRGB color space;
//     * As a presentation attribute, color-interpolation-filters can be used as a CSS property.
//
// Português:
//
//  O atributo color-interpolation-filters especifica o espaço de cores para operações de imagem realizadas por meio de
//  efeitos de filtro.
//
//   Notas:
//     * Esta propriedade afeta apenas as operações de filtro. Portanto, não tem efeito em primitivos de filtro como
//       <feOffset>, <feImage>, <feTile> ou <feFlood>.
//     * color-interpolation-filters tem um valor inicial diferente de color-interpolation. color-interpolation-filters
//       tem um valor inicial de linearRGB, enquanto color-interpolation tem um valor inicial de sRGB. Assim, no caso
//       padrão, as operações de efeitos de filtro ocorrem no espaço de cores linearRGB, enquanto todas as outras
//       interpolações de cores ocorrem por padrão no espaço de cores sRGB.
//     * Não afeta as funções de filtro, que operam no espaço de cores sRGB.
//     * Como atributo de apresentação, os filtros de interpolação de cores podem ser usados como uma propriedade CSS.
func (e *TagSvgGlobal) ColorInterpolationFilters(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color-interpolation-filters", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color-interpolation-filters", value)
	return e
}

// CrossOrigin
//
// English:
//
//  The crossorigin attribute, valid on the <image> element, provides support for CORS, defining how the element handles
//  crossorigin requests, thereby enabling the configuration of the CORS requests for the element's fetched data. It is
//  a CORS settings attribute.
//
// Português:
//
//  The crossorigin attribute, valid on the <image> element, provides support for CORS, defining how the element handles
//  crossorigin requests, thereby enabling the configuration of the CORS requests for the element's fetched data. It is
//  a CORS settings attribute.
func (e *TagSvgGlobal) CrossOrigin(crossOrigin SvgCrossOrigin) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "crossorigin", crossOrigin.String())
	return e
}

// Cursor
//
// English:
//
//  The cursor attribute specifies the mouse cursor displayed when the mouse pointer is over an element.
//
// This attribute behaves exactly like the css cursor property except that if the browser supports the <cursor> element,
// you should be able to use it with the <funciri> notation.
//
// As a presentation attribute, it also can be used as a property directly inside a CSS stylesheet, see css cursor for
// further information.
//
// Português:
//
//  O atributo cursor especifica o cursor do mouse exibido quando o ponteiro do mouse está sobre um elemento.
//
// Este atributo se comporta exatamente como a propriedade cursor css, exceto que, se o navegador suportar o elemento
// <cursor>, você poderá usá-lo com a notação <funciri>.
//
// Como atributo de apresentação, também pode ser usado como propriedade diretamente dentro de uma folha de estilo CSS,
// veja cursor css para mais informações.
func (e *TagSvgGlobal) Cursor(cursor SvgCursor) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "cursor", cursor.String())
	return e
}

// Cx
//
// English:
//
//  The cx attribute define the x-axis coordinate of a center point.
//
// Português:
//
//  O atributo cx define a coordenada do eixo x de um ponto central.
func (e *TagSvgGlobal) Cx(cx float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "cx", cx)
	return e
}

// Cy
//
// English:
//
// The cy attribute define the y-axis coordinate of a center point.
//
// Português:
//
//  O atributo cy define a coordenada do eixo y de um ponto central.
func (e *TagSvgGlobal) Cy(cy float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "cy", cy)
	return e
}

// D
//
// English:
//
//  The d attribute defines a path to be drawn.
//
// A path definition is a list of path commands where each command is composed of a command letter and numbers that
// represent the command parameters. The commands are detailed below.
//
// You can use this attribute with the following SVG elements: <path>, <glyph>, <missing-glyph>.
//
// d is a presentation attribute, and hence can also be used as a CSS property.
//
// Português:
//
//  O atributo d define um caminho a ser desenhado.
//
// Uma definição de caminho é uma lista de comandos de caminho em que cada comando é composto por uma letra de comando
// e números que representam os parâmetros do comando. Os comandos são detalhados abaixo.
//
// Você pode usar este atributo com os seguintes elementos SVG: <path>, <glyph>, <missing-glyph>.
//
// d é um atributo de apresentação e, portanto, também pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) D(d *SvgPath) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "d", d.String())
	return e
}

// DiffuseConstant
//
// English:
//
//  The diffuseConstant attribute represents the kd value in the Phong lighting model. In SVG, this can be any
//  non-negative number.
//
// It's used to determine the final RGB value of a given pixel. The brighter the lighting-color, the smaller this number
// should be.
//
// Português:
//
//  O atributo difusoConstant representa o valor kd no modelo de iluminação Phong. Em SVG, pode ser qualquer número
//  não negativo.
//
// É usado para determinar o valor RGB final de um determinado pixel. Quanto mais brilhante a cor da iluminação, menor
// deve ser esse número.
func (e *TagSvgGlobal) DiffuseConstant(diffuseConstant float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "diffuseConstant", diffuseConstant)
	return e
}

// Direction
//
// English:
//
//  The direction attribute specifies the inline-base direction of a <text> or <tspan> element. It defines the start
//  and end points of a line of text as used by the text-anchor and inline-size properties. It also may affect the
//  direction in which characters are positioned if the unicode-bidi property's value is either embed or bidi-override.
//
// It applies only to glyphs oriented perpendicular to the inline-base direction, which includes the usual case of
// horizontally-oriented Latin or Arabic text and the case of narrow-cell Latin or Arabic characters rotated 90 degrees
// clockwise relative to a top-to-bottom inline-base direction.
//
// In many cases, the bidirectional Unicode algorithm produces the desired result automatically, so this attribute
// doesn't need to be specified in those cases. For other cases, such as when using right-to-left languages, it may be
// sufficient to add the direction attribute to the outermost <svg> element, and allow that direction to inherit to all
// text elements:
//
//   Notes:
//     * As a presentation attribute, direction can be used as a CSS property. See css direction for further
//       information.
//
// Português:
//
//  O atributo direction especifica a direção da base embutida de um elemento <text> ou <tspan>. Ele define os pontos
//  inicial e final de uma linha de texto conforme usado pelas propriedades text-anchor e inline-size.
//  Também pode afetar a direção na qual os caracteres são posicionados se o valor da propriedade unicode-bidi for
//  incorporado ou substituído por bidi.
//
// Aplica-se apenas a glifos orientados perpendicularmente à direção da base em linha, que inclui o caso usual de texto
// latino ou árabe orientado horizontalmente e o caso de caracteres latinos ou árabes de célula estreita girados 90
// graus no sentido horário em relação a um texto de cima para baixo direção de base em linha.
//
// Em muitos casos, o algoritmo Unicode bidirecional produz o resultado desejado automaticamente, portanto, esse
// atributo não precisa ser especificado nesses casos. Para outros casos, como ao usar idiomas da direita para a
// esquerda, pode ser suficiente adicionar o atributo direction ao elemento <svg> mais externo e permitir que essa
// direção herde todos os elementos de texto:
//
//   Notas:
//     * Como atributo de apresentação, a direção pode ser usada como uma propriedade CSS. Veja a direção do CSS para
//       mais informações.
func (e *TagSvgGlobal) Direction(direction SvgDirection) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "direction", direction.String())
	return e
}

// Display
//
// English:
//
//  The display attribute lets you control the rendering of graphical or container elements.
//
// A value of display="none" indicates that the given element and its children will not be rendered. Any value other
// than none or inherit indicates that the given element will be rendered by the browser.
//
// When applied to a container element, setting display to none causes the container and all of its children to be
// invisible; thus, it acts on groups of elements as a group. This means that any child of an element with
// display="none" will never be rendered even if the child has a value for display other than none.
//
// When the display attribute is set to none, then the given element does not become part of the rendering tree. It has
// implications for the <tspan>, <tref>, and <altGlyph> elements, event processing, for bounding box calculations and
// for calculation of clipping paths:
//
//   * If display is set to none on a <tspan>, <tref>, or <altGlyph> element, then the text string is ignored for the
//     purposes of text layout.
//   * Regarding events, if display is set to none, the element receives no events.
//   * The geometry of a graphics element with display set to none is not included in bounding box and clipping paths
//     calculations.
//
// The display attribute only affects the direct rendering of a given element, whereas it does not prevent elements
// from being referenced by other elements. For example, setting it to none on a <path> element will prevent that
// element from getting rendered directly onto the canvas, but the <path> element can still be referenced by a
// <textPath> element; furthermore, its geometry will be used in text-on-a-path processing even if the <path> has a
// display value of none.
//
// This attribute also affects direct rendering into offscreen canvases, such as occurs with masks or clip paths. Thus,
// setting display="none" on a child of a <mask> will prevent the given child element from being rendered as part of the
// mask. Similarly, setting display="none" on a child of a <clipPath> element will prevent the given child element from
// contributing to the clipping path.
//
//  Notes:
//    * As a presentation attribute, display can be used as a CSS property. See css display for further information.
//
// Português:
//
//  O atributo display permite controlar a renderização de elementos gráficos ou de contêiner.
//
// Um valor de display="none" indica que o elemento fornecido e seus filhos não serão renderizados. Qualquer valor
// diferente de none ou herdar indica que o elemento fornecido será renderizado pelo navegador.
//
// Quando aplicado a um elemento de contêiner, definir display como none faz com que o contêiner e todos os seus filhos
// fiquem invisíveis; assim, atua em grupos de elementos como um grupo. Isso significa que qualquer filho de um elemento
// com display="none" nunca será renderizado, mesmo que o filho tenha um valor para exibição diferente de none.
//
// Quando o atributo display é definido como none, o elemento fornecido não se torna parte da árvore de renderização.
// Tem implicações para os elementos <tspan>, <tref> e <altGlyph>, processamento de eventos, para cálculos de caixa
// delimitadora e para cálculo de caminhos de recorte:
//   * Se display for definido como none em um elemento <tspan>, <tref> ou <altGlyph>, a string de texto será ignorada
//     para fins de layout de texto.
//   * Com relação aos eventos, se display estiver definido como none, o elemento não recebe eventos.
//   * A geometria de um elemento gráfico com exibição definida como nenhum não é incluída nos cálculos da caixa
//     delimitadora e dos caminhos de recorte.
//
// O atributo display afeta apenas a renderização direta de um determinado elemento, mas não impede que os elementos
// sejam referenciados por outros elementos. Por exemplo, defini-lo como none em um elemento <path> impedirá que esse
// elemento seja renderizado diretamente na tela, mas o elemento <path> ainda pode ser referenciado por um elemento
// <textPath>; além disso, sua geometria será usada no processamento de texto em um caminho, mesmo que o <caminho>
// tenha um valor de exibição de nenhum.
//
// Esse atributo também afeta a renderização direta em telas fora da tela, como ocorre com máscaras ou caminhos de
// clipe. Assim, definir display="none" em um filho de uma <mask> impedirá que o elemento filho fornecido seja
// renderizado como parte da máscara. Da mesma forma, definir display="none" em um filho de um elemento <clipPath>
// impedirá que o elemento filho fornecido contribua para o caminho de recorte.
//
//  Notas:
//    * Como atributo de apresentação, display pode ser usado como propriedade CSS. Consulte a exibição css para obter
//      mais informações.
func (e *TagSvgGlobal) Display(display SvgDisplay) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "display", display.String())
	return e
}

// Divisor
//
// English:
//
//  The divisor attribute specifies the value by which the resulting number of applying the kernelMatrix of a
//  <feConvolveMatrix> element to the input image color value is divided to yield the destination color value.
//
//   Input:
//     divisor: specifies the divisor value to apply to the original color
//
// A divisor that is the sum of all the matrix values tends to have an evening effect on the overall color intensity of
// the result.
//
// Português:
//
//  O atributo divisor especifica o valor pelo qual o número resultante da aplicação do kernelMatrix de um elemento
//  <feConvolveMatrix> ao valor da cor da imagem de entrada é dividido para gerar o valor da cor de destino.
//
//   Entrada:
//     divisor: especifica o valor do divisor a ser aplicado na cor original
//
//
// A divisor that is the sum of all the matrix values tends to have an evening effect on the overall color intensity of
// the result.
func (e *TagSvgGlobal) Divisor(divisor float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "divisor", divisor)
	return e
}

// DominantBaseline
//
// English:
//
//  The dominant-baseline attribute specifies the dominant baseline, which is the baseline used to align the box's text and inline-level contents. It also indicates the default alignment baseline of any boxes participating in baseline alignment in the box's alignment context.
//
// It is used to determine or re-determine a scaled-baseline-table. A scaled-baseline-table is a compound value with three components:
//
//   1. a baseline-identifier for the dominant-baseline,
//   2. a baseline-table, and
//   3. a baseline-table font-size.
//
// Some values of the property re-determine all three values. Others only re-establish the baseline-table font-size. When the initial value, auto, would give an undesired result, this property can be used to explicitly set the desired scaled-baseline-table.
//
// If there is no baseline table in the nominal font, or if the baseline table lacks an entry for the desired baseline, then the browser may use heuristics to determine the position of the desired baseline.
//
//   Notes:
//     * As a presentation attribute, dominant-baseline can be used as a CSS property.
//
// Português:
//
//
func (e *TagSvgGlobal) DominantBaseline(dominantBaseline SvgDominantBaseline) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "dominant-baseline", dominantBaseline.String())
	return e
}

// Dur
//
// English:
//
//  The dur attribute indicates the simple duration of an animation.
//
//   Input:
//     dur: indicates the simple duration of an animation.
//       KSvgDur... (e.g. KSvgDurIndefinite)
//       time.Duration (e.g. time.Second * 5)
//
//   Notes:
//     * The interpolation will not work if the simple duration is indefinite (although this may still be useful for
//       <set> elements).
//
// Português:
//
//  O atributo dur indica a duração simples de uma animação.
//
//   Entrada:
//     dur: indica a duração simples de uma animação.
//       KSvgDur... (ex. KSvgDurIndefinite)
//       time.Duration (ex. time.Second * 5)
//
//   Notas:
//     * A interpolação não funcionará se a duração simples for indefinida (embora isso ainda possa ser útil para
//       elementos <set>).
func (e *TagSvgGlobal) Dur(dur interface{}) (ref *TagSvgGlobal) {
	switch converted := dur.(type) {
	case time.Duration:
		e.selfElement.Call("setAttribute", "dur", converted.String())
	case SvgDur:
		e.selfElement.Call("setAttribute", "dur", converted.String())
	default:
		e.selfElement.Call("setAttribute", "dur", dur)
	}

	return e
}

// Dx
//
// English:
//
//  The dx attribute indicates a shift along the x-axis on the position of an element or its content.
//
//   Input:
//     dx: indicates a shift along the x-axis on the position of an element or its content.
//       (e.g. "20%", "0 10%" or "0 10% 20%")
//
// Portuguese
//
//  O atributo dx indica um deslocamento ao longo do eixo x na posição de um elemento ou seu conteúdo.
//
//   Entrada:
//     dx: indica um deslocamento ao longo do eixo x na posição de um elemento ou seu conteúdo.
//       (ex. "20%", "0 10%" ou "0 10% 20%")
func (e *TagSvgGlobal) Dx(dx string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "dx", dx)
	return e
}

// Dy
//
// English:
//
//  The dy attribute indicates a shift along the y-axis on the position of an element or its content.
//
//   Input:
//     dy: indicates a shift along the y-axis on the position of an element or its content.
//       (e.g. "50%", "20", "0 10" or "0 10 20")
//
// Portuguese
//
//  O atributo dy indica um deslocamento ao longo do eixo y na posição de um elemento ou seu conteúdo.
//
//   Entrada:
//     dy: indica um deslocamento ao longo do eixo y na posição de um elemento ou seu conteúdo.
//       (ex. "50%", "20", "0 10" ou "0 10 20")
func (e *TagSvgGlobal) Dy(dy string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "dy", dy)
	return e
}

// EdgeMode
//
// English:
//
//  The edgeMode attribute determines how to extend the input image as necessary with color values so that the matrix
//  operations can be applied when the kernel is positioned at or near the edge of the input image.
//
// Portuguese
//
//  O atributo edgeMode determina como estender a imagem de entrada conforme necessário com valores de cor para que as
//  operações de matriz possam ser aplicadas quando o kernel estiver posicionado na borda da imagem de entrada ou
//  próximo a ela.
//
func (e *TagSvgGlobal) EdgeMode(edgeMode SvgEdgeMode) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "edgeMode", edgeMode.String())
	return e
}

// Elevation
//
// English:
//
//  The elevation attribute specifies the direction angle for the light source from the XY plane towards the Z-axis, in
//  degrees. Note that the positive Z-axis points towards the viewer of the content.
//
// Portuguese
//
//  O atributo de elevação especifica o ângulo de direção da fonte de luz do plano XY em direção ao eixo Z, em graus.
//  Observe que o eixo Z positivo aponta para o visualizador do conteúdo.
func (e *TagSvgGlobal) Elevation(elevation float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "elevation", elevation)
	return e
}

// End
//
// English:
//
//  The end attribute defines an end value for the animation that can constrain the active duration.
//
//   Input:
//     end: defines an end value for the animation
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
// Portuguese
//
//  O atributo final define um valor final para a animação que pode restringir a duração ativa.
//
//   Entrada:
//     end: define um valor final para a animação
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
//
func (e *TagSvgGlobal) End(end interface{}) (ref *TagSvgGlobal) {
	if converted, ok := end.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "end", converted.String())
		return e
	}

	if converted, ok := end.(time.Time); ok {
		e.selfElement.Call("setAttribute", "end", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "end", end)
	return e
}

// Exponent
//
// English:
//
//  The exponent attribute defines the exponent of the gamma function.
//
// Portuguese
//
//  O atributo expoente define o expoente da função gama.
func (e *TagSvgGlobal) Exponent(exponent float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "exponent", exponent)
	return e
}

// Fill
//
// English:
//
//  The fill attribute has two different meanings. For shapes and text it's a presentation attribute that defines the
//  color (or any SVG paint servers like gradients or patterns) used to paint the element;
//
// for animation it defines the final state of the animation.
//
// Português:
//
//  O atributo fill tem dois significados diferentes. Para formas e texto, é um atributo de apresentação que define a
//  cor (ou qualquer servidor de pintura SVG, como gradientes ou padrões) usado para pintar o elemento;
//
// para animação, define o estado final da animação.
func (e *TagSvgGlobal) Fill(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "fill", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "fill", value)
	return e
}

// FillOpacity
//
// English:
//
//  The fill-opacity attribute is a presentation attribute defining the opacity of the paint server (color, gradient,
//  pattern, etc) applied to a shape.
//
//   Notes:
//     *As a presentation attribute fill-opacity can be used as a CSS property.
//
// Portuguese
//
//  O atributo fill-opacity é um atributo de apresentação que define a opacidade do servidor de pintura (cor, gradiente,
//  padrão etc.) aplicado a uma forma.
//
//   Notes:
//     *As a presentation attribute fill-opacity can be used as a CSS property.
func (e *TagSvgGlobal) FillOpacity(fillOpacity float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "fill-opacity", fillOpacity)
	return e
}

// FillRule
//
// English:
//
//  The fill-rule attribute is a presentation attribute defining the algorithm to use to determine the inside part of
//  a shape.
//
//   Notes:
//     * As a presentation attribute, fill-rule can be used as a CSS property.
//
// Portuguese
//
//  O atributo fill-rule é um atributo de apresentação que define o algoritmo a ser usado para determinar a parte
//  interna de uma forma.
//
//   Notas:
//     * Como atributo de apresentação, fill-rule pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) FillRule(fillRule SvgFillRule) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "fill-rule", fillRule.String())
	return e
}

// Filter
//
// English:
//
//  The filter attribute specifies the filter effects defined by the <filter> element that shall be applied to its
//  element.
//
//   Notes:
//     * As a presentation attribute, filter can be used as a CSS property. See css filter for further information.
//
// Portuguese
//
//  O atributo filter especifica os efeitos de filtro definidos pelo elemento <filter> que devem ser aplicados ao seu
//  elemento.
//
//   Notas:
//     * Como atributo de apresentação, o filtro pode ser usado como propriedade CSS. Veja filtro css para mais
//       informações.
func (e *TagSvgGlobal) Filter(filter string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "filter", filter)
	return e
}

// FilterUnits
//
// English:
//
//  The filterUnits attribute defines the coordinate system for the attributes x, y, width and height.
//
// Portuguese
//
//   O atributo filterUnits define o sistema de coordenadas para os atributos x, y, largura e altura.
func (e *TagSvgGlobal) FilterUnits(filterUnits SvgFilterUnits) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "filterUnits", filterUnits.String())
	return e
}

// FloodColor
//
// English:
//
//  The flood-color attribute indicates what color to use to flood the current filter primitive subregion.
//
//   Notes:
//     * As a presentation attribute, flood-color can be used as a CSS property.
//
// Portuguese
//
//  The flood-color attribute indicates what color to use to flood the current filter primitive subregion.
//
//   Notes:
//     * As a presentation attribute, flood-color can be used as a CSS property.
func (e *TagSvgGlobal) FloodColor(floodColor interface{}) (ref *TagSvgGlobal) {
	if converted, ok := floodColor.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "flood-color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "flood-color", floodColor)
	return e
}

// FloodOpacity
//
// English:
//
//  The flood-opacity attribute indicates the opacity value to use across the current filter primitive subregion.
//
//   Notes:
//     * As a presentation attribute, flood-opacity can be used as a CSS property.
//
// Portuguese
//
//  O atributo flood-opacity indica o valor de opacidade a ser usado na sub-região primitiva de filtro atual.
//
//   Notas:
//     * Como atributo de apresentação, a opacidade de inundação pode ser usada como uma propriedade CSS.
func (e *TagSvgGlobal) FloodOpacity(floodOpacity float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "flood-opacity", floodOpacity)
	return e
}

//todo: normalizar fonte com html padrão

// FontFamily
//
// English:
//
//  The font-family attribute indicates which font family will be used to render the text, specified as a prioritized
//  list of font family names and/or generic family names.
//
//   Notes:
//     * As a presentation attribute, font-family can be used as a CSS property. See the css font-family property for
//       more information.
//
// Portuguese
//
//  O atributo font-family indica qual família de fontes será usada para renderizar o texto, especificada como uma lista
//  priorizada de nomes de famílias de fontes e ou nomes de famílias genéricos.
//
//   Notas:
//     * Como atributo de apresentação, font-family pode ser usada como propriedade CSS. Consulte a propriedade CSS
//       font-family para obter mais informações.
func (e *TagSvgGlobal) FontFamily(fontFamily string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-family", fontFamily)
	return e
}

// FontSize
//
// English:
//
//  The font-size attribute refers to the size of the font from baseline to baseline when multiple lines of text are set
//  solid in a multiline layout environment.
//
//   Notes:
//     * As a presentation attribute, font-size can be used as a CSS property. See the css font-size property for more
//       information.
//
// Portuguese
//
//  O atributo font-size refere-se ao tamanho da fonte da linha de base a linha de base quando várias linhas de texto
//  são definidas como sólidas em um ambiente de layout de várias linhas.
//
//   Notas:
//     * Como atributo de apresentação, font-size pode ser usado como uma propriedade CSS. Consulte a propriedade CSS
//       font-size para obter mais informações.
func (e *TagSvgGlobal) FontSize(fontSize float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-size", fontSize)
	return e
}

// FontSizeAdjust
//
// English:
//
//  The font-size-adjust attribute allows authors to specify an aspect value for an element that will preserve the
//  x-height of the first choice font in a substitute font.
//
//   Notes:
//     * As a presentation attribute, font-size-adjust can be used as a CSS property. See the css font-size-adjust
//       property for more information.
//
// Portuguese
//
//  O atributo font-size-adjust permite que os autores especifiquem um valor de aspecto para um elemento que preservará
//  a altura x da fonte de primeira escolha em uma fonte substituta.
//
//   Notes:
//     * As a presentation attribute, font-size-adjust can be used as a CSS property. See the css font-size-adjust
//       property for more information.
func (e *TagSvgGlobal) FontSizeAdjust(fontSizeAdjust float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-size-adjust", fontSizeAdjust)
	return e
}

// FontStretch
//
// English:
//
//  The font-stretch attribute indicates the desired amount of condensing or expansion in the glyphs used to render
//  the text.
//
//   Input:
//     fontStretch: indicates the desired amount of condensing or expansion
//       KSvgFontStretch... (e.g. KSvgFontStretchUltraCondensed)
//       percentage (e.g. "50%")
//
//   Notes:
//     * As a presentation attribute, font-stretch can be used as a CSS property. See the css font-stretch property for
//       more information.
//
// Portuguese
//
//  O atributo font-stretch indica a quantidade desejada de condensação ou expansão nos glifos usados para renderizar
//  o texto.
//
//   Entrada:
//     fontStretch: indica a quantidade desejada de condensação ou expansão
//       KSvgFontStretch... (e.g. KSvgFontStretchUltraCondensed)
//       percentage (e.g. "50%")
//
//   Notas:
//     * Como atributo de apresentação, font-stretch pode ser usado como uma propriedade CSS. Consulte a propriedade
//       CSS font-stretch para obter mais informações.
func (e *TagSvgGlobal) FontStretch(fontStretch interface{}) (ref *TagSvgGlobal) {
	if converted, ok := fontStretch.(SvgFontStretch); ok {
		e.selfElement.Call("setAttribute", "font-stretch", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-stretch", fontStretch)
	return e
}

// FontStyle
//
// English:
//
//  The font-style attribute specifies whether the text is to be rendered using a normal, italic, or oblique face.
//
//   Notes:
//     * As a presentation attribute, font-style can be used as a CSS property. See the css font-style property for
//       more information.
//
// Portuguese
//
//  O atributo font-style especifica se o texto deve ser renderizado usando uma face normal, itálica ou oblíqua.
//
//   Notas:
//     * Como atributo de apresentação, font-style pode ser usado como propriedade CSS. Consulte a propriedade CSS
//       font-style para obter mais informações.
func (e *TagSvgGlobal) FontStyle(fontStyle FontStyleRule) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-style", fontStyle.String())
	return e
}

// FontVariant
//
// English:
//
//  The font-variant attribute indicates whether the text is to be rendered using variations of the font's glyphs.
//
//   Notes:
//     * As a presentation attribute, font-variant can be used as a CSS property. See the css font-variant property
//       for more information.
//
// Portuguese
//
//  O atributo font-variant indica se o texto deve ser renderizado usando variações dos glifos da fonte.
//
//   Notas:
//     * Como atributo de apresentação, font-variant pode ser usado como uma propriedade CSS. Consulte a propriedade
//       CSS font-variant para obter mais informações.
func (e *TagSvgGlobal) FontVariant(fontVariant FontVariantRule) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-variant", fontVariant.String())
	return e
}

// FontWeight
//
// English:
//
//  The font-weight attribute refers to the boldness or lightness of the glyphs used to render the text, relative to
//  other fonts in the same font family.
//
//   Notes:
//     * As a presentation attribute, font-weight can be used as a CSS property. See the css font-weight property for
//       more information.
//
// Portuguese
//
//  O atributo font-weight refere-se ao negrito ou leveza dos glifos usados para renderizar o texto, em relação a
//  outras fontes na mesma família de fontes.
//
//   Notas:
//     * Como atributo de apresentação, o peso da fonte pode ser usado como uma propriedade CSS. Consulte a propriedade
//       CSS font-weight para obter mais informações.
func (e *TagSvgGlobal) FontWeight(fontWeight FontWeightRule) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-weight", fontWeight.String())
	return e
}

// From
//
// English:
//
//  The from attribute indicates the initial value of the attribute that will be modified during the animation.
//
// When used with the to attribute, the animation will change the modified attribute from the from value to the to
// value. When used with the by attribute, the animation will change the attribute relatively from the from value by
// the value specified in by.
//
// Português
//
//  O atributo from indica o valor inicial do atributo que será modificado durante a animação.
//
// Quando usado com o atributo to, a animação mudará o atributo modificado do valor from para o valor to. Quando usado
// com o atributo by, a animação mudará o atributo relativamente do valor from pelo valor especificado em by.
func (e *TagSvgGlobal) From(from float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "from", from)
	return e
}

// Fr
//
// English:
//
//  The fr attribute defines the radius of the focal point for the radial gradient.
//
//   Input:
//     fr: defines the radius of the focal point for the radial gradient
//       float64: (e.g. 0.4 = 40%)
//       string: "40%"
//
// Portuguese
//
//  O atributo fr define o raio do ponto focal para o gradiente radial.
//
//   Entrada:
//     fr: define o raio do ponto focal para o gradiente radial.
//       float64: (ex. 0.4 = 40%)
//       string: "40%"
func (e *TagSvgGlobal) Fr(fr interface{}) (ref *TagSvgGlobal) {
	if converted, ok := fr.(float64); ok {
		p := strconv.FormatFloat(100*converted, 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "fr", p)
		return e
	}

	e.selfElement.Call("setAttribute", "fr", fr)
	return e
}

// Fx
//
// English:
//
//  The fx attribute defines the x-axis coordinate of the focal point for a radial gradient.
//
// Portuguese
//
//  O atributo fx define a coordenada do eixo x do ponto focal para um gradiente radial.
func (e *TagSvgGlobal) Fx(fx float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "fx", fx)
	return e
}

// Fy
//
// English:
//
//  The fy attribute defines the y-axis coordinate of the focal point for a radial gradient.
//
// Portuguese
//
//  O atributo fy define a coordenada do eixo y do ponto focal para um gradiente radial.
func (e *TagSvgGlobal) Fy(fy float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "fy", fy)
	return e
}

// GradientTransform
//
// English:
//
//  The gradientTransform attribute contains the definition of an optional additional transformation from the gradient
//  coordinate system onto the target coordinate system (i.e., userSpaceOnUse or objectBoundingBox).
//  This allows for things such as skewing the gradient. This additional transformation matrix is post-multiplied to
//  (i.e., inserted to the right of) any previously defined transformations, including the implicit transformation
//  necessary to convert from object bounding box units to user space.
//
// Portuguese
//
//  O atributo gradientTransform contém a definição de uma transformação adicional opcional do sistema de coordenadas
//  de gradiente para o sistema de coordenadas de destino (ou seja, userSpaceOnUse ou objectBoundingBox).
//  Isso permite coisas como distorcer o gradiente. Essa matriz de transformação adicional é pós-multiplicada para
//  (ou seja, inserida à direita de) quaisquer transformações definidas anteriormente, incluindo a transformação
//  implícita necessária para converter de unidades de caixa delimitadora de objeto para espaço do usuário.
func (e *TagSvgGlobal) GradientTransform(gradientTransform float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "gradientTransform", gradientTransform)
	return e
}

// GradientUnits
//
// English:
//
//  The gradientUnits attribute defines the coordinate system used for attributes specified on the gradient elements.
//
// Portuguese
//
//  O atributo gradientUnits define o sistema de coordenadas usado para atributos especificados nos elementos
//  gradientes.
func (e *TagSvgGlobal) GradientUnits(gradientUnits SvgGradientUnits) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "gradientUnits", gradientUnits)
	return e
}

// Height
//
// English:
//
//  The height attribute defines the vertical length of an element in the user coordinate system.
//
// Português:
//
//  O atributo height define o comprimento vertical de um elemento no sistema de coordenadas do usuário.
func (e *TagSvgGlobal) Height(height interface{}) (ref *TagSvgGlobal) {
	if converted, ok := height.(float64); ok {
		p := strconv.FormatFloat(100*converted, 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "height", p)
		return e
	}

	e.selfElement.Call("setAttribute", "height", height)
	return e
}

// HRef
//
// English:
//
//  The href attribute defines a link to a resource as a reference URL. The exact meaning of that link depends on the
//  context of each element using it.
//
//   Notes:
//     * Specifications before SVG 2 defined an xlink:href attribute, which is now rendered obsolete by the href
//       attribute.
//       If you need to support earlier browser versions, the deprecated xlink:href attribute can be used as a fallback
//       in addition to the href attribute, e.g. <use href="some-id" xlink:href="some-id x="5" y="5" />.
//
// Português
//
//  O atributo href define um link para um recurso como um URL de referência. O significado exato desse link depende do
//  contexto de cada elemento que o utiliza.
//
//   Notas:
//     * As especificações anteriores ao SVG 2 definiam um atributo xlink:href, que agora se torna obsoleto pelo
//       atributo href.
//       Se você precisar oferecer suporte a versões anteriores do navegador, o atributo obsoleto xlink:href pode ser
//       usado como um substituto além do atributo href, por exemplo,
//       <use href="some-id" xlink:href="some-id x="5" y="5" />.
func (e *TagSvgGlobal) HRef(href string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "href", href)
	return e
}

// Id
//
// English:
//
//  The id attribute assigns a unique name to an element.
//
// Portuguese
//
//  O atributo id atribui um nome exclusivo a um elemento.
func (e *TagSvgGlobal) Id(id string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "id", id)
	return e
}

// ImageRendering
//
// English:
//
//  The image-rendering attribute provides a hint to the browser about how to make speed vs. quality tradeoffs as it
//  performs image processing.
//
// The resampling is always done in a truecolor (e.g., 24-bit) color space even if the original data and/or the target
// device is indexed color.
//
//   Notes:
//     * As a presentation attribute, image-rendering can be used as a CSS property. See the css image-rendering
//       property for more information.
//
// Portuguese
//
//  O atributo de renderização de imagem fornece uma dica ao navegador sobre como fazer compensações de velocidade
//  versus qualidade enquanto executa o processamento de imagem.
//
// A reamostragem é sempre feita em um espaço de cores truecolor (por exemplo, 24 bits), mesmo que os dados originais e
// ou o dispositivo de destino sejam cores indexadas.
//
//   Notas:
//     * Como um atributo de apresentação, a renderização de imagem pode ser usada como uma propriedade CSS. Consulte
//       a propriedade de renderização de imagem css para obter mais informações.
func (e *TagSvgGlobal) ImageRendering(imageRendering string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "image-rendering", imageRendering)
	return e
}

// In
//
// English:
//
//  The in attribute identifies input for the given filter primitive.
//
//   Input:
//     in: identifies input for the given filter primitive.
//       KSvgIn... (e.g. KSvgInSourceAlpha)
//
// The value can be either one of the six keywords defined below, or a string which matches a previous result attribute
// value within the same <filter> element. If no value is provided and this is the first filter primitive, then this
// filter primitive will use SourceGraphic as its input. If no value is provided and this is a subsequent filter
// primitive, then this filter primitive will use the result from the previous filter primitive as its input.
//
// If the value for result appears multiple times within a given <filter> element, then a reference to that result will
// use the closest preceding filter primitive with the given value for attribute result.
//
// Portuguese
//
//  O atributo in identifica à entrada para a primitiva de filtro fornecida.
//
//   Entrada:
//     in: identifica à entrada para a primitiva de filtro fornecida.
//       KSvgIn... (e.g. KSvgInSourceAlpha)
//
// O valor pode ser uma das seis palavras-chave definidas abaixo ou uma string que corresponda a um valor de atributo
// de resultado anterior dentro do mesmo elemento <filter>. Se nenhum valor for fornecido e esta for a primeira
// primitiva de filtro, essa primitiva de filtro usará SourceGraphic como sua entrada. Se nenhum valor for fornecido e
// esta for uma primitiva de filtro subsequente, essa primitiva de filtro usará o resultado da primitiva de filtro
// anterior como sua entrada.
//
// Se o valor do resultado aparecer várias vezes em um determinado elemento <filter>, uma referência à esse resultado
// usará a primitiva de filtro anterior mais próxima com o valor fornecido para o resultado do atributo.
func (e *TagSvgGlobal) In(in interface{}) (ref *TagSvgGlobal) {
	if converted, ok := in.(SvgIn); ok {
		e.selfElement.Call("setAttribute", "in", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "in", in)
	return e
}

// In2
//
// English:
//
//  The in2 attribute identifies the second input for the given filter primitive. It works exactly like the in
//  attribute.
//
//   Input:
//     in2: identifies the second input for the given filter primitive.
//       KSvgIn2... (e.g. KSvgIn2SourceAlpha)
//       string: url(#myClip)
//
// Portuguese
//
//  O atributo in2 identifica a segunda entrada para a primitiva de filtro fornecida. Funciona exatamente como o
//  atributo in.
//
//   Entrada:
//     in2: identifica a segunda entrada para a primitiva de filtro fornecida.
//       KSvgIn2... (ex. KSvgIn2SourceAlpha)
//       string: url(#myClip)
func (e *TagSvgGlobal) In2(in2 interface{}) (ref *TagSvgGlobal) {
	if converted, ok := in2.(SvgIn2); ok {
		e.selfElement.Call("setAttribute", "in2", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "in2", in2)
	return e
}

// Intercept
//
// English:
//
//  The intercept attribute defines the intercept of the linear function of color component transfers when the type
//  attribute is set to linear.
//
// Portuguese
//
//  O atributo de interceptação define a interceptação da função linear de transferências de componentes de cor quando
//  o atributo de tipo é definido como linear.
func (e *TagSvgGlobal) Intercept(intercept float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "intercept", intercept)
	return e
}

// K1
//
// English:
//
//  The k1 attribute defines one of the values to be used within the arithmetic operation of the <feComposite>
//  filter primitive.
//
// The pixel composition is computed using the following formula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//  O atributo k1 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K1(k1 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k1", k1)
	return e
}

// K2
//
// English:
//
//  The k2 attribute defines one of the values to be used within the arithmetic operation of the <feComposite> filter
//  primitive.
//
// The pixel composition is computed using the following formula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//  O atributo k2 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K2(k2 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k2", k2)
	return e
}

// K3
//
// English:
//
//  The k3 attribute defines one of the values to be used within the arithmetic operation of the <feComposite>
//  filter primitive.
//
// The pixel composition is computed using the following formula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//  O atributo k3 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K3(k3 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k3", k3)
	return e
}

// K4
//
// English:
//
//  The k4 attribute defines one of the values to be used within the arithmetic operation of the <feComposite>
//  filter primitive.
//
// The pixel composition is computed using the following formula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//  O atributo k4 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//   result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K4(k4 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k4", k4)
	return e
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
