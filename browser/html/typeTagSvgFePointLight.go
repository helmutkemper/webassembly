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

// TagSvgFePointLight
//
// English:
//
// The <fePointLight> filter primitive defines a light source which allows to create a point light effect.
// It that can be used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>.
//
// Português:
//
// A primitiva de filtro <fePointLight> define uma fonte de luz que permite criar um efeito de luz pontual.
// Ele que pode ser usado dentro de uma primitiva de filtro de iluminação: <feDiffuseLighting> ou <feSpecularLighting>.
type TagSvgFePointLight struct {

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
func (e *TagSvgFePointLight) Init() (ref *TagSvgFePointLight) {
	e.listener = new(sync.Map)

	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgFePointLight) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgFePointLight) CreateElement() (ref *TagSvgFePointLight) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "fePointLight")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgFePointLight) AppendToStage() (ref *TagSvgFePointLight) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFePointLight) AppendById(appendId string) (ref *TagSvgFePointLight) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFePointLight) AppendToElement(el js.Value) (ref *TagSvgFePointLight) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgFePointLight) Append(elements ...Compatible) (ref *TagSvgFePointLight) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgFePointLight) Get() (el js.Value) {
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
func (e *TagSvgFePointLight) Id(id string) (ref *TagSvgFePointLight) {
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
func (e *TagSvgFePointLight) Lang(value interface{}) (ref *TagSvgFePointLight) {

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
//   Input:
//     value: control whether an element is focusable
//
// Português:
//
// O atributo tabindex permite controlar se um elemento é focalizável e definir à ordem relativa do elemento para fins
// de navegação de foco sequencial.
//
//   Input:
//     value: controlar se um elemento é focalizável
func (e *TagSvgFePointLight) Tabindex(value int) (ref *TagSvgFePointLight) {
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
func (e *TagSvgFePointLight) XmlLang(value interface{}) (ref *TagSvgFePointLight) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

// X
//
// English:
//
// The x attribute defines an x-axis coordinate in the user coordinate system.
//
//   Input:
//     value: defines an x-axis coordinate
//       []float64: []float64{0.0, 10.0} = "0, 10"
//       []float32: []float64{0.0, 0.1} = "0%, 10%"
//       float32: 0.1 = "10%"
//       any other type: interface{}
//
// Português:
//
// O atributo x define uma coordenada do eixo x no sistema de coordenadas do usuário.
//
//   Entrada:
//     value: define uma coordenada do eixo x
//       []float64: []float64{0.0, 10.0} = "0, 10"
//       []float32: []float64{0.0, 0.1} = "0%, 10%"
//       float32: 0.1 = "10%"
//       qualquer outro tipo: interface{}
func (e *TagSvgFePointLight) X(value interface{}) (ref *TagSvgFePointLight) {
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

// Y
//
// English:
//
// The y attribute defines an y-axis coordinate in the user coordinate system.
//
//   Input:
//     value: defines an y-axis coordinate
//       []float64: []float64{0.0, 10.0} = "0, 10"
//       []float32: []float32{0.0, 0.1} = "0%, 10%"
//       float32: 0.1 = "10%"
//       any other type: interface{}
//
// Português:
//
// O atributo y define uma coordenada do eixo y no sistema de coordenadas do usuário.
//
//   Entrada:
//     value: define uma coordenada do eixo y
//       []float64: []float64{0.0, 10.0} = "0, 10"
//       []float32: []float32{0.0, 0.1} = "0%, 10%"
//       float32: 0.1 = "10%"
//       qualquer outro tipo: interface{}
func (e *TagSvgFePointLight) Y(value interface{}) (ref *TagSvgFePointLight) {
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

// Z
//
// English:
//
// The z attribute defines the location along the z-axis for a light source in the coordinate system established by the
// primitiveUnits attribute on the <filter> element, assuming that, in the initial coordinate system, the positive
// z-axis comes out towards the person viewing the content and assuming that one unit along the z-axis equals one unit
// in x and y.
//
//   Input:
//     value: defines the location along the z-axis
//       float32: 0.1 = "10%"
//       any other type: interface{}
//
// Português:
//
// O atributo z define a localização ao longo do eixo z para uma fonte de luz no sistema de coordenadas estabelecido
// pelo atributo primitivoUnits no elemento <filter>, assumindo que, no sistema de coordenadas inicial, o eixo z
// positivo sai em direção à pessoa visualizar o conteúdo e assumir que uma unidade ao longo do eixo z é igual a uma
// unidade em x e y.
//
//   Entrada:
//     value: define a localização ao longo do eixo z
//       float32: 0.1 = "10%"
//       qualquer outro tipo: interface{}
func (e *TagSvgFePointLight) Z(value interface{}) (ref *TagSvgFePointLight) {
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

// Text
//
// English:
//
// Adds plain text to the tag's content.
//
// Text:
//
// Adiciona um texto simples ao conteúdo da tag.
func (e *TagSvgFePointLight) Text(value string) (ref *TagSvgFePointLight) {
	e.selfElement.Set("textContent", value)
	return e
}

// Html
//
// English:
//
// Adds HTML to the tag's content.
//
// Text:
//
// Adiciona HTML ao conteúdo da tag.
func (e *TagSvgFePointLight) Html(value string) (ref *TagSvgFePointLight) {
	e.selfElement.Set("innerHTML", value)
	return e
}
