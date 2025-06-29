package html

import (
	"image/color"
	"log"
	"strconv"
	"syscall/js"
	"time"
)

// TagSvgFeFuncG
//
// English:
//
// The <feFuncG> SVG filter primitive defines the transfer function for the green component of the input graphic of
// its parent <feComponentTransfer> element.
//
// Português:
//
// A primitiva de filtro SVG <feFuncG> define a função de transferência para o componente verde do gráfico de entrada
// de seu elemento pai <feComponentTransfer>.
type TagSvgFeFuncG struct {

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
func (e *TagSvgFeFuncG) Init() (ref *TagSvgFeFuncG) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgFeFuncG) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgFeFuncG) CreateElement() (ref *TagSvgFeFuncG) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "feFuncG")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgFeFuncG) AppendToStage() (ref *TagSvgFeFuncG) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeFuncG) AppendById(appendId string) (ref *TagSvgFeFuncG) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeFuncG) AppendToElement(el js.Value) (ref *TagSvgFeFuncG) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgFeFuncG) Append(elements ...Compatible) (ref *TagSvgFeFuncG) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgFeFuncG) Get() (el js.Value) {
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
func (e *TagSvgFeFuncG) Id(id string) (ref *TagSvgFeFuncG) {
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
func (e *TagSvgFeFuncG) GetId() (id string) {
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
func (e *TagSvgFeFuncG) Lang(value interface{}) (ref *TagSvgFeFuncG) {

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
func (e *TagSvgFeFuncG) Tabindex(value interface{}) (ref *TagSvgFeFuncG) {
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
func (e *TagSvgFeFuncG) XmlLang(value interface{}) (ref *TagSvgFeFuncG) {
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
//	Input:
//	  value: type of component transfer function
//	    const: KSvgTypeFeFunc... (e.g. KSvgTypeFeFuncIdentity)
//	    any other type: interface{}
//
// Português:
//
// Indica o tipo de função de transferência de componentes.
//
//	Input:
//	  value: tipo de função de transferência de componente
//	    const: KSvgTypeFeFunc... (ex. KSvgTypeFeFuncIdentity)
//	    any other type: interface{}
func (e *TagSvgFeFuncG) Type(value interface{}) (ref *TagSvgFeFuncG) {
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
//	Input:
//	  value: defines a list of numbers
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1) rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0% 10%"
//	    []float64: []float64{0.0, 10.0} = "0 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    any other type: interface{}
//
// Português:
//
// O atributo tableValues define uma lista de números que definem uma tabela de consulta de valores para uma função de
// transferência de componente de cor.
//
//	Entrada:
//	  value: define uma lista de números
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1) rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0% 10%"
//	    []float64: []float64{0.0, 10.0} = "0 10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s 1s"
//	    time.Duration: time.Second = "1s"
//	    float32: 0.1 = "10%"
//	    float64: 10.0 = "10"
//	    color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	    qualquer outro tipo: interface{}
func (e *TagSvgFeFuncG) TableValues(value interface{}) (ref *TagSvgFeFuncG) {
	if converted, ok := value.([]color.RGBA); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += RGBAToJs(v) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "tableValues", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float32); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "% "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "tableValues", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "tableValues", valueStr[:length])
		return e
	}

	if converted, ok := value.([]time.Duration); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += v.String() + " "
		}
		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "tableValues", valueStr[:length])
		return e
	}

	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "tableValues", converted.String())
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "tableValues", p)
		return e
	}

	if converted, ok := value.(float64); ok {
		p := strconv.FormatFloat(converted, 'g', -1, 64)
		e.selfElement.Call("setAttribute", "tableValues", p)
		return e
	}

	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "tableValues", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "tableValues", value)
	return e
}

// Intercept
//
// English:
//
//	The intercept attribute defines the intercept of the linear function of color component transfers when the type
//	attribute is set to linear.
//
// Portuguese
//
//	O atributo de interceptação define a interceptação da função linear de transferências de componentes de cor quando
//	o atributo de tipo é definido como linear.
func (e *TagSvgFeFuncG) Intercept(intercept float64) (ref *TagSvgFeFuncG) {
	e.selfElement.Call("setAttribute", "intercept", intercept)
	return e
}

// Amplitude
//
// English:
//
//	The amplitude attribute controls the amplitude of the gamma function of a component transfer element when its type
//	attribute is gamma.
//
//	 Input:
//	   value: controls the amplitude of the gamma function
//	     []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1) rgba(255,0,0,1)"
//	     []float32: []float64{0.0, 0.1} = "0% 10%"
//	     []float64: []float64{0.0, 10.0} = "0 10"
//	     []time.Duration: []time.Duration{0, time.Second} = "0s 1s"
//	     time.Duration: time.Second = "1s"
//	     float32: 0.1 = "10%"
//	     float64: 10.0 = "10"
//	     color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	     any other type: interface{}
//
// Português:
//
//	O atributo amplitude controla à amplitude da função gama de um elemento de transferência de componente quando seu
//	atributo de tipo é gama.
//
//	 Entrada:
//	   value: controla a amplitude da função de gama
//	     []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1) rgba(255,0,0,1)"
//	     []float32: []float64{0.0, 0.1} = "0% 10%"
//	     []float64: []float64{0.0, 10.0} = "0 10"
//	     []time.Duration: []time.Duration{0, time.Second} = "0s 1s"
//	     time.Duration: time.Second = "1s"
//	     float32: 0.1 = "10%"
//	     float64: 10.0 = "10"
//	     color.RGBA: factoryColor.NewRed() = "rgba(255,0,0,1)"
//	     qualquer outro tipo: interface{}
func (e *TagSvgFeFuncG) Amplitude(value interface{}) (ref *TagSvgFeFuncG) {
	if converted, ok := value.([]color.RGBA); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += RGBAToJs(v) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "amplitude", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float32); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(100.0*float64(v), 'g', -1, 64) + "% "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "amplitude", valueStr[:length])
		return e
	}

	if converted, ok := value.([]float64); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}

		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "amplitude", valueStr[:length])
		return e
	}

	if converted, ok := value.([]time.Duration); ok {
		var valueStr = ""
		for _, v := range converted {
			valueStr += v.String() + " "
		}
		var length = len(valueStr) - 1

		e.selfElement.Call("setAttribute", "amplitude", valueStr[:length])
		return e
	}

	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "amplitude", converted.String())
		return e
	}

	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "amplitude", p)
		return e
	}

	if converted, ok := value.(float64); ok {
		p := strconv.FormatFloat(converted, 'g', -1, 64)
		e.selfElement.Call("setAttribute", "amplitude", p)
		return e
	}

	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "amplitude", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "amplitude", value)
	return e
}

// Exponent
//
// English:
//
//	The exponent attribute defines the exponent of the gamma function.
//
//	 Input:
//	   exponent: defines the exponent of the gamma function
//
// Portuguese
//
//	O atributo expoente define o expoente da função gama.
//
//	 Entrada:
//	   exponent: define o expoente da função gama
func (e *TagSvgFeFuncG) Exponent(exponent float64) (ref *TagSvgFeFuncG) {
	e.selfElement.Call("setAttribute", "exponent", exponent)
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
func (e *TagSvgFeFuncG) Text(value string) (ref *TagSvgFeFuncG) {
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
func (e *TagSvgFeFuncG) Html(value string) (ref *TagSvgFeFuncG) {
	e.selfElement.Set("innerHTML", value)
	return e
}

// Slope
//
// English:
//
// Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already
// been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for
// compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the
// bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
//
// The slope attribute indicates the vertical stroke angle of a font.
//
// Português:
//
// Descontinuado: este recurso não é mais recomendado. Embora alguns navegadores ainda possam suportá-lo, ele pode já
// ter sido removido dos padrões da Web relevantes, pode estar em processo de eliminação ou pode ser mantido apenas para
// fins de compatibilidade. Evite usá-lo e atualize o código existente, se possível; consulte a tabela de
// compatibilidade na parte inferior desta página para orientar sua decisão. Esteja ciente de que esse recurso pode
// deixar de funcionar a qualquer momento.
//
// O atributo slope indica o ângulo de traço vertical de uma fonte.
func (e *TagSvgFeFuncG) Slope(value interface{}) (ref *TagSvgFeFuncG) {
	e.selfElement.Call("setAttribute", "slope", value)
	return e
}

// Offset
//
// English:
//
// This attribute defines where the gradient stop is placed along the gradient vector.
//
// Português:
//
// Este atributo define onde a parada de gradiente é colocada ao longo do vetor de gradiente.
func (e *TagSvgFeFuncG) Offset(value interface{}) (ref *TagSvgFeFuncG) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "offset", p)
		return e
	}

	if converted, ok := value.(float64); ok {
		p := strconv.FormatFloat(converted, 'g', -1, 64)
		e.selfElement.Call("setAttribute", "offset", p)
		return e
	}

	e.selfElement.Call("setAttribute", "offset", value)
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
func (e *TagSvgFeFuncG) Reference(reference **TagSvgFeFuncG) (ref *TagSvgFeFuncG) {
	*reference = e
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
func (e *TagSvgFeFuncG) Remove(elements ...Compatible) (ref *TagSvgFeFuncG) {
	for _, element := range elements {
		e.selfElement.Call("removeChild", element)
	}

	return e
}
