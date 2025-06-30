package html

import (
	"image/color"
	"log"
	"reflect"
	"strconv"
	"syscall/js"
	"time"
)

// TagSvgSet
//
// English:
//
// The SVG <set> element provides a simple means of just setting the value of an attribute for a specified duration.
//
// It supports all attribute types, including those that cannot reasonably be interpolated, such as string and boolean
// values. For attributes that can be reasonably be interpolated, the <animate> is usually preferred.
//
//	Notes:
//	  * The <set> element is non-additive. The additive and accumulate attributes are not allowed, and will be
//	    ignored if specified.
//
// Português:
//
// O elemento SVG <set> fornece um meio simples de apenas definir o valor de um atributo para uma duração especificada.
//
// Ele suporta todos os tipos de atributos, incluindo aqueles que não podem ser interpolados de maneira razoável, como
// valores de string e booleanos. Para atributos que podem ser razoavelmente interpolados, o <animate> geralmente é
// preferido.
//
//	Notas:
//	  * O elemento <set> não é aditivo. Os atributos aditivo e acumular não são permitidos e serão ignorados se
//	    especificados.
type TagSvgSet struct {
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
func (e *TagSvgSet) Init() (ref *TagSvgSet) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgSet) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgSet) CreateElement() (ref *TagSvgSet) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "set")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgSet) AppendToStage() (ref *TagSvgSet) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgSet) AppendById(appendId string) (ref *TagSvgSet) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgSet) AppendToElement(el js.Value) (ref *TagSvgSet) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgSet) Append(elements ...Compatible) (ref *TagSvgSet) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgSet) Get() (el js.Value) {
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
func (e *TagSvgSet) Id(id string) (ref *TagSvgSet) {
	e.id = id
	e.selfElement.Call("setAttribute", "id", id)
	return e
}

// GetId #global
//
// English:
//
//	Return a unique id for an element
//
// The id attribute specifies a unique id for an HTML element (the value must be unique within the
// HTML document).
//
// The id attribute is most used to point to a style in a style sheet, and by JavaScript (via the HTML
// DOM) to manipulate the element with the specific id.
//
// Português:
//
//	Retorna um ID exclusivo para um elemento
//
// O atributo id especifica um id exclusivo para um elemento HTML (o valor deve ser exclusivo no
// documento HTML).
//
// O atributo id é mais usado para apontar para um estilo em uma folha de estilo, e por JavaScript
// (através do HTML DOM) para manipular o elemento com o id específico.
func (e *TagSvgSet) GetId() (id string) {
	return e.id
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
func (e *TagSvgSet) Lang(value interface{}) (ref *TagSvgSet) {

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
func (e *TagSvgSet) Tabindex(value interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) XmlLang(value interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Class(class string) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Style(value string) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "style", value)
	return e
}

// #styling end -------------------------------------------------------------------------------------------------------

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
func (e *TagSvgSet) To(value interface{}) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "to", TypeToString(value, ";", ";"))
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
func (e *TagSvgSet) Begin(begin interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Dur(dur interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) End(end interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Min(value interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Max(value interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Restart(value interface{}) (ref *TagSvgSet) {
	if converted, ok := value.(SvgAnimationRestart); ok {
		e.selfElement.Call("setAttribute", "restart", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "restart", value)
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
func (e *TagSvgSet) RepeatCount(value interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) RepeatDur(value interface{}) (ref *TagSvgSet) {
	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "repeatDur", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "repeatDur", value)
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
func (e *TagSvgSet) Fill(value interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) HRef(href interface{}) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "href", href)
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
func (e *TagSvgSet) AttributeName(attributeName string) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "attributeName", attributeName)
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
func (e *TagSvgSet) CalcMode(value interface{}) (ref *TagSvgSet) {
	if converted, ok := value.(SvgCalcMode); ok {
		e.selfElement.Call("setAttribute", "calcMode", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "calcMode", value)
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
func (e *TagSvgSet) Values(value interface{}) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "values", TypeToString(value, ";", ";"))
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
func (e *TagSvgSet) KeyTimes(value interface{}) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "keyTimes", TypeToString(value, ";", ""))
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
//
// Esse atributo é ignorado, a menos que o atributo calcMode seja definido como spline.
//
// Se houver algum erro na especificação de keySplines (valores incorretos, muitos ou poucos valores), a animação não
// ocorrerá.
func (e *TagSvgSet) KeySplines(value interface{}) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "keySplines", TypeToString(value, " ", ";"))
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
func (e *TagSvgSet) From(value interface{}) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "from", TypeToString(value, ";", ";"))
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
func (e *TagSvgSet) By(by float64) (ref *TagSvgSet) {
	e.selfElement.Call("setAttribute", "by", by)
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
func (e *TagSvgSet) Additive(value interface{}) (ref *TagSvgSet) {
	if converted, ok := value.(SvgAdditive); ok {
		e.selfElement.Call("setAttribute", "additive", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "additive", value)
	return e
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
func (e *TagSvgSet) Accumulate(value interface{}) (ref *TagSvgSet) {
	if converted, ok := value.(SvgAccumulate); ok {
		e.selfElement.Call("setAttribute", "accumulate", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "accumulate", value)
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
func (e *TagSvgSet) XLinkHRef(value interface{}) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Text(value string) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Html(value string) (ref *TagSvgSet) {
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
func (e *TagSvgSet) Reference(reference **TagSvgSet) (ref *TagSvgSet) {
	*reference = e
	return e
}

// ListenerAddReflect
//
// English:
//
//	Add event listener
//
//	Events:
//	  cancel: Fired for <input> and <dialog> elements when the user cancels the currently open dialog by closing it with the Esc key.
//	  change: Fired when the value of an <input>, <select>, or <textarea> element has been changed and committed by the user. Unlike the input event, the change event is not necessarily fired for each alteration to an element's value.
//	  error:  Fired when a resource failed to load, or can't be used.
//	  load:   Fires for elements containing a resource when the resource has successfully loaded.
//
//	Clipboard events
//	  copy:   Fired when the user initiates a copy action through the browser's user interface.
//	  cut:    Fired when the user initiates a cut action through the browser's user interface.
//	  paste:  Fired when the user initiates a paste action through the browser's user interface.
//
//	Drag & drop events
//	  drag:       This event is fired when an element or text selection is being dragged.
//	  dragend:    This event is fired when a drag operation is being ended (by releasing a mouse button or hitting the escape key).
//	  dragenter:  This event is fired when a dragged element or text selection enters a valid drop target.
//	  dragleave:  This event is fired when a dragged element or text selection leaves a valid drop target.
//	  dragover:   This event is fired continuously when an element or text selection is being dragged and the mouse pointer is over a valid drop target (every 50 ms WHEN mouse is not moving ELSE much faster between 5 ms (slow movement) and 1ms (fast movement) approximately. This firing pattern is different than mouseover ).
//	  dragstart:  This event is fired when the user starts dragging an element or text selection.
//	  drop:       This event is fired when an element or text selection is dropped on a valid drop target.
//
//	Popover events
//	  beforetoggle: Fired when the element is a popover, before it is hidden or shown.
//	  toggle:       Fired when the element is a popover, just after it is hidden or shown.
func (e *TagSvgSet) ListenerAddReflect(event string, params []interface{}, functions []reflect.Value, reference any) (ref *TagSvgSet) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.ListenerAddReflect(event, params, functions, reference)
	return e
}

// ListenerRemove
//
// English:
//
//	Remove event listener
//
//	Events:
//	  cancel: Fired for <input> and <dialog> elements when the user cancels the currently open dialog by closing it with the Esc key.
//	  change: Fired when the value of an <input>, <select>, or <textarea> element has been changed and committed by the user. Unlike the input event, the change event is not necessarily fired for each alteration to an element's value.
//	  error:  Fired when a resource failed to load, or can't be used.
//	  load:   Fires for elements containing a resource when the resource has successfully loaded.
//
//	Clipboard events
//	  copy:   Fired when the user initiates a copy action through the browser's user interface.
//	  cut:    Fired when the user initiates a cut action through the browser's user interface.
//	  paste:  Fired when the user initiates a paste action through the browser's user interface.
//
//	Drag & drop events
//	  drag:       This event is fired when an element or text selection is being dragged.
//	  dragend:    This event is fired when a drag operation is being ended (by releasing a mouse button or hitting the escape key).
//	  dragenter:  This event is fired when a dragged element or text selection enters a valid drop target.
//	  dragleave:  This event is fired when a dragged element or text selection leaves a valid drop target.
//	  dragover:   This event is fired continuously when an element or text selection is being dragged and the mouse pointer is over a valid drop target (every 50 ms WHEN mouse is not moving ELSE much faster between 5 ms (slow movement) and 1ms (fast movement) approximately. This firing pattern is different than mouseover ).
//	  dragstart:  This event is fired when the user starts dragging an element or text selection.
//	  drop:       This event is fired when an element or text selection is dropped on a valid drop target.
//
//	Popover events
//	  beforetoggle: Fired when the element is a popover, before it is hidden or shown.
//	  toggle:       Fired when the element is a popover, just after it is hidden or shown.
func (e *TagSvgSet) ListenerRemove(event string) (ref *TagSvgSet) {
	e.commonEvents.ListenerRemove(event)
	return e
}

// Remove
//
// English:
//
//	Removes a child node from the DOM and returns the removed node.
//
// Português:
//
//	Remove um nó filho do DOM e retorna o nó removido.
func (e *TagSvgSet) Remove(elements ...Compatible) (ref *TagSvgSet) {
	for _, element := range elements {
		e.selfElement.Call("removeChild", element.Get())
	}

	return e
}
