package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"image/color"
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

// ================================================================================================================ Y
// Y
//
// English:
//
//	The y attribute defines a y-axis coordinate in the user coordinate system.
//
//	You can use this attribute with the following SVG elements:
//
//	<altGlyph>
//	<cursor>
//	<feBlend>
//	<feColorMatrix>
//	<feComponentTransfer>
//	<feComposite>
//	<feConvolveMatrix>
//	<feDiffuseLighting>
//	<feDisplacementMap>
//	<feDropShadow>
//	<feFlood>
//	<feFuncA>
//	<feFuncB>
//	<feFuncG>
//	<feFuncR>
//	<feGaussianBlur>
//	<feImage>
//	<feMerge>
//	<feMergeNode>
//	<feMorphology>
//	<feOffset>
//	<fePointLight>
//	<feSpecularLighting>
//	<feSpotLight>
//	<feTile>
//	<feTurbulence>
//	<filter>
//	<foreignObject>
//	<glyphRef>
//	<image>
//	<mask>
//	<pattern>
//	<rect>
//	<svg>
//	<symbol>
//	<text>
//	<tref>
//	<tspan>
//	<use>
//
//	Input:
//	   y: defines a y-axis coordinate in the user coordinate system.
//		 feBlend: For <feBlend>, y defines the minimum y coordinate for the rendering area of the primitive.
//		   feColorMatrix: For <feColorMatrix>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feColorMatrix: For <feColorMatrix>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feBlend: For <feBlend>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feColorMatrix: For <feColorMatrix>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feComponentTransfer: For <feComponentTransfer>, y defines the minimum y coordinate for the rendering
//	   	   area of the primitive.
//		 feComposite: For <feComposite>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feConvolveMatrix: For <feConvolveMatrix>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feDiffuseLighting: For <feDiffuseLighting>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feDisplacementMap: For <feDisplacementMap>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feDropShadow: For <feDropShadow>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feFlood: For <feFlood>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feFuncA: For <feFuncA>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feFuncB: For <feFuncB>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feFuncG: For <feFuncG>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feFuncR: For <feFuncR>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feGaussianBlur: For <feGaussianBlur>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feImage: For <feImage>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feMerge: For <feMerge>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feMergeNode: For <feMergeNode>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feMorphology: For <feMorphology>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feOffset: For <feOffset>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 fePointLight: For <fePointLight>, y defines the y location for the light source in the coordinate system
//	  	   defined by the primitiveUnits attribute on the <filter> element.
//		 feSpecularLighting: For <feSpecularLighting>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feSpotLight: For <feSpotLight>, y defines the y location for the light source in the coordinate system defined
//		   by the primitiveUnits attribute on the <filter> element.
//		 feTile: For <feMerge>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 feTurbulence: For <feTurbulence>, y defines the minimum y coordinate for the rendering area of the primitive.
//		 filter: For <filter>, y defines the y coordinate of the upper left corner for the rendering area of the filter.
//		 foreignObject: For <foreignObject>, y defines the y coordinate of the upper left corner of its viewport.
//		 image: For <image>, y defines the y coordinate of the upper left corner of the image.
//		 mask: For <mask>, y defines the y coordinate of the upper left corner of its area of effect. The exact effect
//		   of this attribute is influenced by the maskUnits attribute.
//		 pattern: For <pattern>, y defines the y coordinate of the upper left corner of the tile pattern. The exact
//		   effect of this attribute is influenced by the patternUnits and patternTransform attributes.
//		 rect: For <rect>, y defines the y coordinate of the upper left corner of the shape.
//		 svg: For <svg>, y defines the y coordinate of the upper left corner of its viewport.
//		 text: For <text>, if it contain a single value, y defines the y coordinate on where the content text position
//		   must be placed. The content text position is usually a point on the baseline of the first line of text.
//		   The exact content text position is influenced by some properties like text-anchor, or direction.
//		   If there are multiple values, y defines the y coordinate of each individual glyph from the text.
//		   If there are less values than glyphs, the remaining glyphs are placed in the continuity of the last
//		   positioned glyph. If there are more values than glyphs, extra values are ignored.
//		 tspan: For <tspan>, if it contain a single value, y defines the y coordinate on where the content text position
//		   must be placed. The content text position is usually a point on the baseline of the first line of text.
//		   The exact content text position is influenced by some properties like text-anchor, or direction.
//		   If there are multiple values, y defines the y coordinate of each individual glyph from the text.
//		   If there are less values than glyphs, the remaining glyphs are placed in the continuity of the last positioned
//		   glyph. If there are more values than glyphs, extra values are ignored.
//		 use: For <use>, y defines the y coordinate of the upper left corner of the referenced element.

// Portuguese
//	O atributo y define uma coordenada do eixo y no sistema de coordenadas do usuário.
//
//	Você pode usar este atributo com os seguintes elementos SVG:
//
//	<altGlyph>
//	<cursor>
//	<feBlend>
//	<feColorMatrix>
//	<feComponentTransfer>
//	<feComposite>
//	<feConvolveMatrix>
//	<feDiffuseLighting>
//	<feDisplacementMap>
//	<feDropShadow>
//	<feFlood>
//	<feFuncA>
//	<feFuncB>
//	<feFuncG>
//	<feFuncR>
//	<feGaussianBlur>
//	<feImage>
//	<feMerge>
//	<feMergeNode>
//	<feMorfologia>
//	<feOffset>
//	<fePointLight>
//	<feSpecularLighting>
//	<feSpotLight>
//	<feTile>
//	<feTurbulence>
//	<filtro>
//	<objeto estrangeiro>
//	<glyphRef>
//	<imagem>
//	<máscara>
//	<padrão>
//	<correto>
//	<svg>
//	<símbolo>
//	<texto>
//	<tref>
//	<tspan>
//	<usar>
//
//	Entrada:
//	   y: define uma coordenada do eixo y no sistema de coordenadas do usuário.
//	     feBlend: Para <feBlend>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feColorMatrix: Para <feColorMatrix>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feColorMatrix: Para <feColorMatrix>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feBlend: Para <feBlend>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feColorMatrix: Para <feColorMatrix>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feComponentTransfer: Para <feComponentTransfer>, y define a coordenada y mínima para a renderização
//	       área do primitivo.
//	     feComposite: Para <feComposite>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feConvolveMatrix: Para <feConvolveMatrix>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feDiffuseLighting: Para <feDiffuseLighting>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feDisplacementMap: Para <feDisplacementMap>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feDropShadow: Para <feDropShadow>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feFlood: Para <feFlood>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feFuncA: Para <feFuncA>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feFuncB: Para <feFuncB>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feFuncG: Para <feFuncG>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feFuncR: Para <feFuncR>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feGaussianBlur: Para <feGaussianBlur>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feImage: Para <feImage>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feMerge: Para <feMerge>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feMergeNode: Para <feMergeNode>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feMorphology: Para <feMorphology>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feOffset: Para <feOffset>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     fePointLight: Para <fePointLight>, y define a localização y para a fonte de luz no sistema de coordenadas
//	       definido pelo atributo primitivaUnits no elemento <filter>.
//	     feSpecularLighting: Para <feSpecularLighting>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feSpotLight: Para <feSpotLight>, y define a localização y para a fonte de luz no sistema de coordenadas definido
//	       pelo atributo primitivoUnits no elemento <filter>.
//	     feTile: Para <feMerge>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     feTurbulence: Para <feTurbulence>, y define a coordenada y mínima para a área de renderização da primitiva.
//	     filter: Para <filter>, y define a coordenada y do canto superior esquerdo para a área de renderização do filtro.
//	     ForeignObject: Para <foreignObject>, y define a coordenada y do canto superior esquerdo de sua viewport.
//	     image: Para <image>, y define a coordenada y do canto superior esquerdo da imagem.
//	     mask: Para <mask>, y define a coordenada y do canto superior esquerdo de sua área de efeito. O efeito exato
//	     deste atributo é influenciado pelo atributo maskUnits.
//	     pattern: Para <pattern>, y define a coordenada y do canto superior esquerdo do padrão de ladrilhos. O exato
//	       o efeito desse atributo é influenciado pelos atributos patternUnits e patternTransform.
//	     rect: Para <rect>, y define a coordenada y do canto superior esquerdo da forma.
//	     svg: Para <svg>, y define a coordenada y do canto superior esquerdo de sua viewport.
//	     text: Para <text>, se contiver um único valor, y define a coordenada y em que a posição do texto do conteúdo
//	       deve ser colocado. A posição do texto do conteúdo geralmente é um ponto na linha de base da primeira linha do texto.
//	       A posição exata do texto do conteúdo é influenciada por algumas propriedades, como âncora de texto ou direção.
//	       Se houver vários valores, y define a coordenada y de cada glifo individual do texto.
//	       Se houver menos valores que glifos, os glifos restantes são colocados na continuidade do último
//	       glifo posicionado. Se houver mais valores do que glifos, os valores extras serão ignorados.
//	     tspan: Para <tspan>, se contiver um único valor, y define a coordenada y em que a posição do texto do conteúdo
//	       deve ser colocado. A posição do texto do conteúdo geralmente é um ponto na linha de base da primeira linha do texto.
//	       A posição exata do texto do conteúdo é influenciada por algumas propriedades, como âncora de texto ou direção.
//	       Se houver vários valores, y define a coordenada y de cada glifo individual do texto.
//	       Se houver menos valores que glifos, os glifos restantes são colocados na continuidade do último posicionado
//	     glifo. Se houver mais valores do que glifos, os valores extras serão ignorados.
//	     use: Para <use>, y define a coordenada y do canto superior esquerdo do elemento referenciado.
//
func (e *TagSvgGlobal) Y(y interface{}) (ref *TagSvgGlobal) {
	if converted, ok := y.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "y", converted.String())
		return e
	}

	if converted, ok := y.(time.Time); ok {
		e.selfElement.Call("setAttribute", "y", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "y", y)
	return e
}

// Y1
//
// English:
//
//	The y1 attribute is used to specify the first y-coordinate for drawing an SVG element that requires more than one
//	coordinate. Elements that only need one coordinate use the y attribute instead.
//
//	You can use this attribute with the following SVG elements:
//
// <line>
// <linearGradient>
//
//	Input:
//	    y1: defines the y1 attribute for the first y-coordinate for drawing an SVG element that requires more than one coordinate
//		  line: For <line>, y1 defines the y coordinate of the starting point of the line.
//		  linearGradient: For <linearGradient>, y1 defines the y coordinate of the starting point of the gradient vector
//		    used to map the gradient stop values. The exact behavior of this attribute is influenced by the gradientUnits attributes

//
//
// Portuguese
//
//	O atributo y1 é usado para especificar a primeira coordenada y para desenhar um elemento SVG que requer mais de uma
//	coordenada. Elementos que precisam apenas de uma coordenada usam o atributo y.
//
//	Você pode usar este atributo com os seguintes elementos SVG:
//
// <line>
// <linearGradient>
//
//	Entrada:
//	   y1: define o atributo y1 para a primeira coordenada y para desenhar um elemento SVG que requer mais de uma coordenada
//	     line: Para <line>, y1 define a coordenada y do ponto inicial da linha.
//	     linearGradient: Para <linearGradient>, y1 define a coordenada y do ponto inicial do vetor gradiente
//	       usado para mapear os valores de parada de gradiente. O comportamento exato deste atributo é influenciado pelos atributos gradientUnits
//
func (e *TagSvgGlobal) Y1(y1 interface{}) (ref *TagSvgGlobal) {
	if converted, ok := y1.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "y1", converted.String())
		return e
	}

	if converted, ok := y1.(time.Time); ok {
		e.selfElement.Call("setAttribute", "y1", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "y1", y1)
	return e
}

// Y2
//
// English:
//
//	The y2 attribute is used to specify the second y-coordinate for drawing an SVG element that requires more than
//	one coordinate. Elements that only need one coordinate use the y attribute instead.
//
//	You can use this attribute with the following SVG elements:
//
// <line>
// <linearGradient>
//
//	Input:
//	    y2: defines the y2 attribute for the first y-coordinate for drawing an SVG element that requires more than one coordinate
//		  line: For <line>, y2 defines the y coordinate of the ending point of the line.
//		  linearGradient: For <linearGradient>, y2 defines the y coordinate of the ending point of the gradient vector
//		    used to map the gradient stop values. The exact behavior of this attribute is influenced by the gradientUnits attributes
//
//
// Portuguese
//
//	O atributo y2 é usado para especificar a primeira coordenada y para desenhar um elemento SVG que requer mais de uma
//	coordenada. Elementos que precisam apenas de uma coordenada usam o atributo y.
//
//	Você pode usar este atributo com os seguintes elementos SVG:
//
// <line>
// <linearGradient>
//
//	Entrada:
//	   y2: define o atributo y2 para a primeira coordenada y para desenhar um elemento SVG que requer mais de uma coordenada
//	     line: Para <line>, y2 define a coordenada y do ponto inicial da linha.
//	     linearGradient: Para <linearGradient>, y2 define a coordenada y do ponto inicial do vetor gradiente
//	       usado para mapear os valores de parada de gradiente. O comportamento exato deste atributo é influenciado pelos atributos gradientUnits
//
func (e *TagSvgGlobal) Y2(y2 interface{}) (ref *TagSvgGlobal) {
	if converted, ok := y2.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "y2", converted.String())
		return e
	}

	if converted, ok := y2.(time.Time); ok {
		e.selfElement.Call("setAttribute", "y2", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "y2", y2)
	return e
}

// yChannelSelector
//
// English:
//
//	The yChannelSelector attribute indicates which color channel from in2 to use to displace the pixels in in
//	along the y-axis.
//
//	You can use this attribute with the following SVG elements:
//
// <line>
// <linearGradient>
//
//	Input:
//	    yChannelSelector: attribute indicates which color channel from in2 to use to displace the pixels in in along the y-axis.
//		  R: This keyword specifies that the red color channel of the input image defined in in2 will be used to displace
//		    the pixels of the input image defined in in along the y-axis.
//		  G: This keyword specifies that the green color channel of the input image defined in in2 will be used to
//		    displace the pixels of the input image defined in in along the y-axis.
//		  A: This keyword specifies that the alpha channel of the input image defined in in2 will be used to displace
//		    the pixels of the input image defined in in along the y-axis.
//
//
// Portuguese
//
//	O atributo yChannelSelector indica qual canal de cor de in2 usar para deslocar os pixels em
//	ao longo do eixo y.
//
//	Você pode usar este atributo com os seguintes elementos SVG:
//
//	<linha>
//	<linearGradiente>
//
//	Entrada:
//	    yChannelSelector: atributo indica qual canal de cor de in2 usar para deslocar os pixels ao longo do eixo y.
//	      R: Esta palavra-chave especifica que o canal de cor vermelha da imagem de entrada definida em in2 será usado para deslocar
//	        os pixels da imagem de entrada definidos ao longo do eixo y.
//	      G: Esta palavra-chave especifica que o canal de cor verde da imagem de entrada definida em in2 será usado para
//	        desloca os pixels da imagem de entrada definida ao longo do eixo y.
//	      R: Esta palavra-chave especifica que o canal alfa da imagem de entrada definida em in2 será usado para deslocar
//	        os pixels da imagem de entrada definidos ao longo do eixo y.
//
func (e *TagSvgGlobal) YchannelSelector(yChannelSelector interface{}) (ref *TagSvgGlobal) {
	if converted, ok := yChannelSelector.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "yChannelSelector", converted.String())
		return e
	}

	if converted, ok := yChannelSelector.(time.Time); ok {
		e.selfElement.Call("setAttribute", "yChannelSelector", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "yChannelSelector", yChannelSelector)
	return e
}

// ================================================================================================================ Z
// Z
//
// English:
//
//	The z attribute defines the location along the z-axis for a light source in the coordinate system established by
//	the primitiveUnits attribute on the <filter> element, assuming that, in the initial coordinate system, the positive
//	z-axis comes out towards the person viewing the content and assuming that one unit along the z-axis equals one unit
//	in x and y.
//
//	You can use this attribute with the following SVG elements:
//
//	<fePointLight>
//	<feSpotLight>
//
//	Input:
//	   z: defines the location along the z-axis for a light source in the coordinate system.
//		 fePointLight: For <fePointLight>, z defines the location along the z-axis for the light source in the
//		   coordinate system established by the primitiveUnits attribute on the <filter> element.
//		 feSpotLight: For <feSpotLight>, z defines the location along the z-axis for the light source in the coordinate
//		   system established by the primitiveUnits attribute on the <filter> element.
//
//
// Portuguese
//
//	O atributo z define a localização ao longo do eixo z para uma fonte de luz no sistema de coordenadas estabelecido por
//	o atributo primitivoUnits no elemento <filter>, assumindo que, no sistema de coordenadas inicial, o valor positivo
//	O eixo z sai em direção à pessoa visualizando o conteúdo e assumindo que uma unidade ao longo do eixo z é igual a uma unidade
//	em x e y.
//
//	Você pode usar este atributo com os seguintes elementos SVG:

//	Entrada:
//
//		z: define a localização ao longo do eixo z para uma fonte de luz no sistema de coordenadas.
//		  fePointLight: Para <fePointLight>, z define a localização ao longo do eixo z para a fonte de luz no
//			sistema de coordenadas estabelecido pelo atributo primitivaUnits no elemento <filter>.
//		  feSpotLight: Para <feSpotLight>, z define a localização ao longo do eixo z para a fonte de luz na coordenada
//		    sistema estabelecido pelo atributo primitivaUnits no elemento <filter>.
//
//
func (e *TagSvgGlobal) Z(z interface{}) (ref *TagSvgGlobal) {
	if converted, ok := z.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "z", converted.String())
		return e
	}

	if converted, ok := z.(time.Time); ok {
		e.selfElement.Call("setAttribute", "z", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "z", z)
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
