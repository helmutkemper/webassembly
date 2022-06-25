package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"strconv"
	"sync"
	"syscall/js"
)

// TagSvgFeFuncA
//
// English:
//
// The <feFuncA> SVG filter primitive defines the transfer function for the alpha component of the input graphic of its
// parent <feComponentTransfer> element.
//
// Português:
//
// A primitiva de filtro SVG <feFuncA> define a função de transferência para o componente alfa do gráfico de entrada
// de seu elemento pai <feComponentTransfer>.
type TagSvgFeFuncA struct {

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

// Init
//
// English:
//
//  Initializes the object correctly.
//
// Português:
//
//  Inicializa o objeto corretamente.
func (e *TagSvgFeFuncA) Init(id string) (ref *TagSvgFeFuncA) {
	e.listener = new(sync.Map)

	e.CreateElement()
	e.prepareStageReference()
	e.Id(id)

	return e
}

func (e *TagSvgFeFuncA) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgFeFuncA) CreateElement() (ref *TagSvgFeFuncA) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "feFuncA")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgFeFuncA) AppendToStage() (ref *TagSvgFeFuncA) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeFuncA) AppendById(appendId string) (ref *TagSvgFeFuncA) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeFuncA) AppendToElement(el js.Value) (ref *TagSvgFeFuncA) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgFeFuncA) Get() (el js.Value) {
	return e.selfElement
}

// #core start --------------------------------------------------------------------------------------------------------

// Id
//
// English:
//
//  The id attribute assigns a unique name to an element.
//
// Portuguese
//
//  O atributo id atribui um nome exclusivo a um elemento.
func (e *TagSvgFeFuncA) Id(id string) (ref *TagSvgFeFuncA) {
	e.selfElement.Call("setAttribute", "id", id)
	return e
}

// Lang
//
// English:
//
// The lang attribute specifies the primary language used in contents and attributes containing text content of
// particular elements.
//
//   Input:
//     value: specifies the primary language used in contents
//       const KLanguage... (e.g. KLanguageEnglish)
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
//   Entrada:
//     value: especifica o idioma principal usado no conteúdo
//       const KLanguage... (ex. KLanguagePortuguese)
//
// Há também um atributo xml:lang (com namespace). Se ambos estiverem definidos, aquele com namespace será usado e o
// sem namespace será ignorado.
//
// No SVG 1.1 havia um atributo lang definido com um significado diferente e aplicando-se apenas aos elementos <glyph>.
// Esse atributo especificou uma lista de idiomas de acordo com a RFC 5646: Tags for Identification Languages
// (também conhecido como BCP 47). O glifo deveria ser usado se o atributo xml:lang correspondesse exatamente a um dos
// idiomas fornecidos no valor desse parâmetro, ou se o atributo xml:lang fosse exatamente igual a um prefixo de um dos
// idiomas fornecidos no valor desse parâmetro de modo que o primeiro caractere de tag após o prefixo fosse "-".
func (e *TagSvgFeFuncA) Lang(value interface{}) (ref *TagSvgFeFuncA) {

	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "lang", value)
	return e
}

// Tabindex
//
// English:
//
// The tabindex attribute allows you to control whether an element is focusable and to define the relative order of the
// element for the purposes of sequential focus navigation.
//
// Português:
//
// O atributo tabindex permite controlar se um elemento é focalizável e definir a ordem relativa do elemento para fins
// de navegação de foco sequencial.
func (e *TagSvgFeFuncA) Tabindex(value int) (ref *TagSvgFeFuncA) {
	e.selfElement.Call("setAttribute", "tabindex", value)
	return e
}

// XmlLang
//
// English:
//
// The xml:lang attribute specifies the primary language used in contents and attributes containing text content of
// particular elements.
//
//   Input:
//     value: specifies the primary language
//       const: KLanguage... (e.g. KLanguageEnglish)
//       any other type: interface{}
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
//   Entrada:
//     value: especifica o idioma principal
//       const: KLanguage... (e.g. KLanguagePortuguese)
//       qualquer outro tipo: interface{}
//
// É um atributo universal permitido em todos os dialetos XML para marcar a linguagem humana natural que um elemento
// contém.
//
// Há também um atributo lang (sem namespace). Se ambos estiverem definidos, aquele com namespace será usado e o sem
// namespace será ignorado.
func (e *TagSvgFeFuncA) XmlLang(value interface{}) (ref *TagSvgFeFuncA) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

// Type
//
// English:
//
// Indicates the type of component transfer function.
//
//   Input:
//     value: type of component transfer function
//       const: KSvgTypeFeFunc... (e.g. KSvgTypeFeFuncIdentity)
//       any other type: interface{}
//
// Português:
//
// Indica o tipo de função de transferência de componentes.
//
//   Input:
//     value: tipo de função de transferência de componente
//       const: KSvgTypeFeFunc... (ex. KSvgTypeFeFuncIdentity)
//       any other type: interface{}
func (e *TagSvgFeFuncA) Type(value interface{}) (ref *TagSvgFeFuncA) {
	if converted, ok := value.(SvgTypeFeFunc); ok {
		e.selfElement.Call("setAttribute", "type", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "type", value)
	return e
}

// TableValues
//
// English:
//
// The tableValues attribute defines a list of numbers defining a lookup table of values for a color component transfer
// function.
//
//   Input:
//     value: defines a list of numbers
//       []float64: e.g. []float64{0.0, 1.0} = "0 1"
//       any other type: interface{}
//
// Português:
//
// O atributo tableValues define uma lista de números que definem uma tabela de consulta de valores para uma função de
// transferência de componente de cor.
//
//   Entrada:
//     value: define uma lista de números
//       []float64: ex. []float64{0.0, 1.0} = "0 1"
//       qualquer outro tipo: interface{}
func (e *TagSvgFeFuncA) TableValues(value interface{}) (ref *TagSvgFeFuncA) {
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
func (e *TagSvgFeFuncA) Intercept(intercept float64) (ref *TagSvgFeFuncA) {
	e.selfElement.Call("setAttribute", "intercept", intercept)
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
//       float32: 1.0 = "100%"
//       any other type: interface{}
//
// Português:
//
//  O atributo amplitude controla à amplitude da função gama de um elemento de transferência de componente quando seu
//  atributo de tipo é gama.
//
//   Entrada:
//     amplitude: controla a amplitude da função de gama
//       float32: 1.0 = "100%"
//       qualquer outro tipo: interface{}
func (e *TagSvgFeFuncA) Amplitude(amplitude interface{}) (ref *TagSvgFeFuncA) {
	if converted, ok := amplitude.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "amplitude", p)
		return e
	}

	e.selfElement.Call("setAttribute", "amplitude", amplitude)
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
func (e *TagSvgFeFuncA) Exponent(exponent float64) (ref *TagSvgFeFuncA) {
	e.selfElement.Call("setAttribute", "exponent", exponent)
	return e
}
