package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"syscall/js"
)

// TagSvgStyle
//
// English:
//
// The SVG <style> element allows style sheets to be embedded directly within SVG content.
//
//	Notes:
//	  * SVG's style element has the same attributes as the corresponding element in HTML
//	    (see HTML's <style> element).
//
// Português:
//
// O elemento SVG <style> permite que as folhas de estilo sejam incorporadas diretamente no conteúdo SVG.
//
//	Notas:
//	  * O elemento de estilo SVG tem os mesmos atributos que o elemento correspondente em HTML
//	    (definir elemento HTML <style>).
type TagSvgStyle struct {
	commonEvents commonEvents

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
}

// Init
//
// English:
//
//	Initializes the object correctly.
//
// Português:
//
//	Inicializa o objeto corretamente.
func (e *TagSvgStyle) Init() (ref *TagSvgStyle) {
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
//	The id attribute assigns a unique name to an element.
//
// Portuguese
//
//	O atributo id atribui um nome exclusivo a um elemento.
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
func (e *TagSvgStyle) Tabindex(value interface{}) (ref *TagSvgStyle) {
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
//	Input:
//	  value: content type of the element
//	    any other type: interface{}
//
// Português:
//
// Define o tipo de conteúdo do elemento.
//
//	Input:
//	  value: tipo de conteúdo do elemento.
//	    qualquer outro tipo: interface{}
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

// Reference
//
// English:
//
// Pass the object reference to an external variable.
//
// Português:
//
// Passa a referencia do objeto para uma variável externa.
//
//	Example: / Exemplo:
//	  var circle *html.TagSvgCircle
//	  factoryBrowser.NewTagSvgCircle().Reference(&circle).R(5).Fill(factoryColor.NewRed())
//	  log.Printf("x: %v, y: %v", circle.GetX(), circle.GetY())
func (e *TagSvgStyle) Reference(reference **TagSvgStyle) (ref *TagSvgStyle) {
	*reference = e
	return e
}
func (e *TagSvgStyle) AddListenerAbort(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerAbort() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagSvgStyle) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerAuxclick() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagSvgStyle) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerBeforeinput() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagSvgStyle) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerBeforematch() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagSvgStyle) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerBeforetoggle() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagSvgStyle) AddListenerCancel(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerCancel() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagSvgStyle) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerCanplay() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagSvgStyle) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerCanplaythrough() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagSvgStyle) AddListenerChange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerChange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagSvgStyle) AddListenerClick(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerClick() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagSvgStyle) AddListenerClose(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerClose() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagSvgStyle) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerContextlost() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagSvgStyle) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerContextmenu() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagSvgStyle) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerContextrestored() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagSvgStyle) AddListenerCopy(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerCopy() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagSvgStyle) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerCuechange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagSvgStyle) AddListenerCut(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerCut() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagSvgStyle) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDblclick() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagSvgStyle) AddListenerDrag(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDrag() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagSvgStyle) AddListenerDragend(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDragend() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagSvgStyle) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDragenter() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagSvgStyle) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDragleave() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagSvgStyle) AddListenerDragover(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDragover() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagSvgStyle) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDragstart() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagSvgStyle) AddListenerDrop(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDrop() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagSvgStyle) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerDurationchange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagSvgStyle) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerEmptied() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagSvgStyle) AddListenerEnded(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerEnded() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagSvgStyle) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerFormdata() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagSvgStyle) AddListenerInput(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerInput() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagSvgStyle) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerInvalid() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagSvgStyle) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerKeydown() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagSvgStyle) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerKeypress() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagSvgStyle) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerKeyup() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagSvgStyle) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerLoadeddata() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagSvgStyle) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerLoadedmetadata() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagSvgStyle) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerLoadstart() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagSvgStyle) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMousedown() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagSvgStyle) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMouseenter() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagSvgStyle) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMouseleave() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagSvgStyle) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMousemove() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagSvgStyle) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMouseout() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagSvgStyle) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMouseover() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagSvgStyle) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMouseup() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagSvgStyle) AddListenerPaste(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPaste() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagSvgStyle) AddListenerPause(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPause() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagSvgStyle) AddListenerPlay(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPlay() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagSvgStyle) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPlaying() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagSvgStyle) AddListenerProgress(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerProgress() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagSvgStyle) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerRatechange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagSvgStyle) AddListenerReset(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerReset() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagSvgStyle) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerScrollend() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagSvgStyle) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerSecuritypolicyviolation() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagSvgStyle) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerSeeked() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagSvgStyle) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerSeeking() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagSvgStyle) AddListenerSelect(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerSelect() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagSvgStyle) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerSlotchange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagSvgStyle) AddListenerStalled(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerStalled() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagSvgStyle) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerSubmit() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagSvgStyle) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerSuspend() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagSvgStyle) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerTimeupdate() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagSvgStyle) AddListenerToggle(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerToggle() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagSvgStyle) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerVolumechange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagSvgStyle) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerWaiting() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagSvgStyle) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerWebkitanimationend() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagSvgStyle) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerWebkitanimationiteration() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagSvgStyle) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerWebkitanimationstart() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagSvgStyle) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerWebkittransitionend() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagSvgStyle) AddListenerWheel(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerWheel() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagSvgStyle) AddListenerBlur(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerBlur() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagSvgStyle) AddListenerError(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerError() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagSvgStyle) AddListenerFocus(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerFocus() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagSvgStyle) AddListenerLoad(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerLoad() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagSvgStyle) AddListenerResize(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerResize() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagSvgStyle) AddListenerScroll(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerScroll() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagSvgStyle) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerAfterprint() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagSvgStyle) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerBeforeprint() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagSvgStyle) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerBeforeunload() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagSvgStyle) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerHashchange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagSvgStyle) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerLanguagechange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagSvgStyle) AddListenerMessage(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMessage() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagSvgStyle) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerMessageerror() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagSvgStyle) AddListenerOffline(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerOffline() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagSvgStyle) AddListenerOnline(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerOnline() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagSvgStyle) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPageswap() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagSvgStyle) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPagehide() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagSvgStyle) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPagereveal() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagSvgStyle) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPageshow() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagSvgStyle) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerPopstate() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagSvgStyle) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerRejectionhandled() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagSvgStyle) AddListenerStorage(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerStorage() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagSvgStyle) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerUnhandledrejection() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagSvgStyle) AddListenerUnload(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerUnload() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagSvgStyle) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerReadystatechange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagSvgStyle) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagSvgStyle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagSvgStyle) RemoveListenerVisibilitychange() (ref *TagSvgStyle) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
