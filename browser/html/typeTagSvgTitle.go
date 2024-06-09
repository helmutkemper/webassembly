package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"log"
	"syscall/js"
)

// TagSvgTitle
//
// English:
//
// The <title> element provides an accessible, short-text description of any SVG container element or graphics element.
//
// Text in a <title> element is not rendered as part of the graphic, but browsers usually display it as a tooltip.
// If an element can be described by visible text, it is recommended to reference that text with an aria-labelledby
// attribute rather than using the <title> element.
//
//	Notes:
//	  * For backward compatibility with SVG 1.1, <title> elements should be the first child element of their parent.
//
// Português:
//
// O elemento <title> fornece uma descrição de texto curto acessível de qualquer elemento de contêiner SVG ou elemento
// gráfico.
//
// O texto em um elemento <title> não é renderizado como parte do gráfico, mas os navegadores geralmente o exibem como
// uma dica de ferramenta. Se um elemento puder ser descrito por texto visível, é recomendável fazer referência a esse
// texto com um atributo aria-labelledby em vez de usar o elemento <title>.
//
//	Notas:
//	  * Para compatibilidade com versões anteriores com SVG 1.1, os elementos <title> devem ser o primeiro elemento
//	    filho de seu pai.
type TagSvgTitle struct {
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

	x int
	y int

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

	// fnClick
	//
	// English:
	//
	// Fired when the user clicks the primary pointer button.
	//
	// Português:
	//
	// Acionado quando o usuário clica no botão do ponteiro principal.
	fnClick *js.Func

	// fnMouseOver
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved outside the element.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do elemento.
	fnMouseOver *js.Func

	// fnMouseOut
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved outside the boundary of the element.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento.
	fnMouseOut *js.Func

	// fnMouseMove
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved while over an element.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido sobre um elemento.
	fnMouseMove *js.Func

	// fnMouseLeave
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved outside the boundary of the element and all of its descendants.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento e de todos os seus descendentes.
	fnMouseLeave *js.Func

	// fnMouseEnter
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved inside the boundary of the element or one of its descendants.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para dentro do limite do elemento ou de um de seus descendentes.
	fnMouseEnter *js.Func

	// fnMouseDown
	//
	// English:
	//
	// Fired when the user presses a button on a mouse or other pointing device, while the pointer is over the element.
	//
	// Português:
	//
	// Acionado quando o usuário pressiona um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
	fnMouseDown *js.Func

	// fnMouseUp
	//
	// English:
	//
	// Fired when the user releases a button on a mouse or other pointing device, while the pointer is over the element.
	//
	// Português:
	//
	// Acionado quando o usuário libera um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
	fnMouseUp *js.Func

	// fnMouseWheel
	//
	// English:
	//
	// Fired when the user rotates a mouse wheel or similar user interface component such as a touchpad.
	//
	// Português:
	//
	// Acionado quando o usuário gira a roda do mouse ou um componente de interface de usuário semelhante, como um touchpad.
	fnMouseWheel *js.Func

	// fnDoubleClick
	//
	// English:
	//
	// Fired when the user double-clicks the primary pointer button.
	//
	// Português:
	//
	// Acionado quando o usuário clica duas vezes no botão do ponteiro principal.
	fnDoubleClick *js.Func

	// fnFocusIn
	//
	// English:
	//
	// The focusin event fires when an element is about to receive focus.
	//
	// Português:
	//
	// O evento focusin é acionado quando um elemento está prestes a receber o foco.
	fnFocusIn *js.Func

	// fnFocusOut
	//
	// English:
	//
	// The focusout event fires when an element is about to lose focus.
	//
	// Português:
	//
	// O evento focusout é acionado quando um elemento está prestes a perder o foco.
	fnFocusOut *js.Func
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
func (e *TagSvgTitle) Init() (ref *TagSvgTitle) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgTitle) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgTitle) CreateElement() (ref *TagSvgTitle) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "title")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgTitle) AppendToStage() (ref *TagSvgTitle) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgTitle) AppendById(appendId string) (ref *TagSvgTitle) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgTitle) AppendToElement(el js.Value) (ref *TagSvgTitle) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgTitle) Append(elements ...Compatible) (ref *TagSvgTitle) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgTitle) Get() (el js.Value) {
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
func (e *TagSvgTitle) Id(id string) (ref *TagSvgTitle) {
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
func (e *TagSvgTitle) Lang(value interface{}) (ref *TagSvgTitle) {

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
func (e *TagSvgTitle) Tabindex(value interface{}) (ref *TagSvgTitle) {
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
func (e *TagSvgTitle) XmlLang(value interface{}) (ref *TagSvgTitle) {
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
func (e *TagSvgTitle) Class(class string) (ref *TagSvgTitle) {
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
func (e *TagSvgTitle) Style(value string) (ref *TagSvgTitle) {
	e.selfElement.Call("setAttribute", "style", value)
	return e
}

// #styling end -------------------------------------------------------------------------------------------------------

// Title
//
// English:
//
// Adds plain text to the tag's content.
//
// Text:
//
// Adiciona um texto simples ao conteúdo da tag.
func (e *TagSvgTitle) Title(value string) (ref *TagSvgTitle) {
	e.selfElement.Set("textContent", value)
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
func (e *TagSvgTitle) Text(value string) (ref *TagSvgTitle) {
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
func (e *TagSvgTitle) Html(value string) (ref *TagSvgTitle) {
	e.selfElement.Set("innerHTML", value)
	return e
}

// GetXY
//
// English:
//
//	Returns the X and Y axes in pixels.
//
// Português:
//
//	Retorna os eixos X e Y em pixels.
func (e *TagSvgTitle) GetXY() (x, y int) {
	x = e.x
	y = e.y
	return
}

// GetX
//
// English:
//
//	Returns the X axe in pixels.
//
// Português:
//
//	Retorna o eixo X em pixels.
func (e *TagSvgTitle) GetX() (x int) {
	return e.x
}

// GetY
//
// English:
//
//	Returns the Y axe in pixels.
//
// Português:
//
//	Retorna o eixo Y em pixels.
func (e *TagSvgTitle) GetY() (y int) {
	return e.y
}

// GetTop
//
// English:
//
//	Same as GetX() function, returns the x position of the element.
//
// Português:
//
//	O mesmo que a função GetX(), retorna a posição x do elemento.
func (e *TagSvgTitle) GetTop() (top float64) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	top = coordinate.Get("top").Float()
	return
}

// GetRight
//
// English:
//
//	It is the same as x + width.
//
// Português:
//
//	É o mesmo que x + width.
func (e *TagSvgTitle) GetRight() (right float64) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	right = coordinate.Get("right").Float()
	return
}

// GetBottom
//
// English:
//
//	It is the same as y + height.
//
// Português:
//
//	É o mesmo que y + Height.
func (e *TagSvgTitle) GetBottom() (bottom float64) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	bottom = coordinate.Get("bottom").Float()
	return
}

// GetLeft
//
// English:
//
//	Same as GetY() function, returns the y position of the element.
//
// Português:
//
//	O mesmo que a função GetY(), retorna a posição y do elemento.
func (e *TagSvgTitle) GetLeft() (left float64) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	left = coordinate.Get("left").Float()
	return
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
func (e *TagSvgTitle) Reference(reference **TagSvgTitle) (ref *TagSvgTitle) {
	*reference = e
	return e
}

// AddListenerMouseOver
//
// English:
//
// Adds a mouse over event listener equivalent to the JavaScript command addEventListener('mouseover',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when a mouse or other pointing device is moved outside the element.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse sobre, equivalente ao comando JavaScript addEventListener('mouseover',fn).
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do elemento.
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseOver(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseOver != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseOver, this, args)
		return nil
	})
	e.fnMouseOver = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseover",
		*e.fnMouseOver,
	)
	return e
}

// RemoveListenerMouseOver
//
// English:
//
// Removes a mouse over event listener, equivalent to the JavaScript command RemoveEventListener('mouseover',fn).
//
// Português:
//
// Remove um ouvinte de evento de mouse sobre, equivalente ao comando JavaScript RemoveEventListener('mouseover',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseOver() (ref *TagSvgTitle) {
	if e.fnMouseOver == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseover",
		*e.fnMouseOver,
	)
	e.fnMouseOver = nil
	return e
}

// AddListenerMouseOut
//
// English:
//
// Adds a mouse out event listener equivalent to the JavaScript command addEventListener('mouseout',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse fora, equivalente ao comando JavaScript addEventListener('mouseout',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseOut(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseOut != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseOut, this, args)
		return nil
	})
	e.fnMouseOut = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseout",
		*e.fnMouseOut,
	)
	return e
}

// RemoveListenerMouseOut
//
// English:
//
// Removes a mouse out event listener, equivalent to the JavaScript command RemoveEventListener('mouseout',fn).
//
// Português:
//
// Remove um ouvinte de evento de mouse fora, equivalente ao comando JavaScript RemoveEventListener('mouseout',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseOut() (ref *TagSvgTitle) {
	if e.fnMouseOut == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseout",
		*e.fnMouseOut,
	)
	e.fnMouseOut = nil
	return e
}

// AddListenerMouseMove
//
// English:
//
// Adds a mouse move event listener equivalent to the JavaScript command addEventListener('mousemove',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse move, equivalente ao comando JavaScript addEventListener('mousemove',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseMove(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseMove != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseMove, this, args)
		return nil
	})
	e.fnMouseMove = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousemove",
		*e.fnMouseMove,
	)
	return e
}

// RemoveListenerMouseMove
//
// English:
//
// Removes a mouse move event listener, equivalent to the JavaScript command RemoveEventListener('mousemove',fn).
//
// Português:
//
// Remove um ouvinte de evento de mouse move, equivalente ao comando JavaScript RemoveEventListener('mousemove',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseMove() (ref *TagSvgTitle) {
	if e.fnMouseMove == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousemove",
		*e.fnMouseMove,
	)
	e.fnMouseMove = nil
	return e
}

// AddListenerMouseLeave
//
// English:
//
// Adds a mouse leave event listener equivalent to the JavaScript command addEventListener('mouseleave',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse saiu, equivalente ao comando JavaScript addEventListener('mouseleave',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseLeave(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseLeave != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseLeave, this, args)
		return nil
	})
	e.fnMouseLeave = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseleave",
		e.fnMouseLeave,
	)
	return e
}

// RemoveListenerMouseLeave
//
// English:
//
// Removes a mouse leave event listener, equivalent to the JavaScript command RemoveEventListener('mouseleave',fn).
//
// Português:
//
// Remove um ouvinte de evento de mouse saiu, equivalente ao comando JavaScript RemoveEventListener('mouseleave',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseLeave() (ref *TagSvgTitle) {
	if e.fnMouseLeave == nil {
		return e
	}

	e.selfElement.Call(
		"addEventListener",
		"mouseleave",
		e.fnMouseLeave,
	)
	e.fnMouseLeave = nil
	return e
}

// AddListenerMouseEnter
//
// English:
//
// Adds a mouse enter event listener equivalent to the JavaScript command addEventListener('mouseenter',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse entrou, equivalente ao comando JavaScript addEventListener('mouseenter',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseEnter(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseEnter != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseEnter, this, args)
		return nil
	})
	e.fnMouseEnter = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseenter",
		*e.fnMouseEnter,
	)
	return e
}

// RemoveListenerMouseEnter
//
// English:
//
// Removes a mouse enter event listener, equivalent to the JavaScript command RemoveEventListener('mouseenter',fn).
//
// Português:
//
// Remove um ouvinte de evento de mouse entrou, equivalente ao comando JavaScript RemoveEventListener('mouseenter',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseEnter() (ref *TagSvgTitle) {
	if e.fnMouseEnter == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseenter",
		*e.fnMouseEnter,
	)
	e.fnMouseEnter = nil
	return e
}

// AddListenerMouseDown
//
// English:
//
// Adds a mouse down event listener equivalent to the JavaScript command addEventListener('mousedown',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de botão do mouse precionado, equivalente ao comando JavaScript
// addEventListener('mousedown',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseDown(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseDown != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseDown, this, args)
		return nil
	})
	e.fnMouseDown = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousedown",
		e.fnMouseDown,
	)
	return e
}

// RemoveListenerMouseDown
//
// English:
//
// Removes a mouse down event listener, equivalent to the JavaScript command RemoveEventListener('mousedown',fn).
//
// Português:
//
// Remove um ouvinte de evento de botão do mouse precionado, equivalente ao comando JavaScript RemoveEventListener('mousedown',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseDown() (ref *TagSvgTitle) {
	if e.fnMouseDown == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousedown",
		e.fnMouseDown,
	)
	e.fnMouseDown = nil
	return e
}

// AddListenerMouseUp
//
// English:
//
// Adds a mouse uo event listener equivalent to the JavaScript command addEventListener('mouseup',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de botão do mouse liberado, equivalente ao comando JavaScript
// addEventListener('mouseup',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseUp(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseUp != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseUp, this, args)
		return nil
	})
	e.fnMouseUp = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseup",
		*e.fnMouseUp,
	)
	return e
}

// RemoveListenerMouseUp
//
// English:
//
// Removes a mouse up event listener, equivalent to the JavaScript command RemoveEventListener('mouseup',fn).
//
// Português:
//
// Remove um ouvinte de evento de botão do mouse liberado, equivalente ao comando JavaScript RemoveEventListener('mouseup',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseUp() (ref *TagSvgTitle) {
	if e.fnMouseUp == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseup",
		*e.fnMouseUp,
	)
	e.fnMouseUp = nil
	return e
}

// AddListenerMouseWheel
//
// English:
//
// Adds a mouse wheel event listener equivalent to the JavaScript command addEventListener('mousewheel',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de roda do mouse equivalente ao comando JavaScript addEventListener('mousewheel',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerMouseWheel(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnMouseWheel != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseWheel, this, args)
		return nil
	})
	e.fnMouseWheel = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousewheel",
		*e.fnMouseWheel,
	)
	return e
}

// RemoveListenerMouseWheel
//
// English:
//
// Removes a mouse wheel event listener, equivalent to the JavaScript command RemoveEventListener('mousewheel',fn).
//
// Português:
//
// Remove um ouvinte de evento de roda do mouse, equivalente ao comando JavaScript RemoveEventListener('mousewheel',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerMouseWheel() (ref *TagSvgTitle) {
	if e.fnMouseWheel == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousewheel",
		*e.fnMouseWheel,
	)
	e.fnMouseWheel = nil
	return e
}

// AddListenerDoubleClick
//
// English:
//
// Adds a mouse double click event listener equivalent to the JavaScript command addEventListener('dblclick',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de click duplo do mouse equivalente ao comando JavaScript
// addEventListener('dblclick',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerDoubleClick(mouseEvent chan mouse.Data) (ref *TagSvgTitle) {
	if e.fnDoubleClick != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventDoubleClick, this, args)
		return nil
	})
	e.fnDoubleClick = &fn

	e.selfElement.Call(
		"addEventListener",
		"dblclick",
		*e.fnDoubleClick,
	)
	return e
}

// RemoveListenerDoubleClick
//
// English:
//
// Removes a double click event listener, equivalent to the JavaScript command RemoveEventListener('dblclick',fn).
//
// Português:
//
// Remove um ouvinte de evento de click duplo, equivalente ao comando JavaScript RemoveEventListener('dblclick',fn).
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerDoubleClick() (ref *TagSvgTitle) {
	if e.fnDoubleClick == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"dblclick",
		*e.fnDoubleClick,
	)
	e.fnDoubleClick = nil
	return e
}

// AddListenerFocusIn
//
// English:
//
// Adds a focus event listener equivalent to the JavaScript command addEventListener('focusin',fn).
//
//	Input:
//	  focusEvent: pointer to channel struct{}
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de foco equivalente ao comando JavaScript addEventListener('focusin',fn).
//
//	Entrada:
//	  focusEvent: ponteiro para o channel struct{}
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  focusEvent := make(chan struct{})
//	  factoryBrowser.NewTagSvgCircle().AddListenerFocusIn(&focusEvent) ...
//
//	  go func() {
//	    for {
//	      select {
//	      case <-focusEvent:
//	        log.Printf("focus in")
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerFocusIn(focusEvent chan struct{}) (ref *TagSvgTitle) {
	if e.fnFocusIn != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		focusEvent <- struct{}{}
		return nil
	})
	e.fnFocusIn = &fn

	e.selfElement.Call(
		"addEventListener",
		"focusin",
		*e.fnFocusIn,
	)
	return e
}

// RemoveListenerFocusIn #replicar
//
// English:
//
// Removes a focus event listener equivalent to the JavaScript command removeEventListener('focusin',fn).
//
//	Input:
//	  focusEvent: pointer to channel struct{}
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Remove um ouvinte de evento de foco equivalente ao comando JavaScript removeEventListener('focusin',fn).
//
//	Entrada:
//	  focusEvent: ponteiro para o channel struct{}
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  focusEvent := make(chan struct{})
//	  factoryBrowser.NewTagSvgCircle().AddListenerFocusIn(&focusEvent) ...
//
//	  go func() {
//	    for {
//	      select {
//	      case <-focusEvent:
//	        log.Printf("focus in")
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerFocusIn() (ref *TagSvgTitle) {
	if e.fnFocusIn == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"focusin",
		*e.fnFocusIn,
	)
	e.fnFocusIn = nil
	return e
}

// AddListenerFocusOut #replicar
//
// English:
//
// Adds a focus out event listener equivalent to the JavaScript command addEventListener('focusout',fn).
//
//	Input:
//	  focusEvent: pointer to channel struct{}
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de perda de foco equivalente ao comando JavaScript addEventListener('focusout',fn).
//
//	Entrada:
//	  focusEvent: ponteiro para o channel struct{}
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  focusEvent := make(chan struct{})
//	  factoryBrowser.NewTagSvgCircle().AddListenerFocusOut(&focusEvent) ...
//
//	  go func() {
//	    for {
//	      select {
//	      case <-focusEvent:
//	        log.Printf("focus out")
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerFocusOut(focusEvent chan struct{}) (ref *TagSvgTitle) {
	if e.fnFocusOut != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		focusEvent <- struct{}{}
		return nil
	})
	e.fnFocusOut = &fn

	e.selfElement.Call(
		"addEventListener",
		"focusout",
		*e.fnFocusOut,
	)
	return e
}

// RemoveListenerFocusOut #replicar
//
// English:
//
// Remove a focus out event listener equivalent to the JavaScript command removeEventListener('focusout',fn).
//
//	Input:
//	  focusEvent: pointer to channel struct{}
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Remove um ouvinte de perda de foco equivalente ao comando JavaScript removeEventListener('focusout',fn).
//
//	Entrada:
//	  focusEvent: ponteiro para o channel struct{}
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  focusEvent := make(chan struct{})
//	  factoryBrowser.NewTagSvgCircle().AddListenerFocusOut(&focusEvent) ...
//
//	  go func() {
//	    for {
//	      select {
//	      case <-focusEvent:
//	        log.Printf("focus out")
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerFocusOut() (ref *TagSvgTitle) {
	if e.fnFocusOut == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"focusout",
		*e.fnFocusOut,
	)
	e.fnFocusOut = nil
	return e
}
func (e *TagSvgTitle) AddListenerAbort(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerAbort() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagSvgTitle) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerAuxclick() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagSvgTitle) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerBeforeinput() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagSvgTitle) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerBeforematch() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagSvgTitle) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerBeforetoggle() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagSvgTitle) AddListenerCancel(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerCancel() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagSvgTitle) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerCanplay() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagSvgTitle) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerCanplaythrough() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagSvgTitle) AddListenerChange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerChange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerChange()
	return e
}

// AddListenerClick
//
// English:
//
// Adds a mouse click event listener equivalent to the JavaScript command addEventListener('click',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when the user clicks the primary pointer button.
//
// Português:
//
// Adiciona um ouvinte de evento de click do mouse, equivalente ao comando JavaScript addEventListener('click',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando o usuário clica no botão do ponteiro principal.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) AddListenerClick(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

// RemoveListenerClick
//
// English:
//
// Removes a mouse click event listener, equivalent to the JavaScript command RemoveEventListener('click',fn).
//
// Fired when the user clicks the primary pointer button.
//
// Português:
//
// Remove um ouvinte de evento de click do mouse, equivalente ao comando JavaScript RemoveEventListener('click',fn).
//
// Acionado quando o usuário clica no botão do ponteiro principal.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSvgTitle) RemoveListenerClick() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagSvgTitle) AddListenerClose(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerClose() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagSvgTitle) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerContextlost() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagSvgTitle) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerContextmenu() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagSvgTitle) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerContextrestored() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagSvgTitle) AddListenerCopy(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerCopy() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagSvgTitle) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerCuechange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagSvgTitle) AddListenerCut(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerCut() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagSvgTitle) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDblclick() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagSvgTitle) AddListenerDrag(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDrag() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagSvgTitle) AddListenerDragend(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDragend() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagSvgTitle) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDragenter() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagSvgTitle) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDragleave() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagSvgTitle) AddListenerDragover(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDragover() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagSvgTitle) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDragstart() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagSvgTitle) AddListenerDrop(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDrop() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagSvgTitle) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerDurationchange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagSvgTitle) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerEmptied() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagSvgTitle) AddListenerEnded(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerEnded() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagSvgTitle) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerFormdata() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagSvgTitle) AddListenerInput(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerInput() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagSvgTitle) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerInvalid() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagSvgTitle) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerKeydown() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagSvgTitle) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerKeypress() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagSvgTitle) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerKeyup() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagSvgTitle) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerLoadeddata() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagSvgTitle) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerLoadedmetadata() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagSvgTitle) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerLoadstart() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagSvgTitle) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMousedown() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagSvgTitle) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMouseenter() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagSvgTitle) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMouseleave() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagSvgTitle) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMousemove() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagSvgTitle) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMouseout() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagSvgTitle) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMouseover() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagSvgTitle) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMouseup() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagSvgTitle) AddListenerPaste(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPaste() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagSvgTitle) AddListenerPause(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPause() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagSvgTitle) AddListenerPlay(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPlay() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagSvgTitle) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPlaying() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagSvgTitle) AddListenerProgress(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerProgress() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagSvgTitle) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerRatechange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagSvgTitle) AddListenerReset(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerReset() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagSvgTitle) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerScrollend() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagSvgTitle) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerSecuritypolicyviolation() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagSvgTitle) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerSeeked() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagSvgTitle) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerSeeking() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagSvgTitle) AddListenerSelect(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerSelect() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagSvgTitle) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerSlotchange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagSvgTitle) AddListenerStalled(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerStalled() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagSvgTitle) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerSubmit() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagSvgTitle) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerSuspend() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagSvgTitle) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerTimeupdate() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagSvgTitle) AddListenerToggle(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerToggle() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagSvgTitle) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerVolumechange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagSvgTitle) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerWaiting() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagSvgTitle) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerWebkitanimationend() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagSvgTitle) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerWebkitanimationiteration() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagSvgTitle) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerWebkitanimationstart() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagSvgTitle) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerWebkittransitionend() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagSvgTitle) AddListenerWheel(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerWheel() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagSvgTitle) AddListenerBlur(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerBlur() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagSvgTitle) AddListenerError(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerError() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagSvgTitle) AddListenerFocus(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerFocus() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagSvgTitle) AddListenerLoad(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerLoad() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagSvgTitle) AddListenerResize(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerResize() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagSvgTitle) AddListenerScroll(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerScroll() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagSvgTitle) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerAfterprint() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagSvgTitle) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerBeforeprint() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagSvgTitle) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerBeforeunload() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagSvgTitle) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerHashchange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagSvgTitle) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerLanguagechange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagSvgTitle) AddListenerMessage(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMessage() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagSvgTitle) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerMessageerror() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagSvgTitle) AddListenerOffline(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerOffline() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagSvgTitle) AddListenerOnline(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerOnline() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagSvgTitle) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPageswap() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagSvgTitle) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPagehide() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagSvgTitle) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPagereveal() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagSvgTitle) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPageshow() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagSvgTitle) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerPopstate() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagSvgTitle) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerRejectionhandled() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagSvgTitle) AddListenerStorage(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerStorage() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagSvgTitle) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerUnhandledrejection() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagSvgTitle) AddListenerUnload(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerUnload() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagSvgTitle) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerReadystatechange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagSvgTitle) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagSvgTitle) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagSvgTitle) RemoveListenerVisibilitychange() (ref *TagSvgTitle) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
