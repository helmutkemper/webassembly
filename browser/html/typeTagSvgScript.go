package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"syscall/js"
)

// TagSvgScript
//
// English:
//
// The SVG script element allows to add scripts to an SVG document.
//
//	Notes:
//	  * While SVG's script element is equivalent to the HTML <script> element, it has some discrepancies, like it uses
//	    the href attribute instead of src and it doesn't support ECMAScript modules so far (See browser compatibility
//	    below for details)
//	  * document/window.addEventListener("DOMContentLoaded", (e) => {...}); - didn't work in tests (07/2022)
//
// Português:
//
// O elemento script SVG permite adicionar scripts a um documento SVG.
//
//	Notas:
//	  * Embora o elemento script do SVG seja equivalente ao elemento HTML <script>, ele tem algumas discrepâncias, como
//	    usar o atributo href em vez de src e não suportar módulos ECMAScript até agora (consulte a compatibilidade do
//	    navegador abaixo para obter detalhes)
//	  * document/window.addEventListener("DOMContentLoaded", (e) => {...}); - não funcionou nos testes (07/2022)
type TagSvgScript struct {
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
func (e *TagSvgScript) Init() (ref *TagSvgScript) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgScript) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgScript) CreateElement() (ref *TagSvgScript) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "script")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgScript) AppendToStage() (ref *TagSvgScript) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgScript) AppendById(appendId string) (ref *TagSvgScript) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgScript) AppendToElement(el js.Value) (ref *TagSvgScript) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgScript) Append(elements ...Compatible) (ref *TagSvgScript) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgScript) Get() (el js.Value) {
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
func (e *TagSvgScript) Id(id string) (ref *TagSvgScript) {
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
func (e *TagSvgScript) Lang(value interface{}) (ref *TagSvgScript) {

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
func (e *TagSvgScript) Tabindex(value interface{}) (ref *TagSvgScript) {
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
func (e *TagSvgScript) XmlLang(value interface{}) (ref *TagSvgScript) {
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
func (e *TagSvgScript) Class(class string) (ref *TagSvgScript) {
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
func (e *TagSvgScript) Style(value string) (ref *TagSvgScript) {
	e.selfElement.Call("setAttribute", "style", value)
	return e
}

// #styling end -------------------------------------------------------------------------------------------------------

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
func (e *TagSvgScript) CrossOrigin(value interface{}) (ref *TagSvgScript) {
	if converted, ok := value.(SvgCrossOrigin); ok {
		e.selfElement.Call("setAttribute", "crossorigin", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "crossorigin", value)
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
func (e *TagSvgScript) HRef(href interface{}) (ref *TagSvgScript) {
	e.selfElement.Call("setAttribute", "href", href)
	return e
}

// Type
//
// English:
//
// Defines the content type of the element.
//
//	Input:
//	  value: type of the element
//
// Português:
//
// Define o tipo de conteúdo do elemento.
//
//	Input:
//	  value: tipo de conteúdo do elemento
func (e *TagSvgScript) Type(value interface{}) (ref *TagSvgScript) {
	e.selfElement.Call("setAttribute", "type", value)
	return e
}

// XLinkHRef
//
// English:
//
// Deprecated: use HRef() function
//
// The xlink:href attribute defines a reference to a resource as a reference IRI. The exact meaning of that link depends
// on the context of each element using it.
//
//	Notes:
//	  * SVG 2 removed the need for the xlink namespace, so instead of xlink:href you should use href. If you need to
//	    support earlier browser versions, the deprecated xlink:href attribute can be used as a fallback in addition to
//	    the href attribute, e.g. <use href="some-id" xlink:href="some-id" x="5" y="5" />.
//
// Português:
//
// Obsoleto: use a função HRef()
//
// O atributo xlink:href define uma referência a um recurso como um IRI de referência. O significado exato desse link
// depende do contexto de cada elemento que o utiliza.
//
//	Notas:
//	  * O SVG 2 removeu a necessidade do namespace xlink, então ao invés de xlink:href você deve usar href. Se você
//	    precisar oferecer suporte a versões anteriores do navegador, o atributo obsoleto xlink:href pode ser usado como
//	    um substituto além do atributo href, por exemplo, <use href="some-id" xlink:href="some-id" x="5" y="5" >.
func (e *TagSvgScript) XLinkHRef(value interface{}) (ref *TagSvgScript) {
	e.selfElement.Call("setAttribute", "xlink:href", value)
	return e
}

// Script
//
// English:
//
// Adds plain text to the tag's content.
//
//	Notes:
//	  * document/window.addEventListener("DOMContentLoaded", (e) => {...}); - didn't work in tests (07/2022)
//
// Text:
//
// Adiciona um texto simples ao conteúdo da tag.
//
//	Notras:
//	  * document/window.addEventListener("DOMContentLoaded", (e) => {...}); - não funcionou nos testes (07/2022)
func (e *TagSvgScript) Script(value string) (ref *TagSvgScript) {
	e.selfElement.Set("textContent", value)
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
func (e *TagSvgScript) Reference(reference **TagSvgScript) (ref *TagSvgScript) {
	*reference = e
	return e
}
func (e *TagSvgScript) AddListenerAbort(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerAbort() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagSvgScript) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerAuxclick() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagSvgScript) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerBeforeinput() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagSvgScript) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerBeforematch() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagSvgScript) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerBeforetoggle() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagSvgScript) AddListenerCancel(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerCancel() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagSvgScript) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerCanplay() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagSvgScript) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerCanplaythrough() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagSvgScript) AddListenerChange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerChange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagSvgScript) AddListenerClick(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerClick() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagSvgScript) AddListenerClose(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerClose() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagSvgScript) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerContextlost() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagSvgScript) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerContextmenu() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagSvgScript) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerContextrestored() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagSvgScript) AddListenerCopy(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerCopy() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagSvgScript) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerCuechange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagSvgScript) AddListenerCut(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerCut() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagSvgScript) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDblclick() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagSvgScript) AddListenerDrag(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDrag() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagSvgScript) AddListenerDragend(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDragend() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagSvgScript) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDragenter() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagSvgScript) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDragleave() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagSvgScript) AddListenerDragover(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDragover() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagSvgScript) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDragstart() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagSvgScript) AddListenerDrop(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDrop() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagSvgScript) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerDurationchange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagSvgScript) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerEmptied() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagSvgScript) AddListenerEnded(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerEnded() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagSvgScript) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerFormdata() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagSvgScript) AddListenerInput(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerInput() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagSvgScript) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerInvalid() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagSvgScript) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerKeydown() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagSvgScript) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerKeypress() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagSvgScript) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerKeyup() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagSvgScript) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerLoadeddata() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagSvgScript) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerLoadedmetadata() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagSvgScript) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerLoadstart() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagSvgScript) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMousedown() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagSvgScript) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMouseenter() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagSvgScript) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMouseleave() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagSvgScript) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMousemove() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagSvgScript) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMouseout() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagSvgScript) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMouseover() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagSvgScript) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMouseup() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagSvgScript) AddListenerPaste(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPaste() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagSvgScript) AddListenerPause(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPause() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagSvgScript) AddListenerPlay(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPlay() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagSvgScript) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPlaying() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagSvgScript) AddListenerProgress(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerProgress() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagSvgScript) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerRatechange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagSvgScript) AddListenerReset(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerReset() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagSvgScript) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerScrollend() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagSvgScript) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerSecuritypolicyviolation() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagSvgScript) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerSeeked() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagSvgScript) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerSeeking() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagSvgScript) AddListenerSelect(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerSelect() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagSvgScript) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerSlotchange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagSvgScript) AddListenerStalled(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerStalled() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagSvgScript) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerSubmit() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagSvgScript) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerSuspend() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagSvgScript) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerTimeupdate() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagSvgScript) AddListenerToggle(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerToggle() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagSvgScript) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerVolumechange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagSvgScript) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerWaiting() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagSvgScript) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerWebkitanimationend() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagSvgScript) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerWebkitanimationiteration() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagSvgScript) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerWebkitanimationstart() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagSvgScript) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerWebkittransitionend() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagSvgScript) AddListenerWheel(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerWheel() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagSvgScript) AddListenerBlur(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerBlur() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagSvgScript) AddListenerError(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerError() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagSvgScript) AddListenerFocus(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerFocus() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagSvgScript) AddListenerLoad(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerLoad() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagSvgScript) AddListenerResize(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerResize() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagSvgScript) AddListenerScroll(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerScroll() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagSvgScript) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerAfterprint() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagSvgScript) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerBeforeprint() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagSvgScript) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerBeforeunload() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagSvgScript) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerHashchange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagSvgScript) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerLanguagechange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagSvgScript) AddListenerMessage(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMessage() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagSvgScript) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerMessageerror() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagSvgScript) AddListenerOffline(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerOffline() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagSvgScript) AddListenerOnline(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerOnline() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagSvgScript) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPageswap() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagSvgScript) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPagehide() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagSvgScript) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPagereveal() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagSvgScript) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPageshow() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagSvgScript) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerPopstate() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagSvgScript) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerRejectionhandled() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagSvgScript) AddListenerStorage(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerStorage() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagSvgScript) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerUnhandledrejection() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagSvgScript) AddListenerUnload(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerUnload() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagSvgScript) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerReadystatechange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagSvgScript) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagSvgScript) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagSvgScript) RemoveListenerVisibilitychange() (ref *TagSvgScript) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
