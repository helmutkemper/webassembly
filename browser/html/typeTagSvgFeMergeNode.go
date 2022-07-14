package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"sync"
	"syscall/js"
)

// TagSvgFeMergeNode
//
// English:
//
// The feMergeNode takes the result of another filter to be processed by its parent <feMerge>.
//
// Português:
//
// O feMergeNode recebe o resultado de outro filtro para ser processado por seu pai <feMerge>.
type TagSvgFeMergeNode struct {

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
func (e *TagSvgFeMergeNode) Init() (ref *TagSvgFeMergeNode) {
	e.listener = new(sync.Map)

	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgFeMergeNode) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgFeMergeNode) CreateElement() (ref *TagSvgFeMergeNode) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "feMergeNode")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgFeMergeNode) AppendToStage() (ref *TagSvgFeMergeNode) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeMergeNode) AppendById(appendId string) (ref *TagSvgFeMergeNode) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeMergeNode) AppendToElement(el js.Value) (ref *TagSvgFeMergeNode) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgFeMergeNode) Append(elements ...Compatible) (ref *TagSvgFeMergeNode) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgFeMergeNode) Get() (el js.Value) {
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
func (e *TagSvgFeMergeNode) Id(id string) (ref *TagSvgFeMergeNode) {
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
func (e *TagSvgFeMergeNode) Lang(value interface{}) (ref *TagSvgFeMergeNode) {

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
func (e *TagSvgFeMergeNode) Tabindex(value int) (ref *TagSvgFeMergeNode) {
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
func (e *TagSvgFeMergeNode) XmlLang(value interface{}) (ref *TagSvgFeMergeNode) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

// In
//
// English:
//
//  The in attribute identifies input for the given filter primitive.
//
//   Input:
//     in: identifies input for the given filter primitive.
//       KSvgIn... (e.g. KSvgInSourceAlpha)
//       any other type: interface{}
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
//       qualquer outro tipo: interface{}
//
// O valor pode ser uma das seis palavras-chave definidas abaixo ou uma string que corresponda a um valor de atributo
// de resultado anterior dentro do mesmo elemento <filter>. Se nenhum valor for fornecido e esta for a primeira
// primitiva de filtro, essa primitiva de filtro usará SourceGraphic como sua entrada. Se nenhum valor for fornecido e
// esta for uma primitiva de filtro subsequente, essa primitiva de filtro usará o resultado da primitiva de filtro
// anterior como sua entrada.
//
// Se o valor do resultado aparecer várias vezes em um determinado elemento <filter>, uma referência à esse resultado
// usará a primitiva de filtro anterior mais próxima com o valor fornecido para o resultado do atributo.
func (e *TagSvgFeMergeNode) In(in interface{}) (ref *TagSvgFeMergeNode) {
	if converted, ok := in.(SvgIn); ok {
		e.selfElement.Call("setAttribute", "in", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "in", in)
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
func (e *TagSvgFeMergeNode) Text(value string) (ref *TagSvgFeMergeNode) {
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
func (e *TagSvgFeMergeNode) Html(value string) (ref *TagSvgFeMergeNode) {
	e.selfElement.Set("innerHTML", value)
	return e
}
