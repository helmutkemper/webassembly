package html

import (
	"log"
	"syscall/js"
)

// TagSvgFeDistantLight
//
// English:
//
// The <feDistantLight> filter primitive defines a distant light source that can be used within a lighting filter
// primitive: <feDiffuseLighting> or <feSpecularLighting>.
//
// Português:
//
// A primitiva de filtro <feDistantLight> define uma fonte de luz distante que pode ser usada em uma primitiva de filtro
// de iluminação: <feDiffuseLighting> ou <feSpecularLighting>.
type TagSvgFeDistantLight struct {

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
func (e *TagSvgFeDistantLight) Init() (ref *TagSvgFeDistantLight) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgFeDistantLight) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgFeDistantLight) CreateElement() (ref *TagSvgFeDistantLight) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "feDistantLight")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgFeDistantLight) AppendToStage() (ref *TagSvgFeDistantLight) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeDistantLight) AppendById(appendId string) (ref *TagSvgFeDistantLight) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgFeDistantLight) AppendToElement(el js.Value) (ref *TagSvgFeDistantLight) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgFeDistantLight) Append(elements ...Compatible) (ref *TagSvgFeDistantLight) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgFeDistantLight) Get() (el js.Value) {
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
func (e *TagSvgFeDistantLight) Id(id string) (ref *TagSvgFeDistantLight) {
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
func (e *TagSvgFeDistantLight) Lang(value interface{}) (ref *TagSvgFeDistantLight) {

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
func (e *TagSvgFeDistantLight) Tabindex(value interface{}) (ref *TagSvgFeDistantLight) {
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
func (e *TagSvgFeDistantLight) XmlLang(value interface{}) (ref *TagSvgFeDistantLight) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

// Azimuth
//
// English:
//
//	The azimuth attribute specifies the direction angle for the light source on the XY plane (clockwise), in degrees
//	from the x axis.
//
//	 Input:
//	   azimuth: specifies the direction angle for the light source on the XY plane
//
// Português:
//
//	O atributo azimute especifica o ângulo de direção da fonte de luz no plano XY (sentido horário), em graus a partir
//	do eixo x.
//
//	 Entrada:
//	   azimuth: especifica o ângulo de direção para a fonte de luz no plano XY
func (e *TagSvgFeDistantLight) Azimuth(azimuth float64) (ref *TagSvgFeDistantLight) {
	e.selfElement.Call("setAttribute", "azimuth", azimuth)
	return e
}

// Elevation
//
// English:
//
//	The elevation attribute specifies the direction angle for the light source from the XY plane towards the Z-axis, in
//	degrees. Note that the positive Z-axis points towards the viewer of the content.
//
//	 Input:
//	   elevation: specifies the direction angle for the light source
//
// Portuguese
//
//	O atributo de elevação especifica o ângulo de direção da fonte de luz do plano XY em direção ao eixo Z, em graus.
//	Observe que o eixo Z positivo aponta para o visualizador do conteúdo.
//
//	 Entrada:
//	   elevation: especifica o ângulo de direção para a fonte de luz
func (e *TagSvgFeDistantLight) Elevation(elevation float64) (ref *TagSvgFeDistantLight) {
	e.selfElement.Call("setAttribute", "elevation", elevation)
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
func (e *TagSvgFeDistantLight) Text(value string) (ref *TagSvgFeDistantLight) {
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
func (e *TagSvgFeDistantLight) Html(value string) (ref *TagSvgFeDistantLight) {
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
func (e *TagSvgFeDistantLight) Reference(reference **TagSvgFeDistantLight) (ref *TagSvgFeDistantLight) {
	*reference = e
	return e
}
