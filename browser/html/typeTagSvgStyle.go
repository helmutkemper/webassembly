package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"sync"
	"syscall/js"
)

// TagSvgStyle
//
// English:
//
// The SVG <style> element allows style sheets to be embedded directly within SVG content.
//
//   Notes:
//     * SVG's style element has the same attributes as the corresponding element in HTML
//       (see HTML's <style> element).
//
// Português:
//
// O elemento SVG <style> permite que as folhas de estilo sejam incorporadas diretamente no conteúdo SVG.
//
//   Notas:
//     * O elemento de estilo SVG tem os mesmos atributos que o elemento correspondente em HTML
//       (definir elemento HTML <style>).
type TagSvgStyle struct {

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
func (e *TagSvgStyle) Init() (ref *TagSvgStyle) {
	e.listener = new(sync.Map)

	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgStyle) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgStyle) CreateElement() (ref *TagSvgStyle) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "style")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgStyle) AppendToStage() (ref *TagSvgStyle) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgStyle) AppendById(appendId string) (ref *TagSvgStyle) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgStyle) AppendToElement(el js.Value) (ref *TagSvgStyle) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgStyle) Append(elements ...Compatible) (ref *TagSvgStyle) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgStyle) Get() (el js.Value) {
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
func (e *TagSvgStyle) Id(id string) (ref *TagSvgStyle) {
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
func (e *TagSvgStyle) Lang(value interface{}) (ref *TagSvgStyle) {

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
func (e *TagSvgStyle) Tabindex(value int) (ref *TagSvgStyle) {
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
func (e *TagSvgStyle) XmlLang(value interface{}) (ref *TagSvgStyle) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

// #styling start -----------------------------------------------------------------------------------------------------

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
func (e *TagSvgStyle) Class(class string) (ref *TagSvgStyle) {
	e.selfElement.Call("setAttribute", "class", class)
	return e
}

// Style
//
// English:
//
// The style attribute allows to style an element using CSS declarations.
//
//   Input:
//     value: allows to style an element using CSS declarations
//
// It functions identically to the style attribute in HTML.
//
// Português:
//
// O atributo style permite estilizar um elemento usando declarações CSS.
//
//   Entrada:
//     value: permite estilizar um elemento usando declarações CSS
//
// Funciona de forma idêntica ao atributo style em HTML.
func (e *TagSvgStyle) Style(value string) (ref *TagSvgStyle) {
	e.selfElement.Set("innerHTML", value)
	return e
}

// #styling end -------------------------------------------------------------------------------------------------------

// Type
//
// English:
//
// Defines the content type of the element
//
//   Input:
//     value: content type of the element
//       any other type: interface{}
//
// Português:
//
// Define o tipo de conteúdo do elemento.
//
//   Input:
//     value: tipo de conteúdo do elemento.
//       qualquer outro tipo: interface{}
func (e *TagSvgStyle) Type(value interface{}) (ref *TagSvgStyle) {
	e.selfElement.Call("setAttribute", "type", value)
	return e
}

// Media
//
// English:
//
// The media attribute specifies a media query that must be matched for a style sheet to apply.
//
//   Input:
//     value: specifies a media query that must be matched for a style sheet to apply
//       string: e.g. "all and (min-width: 600px)"
//
// Português:
//
// O atributo de mídia especifica uma consulta de mídia que deve ser correspondida para que uma folha de estilo seja
// aplicada.
//
//   Entrada:
//     value: especifica uma consulta de mídia que deve ser correspondida para que uma folha de estilo seja aplicada
//       string: e.g. "all and (min-width: 600px)"
func (e *TagSvgStyle) Media(value interface{}) (ref *TagSvgStyle) {
	e.selfElement.Call("setAttribute", "media", value)
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
func (e *TagSvgStyle) Text(value string) (ref *TagSvgStyle) {
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
func (e *TagSvgStyle) Html(value string) (ref *TagSvgStyle) {
	e.selfElement.Set("innerHTML", value)
	return e
}
