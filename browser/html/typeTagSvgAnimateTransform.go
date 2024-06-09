package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/animation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"github.com/helmutkemper/iotmaker.webassembly/platform/engine"
	"image/color"
	"log"
	"strconv"
	"syscall/js"
	"time"
)

// TagSvgAnimateTransform
//
// English:
//
// The animateTransform element animates a transformation attribute on its target element, thereby allowing animations
// to control translation, scaling, rotation, and/or skewing.
//
// Português:
//
// O elemento animateTransform anima um atributo de transformação em seu elemento de destino, permitindo assim que as
// animações controlem a tradução, dimensionamento, rotação e ou inclinação.
type TagSvgAnimateTransform struct {
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

	fnBegin  *js.Func
	fnRepeat *js.Func
	fnEnd    *js.Func
	fnMotion *js.Func

	// fnMotionEngineId
	//
	// English:
	//
	// Engine returns an ID for each function added and the ID must be used to remove the function.
	//
	// Português:
	//
	// Engine retorna um ID para cada função adicionada e o ID deve ser usado para remover à função.
	fnMotionEngineId string

	engine engine.IEngine
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
func (e *TagSvgAnimateTransform) Init() (ref *TagSvgAnimateTransform) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgAnimateTransform) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgAnimateTransform) CreateElement() (ref *TagSvgAnimateTransform) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "animateTransform")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgAnimateTransform) AppendToStage() (ref *TagSvgAnimateTransform) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgAnimateTransform) AppendById(appendId string) (ref *TagSvgAnimateTransform) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgAnimateTransform) AppendToElement(el js.Value) (ref *TagSvgAnimateTransform) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgAnimateTransform) Append(elements ...Compatible) (ref *TagSvgAnimateTransform) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgAnimateTransform) Get() (el js.Value) {
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
func (e *TagSvgAnimateTransform) Id(id string) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Lang(value interface{}) (ref *TagSvgAnimateTransform) {

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
func (e *TagSvgAnimateTransform) Tabindex(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) XmlLang(value interface{}) (ref *TagSvgAnimateTransform) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

// #animate start -----------------------------------------------------------------------------------------------------

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
func (e *TagSvgAnimateTransform) HRef(href interface{}) (ref *TagSvgAnimateTransform) {
	e.selfElement.Call("setAttribute", "href", href)
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
func (e *TagSvgAnimateTransform) Begin(begin interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Dur(dur interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) End(end interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Min(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Max(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Restart(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) RepeatCount(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) RepeatDur(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Fill(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) CalcMode(value interface{}) (ref *TagSvgAnimateTransform) {
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
//	    qualquer outro tipo: interface{}
func (e *TagSvgAnimateTransform) Values(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) KeyTimes(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) KeySplines(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) From(value interface{}) (ref *TagSvgAnimateTransform) {
	e.selfElement.Call("setAttribute", "from", TypeToString(value, " ", " "))
	return e
}

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
func (e *TagSvgAnimateTransform) To(value interface{}) (ref *TagSvgAnimateTransform) {
	e.selfElement.Call("setAttribute", "to", TypeToString(value, " ", " "))
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
func (e *TagSvgAnimateTransform) By(by float64) (ref *TagSvgAnimateTransform) {
	e.selfElement.Call("setAttribute", "by", by)
	return e
}

// #animate end -------------------------------------------------------------------------------------------------------

// Type
//
// English:
//
// Defines the type of transformation, whose values change over time.
//
//	Input:
//	  value: type of transformation
//	    const: KSvgTypeTransform... (e.g. KSvgTypeTransformTranslate)
//	    any other type: interface{}
//
// Português:
//
// Define o tipo de transformação, cujos valores mudam ao longo do tempo.
//
//	Entrada:
//	  value: tipo de transformação
//	    const: KSvgTypeTransform... (ex. KSvgTypeTransformTranslate)
//	    qualquer outro tipo: interface{}
func (e *TagSvgAnimateTransform) Type(value interface{}) (ref *TagSvgAnimateTransform) {
	if converted, ok := value.(SvgTypeTransform); ok {
		e.selfElement.Call("setAttribute", "type", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "type", value)
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
func (e *TagSvgAnimateTransform) AttributeName(attributeName string) (ref *TagSvgAnimateTransform) {
	e.selfElement.Call("setAttribute", "attributeName", attributeName)
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
func (e *TagSvgAnimateTransform) XLinkHRef(value interface{}) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Text(value string) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Html(value string) (ref *TagSvgAnimateTransform) {
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
func (e *TagSvgAnimateTransform) Reference(reference **TagSvgAnimateTransform) (ref *TagSvgAnimateTransform) {
	*reference = e
	return e
}

// AddListenerBegin
//
// English:
//
// Adds an animation begin event listener equivalent to the JavaScript command addEventListener('beginEvent',fn).
//
//	Input:
//	  animationEvent: pointer to channel animation.Data
//
//	Notes:
//	  * For more information see the sites https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//	    and https://developer.mozilla.org/en-US/docs/Web/API/SVGAnimationElement
//
// Português:
//
// Adiciona um ouvinte de inicio de animação equivalente ao comando JavaScript addEventListener('beginEvent',fn).
//
//	Entrada:
//	  animationEvent: ponteiro para o channel animation.Data
//
//	Notas:
//	  * Para mais informações veja os sites https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//	    e https://developer.mozilla.org/en-US/docs/Web/API/SVGAnimationElement
//
//	Example: / Exemplo:
//	  animationEvent := make(chan animation.Data)
//	  factoryBrowser.NewTagSvgCircle() ... .Append(
//	    factoryBrowser.NewTagSvgAnimate().AddListenerEnd(&animationEvent)...
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case data := <-animationEvent:
//	        log.Printf("current time (seconds): %v", data.CurrentTime)
//	      }
//	    }
//	  }()
func (e *TagSvgAnimateTransform) AddListenerBegin(animationEvent chan animation.Data) (ref *TagSvgAnimateTransform) {
	if e.fnBegin != nil {
		return e
	}

	var fn js.Func
	js.FuncOf(func(this js.Value, _ []js.Value) interface{} {
		if this.IsNull() == true || this.IsUndefined() == true {
			return nil
		}

		animationEvent <- animation.EventManager(animation.KEventBegin, this)
		return nil
	})
	e.fnBegin = &fn

	e.selfElement.Call(
		"addEventListener",
		"beginEvent",
		*e.fnBegin,
	)
	return e
}

// RemoveListenerBegin
//
// English:
//
// Removes an animation begin event listener, equivalent to the JavaScript command RemoveEventListener('beginEvent',fn).
//
// Português:
//
// Remove um ouvinte de inicio de animação, equivalente ao comando JavaScript RemoveEventListener('beginEvent',fn).
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
func (e *TagSvgAnimateTransform) RemoveListenerBegin() (ref *TagSvgAnimateTransform) {
	if e.fnBegin == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"beginEvent",
		*e.fnBegin,
	)
	e.fnBegin = nil
	return e
}

// AddListenerRepeat
//
// English:
//
// Adds an animation repeat event listener equivalent to the JavaScript command addEventListener('repeatEvent',fn).
//
//	Input:
//	  animationEvent: pointer to channel animation.Data
//
//	Notes:
//	  * For more information see the sites https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//	    and https://developer.mozilla.org/en-US/docs/Web/API/SVGAnimationElement
//	  * On 072022 this event doesn't work on apple safari
//
// Português:
//
// Adiciona um ouvinte de repetição de animação equivalente ao comando JavaScript addEventListener('repeatEvent',fn).
//
//	Entrada:
//	  animationEvent: ponteiro para o channel animation.Data
//
//	Notas:
//	  * Para mais informações veja os sites https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//	    e https://developer.mozilla.org/en-US/docs/Web/API/SVGAnimationElement
//	  * Em 07/2022, este evento não funciona no safari da apple
//
//	Example: / Exemplo:
//	  animationEvent := make(chan animation.Data)
//	  factoryBrowser.NewTagSvgCircle() ... .Append(
//	    factoryBrowser.NewTagSvgAnimate().AddListenerEnd(&animationEvent)...
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case data := <-animationEvent:
//	        log.Printf("current time (seconds): %v", data.CurrentTime)
//	      }
//	    }
//	  }()
func (e *TagSvgAnimateTransform) AddListenerRepeat(animationEvent chan animation.Data) (ref *TagSvgAnimateTransform) {
	if e.fnRepeat != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if this.IsNull() == true || this.IsUndefined() == true {
			return nil
		}

		animationEvent <- animation.EventManager(animation.KEventRepeat, this)
		return nil
	})
	e.fnRepeat = &fn

	e.selfElement.Call(
		"addEventListener",
		"repeatEvent",
		*e.fnRepeat,
	)
	return e
}

// RemoveListenerRepeat
//
// English:
//
// Removes an animation repeat event listener, equivalent to the JavaScript command RemoveEventListener('repeatEvent',fn).
//
// Português:
//
// Remove um ouvinte de repetição de animação, equivalente ao comando JavaScript RemoveEventListener('repeatEvent',fn).
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
func (e *TagSvgAnimateTransform) RemoveListenerRepeat() (ref *TagSvgAnimateTransform) {
	if e.fnRepeat == nil {
		return e
	}

	e.selfElement.Call(
		"addEventListener",
		"repeatEvent",
		*e.fnRepeat,
	)
	e.fnRepeat = nil
	return e
}

// AddListenerEnd
//
// English:
//
// Adds an animation end event listener equivalent to the JavaScript command addEventListener('endEvent',fn).
//
//	Input:
//	  animationEvent: pointer to channel animation.Data
//
//	Notes:
//	  * For more information see the sites https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//	    and https://developer.mozilla.org/en-US/docs/Web/API/SVGAnimationElement
//
// Português:
//
// Adiciona um ouvinte de fim de animação equivalente ao comando JavaScript addEventListener('endEvent',fn).
//
//	Entrada:
//	  animationEvent: ponteiro para o channel animation.Data
//
//	Notas:
//	  * Para mais informações veja os sites https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//	    e https://developer.mozilla.org/en-US/docs/Web/API/SVGAnimationElement
//
//	Example: / Exemplo:
//	  animationEvent := make(chan animation.Data)
//	  factoryBrowser.NewTagSvgCircle() ... .Append(
//	    factoryBrowser.NewTagSvgAnimate().AddListenerEnd(&animationEvent)...
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case data := <-animationEvent:
//	        log.Printf("current time (seconds): %v", data.CurrentTime)
//	      }
//	    }
//	  }()
func (e *TagSvgAnimateTransform) AddListenerEnd(animationEvent chan animation.Data) (ref *TagSvgAnimateTransform) {
	if e.fnEnd != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if this.IsNull() == true || this.IsUndefined() == true {
			return nil
		}

		animationEvent <- animation.EventManager(animation.KEventEnd, this)
		return nil
	})
	e.fnEnd = &fn

	e.selfElement.Call(
		"addEventListener",
		"endEvent",
		*e.fnEnd,
	)
	return e
}

// RemoveListenerEnd
//
// English:
//
// Removes an animation end event listener, equivalent to the JavaScript command RemoveEventListener('endEvent',fn).
//
// Português:
//
// Remove um ouvinte de fim de animação, equivalente ao comando JavaScript RemoveEventListener('endEvent',fn).
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
func (e *TagSvgAnimateTransform) RemoveListenerEnd() (ref *TagSvgAnimateTransform) {
	if e.fnEnd == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"endEvent",
		*e.fnEnd,
	)
	e.fnEnd = nil
	return e
}

// AddListenerMotion
//
// English:
//
// Adds an animation motion event listener equivalent to the JavaScript command addEventListener('motionEvent',fn).
//
//	Input:
//	  animationEvent: pointer to channel animation.Data
//
//	Notes:
//	  * This is not a native browser event or documented by mozilla.
//
// Português:
//
// Adiciona um ouvinte de movimento de animação equivalente ao comando JavaScript addEventListener('motionEvent',fn).
//
//	Entrada:
//	  animationEvent: ponteiro para o channel animation.Data
//
//	Notas:
//	  * Este não é um evento nativo do navegador ou documentado pela mozilla.
//
//	Example: / Exemplo:
//	  animationEvent := make(chan animation.Data)
//	  factoryBrowser.NewTagSvgCircle() ... .Append(
//	    factoryBrowser.NewTagSvgAnimate().AddListenerEnd(&animationEvent)...
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case data := <-animationEvent:
//	        log.Printf("current time (seconds): %v", data.CurrentTime)
//	      }
//	    }
//	  }()
func (e *TagSvgAnimateTransform) AddListenerMotion(animationEvent chan animation.Data) (ref *TagSvgAnimateTransform) {
	if e.fnMotion != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, _ []js.Value) interface{} {
		if this.IsNull() == true || this.IsUndefined() == true {
			return nil
		}
		var create = time.Now()
		e.fnMotionEngineId, _ = e.engine.DrawAddToFunctions(
			func() {
				var now = time.Now()
				var data animation.Data
				data.This = this
				data.CurrentTime = float64(now.Sub(create).Milliseconds()) / 1000.0
				data.EventName = animation.KEventMotion
				animationEvent <- data
			},
		)
		return nil
	})
	e.fnMotion = &fn

	e.selfElement.Call(
		"addEventListener",
		"beginEvent",
		*e.fnMotion,
	)
	e.selfElement.Call(
		"addEventListener",
		"endEvent",
		js.FuncOf(func(this js.Value, _ []js.Value) interface{} {
			e.RemoveListenerMotion()
			return nil
		}),
	)
	return e
}

// RemoveListenerMotion
//
// English:
//
// Removes an animation motion event listener, equivalent to the JavaScript command RemoveEventListener('motionEvent',fn).
//
//	Notes:
//	  * This is not a native browser event or documented by mozilla.
//
// Português:
//
// Remove um ouvinte de movimento de animação, equivalente ao comando JavaScript RemoveEventListener('motionEvent',fn).
//
//	Notas:
//	  * Este não é um evento nativo do navegador ou documentado pela mozilla.
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
func (e *TagSvgAnimateTransform) RemoveListenerMotion() (ref *TagSvgAnimateTransform) {
	if e.fnMotion == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"endEvent",
		js.FuncOf(func(this js.Value, _ []js.Value) interface{} {
			e.RemoveListenerMotion()
			return nil
		}),
	)
	e.selfElement.Call(
		"removeEventListener",
		"beginEvent",
		*e.fnMotion,
	)

	e.engine.DrawDeleteFromFunctions(e.fnMotionEngineId)
	e.fnMotion = nil

	return e
}
func (e *TagSvgAnimateTransform) AddListenerAbort(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerAbort() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerAuxclick() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerBeforeinput() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerBeforematch() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerBeforetoggle() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerCancel(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerCancel() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerCanplay() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerCanplaythrough() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerChange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerChange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerClick(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerClick() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerClose(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerClose() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerContextlost() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerContextmenu() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerContextrestored() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerCopy(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerCopy() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerCuechange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerCut(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerCut() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDblclick() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDrag(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDrag() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDragend(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDragend() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDragenter() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDragleave() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDragover(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDragover() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDragstart() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDrop(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDrop() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerDurationchange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerEmptied() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerEnded(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerEnded() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerFormdata() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerInput(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerInput() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerInvalid() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerKeydown() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerKeypress() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerKeyup() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerLoadeddata() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerLoadedmetadata() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerLoadstart() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMousedown() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMouseenter() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMouseleave() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMousemove() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMouseout() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMouseover() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMouseup() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPaste(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPaste() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPause(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPause() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPlay(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPlay() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPlaying() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerProgress(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerProgress() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerRatechange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerReset(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerReset() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerScrollend() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerSecuritypolicyviolation() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerSeeked() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerSeeking() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerSelect(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerSelect() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerSlotchange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerStalled(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerStalled() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerSubmit() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerSuspend() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerTimeupdate() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerToggle(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerToggle() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerVolumechange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerWaiting() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerWebkitanimationend() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerWebkitanimationiteration() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerWebkitanimationstart() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerWebkittransitionend() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerWheel(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerWheel() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerBlur(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerBlur() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerError(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerError() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerFocus(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerFocus() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerLoad(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerLoad() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerResize(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerResize() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerScroll(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerScroll() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerAfterprint() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerBeforeprint() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerBeforeunload() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerHashchange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerLanguagechange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMessage(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMessage() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerMessageerror() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerOffline(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerOffline() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerOnline(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerOnline() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPageswap() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPagehide() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPagereveal() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPageshow() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerPopstate() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerRejectionhandled() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerStorage(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerStorage() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerUnhandledrejection() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerUnload(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerUnload() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerReadystatechange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagSvgAnimateTransform) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagSvgAnimateTransform) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagSvgAnimateTransform) RemoveListenerVisibilitychange() (ref *TagSvgAnimateTransform) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
