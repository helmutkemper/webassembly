package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"strconv"
	"sync"
	"syscall/js"
	"time"
)

// TagSvgFeConvolveMatrix
//
// English:
//
// The <feConvolveMatrix> SVG filter primitive applies a matrix convolution filter effect. A convolution combines pixels
// in the input image with neighboring pixels to produce a resulting image. A wide variety of imaging operations can be
// achieved through convolutions, including blurring, edge detection, sharpening, embossing and beveling.
//
// A matrix convolution is based on an n-by-m matrix (the convolution kernel) which describes how a given pixel value
// in the input image is combined with its neighboring pixel values to produce a resulting pixel value. Each result
// pixel is determined by applying the kernel matrix to the corresponding source pixel and its neighboring pixels.
//
// The basic convolution formula which is applied to each color value for a given pixel is:
//
//   COLORX,Y = ( SUM I=0 to [orderY-1] { SUM J=0 to [orderX-1] { SOURCE X-targetX+J, Y-targetY+I *
//                kernelMatrixorderX-J-1, orderY-I-1 } } ) / divisor + bias * ALPHAX,Y
//
// where "orderX" and "orderY" represent the X and Y values for the 'order' attribute, "targetX" represents the value of
// the 'targetX' attribute, "targetY" represents the value of the 'targetY' attribute, "kernelMatrix" represents the
// value of the 'kernelMatrix' attribute, "divisor" represents the value of the 'divisor' attribute, and "bias"
// represents the value of the 'bias' attribute.
//
// Note in the above formulas that the values in the kernel matrix are applied such that the kernel matrix is rotated
// 180 degrees relative to the source and destination images in order to match convolution theory as described in many
// computer graphics textbooks.
//
// To illustrate, suppose you have a input image which is 5 pixels by 5 pixels, whose color values for one of the color
// channels are as follows:
//
//   0    20  40 235 235
//   100 120 140 235 235
//   200 220 240 235 235
//   225 225 255 255 255
//   225 225 255 255 255
//
// and you define a 3-by-3 convolution kernel as follows:
//
//   1 2 3
//   4 5 6
//   7 8 9
//
// Let's focus on the color value at the second row and second column of the image (source pixel value is 120). Assuming
// the simplest case (where the input image's pixel grid aligns perfectly with the kernel's pixel grid) and assuming
// default values for attributes 'divisor', 'targetX' and 'targetY', then resulting color value will be:
//
//   (9*  0 + 8* 20 + 7* 40 +
//    6*100 + 5*120 + 4*140 +
//    3*200 + 2*220 + 1*240) /
//   (9+8+7+6+5+4+3+2+1)
//
// Português:
//
// A primitiva de filtro SVG <feConvolveMatrix> aplica um efeito de filtro de convolução de matriz. Uma convolução
// combina pixels na imagem de entrada com pixels vizinhos para produzir uma imagem resultante. Uma ampla variedade de
// operações de imagem pode ser alcançada por meio de convoluções, incluindo desfoque, detecção de bordas, nitidez,
// relevo e chanfro.
//
// Uma convolução de matriz é baseada em uma matriz n por m (o kernel de convolução) que descreve como um determinado
// valor de pixel na imagem de entrada é combinado com seus valores de pixel vizinhos para produzir um valor de pixel
// resultante. Cada pixel resultante é determinado pela aplicação da matriz kernel ao pixel de origem correspondente e
// seus pixels vizinhos. A fórmula de convolução básica que é aplicada a cada valor de cor para um determinado pixel é:
//
//   COLORX,Y = ( SUM I=0 to [orderY-1] { SUM J=0 to [orderX-1] { SOURCE X-targetX+J, Y-targetY+I *
//                kernelMatrixorderX-J-1, orderY-I-1 } } ) / divisor + bias * ALPHAX,Y
//
// onde "orderX" e "orderY" representam os valores X e Y para o atributo 'order', "targetX" representa o valor do
// atributo 'targetX', "targetY" representa o valor do atributo 'targetY', "kernelMatrix" representa o valor do atributo
// 'kernelMatrix', "divisor" representa o valor do atributo 'divisor' e "bias" representa o valor do atributo 'bias'.
//
// Observe nas fórmulas acima que os valores na matriz do kernel são aplicados de tal forma que a matriz do kernel é
// girada 180 graus em relação às imagens de origem e destino para corresponder à teoria de convolução conforme descrito
// em muitos livros de computação gráfica.
//
// Para ilustrar, suponha que você tenha uma imagem de entrada com 5 pixels por 5 pixels, cujos valores de cor para um
// dos canais de cores sejam os seguintes:
//
//   0    20  40 235 235
//   100 120 140 235 235
//   200 220 240 235 235
//   225 225 255 255 255
//   225 225 255 255 255
//
// e você define um kernel de convolução 3 por 3 da seguinte forma:
//
//   1 2 3
//   4 5 6
//   7 8 9
//
// Vamos nos concentrar no valor da cor na segunda linha e na segunda coluna da imagem (o valor do pixel de origem é
// 120). Assumindo o caso mais simples (onde a grade de pixels da imagem de entrada se alinha perfeitamente com a grade
// de pixels do kernel) e assumindo valores padrão para os atributos 'divisor', 'targetX' e 'targetY', o valor da cor
// resultante será:
//
//   (9*  0 + 8* 20 + 7* 40 +
//    6*100 + 5*120 + 4*140 +
//    3*200 + 2*220 + 1*240) /
//   (9+8+7+6+5+4+3+2+1)
type TagSvgFeConvolveMatrix struct {

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
func (e *TagSvgFeConvolveMatrix) Begin(begin interface{}) (ref *TagSvgFeConvolveMatrix) {
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
func (e *TagSvgFeConvolveMatrix) Bias(bias float64) (ref *TagSvgFeConvolveMatrix) {
	e.selfElement.Call("setAttribute", "bias", bias)
	return e
}

// Init
//
// English:
//
//  Initializes the object correctly.
//
// Português:
//
//  Inicializa o objeto corretamente.
func (e *TagSvgFeConvolveMatrix) Init(id string) (ref *TagSvgFeConvolveMatrix) {
	e.listener = new(sync.Map)

	e.CreateElement(KTagSvgFeConvolveMatrix)
	e.prepareStageReference()
	e.Id(id)

	return e
}

func (e *TagSvgFeConvolveMatrix) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgFeConvolveMatrix) Id(id string) (ref *TagSvgFeConvolveMatrix) {
	e.id = id
	e.selfElement.Set("id", id)
	return e
}

func (e *TagSvgFeConvolveMatrix) CreateElement(tag Tag) (ref *TagSvgFeConvolveMatrix) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	return e
}

// KernelMatrix
//
// English:
//
// The kernelMatrix attribute defines the list of numbers that make up the kernel matrix for the <feConvolveMatrix>
// element.
//
// Values are separated by space characters and/or a comma. The number of entries in the list must equal to <orderX>
// by <orderY> as defined in the order attribute.
//
// Português:
//
// O atributo kernelMatrix define a lista de números que compõem a matriz do kernel para o elemento <feConvolveMatrix>.
//
// Os valores são separados por caracteres de espaço e ou por vírgula. O número de entradas na lista deve ser igual a
// <orderX> por <orderY> conforme definido no atributo order.
func (e *TagSvgFeConvolveMatrix) KernelMatrix(kernelMatrix []float64) (ref *TagSvgFeConvolveMatrix) {
	kernelMatrixString := ""

	for _, v := range kernelMatrix {
		kernelMatrixString += strconv.FormatFloat(v, 'g', -1, 64)
		kernelMatrixString += " "
	}

	l := len(kernelMatrixString) - 1

	e.selfElement.Call("setAttribute", "kernelMatrix", kernelMatrixString[:l])
	return e
}

func (e *TagSvgFeConvolveMatrix) AppendToStage() (ref *TagSvgFeConvolveMatrix) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeConvolveMatrix) AppendById(appendId string) (ref *TagSvgFeConvolveMatrix) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}
