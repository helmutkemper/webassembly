package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"syscall/js"
)

// TagSvgMPath
//
// English:
//
// The <mpath> sub-element for the <animateMotion> element provides the ability to reference an external <path> element
// as the definition of a motion path.
//
// Português:
//
// O subelemento <mpath> para o elemento <animateMotion> fornece a capacidade de referenciar um elemento <path> externo
// como a definição de um caminho de movimento.
type TagSvgMPath struct {
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
func (e *TagSvgMPath) Init() (ref *TagSvgMPath) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgMPath) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgMPath) CreateElement() (ref *TagSvgMPath) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "mpath")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgMPath) AppendToStage() (ref *TagSvgMPath) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgMPath) AppendById(appendId string) (ref *TagSvgMPath) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgMPath) AppendToElement(el js.Value) (ref *TagSvgMPath) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgMPath) Append(elements ...Compatible) (ref *TagSvgMPath) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgMPath) Get() (el js.Value) {
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
func (e *TagSvgMPath) Id(id string) (ref *TagSvgMPath) {
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
func (e *TagSvgMPath) Lang(value interface{}) (ref *TagSvgMPath) {

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
func (e *TagSvgMPath) Tabindex(value interface{}) (ref *TagSvgMPath) {
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
func (e *TagSvgMPath) XmlLang(value interface{}) (ref *TagSvgMPath) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

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
func (e *TagSvgMPath) HRef(href interface{}) (ref *TagSvgMPath) {
	e.selfElement.Call("setAttribute", "href", href)
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
func (e *TagSvgMPath) XLinkHRef(value interface{}) (ref *TagSvgMPath) {
	e.selfElement.Call("setAttribute", "xlink:href", value)
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
func (e *TagSvgMPath) Text(value string) (ref *TagSvgMPath) {
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
func (e *TagSvgMPath) Html(value string) (ref *TagSvgMPath) {
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
func (e *TagSvgMPath) Reference(reference **TagSvgMPath) (ref *TagSvgMPath) {
	*reference = e
	return e
}
func (e *TagSvgMPath) AddListenerAbort(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerAbort() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagSvgMPath) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerAuxclick() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagSvgMPath) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerBeforeinput() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagSvgMPath) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerBeforematch() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagSvgMPath) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerBeforetoggle() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagSvgMPath) AddListenerCancel(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerCancel() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagSvgMPath) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerCanplay() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagSvgMPath) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerCanplaythrough() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagSvgMPath) AddListenerChange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerChange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagSvgMPath) AddListenerClick(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerClick() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagSvgMPath) AddListenerClose(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerClose() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagSvgMPath) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerContextlost() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagSvgMPath) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerContextmenu() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagSvgMPath) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerContextrestored() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagSvgMPath) AddListenerCopy(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerCopy() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagSvgMPath) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerCuechange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagSvgMPath) AddListenerCut(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerCut() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagSvgMPath) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDblclick() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagSvgMPath) AddListenerDrag(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDrag() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagSvgMPath) AddListenerDragend(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDragend() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagSvgMPath) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDragenter() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagSvgMPath) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDragleave() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagSvgMPath) AddListenerDragover(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDragover() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagSvgMPath) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDragstart() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagSvgMPath) AddListenerDrop(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDrop() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagSvgMPath) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerDurationchange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagSvgMPath) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerEmptied() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagSvgMPath) AddListenerEnded(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerEnded() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagSvgMPath) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerFormdata() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagSvgMPath) AddListenerInput(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerInput() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagSvgMPath) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerInvalid() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagSvgMPath) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerKeydown() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagSvgMPath) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerKeypress() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagSvgMPath) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerKeyup() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagSvgMPath) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerLoadeddata() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagSvgMPath) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerLoadedmetadata() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagSvgMPath) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerLoadstart() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagSvgMPath) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMousedown() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagSvgMPath) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMouseenter() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagSvgMPath) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMouseleave() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagSvgMPath) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMousemove() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagSvgMPath) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMouseout() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagSvgMPath) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMouseover() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagSvgMPath) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMouseup() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagSvgMPath) AddListenerPaste(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPaste() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagSvgMPath) AddListenerPause(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPause() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagSvgMPath) AddListenerPlay(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPlay() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagSvgMPath) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPlaying() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagSvgMPath) AddListenerProgress(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerProgress() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagSvgMPath) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerRatechange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagSvgMPath) AddListenerReset(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerReset() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagSvgMPath) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerScrollend() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagSvgMPath) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerSecuritypolicyviolation() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagSvgMPath) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerSeeked() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagSvgMPath) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerSeeking() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagSvgMPath) AddListenerSelect(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerSelect() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagSvgMPath) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerSlotchange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagSvgMPath) AddListenerStalled(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerStalled() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagSvgMPath) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerSubmit() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagSvgMPath) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerSuspend() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagSvgMPath) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerTimeupdate() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagSvgMPath) AddListenerToggle(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerToggle() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagSvgMPath) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerVolumechange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagSvgMPath) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerWaiting() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagSvgMPath) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerWebkitanimationend() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagSvgMPath) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerWebkitanimationiteration() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagSvgMPath) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerWebkitanimationstart() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagSvgMPath) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerWebkittransitionend() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagSvgMPath) AddListenerWheel(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerWheel() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagSvgMPath) AddListenerBlur(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerBlur() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagSvgMPath) AddListenerError(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerError() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagSvgMPath) AddListenerFocus(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerFocus() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagSvgMPath) AddListenerLoad(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerLoad() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagSvgMPath) AddListenerResize(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerResize() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagSvgMPath) AddListenerScroll(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerScroll() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagSvgMPath) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerAfterprint() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagSvgMPath) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerBeforeprint() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagSvgMPath) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerBeforeunload() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagSvgMPath) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerHashchange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagSvgMPath) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerLanguagechange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagSvgMPath) AddListenerMessage(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMessage() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagSvgMPath) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerMessageerror() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagSvgMPath) AddListenerOffline(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerOffline() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagSvgMPath) AddListenerOnline(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerOnline() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagSvgMPath) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPageswap() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagSvgMPath) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPagehide() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagSvgMPath) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPagereveal() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagSvgMPath) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPageshow() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagSvgMPath) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerPopstate() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagSvgMPath) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerRejectionhandled() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagSvgMPath) AddListenerStorage(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerStorage() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagSvgMPath) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerUnhandledrejection() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagSvgMPath) AddListenerUnload(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerUnload() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagSvgMPath) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerReadystatechange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagSvgMPath) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagSvgMPath) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagSvgMPath) RemoveListenerVisibilitychange() (ref *TagSvgMPath) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
