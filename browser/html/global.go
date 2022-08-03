package html

import (
	"fmt"
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
//	The accumulate attribute controls whether or not an animation is cumulative.
//
//	 Input:
//	   value: controls whether or not an animation is cumulative
//	     const: KSvgAccumulate... (e.g. KSvgAccumulateSum)
//	     any other type: interface{}
//
// It is frequently useful for repeated animations to build upon the previous results, accumulating with each iteration.
// This attribute said to the animation if the value is added to the previous animated attribute's value on each
// iteration.
//
//	Notes:
//	  * This attribute is ignored if the target attribute value does not support addition, or if the animation element
//	    does not repeat;
//	  * This attribute will be ignored if the animation function is specified with only the to attribute.
//
// Português:
//
//	O atributo acumular controla se uma animação é cumulativa ou não.
//
//	 Entrada:
//	   value: controla se uma animação é cumulativa ou não
//	     const: KSvgAccumulate... (ex. KSvgAccumulateSum)
//	     qualquer outro tipo: interface{}
//
// Frequentemente, é útil que as animações repetidas se baseiem nos resultados anteriores, acumulando a cada iteração.
// Este atributo é dito à animação se o valor for adicionado ao valor do atributo animado anterior em cada iteração.
//
//	Notas:
//	  * Esse atributo será ignorado se o valor do atributo de destino não suportar adição ou se o elemento de animação
//	    não se repetir;
//	  * Este atributo será ignorado se a função de animação for especificada apenas com o atributo to.
func (e *TagSvgGlobal) Accumulate(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgAccumulate); ok {
		e.selfElement.Call("setAttribute", "accumulate", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "accumulate", value)
	return e
}

// Additive
//
// English:
//
//	The additive attribute controls whether or not an animation is additive.
//
//	 Input:
//	   value: controls whether or not an animation is additive
//	     const: KSvgAdditive... (e.g. KSvgAdditiveSum)
//	     any other type: interface{}
//
// It is frequently useful to define animation as an offset or delta to an attribute's value, rather than as
// absolute values.
//
// Português:
//
//	O atributo aditivo controla se uma animação é ou não aditiva.
//
//	 Entrada:
//	   value: controla se uma animação é aditiva ou não
//	     const: KSvgAdditive... (ex. KSvgAdditiveSum)
//	     qualquer outro tipo: interface{}
//
// É frequentemente útil definir a animação como um deslocamento ou delta para o valor de um atributo, em vez de
// valores absolutos.
func (e *TagSvgGlobal) Additive(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgAdditive); ok {
		e.selfElement.Call("setAttribute", "additive", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "additive", value)
	return e
}

// AlignmentBaseline #presentation
//
// English:
//
//	The alignment-baseline attribute specifies how an object is aligned with respect to its parent. This property
//	specifies which baseline of this element is to be aligned with the corresponding baseline of the parent.
//	For example, this allows alphabetic baselines in Roman text to stay aligned across font size changes.
//	It defaults to the baseline with the same name as the computed value of the alignment-baseline property.
//
//	 Input:
//	   alignmentBaseline: specifies how an object is aligned with respect to its parent.
//	     string: url(#myClip)
//	     consts KSvgAlignmentBaseline... (e.g. KSvgAlignmentBaselineTextBeforeEdge)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute alignment-baseline can be used as a CSS property.
//	   * SVG 2 introduces some changes to the definition of this property. In particular: the values auto, before-edge,
//	     and after-edge have been removed. For backwards compatibility, text-before-edge may be mapped to text-top and
//	     text-after-edge to text-bottom. Neither text-before-edge nor text-after-edge should be used with the
//	     vertical-align property.
//
// Português:
//
//	O atributo alinhamento-base especifica como um objeto é alinhado em relação ao seu pai. Esta propriedade especifica
//	qual linha de base deste elemento deve ser alinhada com a linha de base correspondente do pai. Por exemplo, isso
//	permite que as linhas de base alfabéticas em texto romano permaneçam alinhadas nas alterações de tamanho da fonte.
//	O padrão é a linha de base com o mesmo nome que o valor calculado da propriedade de linha de base de alinhamento.
//
//	 Input:
//	   alignmentBaseline: especifica como um objeto é alinhado em relação ao seu pai.
//	     string: url(#myClip)
//	     consts KSvgAlignmentBaseline...  (ex. KSvgAlignmentBaselineTextBeforeEdge)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como um atributo de apresentação, a linha de base de alinhamento pode ser usada como uma propriedade CSS.
//	   * O SVG 2 introduz algumas mudanças na definição desta propriedade. Em particular: os valores auto, before-edge e
//	     after-edge foram removidos. Para compatibilidade com versões anteriores, text-before-edge pode ser mapeado para
//	     text-top e text-after-edge para text-bottom. Nem text-before-edge nem text-after-edge devem ser usados com a
//	     propriedade vertical-align.
func (e *TagSvgGlobal) AlignmentBaseline(alignmentBaseline interface{}) (ref *TagSvgGlobal) {
	if converted, ok := alignmentBaseline.(SvgAlignmentBaseline); ok {
		e.selfElement.Call("setAttribute", "alignment-baseline", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// Amplitude
//
// English:
//
//	The amplitude attribute controls the amplitude of the gamma function of a component transfer element when its type
//	attribute is gamma.
//
//	 Input:
//	   amplitude: controls the amplitude of the gamma function
//	     float32: 1.0 = "100%"
//	     any other type: interface{}
//
// Português:
//
//	O atributo amplitude controla à amplitude da função gama de um elemento de transferência de componente quando seu
//	atributo de tipo é gama.
//
//	 Entrada:
//	   amplitude: controla a amplitude da função de gama
//	     float32: 1.0 = "100%"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Amplitude(amplitude interface{}) (ref *TagSvgGlobal) {
	if converted, ok := amplitude.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "amplitude", p)
		return e
	}

	e.selfElement.Call("setAttribute", "amplitude", amplitude)
	return e
}

// AttributeName
//
// English:
//
//	The attributeName attribute indicates the name of the CSS property or attribute of the target element that is going
//	to be changed during an animation.
//
//	 Input:
//	   attributeName: indicates the name of the CSS property or attribute of the target element
//	     const: KAttribute... (e.g. KAttributeSeed = "seed")
//
// Português:
//
//	O atributo attributeName indica o nome da propriedade CSS ou atributo do elemento de destino que será alterado
//	durante uma animação.
//
//	 Entrada:
//	   attributeName: indica o nome da propriedade CSS ou atributo do elemento de destino
//	     const: KAttribute... (ex. KAttributeSeed = "seed")
func (e *TagSvgGlobal) AttributeName(attributeName string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "attributeName", attributeName)
	return e
}

// Azimuth
//
// English:
//
//	The azimuth attribute specifies the direction angle for the light source on the XY plane (clockwise), in degrees
//	from the x axis.
//
//	 Input:
//	   azimuth: specifies the direction angle for the light source on the XY plane
//
// Português:
//
//	O atributo azimute especifica o ângulo de direção da fonte de luz no plano XY (sentido horário), em graus a partir
//	do eixo x.
//
//	 Input:
//	   azimuth: especifica o ângulo de direção para a fonte de luz no plano XY
func (e *TagSvgGlobal) Azimuth(azimuth float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "azimuth", azimuth)
	return e
}

// BaseFrequency
//
// English:
//
//	The baseFrequency attribute represents the base frequency parameter for the noise function of the <feTurbulence> filter primitive.
//
//	 Input:
//	   baseFrequency: represents the base frequency parameter for the noise function
//
// Português:
//
//	O atributo baseFrequency representa o parâmetro de frequência base para a função de ruído da primitiva de filtro <feTurbulence>.
//
//	 Entrada:
//	   baseFrequency: representa o parâmetro de frequência base para a função de ruído
func (e *TagSvgGlobal) BaseFrequency(baseFrequency float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "baseFrequency", baseFrequency)
	return e
}

// BaselineShift #presentation
//
// English:
//
//	The baseline-shift attribute allows repositioning of the dominant-baseline relative to the dominant-baseline of the
//	parent text content element. The shifted object might be a sub- or superscript.
//
//	 Input:
//	   baselineShift: allows repositioning of the dominant-baseline relative to the dominant-baseline of the parent text
//	   content element.
//	     float32: 0.05 = "5%"
//	     string: "5%"
//	     consts KSvgBaselineShift... (e.g. KSvgBaselineShiftAuto)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute baseline-shift can be used as a CSS property.
//	   * This property is going to be deprecated and authors are advised to use vertical-align instead.
//
// Português:
//
//	O atributo baseline-shift permite o reposicionamento da linha de base dominante em relação à linha de base dominante
//	do elemento de conteúdo de texto pai. O objeto deslocado pode ser um sub ou sobrescrito.
//
//	 Input:
//	   baselineShift: permite o reposicionamento da linha de base dominante em relação à linha de base dominante do
//	   elemento de conteúdo de texto pai.
//	     float32: 0.05 = "5%"
//	     string: "5%"
//	     consts KSvgBaselineShift... (ex. KSvgBaselineShiftAuto)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, baseline-shift pode ser usado como propriedade CSS.
//	   * Essa propriedade será preterida e os autores são aconselhados a usar alinhamento vertical.
func (e *TagSvgGlobal) BaselineShift(baselineShift interface{}) (ref *TagSvgGlobal) {
	if converted, ok := baselineShift.(SvgBaselineShift); ok {
		e.selfElement.Call("setAttribute", "baseline-shift", converted.String())
		return e
	}

	if converted, ok := baselineShift.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "baseline-shift", p)
		return e
	}

	e.selfElement.Call("setAttribute", "baseline-shift", baselineShift)
	return e
}

// Begin
//
// English:
//
//	The begin attribute defines when an animation should begin or when an element should be discarded.
//
//	 Input:
//	   begin: defines when an animation should begin or when an element should be discarded.
//	     offset-value: This value defines a clock-value that represents a point in time relative to the beginning of the
//	       SVG document (usually the load or DOMContentLoaded event). Negative values are valid.
//	       (e.g. time.Second*5 or "5s")
//	     syncbase-value: This value defines a syncbase and an optional offset from that syncbase. The element's
//	       animation start time is defined relative to the begin or active end of another animation.
//	       A valid syncbase-value consists of an ID reference to another animation element followed by a dot and either
//	       begin or end to identify whether to synchronize with the beginning or active end of the referenced animation
//	       element. An optional offset value as defined in <offset-value> can be appended.
//	       (e.g. "0s;third.end", "first.end" or "second.end")
//	     event-value: This value defines an event and an optional offset that determines the time at which the element's
//	       animation should begin. The animation start time is defined relative to the time that the specified event is
//	       fired.
//	       A valid event-value consists of an element ID followed by a dot and one of the supported events for that
//	       element. All valid events (not necessarily supported by all elements) are defined by the DOM and HTML
//	       specifications. Those are: 'focus', 'blur', 'focusin', 'focusout', 'activate', 'auxclick', 'click',
//	       'dblclick', 'mousedown', 'mouseenter', 'mouseleave', 'mousemove', 'mouseout', 'mouseover', 'mouseup',
//	       'wheel', 'beforeinput', 'input', 'keydown', 'keyup', 'compositionstart', 'compositionupdate',
//	       'compositionend', 'load', 'unload', 'abort', 'error', 'select', 'resize', 'scroll', 'beginEvent', 'endEvent',
//	       and 'repeatEvent'. An optional offset value as defined in <offset-value> can be appended.
//	       (e.g. "startButton.click")
//	     repeat-value: This value defines a qualified repeat event. The element animation start time is defined relative
//	       to the time that the repeat event is raised with the specified iteration value.
//	       A valid repeat value consists of an element ID followed by a dot and the function repeat() with an integer
//	       value specifying the number of repetitions as parameter. An optional offset value as defined in
//	       <offset-value> can be appended.
//	       (e.g. "0s;myLoop.end", "myLoop.begin", "myLoop.repeat(1)" or "myLoop.repeat(2)")
//	     accessKey-value: This value defines an access key that should trigger the animation. The element animation will
//	       begin when the user presses the specified key.
//	       A valid accessKey-value consists of the function accessKey() with the character to be input as parameter. An
//	       optional offset value as defined in <offset-value> can be appended.
//	       (e.g. "accessKey(s)")
//	     wallclock-sync-value: This value defines the animation start time as a real-world clock time.
//	       A valid wallclock-sync-value consists of the function wallclock() with a time value as parameter. The time
//	       syntax is based upon the syntax defined in ISO 8601.
//	       (e.g. time.Now() )
//	     indefinite: The begin of the animation will be determined by a beginElement() method call or a hyperlink
//	       targeted to the element.
//	       (e.g. "infinite")
//
// The attribute value is a semicolon separated list of values. The interpretation of a list of start times is detailed
// in the SMIL specification in "Evaluation of begin and end time lists". Each individual value can be one of the
// following: <offset-value>, <syncbase-value>, <event-value>, <repeat-value>, <accessKey-value>, <wallclock-sync-value>
// or the keyword 'indefinite'.
//
// Português:
//
//	O atributo begin define quando uma animação deve começar ou quando um elemento deve ser descartado.
//
//	 Entrada:
//	   begin: define quando uma animação deve começar ou quando um elemento deve ser descartado.
//	     offset-value: Esse valor define um valor de relógio que representa um ponto no tempo relativo ao início do
//	       documento SVG (geralmente o evento load ou DOMContentLoaded). Valores negativos são válidos.
//	       (e.g. time.Second*5 or "5s")
//	     syncbase-value: Esse valor define uma base de sincronização e um deslocamento opcional dessa base de
//	       sincronização. A hora de início da animação do elemento é definida em relação ao início ou fim ativo de outra
//	       animação.
//	       Um valor syncbase válido consiste em uma referência de ID para outro elemento de animação seguido por um
//	       ponto e um início ou fim para identificar se deve ser sincronizado com o início ou o final ativo do elemento
//	       de animação referenciado. Um valor de deslocamento opcional conforme definido em <offset-value> pode ser
//	       anexado.
//	       (e.g. "0s;third.end", "first.end" or "second.end")
//	     event-value: Esse valor define um evento e um deslocamento opcional que determina a hora em que a animação do
//	       elemento deve começar. A hora de início da animação é definida em relação à hora em que o evento especificado
//	       é acionado.
//	       Um valor de evento válido consiste em um ID de elemento seguido por um ponto e um dos eventos com suporte
//	       para esse elemento. Todos os eventos válidos (não necessariamente suportados por todos os elementos) são
//	       definidos pelas especificações DOM e HTML. Esses valores são: 'focus', 'blur', 'focusin', 'focusout',
//	       'activate', 'auxclick', 'click', 'dblclick', 'mousedown', 'mouseenter', 'mouseleave', 'mousemove',
//	       'mouseout', 'mouseover', 'mouseup', 'wheel', 'beforeinput', 'input', 'keydown', 'keyup', 'compositionstart',
//	       'compositionupdate', 'compositionend', 'load', 'unload', 'abort', 'error', 'select', 'resize', 'scroll',
//	       'beginEvent', 'endEvent', e 'repeatEvent'. Um valor de deslocamento opcional conforme definido em
//	       <offset-value> pode ser anexado.
//	       (e.g. "startButton.click")
//	     repeat-value: Esse valor define um evento de repetição qualificado. A hora de início da animação do elemento é
//	       definida em relação à hora em que o evento de repetição é gerado com o valor de iteração especificado.
//	       Um valor de repetição válido consiste em um ID de elemento seguido por um ponto e a função repeat() com um
//	       valor inteiro especificando o número de repetições como parâmetro. Um valor de deslocamento opcional conforme
//	       definido em <offset-value> pode ser anexado.
//	       (e.g. "0s;myLoop.end", "myLoop.begin", "myLoop.repeat(1)" or "myLoop.repeat(2)")
//	     accessKey-value: Este valor define uma chave de acesso que deve acionar a animação. A animação do elemento
//	       começará quando o usuário pressionar a tecla especificada.
//	       Um valor válido de accessKey consiste na função accessKey() com o caractere a ser inserido como parâmetro.
//	       Um valor de deslocamento opcional conforme definido em <valor de deslocamento> pode ser anexado.
//	       (e.g. "accessKey(s)")
//	     wallclock-sync-value: Esse valor define a hora de início da animação como uma hora do relógio do mundo real.
//	       Um valor wallclock-sync válido consiste na função wallclock() com um valor de tempo como parâmetro. A sintaxe
//	       de tempo é baseada na sintaxe definida na ISO 8601.
//	       (e.g. time.Now() )
//	     indefinite: O início da animação será determinado por uma chamada de método beginElement() ou um hiperlink
//	       direcionado ao elemento.
//	       (e.g. "infinite")
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

	if converted, ok := begin.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "baseline-shift", p)
		return e
	}

	e.selfElement.Call("setAttribute", "begin", begin)
	return e
}

// Bias
//
// English:
//
//	The bias attribute shifts the range of the filter. After applying the kernelMatrix of the <feConvolveMatrix> element
//	to the input image to yield a number and applied the divisor attribute, the bias attribute is added to each
//	component. This allows representation of values that would otherwise be clamped to 0 or 1.
//
//	 Input:
//	   bias: shifts the range of the filter
//
// Português:
//
//	O atributo bias muda o intervalo do filtro. Depois de aplicar o kernelMatrix do elemento <feConvolveMatrix> à imagem
//	de entrada para gerar um número e aplicar o atributo divisor, o atributo bias é adicionado a cada componente. Isso
//	permite a representação de valores que de outra forma seriam fixados em 0 ou 1.
//
//	 Entrada:
//	   bias: muda o intervalo do filtro
func (e *TagSvgGlobal) Bias(bias float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "bias", bias)
	return e
}

// By
//
// English:
//
//	The by attribute specifies a relative offset value for an attribute that will be modified during an animation.
//
//	 Input:
//	   by: specifies a relative offset value for an attribute
//
// The starting value for the attribute is either indicated by specifying it as value for the attribute given in the
// attributeName or the from attribute.
//
// Português:
//
//	O atributo by especifica um valor de deslocamento relativo para um atributo que será modificado durante uma
//	animação.
//
//	 Entrada:
//	   by: especifica um valor de deslocamento relativo para um atributo
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
//	The calcMode attribute specifies the interpolation mode for the animation.
//
//	 Input:
//	   KSvgCalcModeDiscrete: This specifies that the animation function will jump from one value to the next without
//	     any interpolation.
//	   KSvgCalcModeLinear: Simple linear interpolation between values is used to calculate the animation function.
//	     Except for <animateMotion>, this is the default value.
//	   KSvgCalcModePaced: Defines interpolation to produce an even pace of change across the animation.
//	   KSvgCalcModeSpline: Interpolates from one value in the values list to the next according to a time function
//	     defined by a cubic Bézier spline. The points of the spline are defined in the keyTimes attribute, and the
//	     control points for each interval are defined in the keySplines attribute.
//
// The default mode is linear, however if the attribute does not support linear interpolation (e.g. for strings), the
// calcMode attribute is ignored and discrete interpolation is used.
//
//	Notes:
//	  Default value: KSvgCalcModePaced
//
// Português:
//
//	O atributo calcMode especifica o modo de interpolação para a animação.
//
//	 Entrada:
//	   KSvgCalcModeDiscrete: Isso especifica que a função de animação saltará de um valor para o próximo sem qualquer
//	     interpolação.
//	   KSvgCalcModeLinear: A interpolação linear simples entre valores é usada para calcular a função de animação.
//	     Exceto para <animateMotion>, este é o valor padrão.
//	   KSvgCalcModePaced: Define a interpolação para produzir um ritmo uniforme de mudança na animação.
//	   KSvgCalcModeSpline: Interpola de um valor na lista de valores para o próximo de acordo com uma função de tempo
//	     definida por uma spline de Bézier cúbica. Os pontos do spline são definidos no atributo keyTimes e os pontos
//	     de controle para cada intervalo são definidos no atributo keySplines.
//
// O modo padrão é linear, no entanto, se o atributo não suportar interpolação linear (por exemplo, para strings), o
// atributo calcMode será ignorado e a interpolação discreta será usada.
//
//	Notas:
//	  * Valor padrão: KSvgCalcModePaced
func (e *TagSvgGlobal) CalcMode(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgCalcMode); ok {
		e.selfElement.Call("setAttribute", "calcMode", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "calcMode", value)
	return e
}

// Class #styling
//
// English:
//
// Assigns a class name or set of class names to an element. You may assign the same class name or names to any number
// of elements, however, multiple class names must be separated by whitespace characters.
//
//	Input:
//	  class: Assigns a class name or set of class names to an element
//
// An element's class name serves two key roles:
//   - As a style sheet selector, for when an author assigns style information to a set of elements.
//   - For general use by the browser.
//
// Português:
//
// Atribui um nome de classe ou um conjunto de nomes de classe à um elemento. Você pode atribuir o mesmo nome ou nomes
// de classe a qualquer número de elementos, no entanto, vários nomes de classe devem ser separados por caracteres de
// espaço em branco.
//
//	Entrada:
//	  class: Atribui um nome de classe ou um conjunto de nomes de classe à um elemento.
//
// O nome de classe de um elemento tem duas funções principais:
//   - Como um seletor de folha de estilo, para quando um autor atribui informações de estilo a um conjunto de
//     elementos.
//   - Para uso geral pelo navegador.
func (e *TagSvgGlobal) Class(class string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "class", class)
	return e
}

// ClipPathUnits
//
// English:
//
//	The clipPathUnits attribute indicates which coordinate system to use for the contents of the <clipPath> element.
//
//	 Input:
//	   clipPathUnits: indicates which coordinate system to used
//	     KSvgClipPathUnits... (e.g. KSvgClipPathUnitsUserSpaceOnUse)
//
// Português:
//
//	O atributo clipPathUnits indica qual sistema de coordenadas deve ser usado para o conteúdo do elemento <clipPath>.
//
//	 Input:
//	   clipPathUnits: indica qual sistema de coordenadas deve ser usado
//	     KSvgClipPathUnits... (ex. KSvgClipPathUnitsUserSpaceOnUse)
func (e *TagSvgGlobal) ClipPathUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgClipPathUnits); ok {
		e.selfElement.Call("setAttribute", "clipPathUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "clipPathUnits", value)
	return e
}

// ClipPath #presentation
//
// English:
//
//	It binds the element it is applied to with a given <clipPath> element.
//
//	 Input:
//	   clipPath: the element it is applied
//	     (e.g. "url(#myClip)", "circle() fill-box", "circle() stroke-box" or "circle() view-box")
//
// Português:
//
//	Ele associa o elemento ao qual é aplicado a um determinado elemento <clipPath>.
//
//	 Entrada:
//	   clipPath: elemento ao qual é aplicado
//	     (ex. "url(#myClip)", "circle() fill-box", "circle() stroke-box" ou "circle() view-box")
func (e *TagSvgGlobal) ClipPath(clipPath string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "clip-path", clipPath)
	return e
}

// ClipRule #presentation
//
// English:
//
//	It indicates how to determine what side of a path is inside a shape in order to know how a <clipPath> should clip
//	its target.
//
//	 Input:
//	   value: side of a path
//	     const: KSvgClipRule... (e.g. KSvgClipRuleNonzero)
//	     any other type: interface{}
//
// Português:
//
//	Ele indica como determinar qual lado de um caminho está dentro de uma forma para saber como um <clipPath> deve
//	recortar seu destino.
//
//	 Input:
//	   value: lado de um caminho
//	     const: KSvgClipRule... (e.g. KSvgClipRuleNonzero)
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) ClipRule(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgClipRule); ok {
		e.selfElement.Call("setAttribute", "clip-rule", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "clip-rule", value)
	return e
}

// Color #presentation
//
// English:
//
//	It provides a potential indirect value (currentcolor) for the fill, stroke, stop-color, flood-color and
//	lighting-color presentation attributes.
//
//	 Input:
//	   value: potential indirect value of color
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, color can be used as a CSS property. See CSS color for further information.
//
// Português:
//
//	Ele fornece um valor indireto potencial (currentcolor) para os atributos de apresentação de preenchimento, traçado,
//	cor de parada, cor de inundação e cor de iluminação.
//
//	 Entrada:
//	   value: valor indireto potencial da cor
//	     string: ex. "black"
//	     factory: ex. factoryColor.NewYellow()
//	     RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, a cor pode ser usada como propriedade CSS. Veja cor CSS para mais informações.
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
// The color-interpolation attribute specifies the color space for gradient interpolations, color animations, and alpha
// compositing.
//
//	Input:
//	  value: specifies the color space for gradient interpolations
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
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
//	Notes:
//	  * For filter effects, the color-interpolation-filters property controls which color space is used.
//	  * As a presentation attribute, color-interpolation can be used as a CSS property.
//
// Português:
//
// O atributo color-interpolation especifica o espaço de cores para interpolações de gradiente, animações de cores e
// composição alfa.
//
//	Entrada:
//	  value: especifica o espaço de cores para interpolações de gradiente
//	    string: ex. "black"
//	    factory: ex. factoryColor.NewYellow()
//	    RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
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
//	Notas:
//	  * Para efeitos de filtro, a propriedade color-interpolation-filters controla qual espaço de cor é usado.
//	  * Como atributo de apresentação, a interpolação de cores pode ser usada como uma propriedade CSS.
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
// The color-interpolation-filters attribute specifies the color space for imaging operations performed via filter
// effects.
//
//	Input:
//	  value: specifies the color space for imaging operations
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
//	Notes:
//	  * This property just has an affect on filter operations. Therefore, it has no effect on filter primitives like
//	    <feOffset>, <feImage>, <feTile> or <feFlood>;
//	  * color-interpolation-filters has a different initial value than color-interpolation. color-interpolation-filters
//	    has an initial value of linearRGB, whereas color-interpolation has an initial value of sRGB. Thus, in the
//	    default case, filter effects operations occur in the linearRGB color space, whereas all other color
//	    interpolations occur by default in the sRGB color space;
//	  * It has no affect on filter functions, which operate in the sRGB color space;
//	  * As a presentation attribute, color-interpolation-filters can be used as a CSS property.
//
// Português:
//
// O atributo color-interpolation-filters especifica o espaço de cores para operações de imagem realizadas por meio de
// efeitos de filtro.
//
//	Entrada:
//	  value: especifica o espaço de cores para operações de imagem
//	    string: ex. "black"
//	    factory: ex. factoryColor.NewYellow()
//	    RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Esta propriedade afeta apenas as operações de filtro. Portanto, não tem efeito em primitivos de filtro como
//	    <feOffset>, <feImage>, <feTile> ou <feFlood>.
//	  * color-interpolation-filters tem um valor inicial diferente de color-interpolation. color-interpolation-filters
//	    tem um valor inicial de linearRGB, enquanto color-interpolation tem um valor inicial de sRGB. Assim, no caso
//	    padrão, as operações de efeitos de filtro ocorrem no espaço de cores linearRGB, enquanto todas as outras
//	    interpolações de cores ocorrem por padrão no espaço de cores sRGB.
//	  * Não afeta as funções de filtro, que operam no espaço de cores sRGB.
//	  * Como atributo de apresentação, os filtros de interpolação de cores podem ser usados como uma propriedade CSS.
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
// The crossorigin attribute, valid on the <image> element, provides support for CORS, defining how the element handles
// crossorigin requests, thereby enabling the configuration of the CORS requests for the element's fetched data.
//
//	Input:
//	  value: provides support for CORS, defining how the element handles crossorigin requests
//	    const: KSvgCrossOrigin... (e.g.: KSvgCrossOriginUseCredentials)
//	    any other type: interface{}
//
// Português:
//
// O atributo crossorigin, válido no elemento <image>, fornece suporte para CORS, definindo como o elemento trata as
// requisições de origem cruzada, permitindo assim a configuração das requisições CORS para os dados buscados do
// elemento.
//
//	Entrada:
//	  value: fornece suporte para CORS, definindo como o elemento lida com solicitações de origem cruzada
//	    const: KSvgCrossOrigin... (ex.: KSvgCrossOriginUseCredentials)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) CrossOrigin(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgCrossOrigin); ok {
		e.selfElement.Call("setAttribute", "crossorigin", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "crossorigin", value)
	return e
}

// Cursor
//
// English:
//
// The cursor attribute specifies the mouse cursor displayed when the mouse pointer is over an element.
//
//	Input:
//	  value: specifies the mouse cursor displayed when the mouse pointer is over an element
//	    const: KSvgCursor... (e.g.: KSvgCursorMove)
//	    any other type: interface{}
//
// This attribute behaves exactly like the css cursor property except that if the browser supports the <cursor> element,
// you should be able to use it with the <funciri> notation.
//
// As a presentation attribute, it also can be used as a property directly inside a CSS stylesheet, see css cursor for
// further information.
//
// Português:
//
// O atributo cursor especifica o cursor do mouse exibido quando o ponteiro do mouse está sobre um elemento.
//
//	Entrada:
//	  value: especifica o cursor do mouse exibido quando o ponteiro do mouse está sobre um elemento
//	    const: KSvgCursor... (ex.: KSvgCursorMove)
//	    qualquer outro tipo: interface{}
//
// Este atributo se comporta exatamente como a propriedade cursor css, exceto que, se o navegador suportar o elemento
// <cursor>, você poderá usá-lo com a notação <funciri>.
//
// Como atributo de apresentação, também pode ser usado como propriedade diretamente dentro de uma folha de estilo CSS,
// veja cursor css para mais informações.
func (e *TagSvgGlobal) Cursor(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgCursor); ok {
		e.selfElement.Call("setAttribute", "cursor", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "cursor", value)
	return e
}

// Cx
//
// English:
//
//	The cx attribute define the x-axis coordinate of a center point.
//
//	 Input:
//	   value: define the x-axis coordinate
//	     float32: 0.05 = "5%"
//	     any other type: interface{}
//
// Português:
//
//	O atributo cx define a coordenada do eixo x de um ponto central.
//
//	 Entrada:
//	   value: define a coordenada do eixo x
//	     float32: 0.05 = "5%"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Cx(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "cx", p)
		return e
	}

	e.selfElement.Call("setAttribute", "cx", value)
	return e
}

// Cy
//
// English:
//
// The cy attribute define the y-axis coordinate of a center point.
//
//	Input:
//	  value: define the y-axis coordinate
//	    float32: 0.05 = "5%"
//	    any other type: interface{}
//
// Português:
//
//	O atributo cy define a coordenada do eixo y de um ponto central.
//
//	 Entrada:
//	   value: define a coordenada do eixo y
//	     float32: 0.05 = "5%"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Cy(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "cy", p)
		return e
	}

	e.selfElement.Call("setAttribute", "cy", value)
	return e
}

// D #presentation
//
// English:
//
//	The d attribute defines a path to be drawn.
//
//	 Input:
//	   d: path to be drawn
//	     factory: factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()
//	     string: "M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z"
//	     any other type: interface{}
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
//	O atributo d define um caminho a ser desenhado.
//
//	 Entrada:
//	   d: caminho a ser desenhado
//	     factory: factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()
//	     string: "M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z"
//	     qualquer outro tipo: interface{}
//
// Uma definição de caminho é uma lista de comandos de caminho em que cada comando é composto por uma letra de comando
// e números que representam os parâmetros do comando. Os comandos são detalhados abaixo.
//
// Você pode usar este atributo com os seguintes elementos SVG: <path>, <glyph>, <missing-glyph>.
//
// d é um atributo de apresentação e, portanto, também pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) D(d interface{}) (ref *TagSvgGlobal) {
	if converted, ok := d.(*SvgPath); ok {
		e.selfElement.Call("setAttribute", "d", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "d", d)
	return e
}

// DiffuseConstant
//
// English:
//
// The diffuseConstant attribute represents the kd value in the Phong lighting model. In SVG, this can be any
// non-negative number.
//
//	Input:
//	  diffuseConstant: represents the kd value in the Phong lighting model
//
// It's used to determine the final RGB value of a given pixel. The brighter the lighting-color, the smaller this number
// should be.
//
// Português:
//
// O atributo difusoConstant representa o valor kd no modelo de iluminação Phong. Em SVG, pode ser qualquer número
// não negativo.
//
//	Entrada:
//	  diffuseConstant: representa o valor kd no modelo de iluminação Phong
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
// The direction attribute specifies the inline-base direction of a <text> or <tspan> element. It defines the start
// and end points of a line of text as used by the text-anchor and inline-size properties. It also may affect the
// direction in which characters are positioned if the unicode-bidi property's value is either embed or bidi-override.
//
//	Input:
//	  value: specifies the inline-base direction of a <text> or <tspan> element
//	    const: KSvgDirection... (e.g. KSvgDirectionRtl)
//	    any other type: interface{}
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
//	Notes:
//	  * As a presentation attribute, direction can be used as a CSS property. See css direction for further
//	    information.
//
// Português:
//
// O atributo direction especifica a direção da base embutida de um elemento <text> ou <tspan>. Ele define os pontos
// inicial e final de uma linha de texto conforme usado pelas propriedades text-anchor e inline-size.
// Também pode afetar a direção na qual os caracteres são posicionados se o valor da propriedade unicode-bidi for
// incorporado ou substituído por bidi.
//
//	Input:
//	  value: especifica a direção da base inline de um elemento <text> ou <tspan>
//	    const: KSvgDirection... (e.g. KSvgDirectionRtl)
//	    qualquer outro tipo: interface{}
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
//	Notas:
//	  * Como atributo de apresentação, a direção pode ser usada como uma propriedade CSS. Veja a direção do CSS para
//	    mais informações.
func (e *TagSvgGlobal) Direction(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgDirection); ok {
		e.selfElement.Call("setAttribute", "direction", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "direction", value)
	return e
}

// Display
//
// English:
//
//	The display attribute lets you control the rendering of graphical or container elements.
//
//	 Input:
//	   value: control the rendering of graphical or container elements
//	     nil: display="none"
//	     const: KSvgDisplay... (e.g. KSvgDisplayBlock)
//	     any other type: interface{}
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
//   - If display is set to none on a <tspan>, <tref>, or <altGlyph> element, then the text string is ignored for the
//     purposes of text layout.
//   - Regarding events, if display is set to none, the element receives no events.
//   - The geometry of a graphics element with display set to none is not included in bounding box and clipping paths
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
//	Notes:
//	  * As a presentation attribute, display can be used as a CSS property. See css display for further information.
//
// Português:
//
//	O atributo display permite controlar a renderização de elementos gráficos ou de contêiner.
//
//	 Entrada:
//	   value: controlar a renderização de elementos gráficos ou de contêiner
//	     nil: display="none"
//	     const: KSvgDisplay... (ex. KSvgDisplayBlock)
//	     qualquer outro tipo: interface{}
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
//   - Se display for definido como none em um elemento <tspan>, <tref> ou <altGlyph>, a string de texto será ignorada
//     para fins de layout de texto.
//   - Com relação aos eventos, se display estiver definido como none, o elemento não recebe eventos.
//   - A geometria de um elemento gráfico com exibição definida como nenhum não é incluída nos cálculos da caixa
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
//	Notas:
//	  * Como atributo de apresentação, display pode ser usado como propriedade CSS. Consulte a exibição css para obter
//	    mais informações.
func (e *TagSvgGlobal) Display(value interface{}) (ref *TagSvgGlobal) {
	if value == nil {
		e.selfElement.Call("setAttribute", "display", "none")
		return e
	}

	if converted, ok := value.(SvgDisplay); ok {
		e.selfElement.Call("setAttribute", "display", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "display", value)
	return e
}

// Divisor
//
// English:
//
//	The divisor attribute specifies the value by which the resulting number of applying the kernelMatrix of a
//	<feConvolveMatrix> element to the input image color value is divided to yield the destination color value.
//
//	 Input:
//	   divisor: specifies the divisor value to apply to the original color
//
// A divisor that is the sum of all the matrix values tends to have an evening effect on the overall color intensity of
// the result.
//
// Português:
//
//	O atributo divisor especifica o valor pelo qual o número resultante da aplicação do kernelMatrix de um elemento
//	<feConvolveMatrix> ao valor da cor da imagem de entrada é dividido para gerar o valor da cor de destino.
//
//	 Entrada:
//	   divisor: especifica o valor do divisor a ser aplicado na cor original
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
// The dominant-baseline attribute specifies the dominant baseline, which is the baseline used to align the box's text
// and inline-level contents. It also indicates the default alignment baseline of any boxes participating in baseline
// alignment in the box's alignment context.
//
//	Input:
//	  value: is the baseline used to align the box's text and inline-level contents
//	    const: KSvgDominantBaseline... (e.g. KSvgDominantBaselineHanging)
//	    any other type: interface{}
//
// It is used to determine or re-determine a scaled-baseline-table. A scaled-baseline-table is a compound value with
// three components:
//
//  1. a baseline-identifier for the dominant-baseline,
//  2. a baseline-table, and
//  3. a baseline-table font-size.
//
// Some values of the property re-determine all three values. Others only re-establish the baseline-table font-size.
// When the initial value, auto, would give an undesired result, this property can be used to explicitly set the desired
// scaled-baseline-table.
//
// If there is no baseline table in the nominal font, or if the baseline table lacks an entry for the desired baseline,
// then the browser may use heuristics to determine the position of the desired baseline.
//
//	Notes:
//	  * As a presentation attribute, dominant-baseline can be used as a CSS property.
//
// Português:
//
// O atributo linha de base dominante especifica a linha de base dominante, que é a linha de base usada para alinhar o
// texto da caixa e o conteúdo do nível embutido. Também indica a linha de base de alinhamento padrão de todas as caixas
// que participam do alinhamento da linha de base no contexto de alinhamento da caixa.
//
//	Entrada:
//	  value: é a linha de base usada para alinhar o texto da caixa e o conteúdo embutido
//	    const: KSvgDominantBaseline... (ex. KSvgDominantBaselineHanging)
//	    qualquer outro tipo: interface{}
//
// Ele é usado para determinar ou re-determinar uma tabela de linha de base dimensionada. Uma tabela de linha de base
// dimensionada é um valor composto com três componentes:
//
//  1. um identificador de linha de base para a linha de base dominante,
//  2. uma tabela de linha de base, e
//  3. um tamanho de fonte da tabela de linha de base.
//
// Alguns valores da propriedade redeterminam todos os três valores. Outros apenas restabelecem o tamanho da fonte da
// tabela de linha de base. Quando o valor inicial, auto, daria um resultado indesejado, essa propriedade pode ser usada
// para definir explicitamente a tabela de linha de base dimensionada desejada.
//
// Se não houver nenhuma tabela de linha de base na fonte nominal, ou se a tabela de linha de base não tiver uma entrada
// para a linha de base desejada, o navegador poderá usar heurística para determinar a posição da linha de base
// desejada.
//
//	Notas:
//	  * Como atributo de apresentação, a linha de base dominante pode ser usada como uma propriedade CSS.
func (e *TagSvgGlobal) DominantBaseline(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgDominantBaseline); ok {
		e.selfElement.Call("setAttribute", "dominant-baseline", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "dominant-baseline", value)
	return e
}

// Dur
//
// English:
//
//	The dur attribute indicates the simple duration of an animation.
//
//	 Input:
//	   dur: indicates the simple duration of an animation.
//	     KSvgDur... (e.g. KSvgDurIndefinite)
//	     time.Duration (e.g. time.Second * 5)
//
//	 Notes:
//	   * The interpolation will not work if the simple duration is indefinite (although this may still be useful for
//	     <set> elements).
//
// Português:
//
//	O atributo dur indica a duração simples de uma animação.
//
//	 Entrada:
//	   dur: indica a duração simples de uma animação.
//	     KSvgDur... (ex. KSvgDurIndefinite)
//	     time.Duration (ex. time.Second * 5)
//
//	 Notas:
//	   * A interpolação não funcionará se a duração simples for indefinida (embora isso ainda possa ser útil para
//	     elementos <set>).
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
//	The dx attribute indicates a shift along the x-axis on the position of an element or its content.
//
//	 Input:
//	   dx: indicates a shift along the x-axis on the position of an element or its content.
//	     []float32: []float64{0.0, 0.1} = "0% 10%"
//	     []float64: []float64{0.0, 10.0} = "0 10"
//	     float32: 0.1 = "10%"
//	     float64: 10.0 = "10"
//	     any other type: interface{}
//
// Portuguese
//
//	O atributo dx indica um deslocamento ao longo do eixo x na posição de um elemento ou seu conteúdo.
//
//	 Entrada:
//	   dx: indica um deslocamento ao longo do eixo x na posição de um elemento ou seu conteúdo.
//	     []float32: []float64{0.0, 0.1} = "0% 10%"
//	     []float64: []float64{0.0, 10.0} = "0 10"
//	     float32: 0.1 = "10%"
//	     float64: 10.0 = "10"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Dx(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]color.RGBA); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += RGBAToJs(v) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dx", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float32); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "% "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dx", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dx", valueStr[:length])
		return e
	}

	if converted, ok := value.([]time.Duration); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += v.String() + " "
		}
		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dx", valueStr[:length])
		return e
	}

	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "dx", converted.String())
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "dx", p)
		return e
	}

	if converted, ok := value.(float64); ok {
		p := strconv.FormatFloat(converted, 'g', -1, 64)
		e.selfElement.Call("setAttribute", "dx", p)
		return e
	}

	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "dx", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "dx", value)
	return e
}

// Dy
//
// English:
//
//	The dy attribute indicates a shift along the y-axis on the position of an element or its content.
//
//	 Input:
//	   dy: indicates a shift along the y-axis on the position of an element or its content.
//	     []float32: []float64{0.0, 0.1} = "0% 10%"
//	     []float64: []float64{0.0, 10.0} = "0 10"
//	     float32: 0.1 = "10%"
//	     float64: 10.0 = "10"
//	     any other type: interface{}
//
// Portuguese
//
//	O atributo dy indica um deslocamento ao longo do eixo y na posição de um elemento ou seu conteúdo.
//
//	 Entrada:
//	   dy: indica um deslocamento ao longo do eixo y na posição de um elemento ou seu conteúdo.
//	     []float32: []float64{0.0, 0.1} = "0% 10%"
//	     []float64: []float64{0.0, 10.0} = "0 10"
//	     float32: 0.1 = "10%"
//	     float64: 10.0 = "10"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Dy(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]color.RGBA); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += RGBAToJs(v) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dy", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float32); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "% "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dy", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dy", valueStr[:length])
		return e
	}

	if converted, ok := value.([]time.Duration); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += v.String() + " "
		}
		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "dy", valueStr[:length])
		return e
	}

	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "dy", converted.String())
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "dy", p)
		return e
	}

	if converted, ok := value.(float64); ok {
		p := strconv.FormatFloat(converted, 'g', -1, 64)
		e.selfElement.Call("setAttribute", "dy", p)
		return e
	}

	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "dy", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "dy", value)
	return e
}

// EdgeMode
//
// English:
//
// The edgeMode attribute determines how to extend the input image as necessary with color values so that the matrix
// operations can be applied when the kernel is positioned at or near the edge of the input image.
//
//	Input:
//	  value: determines how to extend the input image
//	    const: KSvgEdgeMode... (e.g. KSvgEdgeModeWrap)
//	    any other type: interface{}
//
// # Portuguese
//
// O atributo edgeMode determina como estender a imagem de entrada conforme necessário com valores de cor para que as
// operações de matriz possam ser aplicadas quando o kernel estiver posicionado na borda da imagem de entrada ou
// próximo a ela.
//
//	Entrada:
//	  value: determina como estender a imagem de entrada
//	    const: KSvgEdgeMode... (ex. KSvgEdgeModeWrap)
//	    any other type: interface{}
func (e *TagSvgGlobal) EdgeMode(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgEdgeMode); ok {
		e.selfElement.Call("setAttribute", "edgeMode", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "edgeMode", value)
	return e
}

// Elevation
//
// English:
//
//	The elevation attribute specifies the direction angle for the light source from the XY plane towards the Z-axis, in
//	degrees. Note that the positive Z-axis points towards the viewer of the content.
//
//	 Input:
//	   elevation: specifies the direction angle for the light source
//
// Portuguese
//
//	O atributo de elevação especifica o ângulo de direção da fonte de luz do plano XY em direção ao eixo Z, em graus.
//	Observe que o eixo Z positivo aponta para o visualizador do conteúdo.
//
//	 Entrada:
//	   elevation: especifica o ângulo de direção para a fonte de luz
func (e *TagSvgGlobal) Elevation(elevation float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "elevation", elevation)
	return e
}

// End
//
// English:
//
//	The end attribute defines an end value for the animation that can constrain the active duration.
//
//	 Input:
//	   end: defines an end value for the animation
//	     offset-value: This value defines a clock-value that represents a point in time relative to the beginning of the
//	       SVG document (usually the load or DOMContentLoaded event). Negative values are valid.
//	       (e.g. time.Second*5 or "5s")
//	     syncbase-value: This value defines a syncbase and an optional offset from that syncbase. The element's
//	       animation start time is defined relative to the begin or active end of another animation.
//	       A valid syncbase-value consists of an ID reference to another animation element followed by a dot and either
//	       begin or end to identify whether to synchronize with the beginning or active end of the referenced animation
//	       element. An optional offset value as defined in <offset-value> can be appended.
//	       (e.g. "0s;third.end", "first.end" or "second.end")
//	     event-value: This value defines an event and an optional offset that determines the time at which the element's
//	       animation should begin. The animation start time is defined relative to the time that the specified event is
//	       fired.
//	       A valid event-value consists of an element ID followed by a dot and one of the supported events for that
//	       element. All valid events (not necessarily supported by all elements) are defined by the DOM and HTML
//	       specifications. Those are: 'focus', 'blur', 'focusin', 'focusout', 'activate', 'auxclick', 'click',
//	       'dblclick', 'mousedown', 'mouseenter', 'mouseleave', 'mousemove', 'mouseout', 'mouseover', 'mouseup',
//	       'wheel', 'beforeinput', 'input', 'keydown', 'keyup', 'compositionstart', 'compositionupdate',
//	       'compositionend', 'load', 'unload', 'abort', 'error', 'select', 'resize', 'scroll', 'beginEvent', 'endEvent',
//	       and 'repeatEvent'. An optional offset value as defined in <offset-value> can be appended.
//	       (e.g. "startButton.click")
//	     repeat-value: This value defines a qualified repeat event. The element animation start time is defined relative
//	       to the time that the repeat event is raised with the specified iteration value.
//	       A valid repeat value consists of an element ID followed by a dot and the function repeat() with an integer
//	       value specifying the number of repetitions as parameter. An optional offset value as defined in
//	       <offset-value> can be appended.
//	       (e.g. "0s;myLoop.end", "myLoop.begin", "myLoop.repeat(1)" or "myLoop.repeat(2)")
//	     accessKey-value: This value defines an access key that should trigger the animation. The element animation will
//	       begin when the user presses the specified key.
//	       A valid accessKey-value consists of the function accessKey() with the character to be input as parameter. An
//	       optional offset value as defined in <offset-value> can be appended.
//	       (e.g. "accessKey(s)")
//	     wallclock-sync-value: This value defines the animation start time as a real-world clock time.
//	       A valid wallclock-sync-value consists of the function wallclock() with a time value as parameter. The time
//	       syntax is based upon the syntax defined in ISO 8601.
//	       (e.g. time.Now() )
//	     indefinite: The begin of the animation will be determined by a beginElement() method call or a hyperlink
//	       targeted to the element.
//	       (e.g. "infinite")
//
// Portuguese
//
//	O atributo final define um valor final para a animação que pode restringir a duração ativa.
//
//	 Entrada:
//	   end: define um valor final para a animação
//	     offset-value: Esse valor define um valor de relógio que representa um ponto no tempo relativo ao início do
//	       documento SVG (geralmente o evento load ou DOMContentLoaded). Valores negativos são válidos.
//	       (e.g. time.Second*5 or "5s")
//	     syncbase-value: Esse valor define uma base de sincronização e um deslocamento opcional dessa base de
//	       sincronização. A hora de início da animação do elemento é definida em relação ao início ou fim ativo de outra
//	       animação.
//	       Um valor syncbase válido consiste em uma referência de ID para outro elemento de animação seguido por um
//	       ponto e um início ou fim para identificar se deve ser sincronizado com o início ou o final ativo do elemento
//	       de animação referenciado. Um valor de deslocamento opcional conforme definido em <offset-value> pode ser
//	       anexado.
//	       (e.g. "0s;third.end", "first.end" or "second.end")
//	     event-value: Esse valor define um evento e um deslocamento opcional que determina a hora em que a animação do
//	       elemento deve começar. A hora de início da animação é definida em relação à hora em que o evento especificado
//	       é acionado.
//	       Um valor de evento válido consiste em um ID de elemento seguido por um ponto e um dos eventos com suporte
//	       para esse elemento. Todos os eventos válidos (não necessariamente suportados por todos os elementos) são
//	       definidos pelas especificações DOM e HTML. Esses valores são: 'focus', 'blur', 'focusin', 'focusout',
//	       'activate', 'auxclick', 'click', 'dblclick', 'mousedown', 'mouseenter', 'mouseleave', 'mousemove',
//	       'mouseout', 'mouseover', 'mouseup', 'wheel', 'beforeinput', 'input', 'keydown', 'keyup', 'compositionstart',
//	       'compositionupdate', 'compositionend', 'load', 'unload', 'abort', 'error', 'select', 'resize', 'scroll',
//	       'beginEvent', 'endEvent', e 'repeatEvent'. Um valor de deslocamento opcional conforme definido em
//	       <offset-value> pode ser anexado.
//	       (e.g. "startButton.click")
//	     repeat-value: Esse valor define um evento de repetição qualificado. A hora de início da animação do elemento é
//	       definida em relação à hora em que o evento de repetição é gerado com o valor de iteração especificado.
//	       Um valor de repetição válido consiste em um ID de elemento seguido por um ponto e a função repeat() com um
//	       valor inteiro especificando o número de repetições como parâmetro. Um valor de deslocamento opcional conforme
//	       definido em <offset-value> pode ser anexado.
//	       (e.g. "0s;myLoop.end", "myLoop.begin", "myLoop.repeat(1)" or "myLoop.repeat(2)")
//	     accessKey-value: Este valor define uma chave de acesso que deve acionar a animação. A animação do elemento
//	       começará quando o usuário pressionar a tecla especificada.
//	       Um valor válido de accessKey consiste na função accessKey() com o caractere a ser inserido como parâmetro.
//	       Um valor de deslocamento opcional conforme definido em <valor de deslocamento> pode ser anexado.
//	       (e.g. "accessKey(s)")
//	     wallclock-sync-value: Esse valor define a hora de início da animação como uma hora do relógio do mundo real.
//	       Um valor wallclock-sync válido consiste na função wallclock() com um valor de tempo como parâmetro. A sintaxe
//	       de tempo é baseada na sintaxe definida na ISO 8601.
//	       (e.g. time.Now() )
//	     indefinite: O início da animação será determinado por uma chamada de método beginElement() ou um hiperlink
//	       direcionado ao elemento.
//	       (e.g. "infinite")
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
//	The exponent attribute defines the exponent of the gamma function.
//
//	 Input:
//	   exponent: defines the exponent of the gamma function
//
// Portuguese
//
//	O atributo expoente define o expoente da função gama.
//
//	 Entrada:
//	   exponent: define o expoente da função gama
func (e *TagSvgGlobal) Exponent(exponent float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "exponent", exponent)
	return e
}

// Fill
//
// English:
//
// The fill attribute has two different meanings. For shapes and text it's a presentation attribute that defines the
// color (or any SVG paint servers like gradients or patterns) used to paint the element;
//
// for animation it defines the final state of the animation.
//
//	Input:
//	  value: the fill value
//	    nil: fill="none"
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// Português:
//
// O atributo fill tem dois significados diferentes. Para formas e texto, é um atributo de apresentação que define a
// cor (ou qualquer servidor de pintura SVG, como gradientes ou padrões) usado para pintar o elemento;
//
// para animação, define o estado final da animação.
//
//	Input:
//	  value: the fill value
//	    nil: fill="none"
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Fill(value interface{}) (ref *TagSvgGlobal) {
	if value == nil {
		e.selfElement.Call("setAttribute", "fill", "none")
		return e
	}

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
//	The fill-opacity attribute is a presentation attribute defining the opacity of the paint server (color, gradient,
//	pattern, etc) applied to a shape.
//
//	 Input:
//	   value: defining the opacity of the paint
//	     float32: 1.0 = "100%"
//	     any other type: interface{}
//
//	 Notes:
//	   *As a presentation attribute fill-opacity can be used as a CSS property.
//
// Portuguese
//
//	O atributo fill-opacity é um atributo de apresentação que define a opacidade do servidor de pintura (cor, gradiente,
//	padrão etc.) aplicado a uma forma.
//
//	 Entrada:
//	   value: definindo a opacidade da tinta
//	     float32: 1.0 = "100%"
//	     qualquer outro tipo: interface{}
//
//	 Notes:
//	   *As a presentation attribute fill-opacity can be used as a CSS property.
func (e *TagSvgGlobal) FillOpacity(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "fill-opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "fill-opacity", value)
	return e
}

// FillRule
//
// English:
//
//	The fill-rule attribute is a presentation attribute defining the algorithm to use to determine the inside part of
//	a shape.
//
//	 Input:
//	   value: defining the algorithm to use to determine the inside part of a shape.
//	     const: KSvgFillRule... (e.g. KSvgFillRuleEvenOdd)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, fill-rule can be used as a CSS property.
//
// Portuguese
//
//	O atributo fill-rule é um atributo de apresentação que define o algoritmo a ser usado para determinar a parte
//	interna de uma forma.
//
//	 Input:
//	   value: define o algoritmo a ser usado para determinar a parte interna de uma forma.
//	     const: KSvgFillRule... (eg. KSvgFillRuleEvenOdd)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, fill-rule pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) FillRule(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgFillRule); ok {
		e.selfElement.Call("setAttribute", "fill-rule", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "fill-rule", value)
	return e
}

// Filter
//
// English:
//
//	The filter attribute specifies the filter effects defined by the <filter> element that shall be applied to its
//	element.
//
//	 Input:
//	   filter: specifies the filter effects
//
//	 Notes:
//	   * As a presentation attribute, filter can be used as a CSS property. See css filter for further information.
//
// Portuguese
//
//	O atributo filter especifica os efeitos de filtro definidos pelo elemento <filter> que devem ser aplicados ao seu
//	elemento.
//
//	 Entrada:
//	   filter: especifica os efeitos do filtro
//
//	 Notas:
//	   * Como atributo de apresentação, o filtro pode ser usado como propriedade CSS. Veja filtro css para mais
//	     informações.
func (e *TagSvgGlobal) Filter(filter string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "filter", filter)
	return e
}

// FilterUnits
//
// English:
//
// The filterUnits attribute defines the coordinate system for the attributes x, y, width and height.
//
//	Input:
//	  value: defines the coordinate system
//	    const: KSvgFilterUnits... (e.g. KSvgFilterUnitsObjectBoundingBox)
//	    any other type: interface{}
//
// # Portuguese
//
// O atributo filterUnits define o sistema de coordenadas para os atributos x, y, largura e altura.
//
//	Entrada:
//	  value: define o sistema de coordenadas
//	    const: KSvgFilterUnits... (ex. KSvgFilterUnitsObjectBoundingBox)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) FilterUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgFilterUnits); ok {
		e.selfElement.Call("setAttribute", "filterUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "filterUnits", value)
	return e
}

// FloodColor
//
// English:
//
//	The flood-color attribute indicates what color to use to flood the current filter primitive subregion.
//
//	 Input:
//	   floodColor: indicates what color to use to flood the current filter primitive subregion
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, flood-color can be used as a CSS property.
//
// Portuguese
//
//	O atributo flood-color indica qual cor usar para inundar a sub-região primitiva do filtro atual.
//
//	 Entrada:
//	   floodColor: indica qual cor usar para inundar a sub-região primitiva do filtro atual
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, a cor de inundação pode ser usada como uma propriedade CSS.
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
//	The flood-opacity attribute indicates the opacity value to use across the current filter primitive subregion.
//
//	 Input:
//	   floodOpacity: indicates the opacity value
//
//	 Notes:
//	   * As a presentation attribute, flood-opacity can be used as a CSS property.
//
// Portuguese
//
//	O atributo flood-opacity indica o valor de opacidade a ser usado na sub-região primitiva de filtro atual.
//
//	 Entrada:
//	   floodOpacity: indica o valor da opacidade
//
//	 Notas:
//	   * Como atributo de apresentação, a opacidade de inundação pode ser usada como uma propriedade CSS.
func (e *TagSvgGlobal) FloodOpacity(floodOpacity float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "flood-opacity", floodOpacity)
	return e
}

//todo: normalizar fonte com html padrão

// FontFamily
//
// English:
//
// The font-family attribute indicates which font family will be used to render the text, specified as a prioritized
// list of font family names and/or generic family names.
//
//	Input:
//	  fontFamily: indicates which font family will be used
//	    string: e.g. "Verdana, sans-serif"
//	    factory: e.g. factoryFontFamily.NewArial()
//
//	Notes:
//	  * As a presentation attribute, font-family can be used as a CSS property. See the css font-family property for
//	    more information.
//
// # Portuguese
//
// O atributo font-family indica qual família de fontes será usada para renderizar o texto, especificada como uma lista
// priorizada de nomes de famílias de fontes e ou nomes de famílias genéricos.
//
//	Entrada:
//	  fontFamily: indica qual família de fontes será usada
//	    string: ex. "Verdana, sans-serif"
//	    factory: ex. factoryFontFamily.NewArial()
//
//	Notas:
//	  * Como atributo de apresentação, font-family pode ser usada como propriedade CSS. Consulte a propriedade CSS
//	    font-family para obter mais informações.
func (e *TagSvgGlobal) FontFamily(fontFamily string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-family", fontFamily)
	return e
}

// FontSize
//
// English:
//
// The font-size attribute refers to the size of the font from baseline to baseline when multiple lines of text are set
// solid in a multiline layout environment.
//
//	Input:
//	  fontSize: size of the font
//	    string: e.g. "10px","2em"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, font-size can be used as a CSS property. See the css font-size property for more
//	    information.
//
// # Portuguese
//
// O atributo font-size refere-se ao tamanho da fonte da linha de base a linha de base quando várias linhas de texto
// são definidas como sólidas em um ambiente de layout de várias linhas.
//
//	Entrada:
//	  fontSize: tamanho da fonte
//	    string: ex. "10px","2em"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, font-size pode ser usado como uma propriedade CSS. Consulte a propriedade CSS
//	    font-size para obter mais informações.
func (e *TagSvgGlobal) FontSize(fontSize interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-size", fontSize)
	return e
}

// FontSizeAdjust #presentation
//
// English:
//
//	The font-size-adjust attribute allows authors to specify an aspect value for an element that will preserve the
//	x-height of the first choice font in a substitute font.
//
//	 Notes:
//	   * As a presentation attribute, font-size-adjust can be used as a CSS property. See the css font-size-adjust
//	     property for more information.
//
// Portuguese
//
//	O atributo font-size-adjust permite que os autores especifiquem um valor de aspecto para um elemento que preservará
//	a altura x da fonte de primeira escolha em uma fonte substituta.
//
//	 Notes:
//	   * As a presentation attribute, font-size-adjust can be used as a CSS property. See the css font-size-adjust
//	     property for more information.
func (e *TagSvgGlobal) FontSizeAdjust(fontSizeAdjust float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "font-size-adjust", fontSizeAdjust)
	return e
}

// FontStretch #presentation
//
// English:
//
//	The font-stretch attribute indicates the desired amount of condensing or expansion in the glyphs used to render
//	the text.
//
//	 Input:
//	   fontStretch: indicates the desired amount of condensing or expansion
//	     KSvgFontStretch... (e.g. KSvgFontStretchUltraCondensed)
//	     percentage (e.g. "50%")
//
//	 Notes:
//	   * As a presentation attribute, font-stretch can be used as a CSS property. See the css font-stretch property for
//	     more information.
//
// Portuguese
//
//	O atributo font-stretch indica a quantidade desejada de condensação ou expansão nos glifos usados para renderizar
//	o texto.
//
//	 Entrada:
//	   fontStretch: indica a quantidade desejada de condensação ou expansão
//	     KSvgFontStretch... (e.g. KSvgFontStretchUltraCondensed)
//	     percentage (e.g. "50%")
//
//	 Notas:
//	   * Como atributo de apresentação, font-stretch pode ser usado como uma propriedade CSS. Consulte a propriedade
//	     CSS font-stretch para obter mais informações.
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
//	The font-style attribute specifies whether the text is to be rendered using a normal, italic, or oblique face.
//
//	 Input:
//	   value: specifies whether the text is to be rendered using a normal, italic, or oblique face
//	     const: KFontStyleRule... (e.g. KFontStyleRuleItalic)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, font-style can be used as a CSS property. See the css font-style property for
//	     more information.
//
// Portuguese
//
//	O atributo font-style especifica se o texto deve ser renderizado usando uma face normal, itálica ou oblíqua.
//
//	 Entrada:
//	   value: especifica se o texto deve ser renderizado usando uma face normal, itálica ou oblíqua
//	     const: KFontStyleRule... (ex. KFontStyleRuleItalic)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, font-style pode ser usado como propriedade CSS. Consulte a propriedade CSS
//	     font-style para obter mais informações.
func (e *TagSvgGlobal) FontStyle(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(FontStyleRule); ok {
		e.selfElement.Call("setAttribute", "font-style", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-style", value)
	return e
}

// FontVariant
//
// English:
//
//	The font-variant attribute indicates whether the text is to be rendered using variations of the font's glyphs.
//
//	 Input:
//	   value: indicates whether the text is to be rendered
//	     const: KFontVariantRule... (e.g. KFontVariantRuleSmallCaps)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, font-variant can be used as a CSS property. See the css font-variant property
//	     for more information.
//
// Portuguese
//
//	O atributo font-variant indica se o texto deve ser renderizado usando variações dos glifos da fonte.
//
//	 Entrada:
//	   value: indica onde o texto vai ser renderizado.
//	     const: KFontVariantRule... (ex. KFontVariantRuleSmallCaps)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, font-variant pode ser usado como uma propriedade CSS. Consulte a propriedade
//	     CSS font-variant para obter mais informações.
func (e *TagSvgGlobal) FontVariant(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(FontVariantRule); ok {
		e.selfElement.Call("setAttribute", "font-variant", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-variant", value)
	return e
}

// FontWeight
//
// English:
//
//	The font-weight attribute refers to the boldness or lightness of the glyphs used to render the text, relative to
//	other fonts in the same font family.
//
//	 Input:
//	   value: refers to the boldness or lightness of the glyphs used to render the text
//	     const: KFontWeightRule... (e.g. KFontWeightRuleBold)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, font-weight can be used as a CSS property. See the css font-weight property for
//	     more information.
//
// Portuguese
//
//	O atributo font-weight refere-se ao negrito ou leveza dos glifos usados para renderizar o texto, em relação a
//	outras fontes na mesma família de fontes.
//
//	 Entrada:
//	   value: refere-se ao negrito ou leveza dos glifos usados para renderizar o texto
//	     const: KFontWeightRule... (ex. KFontWeightRuleBold)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, o peso da fonte pode ser usado como uma propriedade CSS. Consulte a propriedade
//	     CSS font-weight para obter mais informações.
func (e *TagSvgGlobal) FontWeight(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(FontWeightRule); ok {
		e.selfElement.Call("setAttribute", "font-weight", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-weight", value)
	return e
}

// From
//
// English:
//
// The from attribute indicates the initial value of the attribute that will be modified during the animation.
//
//	Input:
//	  value: initial value of the attribute
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1),rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s, 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    any other type: interface{}
//
// When used with the to attribute, the animation will change the modified attribute from the from value to the to
// value. When used with the by attribute, the animation will change the attribute relatively from the from value by
// the value specified in by.
//
// Português:
//
// O atributo from indica o valor inicial do atributo que será modificado durante a animação.
//
//	Entrada:
//	  value: valor inicial do atributo
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1),rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s, 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    qualquer outro tipo: interface{}
//
// Quando usado com o atributo to, a animação mudará o atributo modificado do valor from para o valor to. Quando usado
// com o atributo by, a animação mudará o atributo relativamente do valor from pelo valor especificado em by.
func (e *TagSvgGlobal) From(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "values", TypeToString(value, ";", ";"))
	return e
}

// Fr
//
// English:
//
//	The fr attribute defines the radius of the focal point for the radial gradient.
//
//	 Input:
//	   fr: defines the radius of the focal point for the radial gradient
//	     float32: (e.g. 0.4 = 40%)
//	     string: "40%"
//
// Portuguese
//
//	O atributo fr define o raio do ponto focal para o gradiente radial.
//
//	 Entrada:
//	   fr: define o raio do ponto focal para o gradiente radial.
//	     float32: (ex. 0.4 = 40%)
//	     string: "40%"
func (e *TagSvgGlobal) Fr(fr interface{}) (ref *TagSvgGlobal) {
	if converted, ok := fr.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
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
//	The fx attribute defines the x-axis coordinate of the focal point for a radial gradient.
//	   value: the x-axis coordinate of the focal point for a radial gradient
//	     float32: 1.0 = "100%"
//	     any other type: interface{}
//
// Portuguese
//
//	O atributo fx define a coordenada do eixo x do ponto focal para um gradiente radial.
//	   value: coordenada do eixo x do ponto focal para um gradiente radial
//	     float32: 1.0 = "100%"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Fx(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "fx", p)
		return e
	}

	e.selfElement.Call("setAttribute", "fx", value)
	return e
}

// Fy
//
// English:
//
//	The fy attribute defines the y-axis coordinate of the focal point for a radial gradient.
//
//	 Input:
//	   value: the y-axis coordinate of the focal point for a radial gradient
//	     float32: 1.0 = "100%"
//	     any other type: interface{}
//
// Portuguese
//
//	O atributo fy define a coordenada do eixo y do ponto focal para um gradiente radial.
//
//	 Entrada:
//	   value: coordenada do eixo y do ponto focal para um gradiente radial
//	     float32: 1.0 = "100%"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Fy(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "fy", p)
		return e
	}

	e.selfElement.Call("setAttribute", "fy", value)
	return e
}

// GradientTransform
//
// English:
//
//	The gradientTransform attribute contains the definition of an optional additional transformation from the gradient
//	coordinate system onto the target coordinate system (i.e., userSpaceOnUse or objectBoundingBox).
//
//	 Input:
//	   value: definition of an optional additional transformation from the gradient coordinate system
//	     Object: &html.TransformFunctions{}
//	     any other type: interface{}
//
//	This allows for things such as skewing the gradient. This additional transformation matrix is post-multiplied to
//	(i.e., inserted to the right of) any previously defined transformations, including the implicit transformation
//	necessary to convert from object bounding box units to user space.
//
// Portuguese
//
//	O atributo gradientTransform contém a definição de uma transformação adicional opcional do sistema de coordenadas
//	de gradiente para o sistema de coordenadas de destino (ou seja, userSpaceOnUse ou objectBoundingBox).
//
//	 Entrada:
//	   value: definição de uma transformação adicional opcional do sistema de coordenadas de gradiente
//	     Object: &html.TransformFunctions{}
//	     qualquer outro tipo: interface{}
//
//	Isso permite coisas como distorcer o gradiente. Essa matriz de transformação adicional é pós-multiplicada para
//	(ou seja, inserida à direita de) quaisquer transformações definidas anteriormente, incluindo a transformação
//	implícita necessária para converter de unidades de caixa delimitadora de objeto para espaço do usuário.
func (e *TagSvgGlobal) GradientTransform(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(*TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "gradientTransform", converted.String())
		return e
	}

	if converted, ok := value.(TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "gradientTransform", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "gradientTransform", value)
	return e
}

// GradientUnits
//
// English:
//
//	The gradientUnits attribute defines the coordinate system used for attributes specified on the gradient elements.
//
//	 Input:
//	   value: defines the coordinate system
//	     const: KSvgGradientUnits... (e.g. KSvgGradientUnitsUserSpaceOnUse)
//	     any other type: interface{}
//
// Portuguese
//
//	O atributo gradientUnits define o sistema de coordenadas usado para atributos especificados nos elementos
//	gradientes.
//
//	 Entrada:
//	   value: define o sistema de coordenadas
//	     const: KSvgGradientUnits... (ex. KSvgGradientUnitsUserSpaceOnUse)
//	     any other type: interface{}
func (e *TagSvgGlobal) GradientUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgGradientUnits); ok {
		e.selfElement.Call("setAttribute", "gradientUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "gradientUnits", value)
	return e
}

// Height
//
// English:
//
//	The height attribute defines the vertical length of an element in the user coordinate system.
//	     float32: 1.0 = "100%"
//	     any other type: interface{}
//
// Português:
//
//	O atributo height define o comprimento vertical de um elemento no sistema de coordenadas do usuário.
//	     float32: 1.0 = "100%"
//	     qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Height(height interface{}) (ref *TagSvgGlobal) {
	if converted, ok := height.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
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
//	The href attribute defines a link to a resource as a reference URL. The exact meaning of that link depends on the
//	context of each element using it.
//
//	 Notes:
//	   * Specifications before SVG 2 defined an xlink:href attribute, which is now rendered obsolete by the href
//	     attribute.
//	     If you need to support earlier browser versions, the deprecated xlink:href attribute can be used as a fallback
//	     in addition to the href attribute, e.g. <use href="some-id" xlink:href="some-id x="5" y="5" />.
//
// Português:
//
//	O atributo href define um link para um recurso como um URL de referência. O significado exato desse link depende do
//	contexto de cada elemento que o utiliza.
//
//	 Notas:
//	   * As especificações anteriores ao SVG 2 definiam um atributo xlink:href, que agora se torna obsoleto pelo
//	     atributo href.
//	     Se você precisar oferecer suporte a versões anteriores do navegador, o atributo obsoleto xlink:href pode ser
//	     usado como um substituto além do atributo href, por exemplo,
//	     <use href="some-id" xlink:href="some-id x="5" y="5" />.
func (e *TagSvgGlobal) HRef(href interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "href", href)
	return e
}

// Id #core
//
// English:
//
//	The id attribute assigns a unique name to an element.
//
// Portuguese
//
//	O atributo id atribui um nome exclusivo a um elemento.
func (e *TagSvgGlobal) Id(id string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "id", id)
	return e
}

// ImageRendering #presentation
//
// English:
//
//	The image-rendering attribute provides a hint to the browser about how to make speed vs. quality tradeoffs as it
//	performs image processing.
//
// The resampling is always done in a truecolor (e.g., 24-bit) color space even if the original data and/or the target
// device is indexed color.
//
//	Notes:
//	  * As a presentation attribute, image-rendering can be used as a CSS property. See the css image-rendering
//	    property for more information.
//
// Portuguese
//
//	O atributo de renderização de imagem fornece uma dica ao navegador sobre como fazer compensações de velocidade
//	versus qualidade enquanto executa o processamento de imagem.
//
// A reamostragem é sempre feita em um espaço de cores truecolor (por exemplo, 24 bits), mesmo que os dados originais e
// ou o dispositivo de destino sejam cores indexadas.
//
//	Notas:
//	  * Como um atributo de apresentação, a renderização de imagem pode ser usada como uma propriedade CSS. Consulte
//	    a propriedade de renderização de imagem css para obter mais informações.
func (e *TagSvgGlobal) ImageRendering(imageRendering string) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "image-rendering", imageRendering)
	return e
}

// In
//
// English:
//
//	The in attribute identifies input for the given filter primitive.
//
//	 Input:
//	   in: identifies input for the given filter primitive.
//	     KSvgIn... (e.g. KSvgInSourceAlpha)
//	     any other type: interface{}
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
//	O atributo in identifica à entrada para a primitiva de filtro fornecida.
//
//	 Entrada:
//	   in: identifica à entrada para a primitiva de filtro fornecida.
//	     KSvgIn... (e.g. KSvgInSourceAlpha)
//	     qualquer outro tipo: interface{}
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
//	The in2 attribute identifies the second input for the given filter primitive. It works exactly like the in
//	attribute.
//
//	 Input:
//	   in2: identifies the second input for the given filter primitive.
//	     KSvgIn2... (e.g. KSvgIn2SourceAlpha)
//	     string: url(#myClip)
//	     any other type: interface{}
//
// Portuguese
//
//	O atributo in2 identifica a segunda entrada para a primitiva de filtro fornecida. Funciona exatamente como o
//	atributo in.
//
//	 Entrada:
//	   in2: identifica a segunda entrada para a primitiva de filtro fornecida.
//	     KSvgIn2... (ex. KSvgIn2SourceAlpha)
//	     string: url(#myClip)
//	     qualquer outro tipo: interface{}
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
//	The intercept attribute defines the intercept of the linear function of color component transfers when the type
//	attribute is set to linear.
//
// Portuguese
//
//	O atributo de interceptação define a interceptação da função linear de transferências de componentes de cor quando
//	o atributo de tipo é definido como linear.
func (e *TagSvgGlobal) Intercept(intercept float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "intercept", intercept)
	return e
}

// K1
//
// English:
//
//	The k1 attribute defines one of the values to be used within the arithmetic operation of the <feComposite>
//	filter primitive.
//
// The pixel composition is computed using the following formula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//	O atributo k1 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K1(k1 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k1", k1)
	return e
}

// K2
//
// English:
//
//	The k2 attribute defines one of the values to be used within the arithmetic operation of the <feComposite> filter
//	primitive.
//
// The pixel composition is computed using the following formula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//	O atributo k2 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K2(k2 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k2", k2)
	return e
}

// K3
//
// English:
//
//	The k3 attribute defines one of the values to be used within the arithmetic operation of the <feComposite>
//	filter primitive.
//
// The pixel composition is computed using the following formula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//	O atributo k3 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K3(k3 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k3", k3)
	return e
}

// K4
//
// English:
//
//	The k4 attribute defines one of the values to be used within the arithmetic operation of the <feComposite>
//	filter primitive.
//
// The pixel composition is computed using the following formula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//
// Portuguese
//
//	O atributo k4 define um dos valores a serem usados na operação aritmética da primitiva de filtro <feComposite>.
//
// A composição de pixels é calculada usando a seguinte fórmula:
//
//	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
func (e *TagSvgGlobal) K4(k4 float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "k4", k4)
	return e
}

// KernelMatrix
//
// English:
//
// The kernelMatrix attribute defines the list of numbers that make up the kernel matrix for the <feConvolveMatrix>
// element.
//
//	Input:
//	  kernelMatrix: list of numbers
//	    []float64: []float64{1, 1, 0, 0, 0, 0, 0, 0, -1} = "1 1 0 0 0 0 0 0 -1"
//	    any other type: interface{}
//
// The list of <number>s that make up the kernel matrix for the convolution. The number of entries in the list
// must equal <orderX> times <orderY>.
// If the result of orderX * orderY is not equal to the number of entries in the value list, the filter primitive
// acts as a pass through filter.
//
// Values are separated by space characters and/or a comma. The number of entries in the list must equal to <orderX>
// by <orderY> as defined in the order attribute.
//
// Português:
//
// O atributo kernelMatrix define a lista de números que compõem a matriz do kernel para o elemento <feConvolveMatrix>.
//
//	Entrada:
//	  kernelMatrix: lista de números
//	    []float64: []float64{1, 1, 0, 0, 0, 0, 0, 0, -1} = "1 1 0 0 0 0 0 0 -1"
//	    qualquer outro tipo: interface{}
//
// A lista de números que compõem a matriz do kernel para a convolução. O número de entradas na lista deve ser
// igual a <orderX> * <orderY>.
// Se o resultado da ordem do pedido não for igual ao número de entradas na lista de valores, a primitiva de
// filtro atua como um filtro de passagem.
//
// Os valores são separados por caracteres de espaço e ou por vírgula. O número de entradas na lista deve ser igual a
// <orderX> por <orderY> conforme definido no atributo order.
func (e *TagSvgGlobal) KernelMatrix(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "kernelMatrix", valueStr[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "kernelMatrix", value)
	return e
}

// KeyPoints
//
// English:
//
// The keyPoints attribute indicates the simple duration of an animation.
//
//	Input:
//	  keyPoints: indicates the simple duration of an animation.
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1);rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%;10%"
//	    []float64: []float64{0.0, 10.0} = "0;10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s;1s"
//	    any other type: interface{}
//
// This value defines a semicolon-separated list of floating point values between 0 and 1 and indicates how far along
// the motion path the object shall move at the moment in time specified by corresponding keyTimes value. The distance
// is calculated along the path specified by the path attribute. Each progress value in the list corresponds to a value
// in the keyTimes attribute list.
// If a list of key points is specified, there must be exactly as many values in the keyPoints list as in the keyTimes
// list.
// If there's a semicolon at the end of the value, optionally followed by white space, both the semicolon and the
// trailing white space are ignored.
// If there are any errors in the value specification (i.e. bad values, too many or too few values), then that's an
// error.
//
// Português:
//
// O atributo keyPoints indica a duração simples de uma animação.
//
//	Entrada:
//	  keyPoints: indica a duração simples de uma animação.
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1);rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%;10%"
//	    []float64: []float64{0.0, 10.0} = "0;10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s;1s"
//	    any other type: interface{}
//
// Este valor define uma lista separada por ponto e vírgula de valores de ponto flutuante entre 0 e 1 e indica a
// distância ao longo do caminho de movimento o objeto deve se mover no momento especificado pelo valor keyTimes
// correspondente. A distância é calculada ao longo do caminho especificado pelo atributo path. Cada valor de progresso
// na lista corresponde a um valor na lista de atributos keyTimes.
// Se uma lista de pontos-chave for especificada, deve haver exatamente tantos valores na lista keyPoints quanto na
// lista keyTimes.
// Se houver um ponto e vírgula no final do valor, opcionalmente seguido por espaço em branco, o ponto e vírgula e o
// espaço em branco à direita serão ignorados.
// Se houver algum erro na especificação do valor (ou seja, valores incorretos, muitos ou poucos valores), isso é um
// erro.
func (e *TagSvgGlobal) KeyPoints(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "keyPoints", TypeToString(value, ";", ";"))
	return e
}

// KeySplines
//
// English:
//
// The keySplines attribute defines a set of Bézier curve control points associated with the keyTimes list, defining a
// cubic Bézier function that controls interval pacing.
//
//	Input:
//	  value: set of Bézier curve control points associated with the keyTimes list
//	    [][]float64: [][]float64{{0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}}
//	    any other type: interface{}
//
// This attribute is ignored unless the calcMode attribute is set to spline.
//
// If there are any errors in the keySplines specification (bad values, too many or too few values), the animation will
// not occur.
//
// Português:
//
// O atributo keySplines define um conjunto de pontos de controle da curva Bézier associados à lista keyTimes,
// definindo uma função Bézier cúbica que controla o ritmo do intervalo.
//
//	Entrada:
//	  value: conjunto de pontos de controle da curva Bézier associados à lista keyTimes
//	    [][]float64: [][]float64{{0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}}
//	    qualquer outro tipo: interface{}
//
// Esse atributo é ignorado, a menos que o atributo calcMode seja definido como spline.
//
// Se houver algum erro na especificação de keySplines (valores incorretos, muitos ou poucos valores), a animação não
// ocorrerá.
func (e *TagSvgGlobal) KeySplines(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "keySplines", TypeToString(value, " ", ";"))
	return e
}

// KeyTimes
//
// English:
//
// The keyTimes attribute represents a list of time values used to control the pacing of the animation.
//
//	Input:
//	  value: list of time values used to control
//	    []float64{0.0, 0.5, 1.0}: values="0; 0.5; 1"
//
// Each time in the list corresponds to a value in the values attribute list, and defines when the value is used in the
// animation.
//
// Each time value in the keyTimes list is specified as a floating point value between 0 and 1 (inclusive), representing
// a proportional offset into the duration of the animation element.
//
// Português:
//
// O atributo keyTimes representa uma lista de valores de tempo usados para controlar o ritmo da animação.
//
//	Entrada:
//	  value: lista de valores de tempo usados para controle
//	    []float64{0.0, 0.5, 1.0}: values="0; 0.5; 1"
//
// Cada vez na lista corresponde a um valor na lista de atributos de valores e define quando o valor é usado na
// animação.
//
// Cada valor de tempo na lista keyTimes é especificado como um valor de ponto flutuante entre 0 e 1 (inclusive),
// representando um deslocamento proporcional à duração do elemento de animação.
func (e *TagSvgGlobal) KeyTimes(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "keyTimes", TypeToString(value, ";", ""))
	return e
}

// Lang
//
// English:
//
// The lang attribute specifies the primary language used in contents and attributes containing text content of
// particular elements.
//
//	Input:
//	  value: specifies the primary language used in contents
//	    const: KLanguage... (e.g. KLanguageEnglishGreatBritain)
//	    any other type: interface{}
//
// There is also an xml:lang attribute (with namespace). If both of them are defined, the one with namespace is used and
// the one without is ignored.
//
// In SVG 1.1 there was a lang attribute defined with a different meaning and only applying to <glyph> elements.
// That attribute specified a list of languages according to RFC 5646: Tags for Identifying Languages
// (also known as BCP 47). The glyph was meant to be used if the xml:lang attribute exactly matched one of the languages
// given in the value of this parameter, or if the xml:lang attribute exactly equaled a prefix of one of the languages
// given in the value of this parameter such that the first tag character following the prefix was "-".
//
// Português:
//
// O atributo lang especifica o idioma principal usado em conteúdos e atributos que contêm conteúdo de texto de
// elementos específicos.
//
//	Entrada:
//	  value: especifica o idioma principal usado no conteúdo
//	    const: KLanguage... (e.g. KLanguagePortugueseBrazil)
//	    qualquer outro tipo: interface{}
//
// Há também um atributo xml:lang (com namespace). Se ambos estiverem definidos, aquele com namespace será usado e o
// sem namespace será ignorado.
//
// No SVG 1.1 havia um atributo lang definido com um significado diferente e aplicando-se apenas aos elementos <glyph>.
// Esse atributo especificou uma lista de idiomas de acordo com a RFC 5646: Tags for Identification Languages
// (também conhecido como BCP 47). O glifo deveria ser usado se o atributo xml:lang correspondesse exatamente a um dos
// idiomas fornecidos no valor desse parâmetro, ou se o atributo xml:lang fosse exatamente igual a um prefixo de um dos
// idiomas fornecidos no valor desse parâmetro de modo que o primeiro caractere de tag após o prefixo fosse "-".
func (e *TagSvgGlobal) Lang(value interface{}) (ref *TagSvgGlobal) {

	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "lang", value)
	return e
}

// LengthAdjust
//
// English:
//
// The lengthAdjust attribute controls how the text is stretched into the length defined by the textLength attribute.
//
//	Input:
//	  value: controls how the text is stretched
//	    KSvgLengthAdjust... (e.g. KSvgLengthAdjustSpacing)
//
// Português:
//
// O atributo lengthAdjust controla como o texto é esticado no comprimento definido pelo atributo textLength.
//
//	Input:
//	  value: controla como o texto é esticado
//	    KSvgLengthAdjust... (e.g. KSvgLengthAdjustSpacing)
func (e *TagSvgGlobal) LengthAdjust(value interface{}) (ref *TagSvgGlobal) {

	if converted, ok := value.(SvgLengthAdjust); ok {
		e.selfElement.Call("setAttribute", "lengthAdjust", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "lengthAdjust", value)
	return e
}

// LetterSpacing #presentation
//
// English:
//
// The letter-spacing attribute controls spacing between text characters, in addition to any spacing from the kerning
// attribute.
//
//	Input:
//	  value: controls spacing between text characters
//
// If the attribute value is a unitless number (like 128), the browser processes it as a <length> in the current user
// coordinate system.
//
// If the attribute value has a unit identifier, such as .25em or 1%, then the browser converts the <length> into its
// corresponding value in the current user coordinate system.
//
// Notes:
//   - As a presentation attribute, letter-spacing can be used as a CSS property.
//     See the css letter-spacing property for more information.
//
// Português:
//
// O atributo letter-spacing controla o espaçamento entre caracteres de texto, além de qualquer espaçamento do atributo
// kerning.
//
//	Input:
//	  value: controla o espaçamento entre caracteres de texto
//
// Se o valor do atributo for um número sem unidade (como 128), o navegador o processará como um <comprimento> no
// sistema de coordenadas do usuário atual.
//
// Se o valor do atributo tiver um identificador de unidade, como .25em ou 1%, o navegador converterá o <comprimento>
// em seu valor correspondente no sistema de coordenadas do usuário atual.
//
// Notas:
//   - Como atributo de apresentação, o espaçamento entre letras pode ser usado como uma propriedade CSS.
//     Consulte a propriedade de espaçamento entre letras do CSS para obter mais informações.
func (e *TagSvgGlobal) LetterSpacing(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "letter-spacing", value)
	return e
}

// LightingColor #presentation
//
// English:
//
// The lighting-color attribute defines the color of the light source for lighting filter primitives.
//
//	Input:
//	  value: defines the color of the light source
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// Português:
//
// O atributo lighting-color define a cor da fonte de luz para as primitivas do filtro de iluminação.
//
//	Input:
//	  value: define a cor da fonte de luz
//	    string: ex. "black"
//	    factory: ex. factoryColor.NewYellow()
//	    RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) LightingColor(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "lighting-color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "lighting-color", value)
	return e
}

// LimitingConeAngle
//
// English:
//
// The limitingConeAngle attribute represents the angle in degrees between the spot light axis (i.e. the axis between
// the light source and the point to which it is pointing at) and the spot light cone. So it defines a limiting cone
// which restricts the region where the light is projected. No light is projected outside the cone.
//
//	Input:
//	  value: represents the angle in degrees between the spot light axis
//
// Português:
//
// O atributo limitConeAngle representa o ângulo em graus entre o eixo de luz spot (ou seja, o eixo entre a fonte de
// luz e o ponto para o qual está apontando) e o cone de luz spot. Assim, define um cone limitador que restringe a
// região onde a luz é projetada. Nenhuma luz é projetada fora do cone.
//
//	Input:
//	  value: representa o ângulo em graus entre o eixo da luz spot
func (e *TagSvgGlobal) LimitingConeAngle(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "limitingConeAngle", value)
	return e
}

// MarkerEnd #presentation
//
// English:
//
// The marker-end attribute defines the arrowhead or polymarker that will be drawn at the final vertex of the given
// shape.
//
//	Input:
//	  value: the arrowhead or polymarker that will be drawn
//	    string: (e.g. "url(#triangle)")
//
// For all shape elements, except <polyline> and <path>, the last vertex is the same as the first vertex. In this case,
// if the value of marker-start and marker-end are both not none, then two markers will be rendered on that final
// vertex.
// For <path> elements, for each closed subpath, the last vertex is the same as the first vertex. marker-end is only
// rendered on the final vertex of the path data.
//
// Notes:
//   - As a presentation attribute, marker-end can be used as a CSS property.
//
// Português:
//
// O atributo marker-end define a ponta de seta ou polimarcador que será desenhado no vértice final da forma dada.
//
//	Entrada:
//	  value: a ponta de seta ou polimarcador que será desenhado
//	    string: (e.g. "url(#triangle)")
//
// Para todos os elementos de forma, exceto <polyline> e <path>, o último vértice é o mesmo que o primeiro vértice.
// Nesse caso, se o valor de marker-start e marker-end não for nenhum, então dois marcadores serão renderizados nesse
// vértice final.
// Para elementos <path>, para cada subcaminho fechado, o último vértice é igual ao primeiro vértice.
// O final do marcador é renderizado apenas no vértice final dos dados do caminho.
//
// Notas:
//   - Como atributo de apresentação, o marker-end pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) MarkerEnd(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "marker-end", value)
	return e
}

// MarkerMid #presentation
//
// English:
//
// The marker-mid attribute defines the arrowhead or polymarker that will be drawn at all interior vertices of the
// given shape.
//
//	Input:
//	  value: defines the arrowhead or polymarker that will be drawn
//	    string: e.g. "url(#circle)"
//
// The marker is rendered on every vertex other than the first and last vertices of the path data.
//
// Notes:
//   - As a presentation attribute, marker-mid can be used as a CSS property.
//
// Português:
//
// O atributo marker-mid define a ponta de seta ou polimarcador que será desenhado em todos os vértices internos da
// forma dada.
//
//	Input:
//	  value: define a ponta de seta ou polimarcador que será desenhado
//	    string: ex. "url(#circle)"
//
// O marcador é renderizado em todos os vértices, exceto no primeiro e no último vértice dos dados do caminho.
//
// Notas:
//   - Como atributo de apresentação, o marker-mid pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) MarkerMid(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "marker-mid", value)
	return e
}

// MarkerStart #presentation
//
// English:
//
// The marker-start attribute defines the arrowhead or polymarker that will be drawn at the first vertex of the given
// shape.
//
//	Input:
//	  value: defines the arrowhead or polymarker that will be drawn
//	    string: e.g. "url(#triangle)"
//
// For all shape elements, except <polyline> and <path>, the last vertex is the same as the first vertex. In this case,
// if the value of marker-start and marker-end are both not none, then two markers will be rendered on that final
// vertex.
// For <path> elements, for each closed subpath, the last vertex is the same as the first vertex. marker-start is only
// rendered on the first vertex of the path data.
//
// Notes:
//   - As a presentation attribute, marker-start can be used as a CSS property.
//
// Português:
//
// O atributo marker-start define a ponta de seta ou polimarcador que será desenhado no primeiro vértice da forma dada.
//
//	Entrada:
//	  value: define a ponta de seta ou polimarcador que será desenhado
//	    string: e.g. "url(#triangle)"
//
// Para todos os elementos de forma, exceto <polyline> e <path>, o último vértice é o mesmo que o primeiro vértice.
// Nesse caso, se o valor de marker-start e marker-end não for nenhum, então dois marcadores serão renderizados nesse
// vértice final.
// Para elementos <path>, para cada subcaminho fechado, o último vértice é igual ao primeiro vértice. O início do
// marcador é renderizado apenas no primeiro vértice dos dados do caminho.
//
// Notas:
//   - Como atributo de apresentação, o início do marcador pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) MarkerStart(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "marker-start", value)
	return e
}

// MarkerHeight
//
// English:
//
// The markerHeight attribute represents the height of the viewport into which the <marker> is to be fitted when it is
// rendered according to the viewBox and preserveAspectRatio attributes.
//
//	Input:
//	  value: represents the height of the viewport
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo markerHeight representa a altura da viewport na qual o <marker> deve ser ajustado quando for renderizado
// de acordo com os atributos viewBox e preserveAspectRatio.
//
//	Entrada:
//	  value: representa a altura da janela de visualização
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) MarkerHeight(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "markerHeight", p)
		return e
	}

	e.selfElement.Call("setAttribute", "markerHeight", value)
	return e
}

// MarkerUnits
//
// English:
//
// The markerUnits attribute defines the coordinate system for the markerWidth and markerHeight attributes and the
// contents of the <marker>.
//
//	Input:
//	  value: defines the coordinate system
//	    const KSvgMarkerUnits... (e.g. KSvgMarkerUnitsUserSpaceOnUse)
//	    any other type: interface{}
//
// Português:
//
// O atributo markerUnits define o sistema de coordenadas para os atributos markerWidth e markerHeight e o conteúdo
// do <marker>.
//
//	Entrada:
//	  value: define o sistema de coordenadas
//	    const KSvgMarkerUnits... (ex. KSvgMarkerUnitsUserSpaceOnUse)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) MarkerUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgMarkerUnits); ok {
		e.selfElement.Call("setAttribute", "markerUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "markerUnits", value)
	return e
}

// MarkerWidth
//
// English:
//
// The markerWidth attribute represents the width of the viewport into which the <marker> is to be fitted when it is
// rendered according to the viewBox and preserveAspectRatio attributes.
//
//	Input:
//	  value: represents the width of the viewport
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo markerWidth representa a largura da viewport na qual o <marker> deve ser ajustado quando for renderizado
// de acordo com os atributos viewBox e preserveAspectRatio.
//
//	Input:
//	  value: representa a largura da janela de visualização
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) MarkerWidth(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "markerWidth", p)
		return e
	}

	e.selfElement.Call("setAttribute", "markerWidth", value)
	return e
}

// Mask #presentation
//
// English:
//
// The mask attribute is a presentation attribute mainly used to bind a given <mask> element with the element the
// attribute belongs to.
//
//	Input:
//	  value: attribute mainly used to bind a given <mask> element
//	    string: "url(#myMask)"
//
// Notes:
//   - As a presentation attribute mask can be used as a CSS property.
//
// Português:
//
// O atributo mask é um atributo de apresentação usado principalmente para vincular um determinado elemento <mask> ao
// elemento ao qual o atributo pertence.
//
//	Entrada:
//	  value: atributo usado principalmente para vincular um determinado elemento <mask>
//	    string: "url(#myMask)"
//
// Notas:
//   - Como uma máscara de atributo de apresentação pode ser usada como uma propriedade CSS.
func (e *TagSvgGlobal) Mask(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "mask", value)
	return e
}

// MaskContentUnits
//
// English:
//
// The maskContentUnits attribute indicates which coordinate system to use for the contents of the <mask> element.
//
//	Input:
//	  value: specifies the coordinate system
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    any other type: interface{}
//
// Português:
//
// O atributo maskContentUnits indica qual sistema de coordenadas usar para o conteúdo do elemento <mask>.
//
//	Entrada:
//	  value: especifica o sistema de coordenadas
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) MaskContentUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgUnits); ok {
		e.selfElement.Call("setAttribute", "maskContentUnits", converted.String())
		return e
	}
	e.selfElement.Call("setAttribute", "maskContentUnits", value)
	return e
}

// MaskUnits
//
// English:
//
// The maskUnits attribute indicates which coordinate system to use for the geometry properties of the <mask> element.
//
//	Input:
//	  value: specifies the coordinate system
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    any other type: interface{}
//
// Português:
//
// O atributo maskUnits indica qual sistema de coordenadas usar para as propriedades geométricas do elemento <mask>.
//
//	Entrada:
//	  value: especifica o sistema de coordenadas
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) MaskUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgUnits); ok {
		e.selfElement.Call("setAttribute", "maskUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "maskUnits", value)
	return e
}

// Max
//
// English:
//
// The max attribute specifies the maximum value of the active animation duration.
//
//	Input:
//	  value: specifies the maximum value
//	    float32: 1.0 = "100%"
//	    time.Duration: 5*time.Second = "5s"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// Português:
//
// O atributo max especifica o valor máximo da duração da animação ativa.
//
//	Entrada:
//	  value: especifica o valor máximo
//	    float32: 1.0 = "100%"
//	    time.Duration: 5*time.Second = "5s"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Max(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "max", RGBAToJs(converted))
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "max", p)
		return e
	}

	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "max", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "max", value)
	return e
}

// Media
//
// English:
//
// The media attribute specifies a media query that must be matched for a style sheet to apply.
//
//	Input:
//	  value: specifies a media query that must be matched for a style sheet to apply
//	    string: e.g. "all and (min-width: 600px)"
//
// Português:
//
// O atributo de mídia especifica uma consulta de mídia que deve ser correspondida para que uma folha de estilo seja
// aplicada.
//
//	Entrada:
//	  value: especifica uma consulta de mídia que deve ser correspondida para que uma folha de estilo seja aplicada
//	    string: e.g. "all and (min-width: 600px)"
func (e *TagSvgGlobal) Media(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "media", value)
	return e
}

// Min
//
// English:
//
// The min attribute specifies the minimum value of the active animation duration.
//
//	Input:
//	  value: specifies the minimum value
//	    float32: 1.0 = "100%"
//	    time.Duration: 5*time.Second = "5s"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// Português:
//
// O atributo min especifica o valor mínimo da duração da animação ativa.
//
//	Input:
//	  value: especifica o valor mínimo
//	    float32: 1.0 = "100%"
//	    time.Duration: 5*time.Second = "5s"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Min(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "min", RGBAToJs(converted))
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "min", p)
		return e
	}

	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "min", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "min", value)
	return e
}

// Mode
//
// English:
//
// The mode attribute defines the blending mode on the <feBlend> filter primitive.
//
//	Input:
//	  value: defines the blending mode
//	    const KSvgMode... (e.g. KSvgModeNormal)
//	    any other type: interface{}
//
// Português:
//
// O atributo mode define o modo de mesclagem na primitiva de filtro <feBlend>.
//
//	Entrada:
//	  value: define o modo de mesclagem
//	    const KSvgMode... (ex. KSvgModeNormal)
//	    qualquer outro tipo: interface{}
//
// todo: exemplos: https://developer.mozilla.org/en-US/docs/Web/CSS/blend-mode
func (e *TagSvgGlobal) Mode(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgMode); ok {
		e.selfElement.Call("setAttribute", "mode", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "mode", value)
	return e
}

// NumOctaves
//
// English:
//
// The numOctaves attribute defines the number of octaves for the noise function of the <feTurbulence> primitive.
//
//	Input:
//	  value: defines the number of octaves for the noise function
//
// An octave is a noise function defined by its frequency and amplitude. A turbulence is built by accumulating several
// octaves with increasing frequencies and decreasing amplitudes. The higher the number of octaves, the more natural
// the noise looks. Though more octaves also require more calculations, resulting in a negative impact on performance.
//
// Português:
//
// O atributo numOctaves define o número de oitavas para a função de ruído da primitiva <feTurbulence>.
//
//	Input:
//	  value: define o número de oitavas para a função de ruído
//
// Uma oitava é uma função de ruído definida por sua frequência e amplitude. Uma turbulência é construída acumulando
// várias oitavas com frequências crescentes e amplitudes decrescentes. Quanto maior o número de oitavas, mais natural
// o ruído parece. Embora mais oitavas também exijam mais cálculos, resultando em um impacto negativo no desempenho.
func (e *TagSvgGlobal) NumOctaves(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "numOctaves", value)
	return e
}

// Opacity #presentation
//
// English:
//
// The opacity attribute specifies the transparency of an object or of a group of objects, that is, the degree to which
// the background behind the element is overlaid.
//
//	Input:
//	  value: specifies the transparency of an object
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, opacity can be used as a CSS property. See the css opacity property for more
//	    information.
//
// Português:
//
// O atributo opacity especifica a transparência de um objeto ou de um grupo de objetos, ou seja, o grau em que o fundo
// atrás do elemento é sobreposto.
//
//	Entrada:
//	  value: especifica a transparência de um objeto
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
//	Notes:
//	  * Como atributo de apresentação, a opacidade pode ser usada como uma propriedade CSS. Consulte a propriedade de
//	    opacidade do CSS para obter mais informações.
func (e *TagSvgGlobal) Opacity(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "opacity", value)
	return e
}

// Operator
//
// English:
//
// The operator attribute has two meanings based on the context it's used in. Either it defines the compositing or
// morphing operation to be performed.
//
//	Input:
//	  value: defines the compositing or morphing
//	    const: KSvgOperatorFeComposite... (e.g. KSvgOperatorFeCompositeOver)
//	    const: KSvgOperatorFeMorphology... (e.g. KKSvgOperatorFeCompositeErode)
//
// Português:
//
// O atributo operador tem dois significados com base no contexto em que é usado. Ele define a operação de composição
// ou transformação a ser executada.
//
//	Entrada:
//	  value: define a composição ou morphing
//	    const: KSvgOperatorFeComposite... (e.g. KSvgOperatorFeCompositeOver)
//	    const: KSvgOperatorFeMorphology... (e.g. KKSvgOperatorFeCompositeErode)
//
// fixme: separar quando colocar em <feComposite> e <feMorphology>
func (e *TagSvgGlobal) Operator(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgOperatorFeComposite); ok {
		e.selfElement.Call("setAttribute", "operator", converted.String())
		return e
	}

	if converted, ok := value.(SvgOperatorFeMorphology); ok {
		e.selfElement.Call("setAttribute", "operator", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "operator", value)
	return e
}

// Order
//
// English:
//
// The order attribute indicates the size of the matrix to be used by a <feConvolveMatrix> element.
//
//	Input:
//	  value: indicates the size of the matrix.
//	    []float64: []float64{1.0, 1.0, 1.0} = "1 1 1"
//	    any other type: interface{}
//
// Português:
//
// O atributo order indica o tamanho da matriz a ser usada por um elemento <feConvolveMatrix>.
//
//	Entrada:
//	  value: indica o tamanho da matriz.
//	    []float64: []float64{1.0, 1.0, 1.0} = "1 1 1"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Order(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		var order = ""
		for _, v := range converted {
			order += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}

		length := len(order) - 1
		e.selfElement.Call("setAttribute", "order", order[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "order", value)
	return e
}

// Orient
//
// English:
//
// The orient attribute indicates how a marker is rotated when it is placed at its position on the shape.
//
//	Input:
//	  value: indicates how a marker is rotated
//	    const: KSvgOrient... (e.g. KSvgOrientAuto)
//	    Degrees: Degrees(-65) = "-65deg"
//	    any other type: interface{}
//
// Português:
//
// O atributo orient indica como um marcador é girado quando é colocado em sua posição na forma.
//
//	Entrada:
//	  value: indica como um marcador é girado
//	    const: KSvgOrient... (ex. KSvgOrientAuto)
//	    Degrees: Degrees(-65) = "-65deg"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Orient(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(Degrees); ok {
		e.selfElement.Call("setAttribute", "orient", converted.String())
		return e
	}

	if converted, ok := value.(SvgOrient); ok {
		e.selfElement.Call("setAttribute", "orient", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "orient", value)
	return e
}

// Origin
//
// English:
//
//	Input:
//	  value: specifies the origin of motion for an animation
//
// The origin attribute specifies the origin of motion for an animation. It has no effect in SVG.
//
// Português:
//
// O atributo origin especifica a origem do movimento de uma animação. Não tem efeito em SVG.
//
//	Entrada:
//	  value: especifica a origem do movimento de uma animação
func (e *TagSvgGlobal) Origin(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "origin", value)
	return e
}

// Overflow
//
// English:
//
// The overflow attribute sets what to do when an element's content is too big to fit in its block formatting context.
//
//	Input:
//	  value: sets what to do when an element's content is too big to fit in its block formatting context
//	    const: KOverflow... (e.g. KOverflowHidden)
//	    any other type: interface{}
//
// This attribute has the same parameter values and meaning as the css overflow property, however, the following
// additional points apply:
//
//   - If it has a value of visible, the attribute has no effect (i.e., a clipping rectangle is not created).
//
//   - If the overflow property has the value hidden or scroll, a clip of the exact size of the SVG viewport is applied.
//
//   - When scroll is specified on an <svg> element, a scrollbar or panner is normally shown for the SVG viewport
//     whether or not any of its content is clipped.
//
//   - Within SVG content, the value auto implies that all rendered content for child elements must be visible, either
//     through a scrolling mechanism, or by rendering with no clip.
//
//     Notes:
//
//   - Although the initial value for overflow is auto, it is overwritten in the User Agent style sheet for the <svg>
//     element when it is not the root element of a stand-alone document, the <pattern> element, and the <marker>
//     element to be hidden by default.
//
//   - As a presentation attribute, overflow can be used as a CSS property. See the CSS overflow property for more
//     information.
//
// Português:
//
// O atributo overflow define o que fazer quando o conteúdo de um elemento é muito grande para caber em seu contexto
// de formatação de bloco.
//
//	Entrada:
//	  value: define o que fazer quando o conteúdo de um elemento é muito grande para caber em seu contexto de
//	      formatação de bloco
//	    const: KOverflow... (e.g. KOverflowHidden)
//	    qualquer outro tipo: interface{}
//
// Este atributo tem os mesmos valores de parâmetro e significado que a propriedade CSS overflow, no entanto, os
// seguintes pontos adicionais se aplicam:
//
//   - Se tiver um valor de visible, o atributo não terá efeito (ou seja, um retângulo de recorte não será criado).
//
//   - Se a propriedade overflow tiver o valor oculto ou rolar, um clipe do tamanho exato da janela de visualização SVG
//     será aplicado.
//
//   - Quando a rolagem é especificada em um elemento <svg>, uma barra de rolagem ou panner normalmente é mostrado para
//     a janela de visualização SVG, independentemente de seu conteúdo estar ou não recortado.
//
//   - No conteúdo SVG, o valor auto implica que o conteúdo renderizado para elementos filho deve ser visível por
//     completo, seja por meio de um mecanismo de rolagem ou renderizando sem clipe.
//
//     Notas:
//
//   - Embora o valor inicial para estouro seja auto, ele é substituído na folha de estilo do User Agent para o
//     elemento <svg> quando não é o elemento raiz de um documento autônomo, o elemento <pattern> e o elemento
//     <marker> para ser ocultado por padrão.
//
//   - Como atributo de apresentação, overflow pode ser usado como propriedade CSS. Consulte a propriedade CSS
//     overflow para obter mais informações.
func (e *TagSvgGlobal) Overflow(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(Overflow); ok {
		e.selfElement.Call("setAttribute", "overflow", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "overflow", value)
	return e
}

// PaintOrder
//
// English:
//
// The paint-order attribute specifies the order that the fill, stroke, and markers of a given shape or text element
// are painted.
//
//	Input:
//	  value: specifies the order that the fill, stroke, and markers
//	    const: SvgPaintOrder... (e.g. KSvgPaintOrderStroke)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, paint-order can be used as a CSS property.
//
// Português:
//
// O atributo paint-order especifica a ordem em que o preenchimento, o traçado e os marcadores de uma determinada forma
// ou elemento de texto são pintados.
//
//	Entrada:
//	  value: especifica a ordem em que o preenchimento, o traçado e os marcadores
//	    const: SvgPaintOrder... (e.g. KSvgPaintOrderStroke)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, paint-order pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) PaintOrder(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgPaintOrder); ok {
		e.selfElement.Call("setAttribute", "paint-order", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "paint-order", value)
	return e
}

// Path
//
// English:
//
// The path attribute has two different meanings, either it defines a text path along which the characters of a text
// are rendered, or a motion path along which a referenced element is animated.
//
//	Input:
//	  value: defines a text path along which the characters of a text are rendered, or a motion path along which a
//	      referenced element is animated
//	    factory: factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()
//	    string: "M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z"
//	    any other type: interface{}
//
// Português:
//
// O atributo path tem dois significados diferentes: define um caminho de texto ao longo do qual os caracteres de um
// texto são renderizados ou um caminho de movimento ao longo do qual um elemento referenciado é animado.
//
//	Entrada:
//	  value: define um caminho de texto ao longo do qual os caracteres de um texto são renderizados ou um caminho de
//	      movimento ao longo do qual um elemento referenciado é animado
//	    factory: factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()
//	    string: "M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Path(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(*SvgPath); ok {
		e.selfElement.Call("setAttribute", "path", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "path", value)
	return e
}

// PathLength
//
// English:
//
// The pathLength attribute lets authors specify a total length for the path, in user units.
//
//	Input:
//	  value: lets authors specify a total length for the path
//
// This value is then used to calibrate the browser's distance calculations with those of the author, by scaling all
// distance computations using the ratio pathLength/(computed value of path length).
//
// This can affect the actual rendered lengths of paths; including text paths, animation paths, and various stroke
// operations. Basically, all computations that require the length of the path. stroke-dasharray, for example, will
// assume the start of the path being 0 and the end point the value defined in the pathLength attribute.
//
// Português:
//
// O atributo pathLength permite que os autores especifiquem um comprimento total para o caminho, em unidades de
// usuário.
//
//	Entrada:
//	  value: permite que os autores especifiquem um comprimento total para o caminho
//
// Este valor é então usado para calibrar os cálculos de distância do navegador com os do autor, escalando todos os
// cálculos de distância usando a razão pathLength (valor calculado do comprimento do caminho).
//
// Isso pode afetar os comprimentos reais dos caminhos renderizados; incluindo caminhos de texto, caminhos de animação
// e várias operações de traçado. Basicamente, todos os cálculos que exigem o comprimento do caminho. stroke-dasharray,
// por exemplo, assumirá o início do caminho sendo 0 e o ponto final o valor definido no atributo pathLength.
func (e *TagSvgGlobal) PathLength(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "pathLength", value)
	return e
}

// PatternContentUnits
//
// English:
//
// The patternContentUnits attribute indicates which coordinate system to use for the contents of the <pattern> element.
//
//	Input:
//	  value: specifies the coordinate system
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    any other type: interface{}
//
//	Notes:
//	  * That this attribute has no effect if attribute viewBox is specified on the <pattern> element.
//
// Português:
//
// O atributo patternContentUnits indica qual sistema de coordenadas deve ser usado para o conteúdo do elemento
// <pattern>.
//
//	Entrada:
//	  value: especifica o sistema de coordenadas
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Que este atributo não tem efeito se o atributo viewBox for especificado no elemento <pattern>.
func (e *TagSvgGlobal) PatternContentUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgUnits); ok {
		e.selfElement.Call("setAttribute", "patternContentUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "patternContentUnits", value)
	return e
}

// PatternTransform
//
// English:
//
// The patternTransform attribute defines a list of transform definitions that are applied to a pattern tile.
//
//	Input:
//	  value: defines a list of transform definitions that are applied to a pattern tile
//	    factory: e.g. factoryBrowser.NewTransform().RotateAngle(20).SkewX(30).Scale(1, 0.5)
//	    string: e.g. "rotate(20) skewX(30) scale(1 0.5)"
//	    any other type: interface{}
//
// Português:
//
// O atributo patternTransform define uma lista de definições de transformação que são aplicadas a um tile.
//
//	Entrada:
//	  value: define uma lista de definições de transformação que são aplicadas a um tile
//	    factory: ex. factoryBrowser.NewTransform().RotateAngle(20).SkewX(30).Scale(1, 0.5)
//	    string: ex. "rotate(20) skewX(30) scale(1 0.5)"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) PatternTransform(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(*TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "patternTransform", converted.String())
		return e
	}

	if converted, ok := value.(TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "patternTransform", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "patternTransform", value)
	return e
}

// PatternUnits
//
// English:
//
// The patternUnits attribute indicates which coordinate system to use for the geometry properties of the <pattern>
// element.
//
//	Input:
//	  value: specifies the coordinate system
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    any other type: interface{}
//
// Português:
//
// O atributo patternUnits indica qual sistema de coordenadas deve ser usado para as propriedades geométricas do
// elemento <pattern>.
//
//	Entrada:
//	  value: especifica o sistema de coordenadas
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) PatternUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgUnits); ok {
		e.selfElement.Call("setAttribute", "patternUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "patternUnits", value)
	return e
}

// PointerEvents
//
// English:
//
// The pointer-events attribute is a presentation attribute that allows defining whether or when an element may be the
// target of a mouse event.
//
//	Input:
//	  value: defining whether or when an element may be the target of a mouse event
//	    const: KSvgPointerEvents... (e.g. KSvgPointerEventsVisibleStroke)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute pointer-events can be used as a CSS property.
//
// Português:
//
// O atributo pointer-events é um atributo de apresentação que permite definir se ou quando um elemento pode ser alvo
// de um evento de mouse.
//
//	Entrada:
//	  value: define se ou quando um elemento pode ser alvo de um evento de mouse.
//	    const: KSvgPointerEvents... (e.g. KSvgPointerEventsVisibleStroke)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um atributo de apresentação, os eventos de ponteiro podem ser usados como uma propriedade CSS.
func (e *TagSvgGlobal) PointerEvents(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgPointerEvents); ok {
		e.selfElement.Call("setAttribute", "pointer-events", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "pointer-events", value)
	return e
}

// Points
//
// English:
//
// The points attribute defines a list of points. Each point is defined by a pair of number representing a X and a Y
// coordinate in the user coordinate system. If the attribute contains an odd number of coordinates, the last one will
// be ignored.
//
//	Input:
//	  value: list of points representing coordinates X and Y
//	    [][]float64: [][]float64{{0,0},{1,1},{2,2}} = "0,0 1,1 2,2"
//	    any other type: interface{}
//
// Português:
//
// O atributo points define uma lista de pontos. Cada ponto é definido por um par de números representando uma
// coordenada X e Y no sistema de coordenadas do usuário. Se o atributo contiver um número ímpar de coordenadas, a
// última será ignorada.
//
//	Entrada:
//	  value: lista de pontos representando as coordenadas X e Y
//	    [][]float64: [][]float64{{0,0},{1,1},{2,2}} = "0,0 1,1 2,2"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Points(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "points", TypeToString(value, ",", " "))
	return e
}

// PointsAtX
//
// English:
//
// The pointsAtX attribute represents the x location in the coordinate system established by attribute primitiveUnits
// on the <filter> element of the point at which the light source is pointing.
//
//	Input:
//	  value: represents the x location in the coordinate system
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo pointsAtX representa a localização x no sistema de coordenadas estabelecido pelo atributo primitivaUnits
// no elemento <filter> do ponto para o qual a fonte de luz está apontando.
//
//	Entrada:
//	  value: representa a localização x no sistema de coordenadas
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) PointsAtX(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "pointsAtX", p)
		return e
	}

	e.selfElement.Call("setAttribute", "pointsAtX", value)
	return e
}

// PointsAtY
//
// English:
//
// The pointsAtY attribute represents the y location in the coordinate system established by attribute primitiveUnits
// on the <filter> element of the point at which the light source is pointing.
//
//	Input:
//	  value: represents the y location in the coordinate system
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo pointsAtY representa a localização y no sistema de coordenadas estabelecido pelo atributo primitivaUnits
// no elemento <filter> do ponto para o qual a fonte de luz está apontando.
//
//	Entrada:
//	  value: representa a localização y no sistema de coordenadas
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) PointsAtY(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "pointsAtY", p)
		return e
	}

	e.selfElement.Call("setAttribute", "pointsAtY", value)
	return e
}

// PointsAtZ
//
// English:
//
// The pointsAtZ attribute represents the y location in the coordinate system established by attribute primitiveUnits
// on the <filter> element of the point at which the light source is pointing, assuming that, in the initial local
// coordinate system, the positive z-axis comes out towards the person viewing the content and assuming that one unit
// along the z-axis equals one unit in x and y.
//
//	Input:
//	  value: represents the y location in the coordinate system
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo pointsAtZ representa a localização y no sistema de coordenadas estabelecido pelo atributo primitivaUnits
// no elemento <filter> do ponto em que a fonte de luz está apontando, assumindo que, no sistema de coordenadas local
// inicial, o eixo z positivo sai em direção a pessoa visualizando o conteúdo e assumindo que uma unidade ao longo do
// eixo z é igual a uma unidade em x e y.
//
//	Input:
//	  value: representa a localização y no sistema de coordenadas
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) PointsAtZ(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "pointsAtZ", p)
		return e
	}

	e.selfElement.Call("setAttribute", "pointsAtZ", value)
	return e
}

// PreserveAlpha
//
// English:
//
// The preserveAlpha attribute indicates how a <feConvolveMatrix> element handled alpha transparency.
//
//	Input:
//	  value: indicates how handled alpha transparency.
//
// Português:
//
// O atributo preserveAlpha indica como um elemento <feConvolveMatrix> trata a transparência alfa.
//
//	Input:
//	  value: indica como a transparência alfa é tratada.
func (e *TagSvgGlobal) PreserveAlpha(value bool) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "preserveAlpha", value)
	return e
}

// PreserveAspectRatio
//
// English:
//
//	The preserveAspectRatio attribute indicates how an element with a viewBox providing a given aspect ratio must fit
//	into a viewport with a different aspect ratio.
//
//	 Input:
//	   ratio: Indicates how an element with a viewBox providing a given aspect ratio.
//	     const: KRatio... (e.g. KRatioXMinYMin)
//	     any other type: interface{}
//	   meet: The meet or slice reference
//	     const: KMeetOrSliceReference... (e.g. KMeetOrSliceReferenceSlice)
//	     any other type: interface{}
//
// Because the aspect ratio of an SVG image is defined by the viewBox attribute, if this attribute isn't set, the
// preserveAspectRatio attribute has no effect (with one exception, the <image> element, as described below).
//
// Português:
//
//	O atributo preserveAspectRatio indica como um elemento com uma viewBox fornecendo uma determinada proporção deve
//	caber em uma viewport com uma proporção diferente.
//
//	 Input:
//	   ratio: Indica como um elemento com uma viewBox fornece uma determinada proporção.
//	     const: KRatio... (ex. KRatioXMinYMin)
//	     qualquer outro tipo: interface{}
//	   meet: A referência de encontro ou fatia
//	     const: KMeetOrSliceReference... (ex. KMeetOrSliceReferenceSlice)
//	     qualquer outro tipo: interface{}
//
// Como a proporção de uma imagem SVG é definida pelo atributo viewBox, se esse atributo não estiver definido, o
// atributo preserveAspectRatio não terá efeito (com uma exceção, o elemento <image>, conforme descrito abaixo).
func (e *TagSvgGlobal) PreserveAspectRatio(ratio, meet interface{}) (ref *TagSvgGlobal) {
	if converted, ok := ratio.(Ratio); ok {
		ratio = converted.String()
	}

	if converted, ok := meet.(MeetOrSliceReference); ok {
		meet = converted.String()
	}

	e.selfElement.Call("setAttribute", "preserveAspectRatio", fmt.Sprintf("%v %v", ratio, meet))
	return e
}

// PrimitiveUnits
//
// English:
//
// The primitiveUnits attribute specifies the coordinate system for the various length values within the filter
// primitives and for the attributes that define the filter primitive subregion.
//
//	Input:
//	  value: specifies the coordinate system
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    any other type: interface{}
//
// Português:
//
// O atributo primitivaUnits especifica o sistema de coordenadas para os vários valores de comprimento dentro das
// primitivas de filtro e para os atributos que definem a sub-região da primitiva de filtro.
//
//	Entrada:
//	  value: especifica o sistema de coordenadas
//	    const KSvgUnits... (e.g. KSvgUnitsObjectBoundingBox)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) PrimitiveUnits(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgUnits); ok {
		e.selfElement.Call("setAttribute", "primitiveUnits", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "primitiveUnits", value)
	return e
}

// R
//
// English:
//
// The r attribute defines the radius of a circle.
//
//	Input:
//	  value: radius of a circle
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo r define o raio de um círculo.
//
//	Input:
//	  value: raio de um círculo
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) R(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "r", p)
		return e
	}

	e.selfElement.Call("setAttribute", "r", value)
	return e
}

// Radius
//
// English:
//
// The radius attribute represents the radius (or radii) for the operation on a given <feMorphology> filter primitive.
//
//	Input:
//	  value: represents the radius (or radii) for the operation
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// If two numbers are provided, the first number represents the x-radius and the second one the y-radius. If one number
// is provided, then that value is used for both x and y. The values are in the coordinate system established by the
// primitiveUnits attribute on the <filter> element.
//
// A negative or zero value disables the effect of the given filter primitive (i.e., the result is the filter input
// image).
//
// Português:
//
// O atributo radius representa o raio (ou raios) para a operação em uma determinada primitiva de filtro <feMorphology>.
//
//	Entrada:
//	  value: representa o raio (ou raios) para à operação
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// Se dois números forem fornecidos, o primeiro número representa o raio x e o segundo o raio y. Se um número for
// fornecido, esse valor será usado para x e y. Os valores estão no sistema de coordenadas estabelecido pelo atributo
// primitivaUnits no elemento <filter>.
//
// Um valor negativo ou zero desativa o efeito da primitiva de filtro fornecida (ou seja, o resultado é a imagem de
// entrada do filtro).
func (e *TagSvgGlobal) Radius(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "radius", p)
		return e
	}

	e.selfElement.Call("setAttribute", "radius", value)
	return e
}

// RefX
//
// English:
//
// The refX attribute defines the x coordinate of an element's reference point.
//
//	Input:
//	  value: defines the x coordinate of an element's reference point
//	    float32: 1.0 = "100%"
//	    const: KPositionHorizontal... (e.g. KPositionHorizontalLeft)
//	    any other type: interface{}
//
// Português:
//
// O atributo refX define a coordenada x do ponto de referência de um elemento.
//
//	Entrada:
//	  value: define a coordenada x do ponto de referência de um elemento
//	    float32: 1.0 = "100%"
//	    const: KPositionHorizontal... (ex. KPositionHorizontalLeft)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) RefX(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "refX", p)
		return e
	}

	if converted, ok := value.(PositionHorizontal); ok {
		e.selfElement.Call("setAttribute", "refX", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "refX", value)
	return e
}

// RefY
//
// English:
//
// The refX attribute defines the y coordinate of an element's reference point.
//
//	Input:
//	  value: defines the y coordinate of an element's reference point
//	    float32: 1.0 = "100%"
//	    const: KPositionVertical... (e.g. KPositionVerticalTop)
//	    any other type: interface{}
//
// Português:
//
// O atributo refX define a coordenada y do ponto de referência de um elemento.
//
//	Entrada:
//	  value: define a coordenada y do ponto de referência de um elemento
//	    float32: 1.0 = "100%"
//	    const: KPositionVertical... (ex. KPositionVerticalTop)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) RefY(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "refY", p)
		return e
	}

	if converted, ok := value.(PositionVertical); ok {
		e.selfElement.Call("setAttribute", "refY", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "refY", value)
	return e
}

// RepeatCount
//
// English:
//
// The repeatCount attribute indicates the number of times an animation will take place.
//
//	Input:
//	  value: indicates the number of times an animation will take place
//	    int: number of times
//	    const: KSvgDurIndefinite
//	    any other type: interface{}
//
// Português:
//
// O atributo repeatCount indica o número de vezes que uma animação ocorrerá.
//
//	Input:
//	  value: indica o número de vezes que uma animação ocorrerá
//	    int: número de vezes
//	    const: KSvgDurIndefinite
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) RepeatCount(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgDur); ok {
		e.selfElement.Call("setAttribute", "repeatCount", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "repeatCount", value)
	return e
}

// RepeatDur
//
// English:
//
// The repeatDur attribute specifies the total duration for repeating an animation.
//
//	Input:
//	  value: specifies the total duration for repeating an animation
//	    string: "5s"
//	    time.Duration: 5*time.Second = "5s"
//	    const: KSvgDurIndefinite
//	    any other type: interface{}
//
// Português:
//
// O atributo repeatDur especifica a duração total para repetir uma animação.
//
//	Entrada:
//	  value: especifica a duração total para repetir uma animação
//	    string: "5s"
//	    time.Duration: 5*time.Second = "5s"
//	    const: KSvgDurIndefinite
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) RepeatDur(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "repeatDur", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "repeatDur", value)
	return e
}

// Restart
//
// English:
//
// The restart attribute specifies whether or not an animation can restart.
//
//	Input:
//	  value: especifica se uma animação pode ou não reiniciar
//	    const: KSvgAnimationRestart... (e.g. KSvgAnimationRestartAlways)
//	    any other type: interface{}
//
// Português:
//
// O atributo restart especifica se uma animação pode ou não reiniciar.
//
//	Entrada:
//	  value: especifica se uma animação pode ou não reiniciar
//	    const: KSvgAnimationRestart... (ex. KSvgAnimationRestartAlways)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Restart(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgAnimationRestart); ok {
		e.selfElement.Call("setAttribute", "restart", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "restart", value)
	return e
}

// Result
//
// English:
//
// The result attribute defines the assigned name for this filter primitive.
//
//	Input:
//	  value: defines the assigned name for this filter primitive
//
// If supplied, then graphics that result from processing this filter primitive can be referenced by an in attribute on
// a subsequent filter primitive within the same <filter> element. If no value is provided, the output will only be
// available for re-use as the implicit input into the next filter primitive if that filter primitive provides no value
// for its in attribute.
//
// Português:
//
// O atributo result define o nome atribuído para esta primitiva de filtro.
//
//	Entrada:
//	  value: define o nome atribuído para esta primitiva de filtro
//
// Se fornecido, os gráficos resultantes do processamento dessa primitiva de filtro podem ser referenciados por um
// atributo in em uma primitiva de filtro subsequente dentro do mesmo elemento <filter>. Se nenhum valor for fornecido,
// a saída só estará disponível para reutilização como entrada implícita na próxima primitiva de filtro se essa
// primitiva de filtro não fornecer valor para seu atributo in.
func (e *TagSvgGlobal) Result(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "result", value)
	return e
}

// Rotate
//
// English:
//
// The rotate attribute specifies how the animated element rotates as it travels along a path specified in an
// <animateMotion> element.
//
//	Input:
//	  value: specifies how the animated element rotates
//	    const: KSvgRotate... (e.g. KSvgRotateAutoReverse)
//	    any other type: interface{}
//
// Português:
//
// O atributo de rotação especifica como o elemento animado gira enquanto percorre um caminho especificado em um
// elemento <animateMotion>.
//
//	Entrada:
//	  value: especifica como o elemento animado gira
//	    const: KSvgRotate... (e.g. KSvgRotateAutoReverse)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Rotate(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgRotate); ok {
		e.selfElement.Call("setAttribute", "rotate", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "rotate", value)
	return e
}

// Rx
//
// English:
//
// The rx attribute defines a radius on the x-axis.
//
//	Input:
//	  value: defines a radius on the x-axis
//
// Português:
//
// O atributo rx define um raio no eixo x.
//
//	Entrada:
//	  value: defines a radius on the x-axis
func (e *TagSvgGlobal) Rx(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "rx", value)
	return e
}

// Ry
//
// English:
//
// The ry attribute defines a radius on the y-axis.
//
//	Input:
//	  value: defines a radius on the y-axis
//
// Português:
//
// O atributo ry define um raio no eixo y.
//
//	Entrada:
//	  value: define um raio no eixo y
func (e *TagSvgGlobal) Ry(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "ry", value)
	return e
}

// Scale
//
// English:
//
// The scale attribute defines the displacement scale factor to be used on a <feDisplacementMap> filter primitive.
//
//	Input:
//	  value: defines the displacement scale factor
//
// The amount is expressed in the coordinate system established by the primitiveUnits attribute on the <filter> element.
//
// Português:
//
// O atributo scale define o fator de escala de deslocamento a ser usado em uma primitiva de filtro <feDisplacementMap>.
//
//	Entrada:
//	  value: define o fator de escala de deslocamento
//
// A quantidade é expressa no sistema de coordenadas estabelecido pelo atributo primitivaUnits no elemento <filter>.
func (e *TagSvgGlobal) Scale(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "scale", value)
	return e
}

// Seed
//
// English:
//
// The seed attribute represents the starting number for the pseudo random number generator of the <feTurbulence> filter
// primitive.
//
//	Input:
//	  value: represents the starting number for the pseudo random number generator
//
// Português:
//
// O atributo seed representa o número inicial para o gerador de números pseudo aleatórios da primitiva de filtro
// <feTurbulence>.
//
//	Entrada:
//	  value: representa o número inicial para o gerador de números pseudo aleatórios
func (e *TagSvgGlobal) Seed(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "seed", value)
	return e
}

// ShapeRendering
//
// English:
//
// The shape-rendering attribute provides hints to the renderer about what tradeoffs to make when rendering shapes like
// paths, circles, or rectangles.
//
//	Input:
//	  value: provides hints to the renderer
//	    const: KSvgShapeRendering... (e.g. KSvgShapeRenderingAuto)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, shape-rendering can be used as a CSS property.
//
// Português:
//
// O atributo shape-rendering fornece dicas ao renderizador sobre quais compensações fazer ao renderizar formas como
// caminhos, círculos ou retângulos.
//
//	Entrada:
//	  value: fornece dicas para o renderizador
//	    const: KSvgShapeRendering... (ex. KSvgShapeRenderingAuto)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um atributo de apresentação, a renderização de forma pode ser usada como uma propriedade CSS.
func (e *TagSvgGlobal) ShapeRendering(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgShapeRendering); ok {
		e.selfElement.Call("setAttribute", "shape-rendering", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "shape-rendering", value)
	return e
}

// Side
//
// English:
//
// The side attribute determines the side of a path the text is placed on (relative to the path direction).
//
//	Input:
//	  value: side of a path the text is placed
//	    const: KSvgSide... (e.g. KSvgSideRight)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, shape-rendering can be used as a CSS property.
//
// Português:
//
// O atributo side determina o lado de um caminho em que o texto é colocado (em relação à direção do caminho).
//
//	Entrada:
//	  value: lado de um caminho em que o texto é colocado
//	    const: KSvgSide... (e.g. KSvgSideRight)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um atributo de apresentação, a renderização de forma pode ser usada como uma propriedade CSS.
func (e *TagSvgGlobal) Side(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgSide); ok {
		e.selfElement.Call("setAttribute", "side", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "side", value)
	return e
}

// Spacing
//
// English:
//
// The spacing attribute indicates how the user agent should determine the spacing between typographic characters that
// are to be rendered along a path.
//
//	Input:
//	  value: indicates how the user agent should determine the spacing
//	    const: KSvgSpacing... (e.g. KSvgSpacingExact)
//	    any other type: interface{}
//
// Português:
//
// O atributo spacing indica como o agente do usuário deve determinar o espaçamento entre os caracteres tipográficos que
// devem ser renderizados ao longo de um caminho.
//
//	Entrada:
//	  value: indica como o agente do usuário deve determinar o espaçamento
//	    const: KSvgSpacing... (ex. KSvgSpacingExact)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Spacing(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgSpacing); ok {
		e.selfElement.Call("setAttribute", "spacing", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "spacing", value)
	return e
}

// SpecularConstant
//
// English:
//
// The specularConstant attribute controls the ratio of reflection of the specular lighting. It represents the ks value
// in the Phong lighting model. The bigger the value the stronger the reflection.
//
//	Input:
//	  value: controls the ratio of reflection of the specular lighting
//
// Português:
//
// O atributo specularConstant controla a proporção de reflexão da iluminação especular. Ele representa o valor ks no
// modelo de iluminação Phong. Quanto maior o valor, mais forte a reflexão.
//
//	Entrada:
//	  value: controls the ratio of reflection of the specular lighting
func (e *TagSvgGlobal) SpecularConstant(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "specularConstant", value)
	return e
}

// SpecularExponent
//
// English:
//
// The specularExponent attribute controls the focus for the light source. The bigger the value the brighter the light.
//
//	Input:
//	  value: controls the focus for the light source
//
// Português:
//
// O atributo specularExponent controla o foco da fonte de luz. Quanto maior o valor, mais brilhante é a luz.
//
//	Entrada:
//	  value: controla o foco para a fonte de luz
func (e *TagSvgGlobal) SpecularExponent(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "specularExponent", value)
	return e
}

// SpreadMethod
//
// English:
//
// The spreadMethod attribute determines how a shape is filled beyond the defined edges of a gradient.
//
//	Input:
//	  value: determines how a shape is filled
//	    const: KSvgSpreadMethod... (e.g. KSvgSpreadMethodReflect)
//	    any other type: interface{}
//
// Português:
//
// O atributo spreadMethod determina como uma forma é preenchida além das bordas definidas de um gradiente.
//
//	Entrada:
//	  value: determina como uma forma é preenchida
//	    const: KSvgSpreadMethod... (e.g. KSvgSpreadMethodReflect)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) SpreadMethod(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgSpreadMethod); ok {
		e.selfElement.Call("setAttribute", "spreadMethod", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "spreadMethod", value)
	return e
}

// StartOffset
//
// English:
//
// The startOffset attribute defines an offset from the start of the path for the initial current text position along
// the path after converting the path to the <textPath> element's coordinate system.
//
//	Input:
//	  value: defines an offset from the start
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo startOffset define um deslocamento do início do caminho para a posição inicial do texto atual ao longo do
// caminho após a conversão do caminho para o sistema de coordenadas do elemento <textPath>.
//
//	Entrada:
//	  value: define um deslocamento desde o início
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) StartOffset(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "startOffset", p)
		return e
	}

	e.selfElement.Call("setAttribute", "startOffset", value)
	return e
}

// StitchTiles
//
// English:
//
// The stitchTiles attribute defines how the Perlin Noise tiles behave at the border.
//
//	Input:
//	  value: defines how the Perlin Noise tiles behave at the border
//	    const: KSvgStitchTiles... (e.g. KSvgStitchTilesNoStitch)
//	    any other type: interface{}
//
// Português:
//
// O atributo stitchTiles define como os blocos Perlin Noise se comportam na borda.
//
//	Entrada:
//	  value: define como os blocos Perlin Noise se comportam na borda
//	    const: KSvgStitchTiles... (ex. KSvgStitchTilesNoStitch)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) StitchTiles(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgStitchTiles); ok {
		e.selfElement.Call("setAttribute", "stitchTiles", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "stitchTiles", value)
	return e
}

// StdDeviation
//
// English:
//
// The stdDeviation attribute defines the standard deviation for the blur operation.
//
//	Input:
//	  value: defines the standard deviation
//	    []float64: []float64{2,5} = "2 5"
//	    any other type: interface{}
//
// Português:
//
// O atributo stdDeviation define o desvio padrão para a operação de desfoque.
//
//	Input:
//	  value: define o desvio padrão
//	    []float64: []float64{2,5} = "2 5"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) StdDeviation(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		str := ""
		for _, v := range converted {
			str += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}
		length := len(str) - 1

		e.selfElement.Call("setAttribute", "stdDeviation", str[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "stdDeviation", value)
	return e
}

// StopColor #presentation
//
// English:
//
// The stop-color attribute indicates what color to use at a gradient stop.
//
//	Input:
//	  value: indicates what color to use at a gradient stop
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//
//	Notes:
//	  * With respect to gradients, SVG treats the transparent keyword differently than CSS. SVG does not calculate
//	    gradients in pre-multiplied space, so transparent really means transparent black. So, specifying a stop-color
//	    with the value transparent is equivalent to specifying a stop-color with the value black and a stop-opacity
//	    with the value 0.
//	  * As a presentation attribute, stop-color can be used as a CSS property.
//
// Português:
//
// O atributo stop-color indica qual cor usar em uma parada de gradiente.
//
//	Entrada:
//	  value: indica qual cor usar em um fim de gradiente
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//
//	Notss:
//	  * Com relação aos gradientes, o SVG trata a palavra-chave transparente de maneira diferente do CSS. O SVG não
//	    calcula gradientes no espaço pré-multiplicado, portanto, transparente realmente significa preto transparente.
//	    Assim, especificar uma stop-color com o valor transparente é equivalente a especificar uma stop-color com o
//	    valor black e uma stop-opacity com o valor 0.
//	  * Como atributo de apresentação, stop-color pode ser usado como propriedade CSS.
func (e *TagSvgGlobal) StopColor(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "stop-color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "stop-color", value)
	return e
}

// StopOpacity #presentation
//
// English:
//
// The stop-opacity attribute defines the opacity of a given color gradient stop.
//
//	Input:
//	  value: defines the opacity of a given color gradient stop
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// The opacity value used for the gradient calculation is the product of the value of stop-opacity and the opacity of
// the value of the stop-color attribute. For stop-color values that don't include explicit opacity information, the
// opacity is treated as 1.
//
//	Notes:
//	  * As a presentation attribute, stop-opacity can be used as a CSS property.
//
// Português:
//
// O atributo stop-opacity define a opacidade de uma determinada parada de gradiente de cor.
//
//	Entrada:
//	  value: define a opacidade de uma determinada parada de gradiente de cor
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// O valor de opacidade usado para o cálculo do gradiente é o produto do valor de stop-opacity e a opacidade do valor
// do atributo stop-color. Para valores de stop-color que não incluem informações explícitas de opacidade, a opacidade
// é tratada como 1.
//
//	Notas:
//	  * Como atributo de apresentação, stop-opacity pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) StopOpacity(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "stop-opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "stop-opacity", value)
	return e
}

// Stroke #presentation
//
// English:
//
// The stroke attribute is a presentation attribute defining the color (or any SVG paint servers like gradients or
// patterns) used to paint the outline of the shape
//
//	Input:
//	  value: presentation attribute defining the color
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute stroke can be used as a CSS property.
//
// Português:
//
// O atributo de traço é um atributo de apresentação que define a cor (ou qualquer servidor de pintura SVG, como
// gradientes ou padrões) usado para pintar o contorno da forma
//
//	Entrada:
//	  value: atributo de apresentação que define a cor
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um traço de atributo de apresentação pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) Stroke(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "stroke", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "stroke", value)
	return e
}

// StrokeDasharray #presentation
//
// English:
//
// The stroke-dasharray attribute is a presentation attribute defining the pattern of dashes and gaps used to paint the
// outline of the shape
//
//	Input:
//	  value: presentation attribute defining the pattern of dashes
//	    []float64: (e.g. []float64{4, 1, 2}) = "4 1 2"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, stroke-dasharray can be used as a CSS property.
//
// Português:
//
// O atributo stroke-dasharray é um atributo de apresentação que define o padrão de traços e lacunas usados para pintar
// o contorno da forma
//
//	Entrada:
//	  value: atributo de apresentação que define o padrão de traços
//	    []float64: (e.g. []float64{4, 1, 2}) = "4 1 2"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o stroke-dasharray pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) StrokeDasharray(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		str := ""
		for _, v := range converted {
			str += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}
		length := len(str) - 1

		e.selfElement.Call("setAttribute", "stroke-dasharray", str[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-dasharray", value)
	return e
}

// StrokeDashOffset
//
// English:
//
// The stroke-dasharray attribute is a presentation attribute defining the pattern of dashes and gaps used to paint the
// outline of the shape
//
//	Input:
//	  value: presentation attribute defining the pattern of dashes
//	    float32: 0.1 = "10%"
//	    []float32: (e.g. []float32{0.04, 0.01, 0.02}) = "4% 1% 2%"
//	    []float64: (e.g. []float64{4, 1, 2}) = "4 1 2"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, stroke-dasharray can be used as a CSS property.
//
// Português:
//
// O atributo stroke-dasharray é um atributo de apresentação que define o padrão de traços e lacunas usados para pintar
// o contorno da forma
//
//	Entrada:
//	  value: atributo de apresentação que define o padrão de traços
//	    float32: 0.1 = "10%"
//	    []float32: (e.g. []float32{0.04, 0.01, 0.02}) = "4% 1% 2%"
//	    []float64: (e.g. []float64{4, 1, 2}) = "4 1 2"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o stroke-dasharray pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) StrokeDashOffset(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float32); ok {
		str := ""
		for _, v := range converted {
			str += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "% "
		}
		length := len(str) - 1

		e.selfElement.Call("setAttribute", "stroke-dashoffset", str[:length])
		return e
	}

	if converted, ok := value.([]float64); ok {
		str := ""
		for _, v := range converted {
			str += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}
		length := len(str) - 1

		e.selfElement.Call("setAttribute", "stroke-dashoffset", str[:length])
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "stroke-dashoffset", p)
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-dashoffset", value)
	return e
}

// StrokeLineCap #presentation
//
// English:
//
// The stroke-linecap attribute is a presentation attribute defining the shape to be used at the end of open subpaths
// when they are stroked.
//
//	Input:
//	  value: presentation attribute defining the shape to be used at the end of open subpaths
//	    const: KSvgStrokeLinecap... (e.g. KSvgStrokeLinecapRound)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute stroke-linecap can be used as a CSS property.
//
// Português:
//
// O atributo stroke-linecap é um atributo de apresentação que define a forma a ser usada no final de subcaminhos
// abertos quando eles são traçados.
//
//	Input:
//	  value: atributo de apresentação que define a forma a ser usada no final de subcaminhos
//	    const: KSvgStrokeLinecap... (e.g. KSvgStrokeLinecapRound)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o traço-linecap pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) StrokeLineCap(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgStrokeLinecap); ok {
		e.selfElement.Call("setAttribute", "stroke-linecap", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-linecap", value)
	return e
}

// StrokeLineJoin #presentation
//
// English:
//
// The stroke-linejoin attribute is a presentation attribute defining the shape to be used at the corners of paths when
// they are stroked.
//
//	Input:
//	  value: defining the shape to be used at the corners of paths
//	    const: KSvgStrokeLinejoin... (e.g. KSvgStrokeLinejoinBevel)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute stroke-linejoin can be used as a CSS property.
//
// Português:
//
// O atributo stroke-linejoin é um atributo de apresentação que define a forma a ser usada nos cantos dos caminhos
// quando eles são traçados.
//
//	Entrada:
//	  value: define a forma a ser usada nos cantos dos caminhos
//	    const: KSvgStrokeLinejoin... (ex. KSvgStrokeLinejoinBevel)
//	    any other type: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, stroke-linejoin pode ser usado como propriedade CSS.
func (e *TagSvgGlobal) StrokeLineJoin(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgStrokeLinejoin); ok {
		e.selfElement.Call("setAttribute", "stroke-linejoin", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-linejoin", value)
	return e
}

// StrokeMiterLimit
//
// English:
//
// The stroke-miterlimit attribute is a presentation attribute defining a limit on the ratio of the miter length to the
// stroke-width used to draw a miter join. When the limit is exceeded, the join is converted from a miter to a bevel.
//
//	Input:
//	  value: defining a limit on the ratio of the miter length
//
//	Notes:
//	  * As a presentation attribute stroke-miterlimit can be used as a CSS property.
//
// Português:
//
// O atributo stroke-miterlimit é um atributo de apresentação que define um limite na proporção do comprimento da mitra
// para a largura do traço usado para desenhar uma junção de mitra. Quando o limite é excedido, a junção é convertida
// de uma mitra para um chanfro.
//
//	Entrada:
//	  value: definindo um limite na proporção do comprimento da mitra
//
//	Notas:
//	  * Como atributo de apresentação, stroke-miterlimit pode ser usado como propriedade CSS.
func (e *TagSvgGlobal) StrokeMiterLimit(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "stroke-miterlimit", value)
	return e
}

// StrokeOpacity #presentation
//
// English:
//
// The stroke-opacity attribute is a presentation attribute defining the opacity of the paint server (color, gradient,
// pattern, etc) applied to the stroke of a shape.
//
//	Input:
//	  value: defining the opacity of the paint
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute stroke-opacity can be used as a CSS property.
//
// Português:
//
// O atributo de opacidade do traçado é um atributo de apresentação que define a opacidade do servidor de pintura (cor,
// gradiente, padrão etc.) aplicado ao traçado de uma forma.
//
//	Entrada:
//	  value: definindo a opacidade da tinta
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, a opacidade do traço pode ser usada como uma propriedade CSS.
func (e *TagSvgGlobal) StrokeOpacity(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "stroke-opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-opacity", value)
	return e
}

// StrokeWidth #presentation
//
// English:
//
// The stroke-width attribute is a presentation attribute defining the width of the stroke to be applied to the shape.
//
//	Input:
//	  value: defining the width of the stroke
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo stroke-width é um atributo de apresentação que define a largura do traço a ser aplicado à forma.
//
//	Entrada:
//	  value: definindo a largura do traço
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) StrokeWidth(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "stroke-width", p)
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-width", value)
	return e
}

// Style
//
// English:
//
// The style attribute allows to style an element using CSS declarations.
//
//	Input:
//	  value: allows to style an element using CSS declarations
//
// It functions identically to the style attribute in HTML.
//
// Português:
//
// O atributo style permite estilizar um elemento usando declarações CSS.
//
//	Entrada:
//	  value: permite estilizar um elemento usando declarações CSS
//
// Funciona de forma idêntica ao atributo style em HTML.
func (e *TagSvgGlobal) Style(value string) (ref *TagSvgGlobal) {
	e.selfElement.Set("innerHTML", value)
	return e
}

// SurfaceScale
//
// English:
//
// The surfaceScale attribute represents the height of the surface for a light filter primitive.
//
//	Input:
//	  value: represents the height of the surface for a light filter primitive
//
// Português:
//
// O atributo surfaceScale representa a altura da superfície para uma primitiva de filtro de luz.
//
//	Entrada:
//	  value: representa a altura da superfície para uma primitiva de filtro de luz
func (e *TagSvgGlobal) SurfaceScale(value float64) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "surfaceScale", value)
	return e
}

// SystemLanguage
//
// English:
//
// The systemLanguage attribute represents a list of supported language tags. This list is matched against the language
// defined in the user preferences.
//
//	Input:
//	  value: list of supported language tags
//	    const: KLanguage... (e.g. KLanguageEnglishGreatBritain)
//	    []Language: e.g. []Language{KLanguageEnglishAustralia, KLanguageEnglishAustralia, KLanguageEnglishCanada}
//	    string: e.g. "en-gb, en-us"
//
// Português:
//
// O atributo systemLanguage representa uma lista de tags de idioma com suporte. Esta lista é comparada com o idioma
// definido nas preferências do usuário.
//
//	Entrada:
//	  value: lista de tags de idioma com suporte
//	    const: KLanguage... (e.g. KLanguagePortugueseBrazil)
//	    []Language: e.g. []Language{KLanguagePortugueseBrazil, KLanguagePortuguesePortugal}
//	    string: e.g. "pt-br, pt-pt"
func (e *TagSvgGlobal) SystemLanguage(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "systemLanguage", converted.String())
		return e
	}

	if converted, ok := value.([]Language); ok {
		tags := ""
		for _, v := range converted {
			tags += v.String() + ", "
		}
		length := len(tags) - 2

		e.selfElement.Call("setAttribute", "systemLanguage", tags[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "systemLanguage", value)
	return e
}

// Tabindex
//
// English:
//
// The tabindex attribute allows you to control whether an element is focusable and to define the relative order of the
// element for the purposes of sequential focus navigation.
//
//	Input:
//	  value: control whether an element is focusable
//	    int: focus order
//	    nil: focus disabled
//	    any other type: interface{}
//
// Português:
//
// O atributo tabindex permite controlar se um elemento é focalizável e definir à ordem relativa do elemento para fins
// de navegação de foco sequencial.
//
//	Input:
//	  value: controlar se um elemento é focalizável
//	    int: ordem do foco
//	    nil: disabilita o foco
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Tabindex(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "tabindex", value)
	return e
}

// TableValues
//
// English:
//
// The tableValues attribute defines a list of numbers defining a lookup table of values for a color component transfer
// function.
//
//	Input:
//	  value: defines a list of numbers
//	    []float64: e.g. []float64{0.0, 1.0} = "0 1"
//	    any other type: interface{}
//
// Português:
//
// O atributo tableValues define uma lista de números que definem uma tabela de consulta de valores para uma função de
// transferência de componente de cor.
//
//	Entrada:
//	  value: define uma lista de números
//	    []float64: ex. []float64{0.0, 1.0} = "0 1"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) TableValues(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		tags := ""
		for _, v := range converted {
			tags += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}
		length := len(tags) - 1

		e.selfElement.Call("setAttribute", "tableValues", tags[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "tableValues", value)
	return e
}

// Target
//
// English:
//
// This attribute specifies the name of the browsing context (e.g., a browser tab or an (X)HTML iframe or object
// element) into which a document is to be opened when the link is activated
//
//	Input:
//	  value: specifies the name of the browsing context
//	    const: KTarget... (e.g. KTargetSelf)
//	   any other type: interface{}
//
// The target attribute should be used when there are multiple possible targets for the ending resource, such as when
// the parent document is embedded within an HTML or XHTML document, or is viewed with a tabbed browser.
//
// Português:
//
// Este atributo especifica o nome do contexto de navegação (por exemplo, uma guia do navegador ou um iframe ou elemento
// de objeto (X)HTML) no qual um documento deve ser aberto quando o link é ativado
//
//	Entrada:
//	  value: especifica o nome do contexto de navegação
//	    const: KTarget... (e.g. KTargetSelf)
//	    qualquer outro tipo: interface{}
//
// O atributo target deve ser usado quando houver vários destinos possíveis para o recurso final, como quando o
// documento pai estiver incorporado em um documento HTML ou XHTML ou for visualizado em um navegador com guias.
func (e *TagSvgGlobal) Target(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(Target); ok {
		e.selfElement.Call("setAttribute", "target", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "target", value)
	return e
}

// TargetX
//
// English:
//
// The targetX attribute determines the positioning in horizontal direction of the convolution matrix relative to a
// given target pixel in the input image. The leftmost column of the matrix is column number zero.
// The value must be such that: 0 <= targetX < orderX.
//
//	Input:
//	  value: determines the positioning in horizontal direction
//
// Português:
//
// O atributo targetX determina o posicionamento na direção horizontal da matriz de convolução em relação a um
// determinado pixel alvo na imagem de entrada. A coluna mais à esquerda da matriz é a coluna número zero.
// O valor deve ser tal que: 0 <= targetX < orderX.
//
//	Entrada:
//	  value: determina o posicionamento na direção horizontal
func (e *TagSvgGlobal) TargetX(value int) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "targetX", value)
	return e
}

// TargetY
//
// English:
//
// The targetY attribute determines the positioning in vertical direction of the convolution matrix relative to a given
// target pixel in the input image. The topmost row of the matrix is row number zero.
// The value must be such that: 0 <= targetY < orderY.
//
//	Input:
//	  value: determines the positioning in vertical direction
//
// Português:
//
// O atributo targetY determina o posicionamento na direção vertical da matriz de convolução em relação a um determinado
// pixel alvo na imagem de entrada. A linha superior da matriz é a linha número zero.
// O valor deve ser tal que: 0 <= targetY < orderY.
//
//	Entrada:
//	  value: determines the positioning in vertical direction
func (e *TagSvgGlobal) TargetY(value int) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "targetY", value)
	return e
}

// TextAnchor #presentation
//
// English:
//
// The text-anchor attribute is used to align (start-, middle- or end-alignment) a string of pre-formatted text or
// auto-wrapped text where the wrapping area is determined from the inline-size property relative to a given point.
//
//	Input:
//	  value: used to align a string
//	    const: KSvgTextAnchor... (e.g. KSvgTextAnchorStart)
//	    any other type: interface{}
//
// This attribute is not applicable to other types of auto-wrapped text. For those cases you should use text-align.
// For multi-line text, the alignment takes place for each line.
//
// The text-anchor attribute is applied to each individual text chunk within a given <text> element. Each text chunk
// has an initial current text position, which represents the point in the user coordinate system resulting from
// (depending on context) application of the x and y attributes on the <text> element, any x or y attribute values on a
// <tspan>, <tref> or <altGlyph> element assigned explicitly to the first rendered character in a text chunk, or
// determination of the initial current text position for a <textPath> element.
//
//	Notes:
//	  * As a presentation attribute, text-anchor can be used as a CSS property.
//
// Português:
//
// O atributo text-anchor é usado para alinhar (alinhamento inicial, intermediário ou final) uma string de texto
// pré-formatado ou texto com quebra automática onde a área de quebra é determinada a partir da propriedade inline-size
// relativa a um determinado ponto.
//
//	Entrada:
//	  value: usado para alinhar uma string
//	    const: KSvgTextAnchor... (e.g. KSvgTextAnchorStart)
//	    qualquer outro tipo: interface{}
//
// Este atributo não se aplica a outros tipos de texto com quebra automática. Para esses casos, você deve usar
// text-align. Para texto de várias linhas, o alinhamento ocorre para cada linha.
//
// O atributo text-anchor é aplicado a cada fragmento de texto individual dentro de um determinado elemento <text>.
// Cada pedaço de texto tem uma posição inicial de texto atual, que representa o ponto no sistema de coordenadas do
// usuário resultante (dependendo do contexto) da aplicação dos atributos x e y no elemento <text>, quaisquer valores
// de atributo x ou y em um <tspan >, elemento <tref> ou <altGlyph> atribuído explicitamente ao primeiro caractere
// renderizado em um pedaço de texto, ou determinação da posição inicial do texto atual para um elemento <textPath>.
//
//	Notes:
//	  * As a presentation attribute, text-anchor can be used as a CSS property.
func (e *TagSvgGlobal) TextAnchor(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgTextAnchor); ok {
		e.selfElement.Call("setAttribute", "text-anchor", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "text-anchor", value)
	return e
}

// TextDecoration #presentation
//
// English:
//
// The text-decoration attribute defines whether text is decorated with an underline, overline and/or strike-through.
// It is a shorthand for the text-decoration-line and text-decoration-style properties.
//
//	Input:
//	  value: defines whether text is decorated
//	    const: KSvgTextDecorationLine... (e.g. KSvgTextDecorationLineUnderline)
//	    const: KSvgTextDecorationStyle... (e.g. KSvgTextDecorationStyleDouble)
//	    string: e.g. "black", "line-through"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// The fill and stroke of the text decoration are given by the fill and stroke of the text at the point where the text
// decoration is declared.
//
// The paint order of the text decoration, i.e. the fill and stroke, is determined by the value of the paint-order
// attribute at the point where the text decoration is declared.
//
//	Notes:
//	  * As a presentation attribute, text-decoration can be used as a CSS property. See the css text-decoration
//	    property for more information.
//
// Português:
//
// O atributo text-decoration define se o texto é decorado com sublinhado, overline e ou tachado.
// É um atalho para as propriedades text-decoration-line e text-decoration-style.
//
//	Entrada:
//	  value: define se o texto é decorado
//	    const: KSvgTextDecorationLine... (ex. KSvgTextDecorationLineUnderline)
//	    const: KSvgTextDecorationStyle... (ex. KSvgTextDecorationStyleDouble)
//	    string: e.g. "black", "line-through"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
//
// O preenchimento e o traçado da decoração de texto são dados pelo preenchimento e traçado do texto no ponto em que a
// decoração de texto é declarada.
//
// A ordem de pintura da decoração do texto, ou seja, o preenchimento e o traço, é determinada pelo valor do atributo
// paint-order no ponto em que a decoração do texto é declarada.
//
//	Notas:
//	  * Como atributo de apresentação, a decoração de texto pode ser usada como uma propriedade CSS. Consulte a
//	    propriedade CSS text-decoration para obter mais informações.
func (e *TagSvgGlobal) TextDecoration(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "text-decoration", RGBAToJs(converted))
		return e
	}

	if converted, ok := value.(SvgTextDecorationLine); ok {
		e.selfElement.Call("setAttribute", "text-decoration", converted.String())
		return e
	}

	if converted, ok := value.(SvgTextDecorationStyle); ok {
		e.selfElement.Call("setAttribute", "text-decoration", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "text-decoration", value)
	return e
}

// TextRendering #presentation
//
// English:
//
// The text-rendering attribute provides hints to the renderer about what tradeoffs to make when rendering text.
//
//	Notes:
//	  * As a presentation attribute, text-rendering can be used as a CSS property.
//	    See the css text-rendering property for more information.
//
// Português:
//
// O atributo text-rendering fornece dicas ao renderizador sobre quais compensações fazer ao renderizar o texto.
//
//	Notas:
//	  * Como um atributo de apresentação, a renderização de texto pode ser usada como uma propriedade CSS.
//	    Consulte a propriedade de renderização de texto css para obter mais informações.
func (e *TagSvgGlobal) TextRendering(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgTextRendering); ok {
		e.selfElement.Call("setAttribute", "text-rendering", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "text-rendering", value)
	return e
}

// TextLength
//
// English:
//
// The textLength attribute, available on SVG <text> and <tspan> elements, lets you specify the width of the space into
// which the text will draw. The user agent will ensure that the text does not extend farther than that distance, using
// the method or methods specified by the lengthAdjust attribute. By default, only the spacing between characters is
// adjusted, but the glyph size can also be adjusted if you change lengthAdjust.
//
//	Input:
//	  value: specify the width of the space into which the text will draw
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// By using textLength, you can ensure that your SVG text displays at the same width regardless of conditions including
// web fonts failing to load (or not having loaded yet).
//
// Português:
//
// O atributo textLength, disponível nos elementos SVG <text> e <tspan>, permite especificar a largura do espaço no qual
// o texto será desenhado. O agente do usuário garantirá que o texto não se estenda além dessa distância, usando o
// método ou métodos especificados pelo atributo lengthAdjust. Por padrão, apenas o espaçamento entre os caracteres é
// ajustado, mas o tamanho do glifo também pode ser ajustado se você alterar o lengthAdjust.
//
//	Input:
//	  value: especifique a largura do espaço no qual o texto será desenhado
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// Ao usar textLength, você pode garantir que seu texto SVG seja exibido na mesma largura, independentemente das
// condições, incluindo fontes da Web que não carregam (ou ainda não foram carregadas).
func (e *TagSvgGlobal) TextLength(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "textLength", p)
		return e
	}

	e.selfElement.Call("setAttribute", "textLength", value)
	return e
}

// To
//
// English:
//
// The to attribute indicates the final value of the attribute that will be modified during the animation.
//
//	Input:
//	  value: final value of the attribute
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1),rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s, 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    any other type: interface{}
//
// The value of the attribute will change between the from attribute value and this value.
//
// Português:
//
// O atributo to indica o valor final do atributo que será modificado durante a animação.
//
//	Entrada:
//	  value: valor final do atributo
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1),rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s, 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    qualquer outro tipo: interface{}
//
// O valor do atributo mudará entre o valor do atributo from e este valor.
func (e *TagSvgGlobal) To(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "values", TypeToString(value, ";", ";"))
	return e
}

// Transform
//
// English:
//
// The transform attribute defines a list of transform definitions that are applied to an element and the element's
// children.
//
//	Input:
//	  value: defines a list of transform definitions
//	    factory: e.g. factoryBrowser.NewTransform().Translate(100, 0).Scale(4, 1)
//	    string: e.g. "translate(300,0) scale(4,1)"
//	    any other type: interface{}
//
//	Notes:
//	  * As of SVG2, transform is a presentation attribute, meaning it can be used as a CSS property. However, be aware
//	    that there are some differences in syntax between the CSS property and the attribute. See the documentation for
//	    the CSS property transform for the specific syntax to use in that case.
//
// Português:
//
// O atributo transform define uma lista de definições de transformação que são aplicadas a um elemento e aos filhos do
// elemento.
//
//	Entrada:
//	  value: define uma lista de definições de transformação
//	    factory: ex. factoryBrowser.NewTransform().Translate(100, 0).Scale(4, 1)
//	    string: ex. "translate(300,0) scale(4,1)"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * A partir do SVG2, transform é um atributo de apresentação, o que significa que pode ser usado como uma
//	    propriedade CSS. No entanto, esteja ciente de que existem algumas diferenças na sintaxe entre a propriedade CSS
//	    e o atributo. Consulte a documentação da transformação da propriedade CSS para obter a sintaxe específica a ser
//	    usada nesse caso.
func (e *TagSvgGlobal) Transform(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(*TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "transform", converted.String())
		return e
	}

	if converted, ok := value.(TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "transform", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "transform", value)
	return e
}

// TransformOrigin
//
// English:
//
// The transform-origin SVG attribute sets the origin for an item's transformations.
//
//	Input:
//	  valueA, valueB: the origin for an item's transformations
//	    const: KSvgTransformOrigin... (e.g. KSvgTransformOriginLeft)
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute in SVG, transform-origin corresponds in syntax and behavior to the transform-origin
//	    property in CSS, and can be used as CSS property to style SVG. See the CSS transform-origin property for more
//	    information.
//
// Português:
//
// O atributo SVG transform-origin define a origem das transformações de um item.
//
//	Entrada:
//	  valueA, valueB: a origem das transformações de um item
//	    const: KSvgTransformOrigin... (ex. KSvgTransformOriginLeft)
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um atributo de apresentação em SVG, transform-origin corresponde em sintaxe e comportamento à propriedade
//	    transform-origin em CSS e pode ser usado como propriedade CSS para estilizar SVG. Consulte a propriedade
//	    transform-origin do CSS para obter mais informações.
func (e *TagSvgGlobal) TransformOrigin(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgTransformOrigin); ok {
		e.selfElement.Call("setAttribute", "transform-origin", converted.String())
		return e
	}

	if converted, ok := value.([]SvgTransformOrigin); ok {
		switch len(converted) {
		case 1:
			e.selfElement.Call("setAttribute", "transform-origin", converted[0].String())
			return e

		case 2:
			e.selfElement.Call("setAttribute", "transform-origin", fmt.Sprintf("%v %v", converted[0].String(), converted[1].String()))
			return e
		}
	}

	e.selfElement.Call("setAttribute", "transform-origin", TypeToString(value, ";", ";"))
	return e
}

// Type
//
// English:
//
// fixme: comentar
//
//	Input:
//	  value:
//	    fixme: comentar
//	    any other type: interface{}
//
// For the <animateTransform> element, it defines the type of transformation, whose values change over time.
// For the <feColorMatrix> element, it indicates the type of matrix operation. The keyword matrix indicates that a full 5x4 matrix of values will be provided. The other keywords represent convenience shortcuts to allow commonly used color operations to be performed without specifying a complete matrix.
// For the <feFuncR>, <feFuncG>, <feFuncB>, and <feFuncA> elements, it Indicates the type of component transfer function.
// For the <feTurbulence> element, it indicates whether the filter primitive should perform a noise or turbulence function.
// For the <style> and <script> elements, it defines the content type of the element.
//
// Português:
//
// fixme: comentar
//
//	Input:
//	  value:
//	    fixme: comentar
//	    any other type: interface{}
//
// Para o elemento <animateTransform>, define o tipo de transformação, cujos valores mudam ao longo do tempo.
// Para o elemento <feColorMatrix>, indica o tipo de operação da matriz. A matriz de palavras-chave indica que uma matriz de valores 5x4 completa será fornecida. As outras palavras-chave representam atalhos de conveniência para permitir que as operações de cores comumente usadas sejam executadas sem especificar uma matriz completa.
// Para os elementos <feFuncR>, <feFuncG>, <feFuncB> e <feFuncA>, indica o tipo de função de transferência do componente.
// Para o elemento <feTurbulence>, indica se a primitiva do filtro deve executar uma função de ruído ou turbulência.
// Para os elementos <style> e <script>, define o tipo de conteúdo do elemento.
func (e *TagSvgGlobal) Type(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgTypeTransform); ok { //fixme: fazer
		e.selfElement.Call("setAttribute", "type", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "type", value)
	return e
}

// UnderlinePosition
//
// English:
//
// The underline-position attribute represents the ideal vertical position of the underline. The underline position is
// expressed in the font's coordinate system.
//
//	Input:
//	  value: represents the ideal vertical position of the underline
//
// Português:
//
// O atributo underline-position representa a posição vertical ideal do sublinhado. A posição do sublinhado é expressa
// no sistema de coordenadas da fonte.
//
//	Entrada:
//	  value: representa a posição vertical ideal do sublinhado
func (e *TagSvgGlobal) UnderlinePosition(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "underline-position", value)
	return e
}

// UnicodeBidi #presentation
//
// English:
//
// The unicode-bidi attribute specifies how the accumulation of the background image is managed.
//
//	Input:
//	  value: specifies how the accumulation of the background image is managed
//	    const: KSvgTransformOrigin... (e.g. KSvgTransformOriginLeft)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, unicode-bidi can be used as a CSS property. See the CSS unicode-bidi property for
//	    more information.
//
// Português:
//
// O atributo unicode-bidi especifica como o acúmulo da imagem de fundo é gerenciado.
//
//	Entrada:
//	  value: especifica como o acúmulo da imagem de fundo é gerenciado
//	    const: KSvgTransformOrigin... (e.g. KSvgTransformOriginLeft)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o unicode-bidi pode ser usado como uma propriedade CSS. Consulte a propriedade
//	    CSS unicode-bidi para obter mais informações.
func (e *TagSvgGlobal) UnicodeBidi(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgTransformOrigin); ok {
		e.selfElement.Call("setAttribute", "unicode-bidi", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "unicode-bidi", value)
	return e
}

// Values
//
// English:
//
// The values attribute has different meanings, depending upon the context where it's used, either it defines a sequence
// of values used over the course of an animation, or it's a list of numbers for a color matrix, which is interpreted
// differently depending on the type of color change to be performed.
//
//	Input:
//	  value: list of values
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1),rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s, 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    any other type: interface{}
//
// Português:
//
// O atributo values tem significados diferentes, dependendo do contexto em que é usado, ou define uma sequência de
// valores usados ao longo de uma animação, ou é uma lista de números para uma matriz de cores, que é interpretada de
// forma diferente dependendo do tipo de mudança de cor a ser executada.
//
//	Input:
//	  value: lista de valores
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1),rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s, 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    any other type: interface{}
func (e *TagSvgGlobal) Values(value interface{}) (ref *TagSvgGlobal) {
	e.selfElement.Call("setAttribute", "values", TypeToString(value, ";", ";"))
	return e
}

// VectorEffect #presentation
//
// English:
//
// The vector-effect property specifies the vector effect to use when drawing an object.
//
//	Input:
//	  value: specifies the vector effect
//	    const: KSvgVectorEffect... (e.g. KSvgVectorEffectNonScalingStroke)
//
// Vector effects are applied before any of the other compositing operations, i.e. filters, masks and clips.
//
//	Notes:
//	  * As a presentation attribute, vector-effect can be used as a CSS property.
//
// Português:
//
// A propriedade vector-effect especifica o efeito vetorial a ser usado ao desenhar um objeto.
//
//	Entrada:
//	  value: especifica o efeito vetorial
//	    const: KSvgVectorEffect... (ex. KSvgVectorEffectNonScalingStroke)
//
// Os efeitos vetoriais são aplicados antes de qualquer outra operação de composição, ou seja, filtros, máscaras e
// clipes.
//
//	Notas:
//	  * Como atributo de apresentação, o efeito vetorial pode ser usado como uma propriedade CSS.
func (e *TagSvgGlobal) VectorEffect(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgVectorEffect); ok {
		e.selfElement.Call("setAttribute", "vector-effect", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "vector-effect", value)
	return e
}

// ViewBox
//
// English:
//
// The viewBox attribute defines the position and dimension, in user space, of an SVG viewport.
//
//	Input:
//	  value: defines the position and dimension, in user space, of an SVG viewport
//	    []float64: ex. []float64{0.0, 0.0, 10.0, 10.0} = "0 0 10 10"
//	    any other type: interface{}
//
// The value of the viewBox attribute is a list of four numbers: min-x, min-y, width and height.
// The numbers, which are separated by whitespace and/or a comma, specify a rectangle in user space which is mapped to
// the bounds of the viewport established for the associated SVG element (not the browser viewport).
//
// Português:
//
// O atributo viewBox define a posição e a dimensão, no espaço do usuário, de uma viewport SVG.
//
//	Input:
//	  value: define a posição e dimensão, no espaço do usuário, de uma viewport SVG
//	    []float64: ex. []float64{0.0, 0.0, 10.0, 10.0} = "0 0 10 10"
//	    qualquer outro tipo: interface{}
//
// O valor do atributo viewBox é uma lista de quatro números: min-x, min-y, largura e altura.
// Os números, que são separados por espaço em branco e ou vírgula, especificam um retângulo no espaço do usuário que é
// mapeado para os limites da janela de visualização estabelecida para o elemento SVG associado (não a janela de
// visualização do navegador).
func (e *TagSvgGlobal) ViewBox(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "viewBox", valueStr[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "viewBox", value)
	return e
}

// Visibility #presentation
//
// English:
//
// The visibility attribute lets you control the visibility of graphical elements.
//
//	Input:
//	  value: lets you control the visibility
//	    const: KSvgVisibility... (e.g. KSvgVisibilityHidden)
//	    any other type: interface{}
//
// With a value of hidden or collapse the current graphics element is invisible.
//
// Depending on the value of attribute pointer-events, graphics elements which have their visibility attribute set to
// hidden still might receive events.
//
//	Notes:
//	  * If the visibility attribute is set to hidden on a text element, then the text is invisible but still takes up
//	    space in text layout calculations;
//	  * As a presentation attribute, visibility can be used as a CSS property. See the css visibility property for
//	    more information.
//
// Português:
//
// O atributo de visibilidade permite controlar a visibilidade dos elementos gráficos.
//
//	Entrada:
//	  value: permite controlar a visibilidade
//	    const: KSvgVisibility... (e.g. KSvgVisibilityHidden)
//	    qualquer outro tipo: interface{}
//
// Com um valor oculto ou recolhido, o elemento gráfico atual fica invisível.
//
// Dependendo do valor do atributo pointer-events, os elementos gráficos que têm seu atributo de visibilidade definido
// como oculto ainda podem receber eventos.
//
//	Notas:
//	  * Se o atributo de visibilidade estiver definido como oculto em um elemento de texto, o texto ficará invisível,
//	    mas ainda ocupará espaço nos cálculos de layout de texto;
//	  * Como atributo de apresentação, a visibilidade pode ser usada como propriedade CSS. Consulte a propriedade de
//	    visibilidade do CSS para obter mais informações.
func (e *TagSvgGlobal) Visibility(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgVisibility); ok {
		e.selfElement.Call("setAttribute", "visibility", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "visibility", value)
	return e
}

// Width
//
// English:
//
// The width attribute defines the horizontal length of an element in the user coordinate system.
//
//	Input:
//	  value: the horizontal length of an element
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo largura define o comprimento horizontal de um elemento no sistema de coordenadas do usuário.
//
//	Entrada:
//	  value: o comprimento horizontal de um elemento
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Width(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "width", p)
		return e
	}

	e.selfElement.Call("setAttribute", "width", value)
	return e
}

// WordSpacing #presentation
//
// English:
//
// The word-spacing attribute specifies spacing behavior between words.
//
//	Input:
//	  value: specifies spacing behavior between words
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// If a <length> is provided without a unit identifier (e.g. an unqualified number such as 128), the browser processes
// the <length> as a width value in the current user coordinate system.
//
// If a <length> is provided with one of the unit identifiers (e.g. .25em or 1%), then the browser converts the <length>
// into a corresponding value in the current user coordinate system.
//
//	Notes:
//	  * As a presentation attribute, word-spacing can be used as a CSS property. See the css word-spacing property for
//	    more information.
//
// Português:
//
// O atributo word-spacing especifica o comportamento do espaçamento entre as palavras.
//
//	Entrada:
//	  value: especifica o comportamento de espaçamento entre palavras
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// Se um <comprimento> for fornecido sem um identificador de unidade (por exemplo, um número não qualificado como 128),
// o navegador processará o <comprimento> como um valor de largura no sistema de coordenadas do usuário atual.
//
// Se um <comprimento> for fornecido com um dos identificadores de unidade (por exemplo, .25em ou 1%), o navegador
// converterá o <comprimento> em um valor correspondente no sistema de coordenadas do usuário atual.
//
//	Notas:
//	  * Como atributo de apresentação, o espaçamento entre palavras pode ser usado como uma propriedade CSS.
//	    Consulte a propriedade de espaçamento entre palavras do CSS para obter mais informações.
func (e *TagSvgGlobal) WordSpacing(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "word-spacing", p)
		return e
	}

	e.selfElement.Call("setAttribute", "word-spacing", value)
	return e
}

// WritingMode #presentation
//
// English:
//
// The writing-mode attribute specifies whether the initial inline-progression-direction for a <text> element shall be
// left-to-right, right-to-left, or top-to-bottom. The writing-mode attribute applies only to <text> elements;
// the attribute is ignored for <tspan>, <tref>, <altGlyph> and <textPath> sub-elements. (Note that the
// inline-progression-direction can change within a <text> element due to the Unicode bidirectional algorithm and
// properties direction and unicode-bidi.)
//
//	Input:
//	  value: specifies whether the initial inline-progression-direction
//	    const: KSvgWritingMode... (e.g. KSvgWritingModeHorizontalTb)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, writing-mode can be used as a CSS property. See the CSS writing-mode property for
//	    more information.
//
// Português:
//
// O atributo write-mode especifica se a direção de progressão inline inicial para um elemento <text> deve ser da
// esquerda para a direita, da direita para a esquerda ou de cima para baixo. O atributo write-mode aplica-se apenas a
// elementos <text>; o atributo é ignorado para os subelementos <tspan>, <tref>, <altGlyph> e <textPath>.
// (Observe que a direção de progressão em linha pode mudar dentro de um elemento <text> devido ao algoritmo
// bidirecional Unicode e direção de propriedades e unicode-bidi.)
//
//	Entrada:
//	  value: especifica se a direção de progressão em linha inicial
//	    const: KSvgWritingMode... (ex. KSvgWritingModeHorizontalTb)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o modo de escrita pode ser usado como uma propriedade CSS. Consulte a
//	    propriedade do modo de gravação CSS para obter mais informações.
func (e *TagSvgGlobal) WritingMode(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgWritingMode); ok {
		e.selfElement.Call("setAttribute", "writing-mode", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "writing-mode", value)
	return e
}

// X
//
// English:
//
// The x attribute defines an x-axis coordinate in the user coordinate system.
//
//	Input:
//	  value: defines an x-axis coordinate
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    float32: 0.1 = "10%"
//	    any other type: interface{}
//
// Português:
//
// O atributo x define uma coordenada do eixo x no sistema de coordenadas do usuário.
//
//	Entrada:
//	  value: define uma coordenada do eixo x
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []float32: []float64{0.0, 0.1} = "0%, 10%"
//	    float32: 0.1 = "10%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) X(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + ", "
		}

		var length = len(valueStr) - 2

		e.selfElement.Call("setAttribute", "x", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float32); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "%, "
		}

		var length = len(valueStr) - 2

		e.selfElement.Call("setAttribute", "x", valueStr[:length])
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "x", p)
		return e
	}

	e.selfElement.Call("setAttribute", "x", value)
	return e
}

// X1
//
// English:
//
// The x1 attribute is used to specify the first x-coordinate for drawing an SVG element that requires more than one
// coordinate.
//
//	Input:
//	  value: specify the first x-coordinate
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Elements that only need one coordinate use the x attribute instead.
//
// Português:
//
// O atributo x1 é usado para especificar a primeira coordenada x para desenhar um elemento SVG que requer mais de uma
// coordenada.
//
//	Input:
//	  value: especifique a primeira coordenada x
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// Elementos que precisam apenas de uma coordenada usam o atributo x.
func (e *TagSvgGlobal) X1(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "x1", p)
		return e
	}

	e.selfElement.Call("setAttribute", "x1", value)
	return e
}

// X2
//
// English:
//
// The x2 attribute is used to specify the second x-coordinate for drawing an SVG element that requires more than one
// coordinate. Elements that only need one coordinate use the x attribute instead.
//
//	Input:
//	  value: specify the second x-coordinate
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo x2 é usado para especificar a segunda coordenada x para desenhar um elemento SVG que requer mais de uma
// coordenada. Elementos que precisam apenas de uma coordenada usam o atributo x.
//
//	Entrada:
//	  value: especifique a segunda coordenada x
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) X2(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "x2", p)
		return e
	}

	e.selfElement.Call("setAttribute", "x2", value)
	return e
}

// XChannelSelector
//
// English:
//
// The xChannelSelector attribute indicates which color channel from in2 to use to displace the pixels in in along the
// x-axis.
//
//	Input:
//	  value: indicates which color channel from in2
//	    const: KSvgChannelSelector... (e.g. KSvgChannelSelectorR)
//	    any other type: interface{}
//
// Português:
//
// O atributo xChannelSelector indica qual canal de cor de in2 usar para deslocar os pixels ao longo do eixo x.
//
//	Entrada:
//	  value: indica qual canal de cor da in2
//	    const: KSvgChannelSelector... (ex. KSvgChannelSelectorR)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) XChannelSelector(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgChannelSelector); ok {
		e.selfElement.Call("setAttribute", "xChannelSelector", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xChannelSelector", value)
	return e
}

// XmlLang #core
//
// English:
//
// The xml:lang attribute specifies the primary language used in contents and attributes containing text content of
// particular elements.
//
//	Input:
//	  value: specifies the primary language
//	    const: KLanguage... (e.g. KLanguageEnglish)
//	    any other type: interface{}
//
// It is a universal attribute allowed in all XML dialects to mark up the natural human language that an element
// contains.
//
// There is also a lang attribute (without namespace). If both of them are defined, the one with namespace is used and
// the one without is ignored.
//
// Português:
//
// O atributo xml:lang especifica o idioma principal usado em conteúdos e atributos que contêm conteúdo de texto de
// elementos específicos.
//
//	Entrada:
//	  value: especifica o idioma principal
//	    const: KLanguage... (e.g. KLanguagePortuguese)
//	    qualquer outro tipo: interface{}
//
// É um atributo universal permitido em todos os dialetos XML para marcar a linguagem humana natural que um elemento
// contém.
//
// Há também um atributo lang (sem namespace). Se ambos estiverem definidos, aquele com namespace será usado e o sem
// namespace será ignorado.
func (e *TagSvgGlobal) XmlLang(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// Y
//
// English:
//
// The y attribute defines an y-axis coordinate in the user coordinate system.
//
//	Input:
//	  value: defines an y-axis coordinate
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []float32: []float32{0.0, 0.1} = "0%, 10%"
//	    float32: 0.1 = "10%"
//	    any other type: interface{}
//
// Português:
//
// O atributo y define uma coordenada do eixo y no sistema de coordenadas do usuário.
//
//	Entrada:
//	  value: define uma coordenada do eixo y
//	    []float64: []float64{0.0, 10.0} = "0, 10"
//	    []float32: []float32{0.0, 0.1} = "0%, 10%"
//	    float32: 0.1 = "10%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Y(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + ", "
		}

		var length = len(valueStr) - 2

		e.selfElement.Call("setAttribute", "y", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float32); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "%, "
		}

		var length = len(valueStr) - 2

		e.selfElement.Call("setAttribute", "y", valueStr[:length])
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "y", p)
		return e
	}

	e.selfElement.Call("setAttribute", "y", value)
	return e
}

// Y1
//
// English:
//
// The y1 attribute is used to specify the first y-coordinate for drawing an SVG element that requires more than one
// coordinate.
//
//	Input:
//	  value: specify the first y-coordinate
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Elements that only need one coordinate use the y attribute instead.
//
// Português:
//
// O atributo y1 é usado para especificar a primeira coordenada y para desenhar um elemento SVG que requer mais de uma
// coordenada.
//
//	Input:
//	  value: especifique a primeira coordenada y
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// Elementos que precisam apenas de uma coordenada usam o atributo y.
func (e *TagSvgGlobal) Y1(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "y1", p)
		return e
	}

	e.selfElement.Call("setAttribute", "y1", value)
	return e
}

// Y2
//
// English:
//
// The y2 attribute is used to specify the second y-coordinate for drawing an SVG element that requires more than one
// coordinate.
//
//	Input:
//	  value: specify the second x-coordinate
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Elements that only need one coordinate use the x attribute instead.
//
// Português:
//
// O atributo y2 é usado para especificar a segunda coordenada y para desenhar um elemento SVG que requer mais de uma
// coordenada.
//
//	Entrada:
//	  value: especifique a segunda coordenada x
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// Elementos que precisam apenas de uma coordenada usam o atributo y.
func (e *TagSvgGlobal) Y2(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "y2", p)
		return e
	}

	e.selfElement.Call("setAttribute", "y2", value)
	return e
}

// YChannelSelector
//
// English:
//
// The yChannelSelector attribute indicates which color channel from in2 to use to displace the pixels in along the
// y-axis.
//
//	Input:
//	  value: indicates which color channel from in2
//	    const: KSvgChannelSelector... (e.g. KSvgChannelSelectorR)
//	    any other type: interface{}
//
// Português:
//
// O atributo yChannelSelector indica qual canal de cor de in2 usar para deslocar os pixels ao longo do eixo y.
//
//	Entrada:
//	  value: indica qual canal de cor da in2
//	    const: KSvgChannelSelector... (ex. KSvgChannelSelectorR)
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) YChannelSelector(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.(SvgChannelSelector); ok {
		e.selfElement.Call("setAttribute", "yChannelSelector", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "yChannelSelector", value)
	return e
}

// Z
//
// English:
//
// The z attribute defines the location along the z-axis for a light source in the coordinate system established by the
// primitiveUnits attribute on the <filter> element, assuming that, in the initial coordinate system, the positive
// z-axis comes out towards the person viewing the content and assuming that one unit along the z-axis equals one unit
// in x and y.
//
//	Input:
//	  value: defines the location along the z-axis
//	    float32: 0.1 = "10%"
//	    any other type: interface{}
//
// Português:
//
// O atributo z define a localização ao longo do eixo z para uma fonte de luz no sistema de coordenadas estabelecido
// pelo atributo primitivoUnits no elemento <filter>, assumindo que, no sistema de coordenadas inicial, o eixo z
// positivo sai em direção à pessoa visualizar o conteúdo e assumir que uma unidade ao longo do eixo z é igual a uma
// unidade em x e y.
//
//	Entrada:
//	  value: define a localização ao longo do eixo z
//	    float32: 0.1 = "10%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgGlobal) Z(value interface{}) (ref *TagSvgGlobal) {
	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + ", "
		}

		var length = len(valueStr) - 2

		e.selfElement.Call("setAttribute", "z", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float32); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "%, "
		}

		var length = len(valueStr) - 2

		e.selfElement.Call("setAttribute", "z", valueStr[:length])
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "z", p)
		return e
	}

	e.selfElement.Call("setAttribute", "z", value)
	return e
}
