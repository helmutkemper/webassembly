package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/animation"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/engine"
	"image/color"
	"log"
	"strconv"
	"sync"
	"syscall/js"
	"time"
)

// TagSvgAnimateMotion
//
// English:
//
// The SVG <animateMotion> element provides a way to define how an element moves along a motion path.
//
//	Notes:
//	  * To reuse an existing path, it will be necessary to use an <mpath> element inside the <animateMotion> element
//	    instead of the path attribute.
//
// Português:
//
// O elemento SVG <animateMotion> fornece uma maneira de definir como um elemento se move ao longo de um caminho
// de movimento.
//
//	Notas:
//	  * Para reutilizar um caminho existente, será necessário usar um elemento <mpath> dentro do elemento
//	    <animateMotion> ao invés do atributo path.
type TagSvgAnimateMotion struct {

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

func (e *TagSvgAnimateMotion) Engine(engine engine.IEngine) {
	e.engine = engine
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
func (e *TagSvgAnimateMotion) Init() (ref *TagSvgAnimateMotion) {
	e.listener = new(sync.Map)

	e.CreateElement()
	e.prepareStageReference()

	return e
}

func (e *TagSvgAnimateMotion) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgAnimateMotion) CreateElement() (ref *TagSvgAnimateMotion) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "animateMotion")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

func (e *TagSvgAnimateMotion) AppendToStage() (ref *TagSvgAnimateMotion) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgAnimateMotion) AppendById(appendId string) (ref *TagSvgAnimateMotion) {
	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgAnimateMotion) AppendToElement(el js.Value) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgAnimateMotion) Append(elements ...Compatible) (ref *TagSvgAnimateMotion) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagSvgAnimateMotion) Get() (el js.Value) {
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
func (e *TagSvgAnimateMotion) Id(id string) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Lang(value interface{}) (ref *TagSvgAnimateMotion) {

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
func (e *TagSvgAnimateMotion) Tabindex(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) XmlLang(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(Language); ok {
		e.selfElement.Call("setAttribute", "xml:lang", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "xml:lang", value)
	return e
}

// #core end ----------------------------------------------------------------------------------------------------------

// #presentation start ------------------------------------------------------------------------------------------------

// BaselineShift
//
// English:
//
//	The baseline-shift attribute allows repositioning of the dominant-baseline relative to the dominant-baseline of the
//	parent text content element. The shifted object might be a sub- or superscript.
//
//	 Input:
//	   baselineShift: allows repositioning of the dominant-baseline relative to the dominant-baseline of the parent text
//	   content element.
//	     float32: 0.05 = "5%"
//	     string: "5%"
//	     consts KSvgBaselineShift... (e.g. KSvgBaselineShiftAuto)
//
//	 Notes:
//	   * As a presentation attribute baseline-shift can be used as a CSS property.
//	   * This property is going to be deprecated and authors are advised to use vertical-align instead.
//
// Português:
//
//	O atributo baseline-shift permite o reposicionamento da linha de base dominante em relação à linha de base dominante
//	do elemento de conteúdo de texto pai. O objeto deslocado pode ser um sub ou sobrescrito.
//
//	 Input:
//	   baselineShift: permite o reposicionamento da linha de base dominante em relação à linha de base dominante do
//	   elemento de conteúdo de texto pai.
//	     float32: 0.05 = "5%"
//	     string: "5%"
//	     consts KSvgBaselineShift... (ex. KSvgBaselineShiftAuto)
//
//	 Notas:
//	   * Como atributo de apresentação, baseline-shift pode ser usado como propriedade CSS.
//	   * Essa propriedade será preterida e os autores são aconselhados a usar alinhamento vertical.
func (e *TagSvgAnimateMotion) BaselineShift(baselineShift interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := baselineShift.(SvgBaselineShift); ok {
		e.selfElement.Call("setAttribute", "baseline-shift", converted.String())
		return e
	}

	if converted, ok := baselineShift.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "baseline-shift", p)
		return e
	}

	e.selfElement.Call("setAttribute", "baseline-shift", baselineShift)
	return e
}

// ClipPath
//
// English:
//
//	It binds the element it is applied to with a given <clipPath> element.
//
//	 Input:
//	   clipPath: the element it is applied
//	     (e.g. "url(#myClip)", "circle() fill-box", "circle() stroke-box" or "circle() view-box")
//
// Português:
//
//	Ele associa o elemento ao qual é aplicado a um determinado elemento <clipPath>.
//
//	 Entrada:
//	   clipPath: elemento ao qual é aplicado
//	     (ex. "url(#myClip)", "circle() fill-box", "circle() stroke-box" ou "circle() view-box")
func (e *TagSvgAnimateMotion) ClipPath(clipPath string) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "clip-path", clipPath)
	return e
}

// ClipRule
//
// English:
//
//	It indicates how to determine what side of a path is inside a shape in order to know how a <clipPath> should clip
//	its target.
//
//	 Input:
//	   value: side of a path
//	     const: KSvgClipRule... (e.g. KSvgClipRuleNonzero)
//	     any other type: interface{}
//
// Português:
//
//	Ele indica como determinar qual lado de um caminho está dentro de uma forma para saber como um <clipPath> deve
//	recortar seu destino.
//
//	 Input:
//	   value: lado de um caminho
//	     const: KSvgClipRule... (e.g. KSvgClipRuleNonzero)
//	     qualquer outro tipo: interface{}
func (e *TagSvgAnimateMotion) ClipRule(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgClipRule); ok {
		e.selfElement.Call("setAttribute", "clip-rule", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "clip-rule", value)
	return e
}

// Color
//
// English:
//
//	It provides a potential indirect value (currentcolor) for the fill, stroke, stop-color, flood-color and
//	lighting-color presentation attributes.
//
//	 Input:
//	   value: potential indirect value of color
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, color can be used as a CSS property. See CSS color for further information.
//
// Português:
//
//	Ele fornece um valor indireto potencial (currentcolor) para os atributos de apresentação de preenchimento, traçado,
//	cor de parada, cor de inundação e cor de iluminação.
//
//	 Entrada:
//	   value: valor indireto potencial da cor
//	     string: ex. "black"
//	     factory: ex. factoryColor.NewYellow()
//	     RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, a cor pode ser usada como propriedade CSS. Veja cor CSS para mais informações.
func (e *TagSvgAnimateMotion) Color(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color", value)
	return e
}

// ColorInterpolation
//
// English:
//
// The color-interpolation attribute specifies the color space for gradient interpolations, color animations, and alpha
// compositing.
//
//	Input:
//	  value: specifies the color space for gradient interpolations
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// The color-interpolation property chooses between color operations occurring in the sRGB color space or in a (light
// energy linear) linearized RGB color space. Having chosen the appropriate color space, component-wise linear
// interpolation is used.
//
// When a child element is blended into a background, the value of the color-interpolation property on the child
// determines the type of blending, not the value of the color-interpolation on the parent.
// For gradients which make use of the href or the deprecated xlink:href attribute to reference another gradient, the
// gradient uses the property's value from the gradient element which is directly referenced by the fill or stroke
// property. When animating colors, color interpolation is performed according to the value of the color-interpolation
// property on the element being animated.
//
//	Notes:
//	  * For filter effects, the color-interpolation-filters property controls which color space is used.
//	  * As a presentation attribute, color-interpolation can be used as a CSS property.
//
// Português:
//
// O atributo color-interpolation especifica o espaço de cores para interpolações de gradiente, animações de cores e
// composição alfa.
//
//	Entrada:
//	  value: especifica o espaço de cores para interpolações de gradiente
//	    string: ex. "black"
//	    factory: ex. factoryColor.NewYellow()
//	    RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
//
// A propriedade de interpolação de cores escolhe entre operações de cores que ocorrem no espaço de cores sRGB ou em um
// espaço de cores RGB linearizado (energia de luz linear). Tendo escolhido o espaço de cor apropriado, a interpolação
// linear de componentes é usada.
//
// Quando um elemento filho é mesclado em um plano de fundo, o valor da propriedade color-interpolation no filho
// determina o tipo de mesclagem, não o valor da interpolação de cores no pai.
// Para gradientes que usam o href ou o atributo obsoleto xlink:href para referenciar outro gradiente, o gradiente usa
// o valor da propriedade do elemento gradiente que é diretamente referenciado pela propriedade fill ou stroke.
// Ao animar cores, à interpolação de cores é executada de acordo com o valor da propriedade color-interpolation no
// elemento que está sendo animado.
//
//	Notas:
//	  * Para efeitos de filtro, a propriedade color-interpolation-filters controla qual espaço de cor é usado.
//	  * Como atributo de apresentação, a interpolação de cores pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateMotion) ColorInterpolation(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color-interpolation", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color-interpolation", value)
	return e
}

// ColorInterpolationFilters
//
// English:
//
// The color-interpolation-filters attribute specifies the color space for imaging operations performed via filter
// effects.
//
//	Input:
//	  value: specifies the color space for imaging operations
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
//	Notes:
//	  * This property just has an affect on filter operations. Therefore, it has no effect on filter primitives like
//	    <feOffset>, <feImage>, <feTile> or <feFlood>;
//	  * color-interpolation-filters has a different initial value than color-interpolation. color-interpolation-filters
//	    has an initial value of linearRGB, whereas color-interpolation has an initial value of sRGB. Thus, in the
//	    default case, filter effects operations occur in the linearRGB color space, whereas all other color
//	    interpolations occur by default in the sRGB color space;
//	  * It has no affect on filter functions, which operate in the sRGB color space;
//	  * As a presentation attribute, color-interpolation-filters can be used as a CSS property.
//
// Português:
//
// O atributo color-interpolation-filters especifica o espaço de cores para operações de imagem realizadas por meio de
// efeitos de filtro.
//
//	Entrada:
//	  value: especifica o espaço de cores para operações de imagem
//	    string: ex. "black"
//	    factory: ex. factoryColor.NewYellow()
//	    RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Esta propriedade afeta apenas as operações de filtro. Portanto, não tem efeito em primitivos de filtro como
//	    <feOffset>, <feImage>, <feTile> ou <feFlood>.
//	  * color-interpolation-filters tem um valor inicial diferente de color-interpolation. color-interpolation-filters
//	    tem um valor inicial de linearRGB, enquanto color-interpolation tem um valor inicial de sRGB. Assim, no caso
//	    padrão, as operações de efeitos de filtro ocorrem no espaço de cores linearRGB, enquanto todas as outras
//	    interpolações de cores ocorrem por padrão no espaço de cores sRGB.
//	  * Não afeta as funções de filtro, que operam no espaço de cores sRGB.
//	  * Como atributo de apresentação, os filtros de interpolação de cores podem ser usados como uma propriedade CSS.
func (e *TagSvgAnimateMotion) ColorInterpolationFilters(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color-interpolation-filters", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color-interpolation-filters", value)
	return e
}

// Cursor
//
// English:
//
// The cursor attribute specifies the mouse cursor displayed when the mouse pointer is over an element.
//
//	Input:
//	  value: specifies the mouse cursor displayed when the mouse pointer is over an element
//	    const: KSvgCursor... (e.g.: KSvgCursorMove)
//	    any other type: interface{}
//
// This attribute behaves exactly like the css cursor property except that if the browser supports the <cursor> element,
// you should be able to use it with the <funciri> notation.
//
// As a presentation attribute, it also can be used as a property directly inside a CSS stylesheet, see css cursor for
// further information.
//
// Português:
//
// O atributo cursor especifica o cursor do mouse exibido quando o ponteiro do mouse está sobre um elemento.
//
//	Entrada:
//	  value: especifica o cursor do mouse exibido quando o ponteiro do mouse está sobre um elemento
//	    const: KSvgCursor... (ex.: KSvgCursorMove)
//	    qualquer outro tipo: interface{}
//
// Este atributo se comporta exatamente como a propriedade cursor css, exceto que, se o navegador suportar o elemento
// <cursor>, você poderá usá-lo com a notação <funciri>.
//
// Como atributo de apresentação, também pode ser usado como propriedade diretamente dentro de uma folha de estilo CSS,
// veja cursor css para mais informações.
func (e *TagSvgAnimateMotion) Cursor(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgCursor); ok {
		e.selfElement.Call("setAttribute", "cursor", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "cursor", value)
	return e
}

// Direction
//
// English:
//
// The direction attribute specifies the inline-base direction of a <text> or <tspan> element. It defines the start
// and end points of a line of text as used by the text-anchor and inline-size properties. It also may affect the
// direction in which characters are positioned if the unicode-bidi property's value is either embed or bidi-override.
//
//	Input:
//	  value: specifies the inline-base direction of a <text> or <tspan> element
//	    const: KSvgDirection... (e.g. KSvgDirectionRtl)
//	    any other type: interface{}
//
// It applies only to glyphs oriented perpendicular to the inline-base direction, which includes the usual case of
// horizontally-oriented Latin or Arabic text and the case of narrow-cell Latin or Arabic characters rotated 90 degrees
// clockwise relative to a top-to-bottom inline-base direction.
//
// In many cases, the bidirectional Unicode algorithm produces the desired result automatically, so this attribute
// doesn't need to be specified in those cases. For other cases, such as when using right-to-left languages, it may be
// sufficient to add the direction attribute to the outermost <svg> element, and allow that direction to inherit to all
// text elements:
//
//	Notes:
//	  * As a presentation attribute, direction can be used as a CSS property. See css direction for further
//	    information.
//
// Português:
//
// O atributo direction especifica a direção da base embutida de um elemento <text> ou <tspan>. Ele define os pontos
// inicial e final de uma linha de texto conforme usado pelas propriedades text-anchor e inline-size.
// Também pode afetar a direção na qual os caracteres são posicionados se o valor da propriedade unicode-bidi for
// incorporado ou substituído por bidi.
//
//	Input:
//	  value: especifica a direção da base inline de um elemento <text> ou <tspan>
//	    const: KSvgDirection... (e.g. KSvgDirectionRtl)
//	    qualquer outro tipo: interface{}
//
// Aplica-se apenas a glifos orientados perpendicularmente à direção da base em linha, que inclui o caso usual de texto
// latino ou árabe orientado horizontalmente e o caso de caracteres latinos ou árabes de célula estreita girados 90
// graus no sentido horário em relação a um texto de cima para baixo direção de base em linha.
//
// Em muitos casos, o algoritmo Unicode bidirecional produz o resultado desejado automaticamente, portanto, esse
// atributo não precisa ser especificado nesses casos. Para outros casos, como ao usar idiomas da direita para a
// esquerda, pode ser suficiente adicionar o atributo direction ao elemento <svg> mais externo e permitir que essa
// direção herde todos os elementos de texto:
//
//	Notas:
//	  * Como atributo de apresentação, a direção pode ser usada como uma propriedade CSS. Veja a direção do CSS para
//	    mais informações.
func (e *TagSvgAnimateMotion) Direction(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgDirection); ok {
		e.selfElement.Call("setAttribute", "direction", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "direction", value)
	return e
}

// Display
//
// English:
//
//	The display attribute lets you control the rendering of graphical or container elements.
//
//	 Input:
//	   value: control the rendering of graphical or container elements
//	     nil: display="none"
//	     const: KSvgDisplay... (e.g. KSvgDisplayBlock)
//	     any other type: interface{}
//
// A value of display="none" indicates that the given element and its children will not be rendered. Any value other
// than none or inherit indicates that the given element will be rendered by the browser.
//
// When applied to a container element, setting display to none causes the container and all of its children to be
// invisible; thus, it acts on groups of elements as a group. This means that any child of an element with
// display="none" will never be rendered even if the child has a value for display other than none.
//
// When the display attribute is set to none, then the given element does not become part of the rendering tree. It has
// implications for the <tspan>, <tref>, and <altGlyph> elements, event processing, for bounding box calculations and
// for calculation of clipping paths:
//
//   - If display is set to none on a <tspan>, <tref>, or <altGlyph> element, then the text string is ignored for the
//     purposes of text layout.
//   - Regarding events, if display is set to none, the element receives no events.
//   - The geometry of a graphics element with display set to none is not included in bounding box and clipping paths
//     calculations.
//
// The display attribute only affects the direct rendering of a given element, whereas it does not prevent elements
// from being referenced by other elements. For example, setting it to none on a <path> element will prevent that
// element from getting rendered directly onto the canvas, but the <path> element can still be referenced by a
// <textPath> element; furthermore, its geometry will be used in text-on-a-path processing even if the <path> has a
// display value of none.
//
// This attribute also affects direct rendering into offscreen canvases, such as occurs with masks or clip paths. Thus,
// setting display="none" on a child of a <mask> will prevent the given child element from being rendered as part of the
// mask. Similarly, setting display="none" on a child of a <clipPath> element will prevent the given child element from
// contributing to the clipping path.
//
//	Notes:
//	  * As a presentation attribute, display can be used as a CSS property. See css display for further information.
//
// Português:
//
//	O atributo display permite controlar a renderização de elementos gráficos ou de contêiner.
//
//	 Entrada:
//	   value: controlar a renderização de elementos gráficos ou de contêiner
//	     nil: display="none"
//	     const: KSvgDisplay... (ex. KSvgDisplayBlock)
//	     qualquer outro tipo: interface{}
//
// Um valor de display="none" indica que o elemento fornecido e seus filhos não serão renderizados. Qualquer valor
// diferente de none ou herdar indica que o elemento fornecido será renderizado pelo navegador.
//
// Quando aplicado a um elemento de contêiner, definir display como none faz com que o contêiner e todos os seus filhos
// fiquem invisíveis; assim, atua em grupos de elementos como um grupo. Isso significa que qualquer filho de um elemento
// com display="none" nunca será renderizado, mesmo que o filho tenha um valor para exibição diferente de none.
//
// Quando o atributo display é definido como none, o elemento fornecido não se torna parte da árvore de renderização.
// Tem implicações para os elementos <tspan>, <tref> e <altGlyph>, processamento de eventos, para cálculos de caixa
// delimitadora e para cálculo de caminhos de recorte:
//   - Se display for definido como none em um elemento <tspan>, <tref> ou <altGlyph>, a string de texto será ignorada
//     para fins de layout de texto.
//   - Com relação aos eventos, se display estiver definido como none, o elemento não recebe eventos.
//   - A geometria de um elemento gráfico com exibição definida como nenhum não é incluída nos cálculos da caixa
//     delimitadora e dos caminhos de recorte.
//
// O atributo display afeta apenas a renderização direta de um determinado elemento, mas não impede que os elementos
// sejam referenciados por outros elementos. Por exemplo, defini-lo como none em um elemento <path> impedirá que esse
// elemento seja renderizado diretamente na tela, mas o elemento <path> ainda pode ser referenciado por um elemento
// <textPath>; além disso, sua geometria será usada no processamento de texto em um caminho, mesmo que o <caminho>
// tenha um valor de exibição de nenhum.
//
// Esse atributo também afeta a renderização direta em telas fora da tela, como ocorre com máscaras ou caminhos de
// clipe. Assim, definir display="none" em um filho de uma <mask> impedirá que o elemento filho fornecido seja
// renderizado como parte da máscara. Da mesma forma, definir display="none" em um filho de um elemento <clipPath>
// impedirá que o elemento filho fornecido contribua para o caminho de recorte.
//
//	Notas:
//	  * Como atributo de apresentação, display pode ser usado como propriedade CSS. Consulte a exibição css para obter
//	    mais informações.
func (e *TagSvgAnimateMotion) Display(value interface{}) (ref *TagSvgAnimateMotion) {
	if value == nil {
		e.selfElement.Call("setAttribute", "display", "none")
		return e
	}

	if converted, ok := value.(SvgDisplay); ok {
		e.selfElement.Call("setAttribute", "display", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "display", value)
	return e
}

// DominantBaseline
//
// English:
//
// The dominant-baseline attribute specifies the dominant baseline, which is the baseline used to align the box's text
// and inline-level contents. It also indicates the default alignment baseline of any boxes participating in baseline
// alignment in the box's alignment context.
//
//	Input:
//	  value: is the baseline used to align the box's text and inline-level contents
//	    const: KSvgDominantBaseline... (e.g. KSvgDominantBaselineHanging)
//	    any other type: interface{}
//
// It is used to determine or re-determine a scaled-baseline-table. A scaled-baseline-table is a compound value with
// three components:
//
//  1. a baseline-identifier for the dominant-baseline,
//  2. a baseline-table, and
//  3. a baseline-table font-size.
//
// Some values of the property re-determine all three values. Others only re-establish the baseline-table font-size.
// When the initial value, auto, would give an undesired result, this property can be used to explicitly set the desired
// scaled-baseline-table.
//
// If there is no baseline table in the nominal font, or if the baseline table lacks an entry for the desired baseline,
// then the browser may use heuristics to determine the position of the desired baseline.
//
//	Notes:
//	  * As a presentation attribute, dominant-baseline can be used as a CSS property.
//
// Português:
//
// O atributo linha de base dominante especifica a linha de base dominante, que é a linha de base usada para alinhar o
// texto da caixa e o conteúdo do nível embutido. Também indica a linha de base de alinhamento padrão de todas as caixas
// que participam do alinhamento da linha de base no contexto de alinhamento da caixa.
//
//	Entrada:
//	  value: é a linha de base usada para alinhar o texto da caixa e o conteúdo embutido
//	    const: KSvgDominantBaseline... (ex. KSvgDominantBaselineHanging)
//	    qualquer outro tipo: interface{}
//
// Ele é usado para determinar ou re-determinar uma tabela de linha de base dimensionada. Uma tabela de linha de base
// dimensionada é um valor composto com três componentes:
//
//  1. um identificador de linha de base para a linha de base dominante,
//  2. uma tabela de linha de base, e
//  3. um tamanho de fonte da tabela de linha de base.
//
// Alguns valores da propriedade redeterminam todos os três valores. Outros apenas restabelecem o tamanho da fonte da
// tabela de linha de base. Quando o valor inicial, auto, daria um resultado indesejado, essa propriedade pode ser usada
// para definir explicitamente a tabela de linha de base dimensionada desejada.
//
// Se não houver nenhuma tabela de linha de base na fonte nominal, ou se a tabela de linha de base não tiver uma entrada
// para a linha de base desejada, o navegador poderá usar heurística para determinar a posição da linha de base
// desejada.
//
//	Notas:
//	  * Como atributo de apresentação, a linha de base dominante pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateMotion) DominantBaseline(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgDominantBaseline); ok {
		e.selfElement.Call("setAttribute", "dominant-baseline", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "dominant-baseline", value)
	return e
}

// FillOpacity
//
// English:
//
//	The fill-opacity attribute is a presentation attribute defining the opacity of the paint server (color, gradient,
//	pattern, etc) applied to a shape.
//
//	 Input:
//	   value: defining the opacity of the paint
//	     float32: 1.0 = "100%"
//	     any other type: interface{}
//
//	 Notes:
//	   *As a presentation attribute fill-opacity can be used as a CSS property.
//
// Portuguese
//
//	O atributo fill-opacity é um atributo de apresentação que define a opacidade do servidor de pintura (cor, gradiente,
//	padrão etc.) aplicado a uma forma.
//
//	 Entrada:
//	   value: definindo a opacidade da tinta
//	     float32: 1.0 = "100%"
//	     qualquer outro tipo: interface{}
//
//	 Notes:
//	   *As a presentation attribute fill-opacity can be used as a CSS property.
func (e *TagSvgAnimateMotion) FillOpacity(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "fill-opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "fill-opacity", value)
	return e
}

// FillRule
//
// English:
//
//	The fill-rule attribute is a presentation attribute defining the algorithm to use to determine the inside part of
//	a shape.
//
//	 Input:
//	   value: defining the algorithm to use to determine the inside part of a shape.
//	     const: KSvgFillRule... (e.g. KSvgFillRuleEvenOdd)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, fill-rule can be used as a CSS property.
//
// Portuguese
//
//	O atributo fill-rule é um atributo de apresentação que define o algoritmo a ser usado para determinar a parte
//	interna de uma forma.
//
//	 Input:
//	   value: define o algoritmo a ser usado para determinar a parte interna de uma forma.
//	     const: KSvgFillRule... (eg. KSvgFillRuleEvenOdd)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, fill-rule pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) FillRule(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgFillRule); ok {
		e.selfElement.Call("setAttribute", "fill-rule", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "fill-rule", value)
	return e
}

// Filter
//
// English:
//
//	The filter attribute specifies the filter effects defined by the <filter> element that shall be applied to its
//	element.
//
//	 Input:
//	   filter: specifies the filter effects
//
//	 Notes:
//	   * As a presentation attribute, filter can be used as a CSS property. See css filter for further information.
//
// Portuguese
//
//	O atributo filter especifica os efeitos de filtro definidos pelo elemento <filter> que devem ser aplicados ao seu
//	elemento.
//
//	 Entrada:
//	   filter: especifica os efeitos do filtro
//
//	 Notas:
//	   * Como atributo de apresentação, o filtro pode ser usado como propriedade CSS. Veja filtro css para mais
//	     informações.
func (e *TagSvgAnimateMotion) Filter(filter string) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "filter", filter)
	return e
}

// FloodColor
//
// English:
//
//	The flood-color attribute indicates what color to use to flood the current filter primitive subregion.
//
//	 Input:
//	   floodColor: indicates what color to use to flood the current filter primitive subregion
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, flood-color can be used as a CSS property.
//
// Portuguese
//
//	O atributo flood-color indica qual cor usar para inundar a sub-região primitiva do filtro atual.
//
//	 Entrada:
//	   floodColor: indica qual cor usar para inundar a sub-região primitiva do filtro atual
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, a cor de inundação pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateMotion) FloodColor(floodColor interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := floodColor.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "flood-color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "flood-color", floodColor)
	return e
}

// FloodOpacity
//
// English:
//
//	The flood-opacity attribute indicates the opacity value to use across the current filter primitive subregion.
//
//	 Input:
//	   floodOpacity: indicates the opacity value
//
//	 Notes:
//	   * As a presentation attribute, flood-opacity can be used as a CSS property.
//
// Portuguese
//
//	O atributo flood-opacity indica o valor de opacidade a ser usado na sub-região primitiva de filtro atual.
//
//	 Entrada:
//	   floodOpacity: indica o valor da opacidade
//
//	 Notas:
//	   * Como atributo de apresentação, a opacidade de inundação pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateMotion) FloodOpacity(floodOpacity float64) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "flood-opacity", floodOpacity)
	return e
}

// FontFamily
//
// English:
//
// The font-family attribute indicates which font family will be used to render the text, specified as a prioritized
// list of font family names and/or generic family names.
//
//	Input:
//	  fontFamily: indicates which font family will be used
//	    string: e.g. "Verdana, sans-serif"
//	    factory: e.g. factoryFontFamily.NewArial()
//
//	Notes:
//	  * As a presentation attribute, font-family can be used as a CSS property. See the css font-family property for
//	    more information.
//
// # Portuguese
//
// O atributo font-family indica qual família de fontes será usada para renderizar o texto, especificada como uma lista
// priorizada de nomes de famílias de fontes e ou nomes de famílias genéricos.
//
//	Entrada:
//	  fontFamily: indica qual família de fontes será usada
//	    string: ex. "Verdana, sans-serif"
//	    factory: ex. factoryFontFamily.NewArial()
//
//	Notas:
//	  * Como atributo de apresentação, font-family pode ser usada como propriedade CSS. Consulte a propriedade CSS
//	    font-family para obter mais informações.
func (e *TagSvgAnimateMotion) FontFamily(fontFamily string) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "font-family", fontFamily)
	return e
}

// FontSize
//
// English:
//
// The font-size attribute refers to the size of the font from baseline to baseline when multiple lines of text are set
// solid in a multiline layout environment.
//
//	Input:
//	  fontSize: size of the font
//	    string: e.g. "10px","2em"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, font-size can be used as a CSS property. See the css font-size property for more
//	    information.
//
// # Portuguese
//
// O atributo font-size refere-se ao tamanho da fonte da linha de base a linha de base quando várias linhas de texto
// são definidas como sólidas em um ambiente de layout de várias linhas.
//
//	Entrada:
//	  fontSize: tamanho da fonte
//	    string: ex. "10px","2em"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, font-size pode ser usado como uma propriedade CSS. Consulte a propriedade CSS
//	    font-size para obter mais informações.
func (e *TagSvgAnimateMotion) FontSize(fontSize interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "font-size", fontSize)
	return e
}

// FontSizeAdjust
//
// English:
//
//	The font-size-adjust attribute allows authors to specify an aspect value for an element that will preserve the
//	x-height of the first choice font in a substitute font.
//
//	 Notes:
//	   * As a presentation attribute, font-size-adjust can be used as a CSS property. See the css font-size-adjust
//	     property for more information.
//
// Portuguese
//
//	O atributo font-size-adjust permite que os autores especifiquem um valor de aspecto para um elemento que preservará
//	a altura x da fonte de primeira escolha em uma fonte substituta.
//
//	 Notes:
//	   * As a presentation attribute, font-size-adjust can be used as a CSS property. See the css font-size-adjust
//	     property for more information.
func (e *TagSvgAnimateMotion) FontSizeAdjust(fontSizeAdjust float64) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "font-size-adjust", fontSizeAdjust)
	return e
}

// FontStretch
//
// English:
//
//	The font-stretch attribute indicates the desired amount of condensing or expansion in the glyphs used to render
//	the text.
//
//	 Input:
//	   fontStretch: indicates the desired amount of condensing or expansion
//	     KSvgFontStretch... (e.g. KSvgFontStretchUltraCondensed)
//	     percentage (e.g. "50%")
//
//	 Notes:
//	   * As a presentation attribute, font-stretch can be used as a CSS property. See the css font-stretch property for
//	     more information.
//
// Portuguese
//
//	O atributo font-stretch indica a quantidade desejada de condensação ou expansão nos glifos usados para renderizar
//	o texto.
//
//	 Entrada:
//	   fontStretch: indica a quantidade desejada de condensação ou expansão
//	     KSvgFontStretch... (e.g. KSvgFontStretchUltraCondensed)
//	     percentage (e.g. "50%")
//
//	 Notas:
//	   * Como atributo de apresentação, font-stretch pode ser usado como uma propriedade CSS. Consulte a propriedade
//	     CSS font-stretch para obter mais informações.
func (e *TagSvgAnimateMotion) FontStretch(fontStretch interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := fontStretch.(SvgFontStretch); ok {
		e.selfElement.Call("setAttribute", "font-stretch", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-stretch", fontStretch)
	return e
}

// FontStyle
//
// English:
//
// The font-style attribute specifies whether the text is to be rendered using a normal, italic, or oblique face.
//
//	Input:
//	  value: specifies whether the text is to be rendered using a normal, italic, or oblique face
//	    const: KFontStyleRule... (e.g. KFontStyleRuleItalic)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, font-style can be used as a CSS property. See the css font-style property for
//	    more information.
//
// # Portuguese
//
// O atributo font-style especifica se o texto deve ser renderizado usando uma face normal, itálica ou oblíqua.
//
//	Entrada:
//	  value: especifica se o texto deve ser renderizado usando uma face normal, itálica ou oblíqua
//	    const: KFontStyleRule... (ex. KFontStyleRuleItalic)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, font-style pode ser usado como propriedade CSS. Consulte a propriedade CSS
//	    font-style para obter mais informações.
func (e *TagSvgAnimateMotion) FontStyle(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(FontStyleRule); ok {
		e.selfElement.Call("setAttribute", "font-style", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-style", value)
	return e
}

// FontVariant
//
// English:
//
//	The font-variant attribute indicates whether the text is to be rendered using variations of the font's glyphs.
//
//	 Input:
//	   value: indicates whether the text is to be rendered
//	     const: KFontVariantRule... (e.g. KFontVariantRuleSmallCaps)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, font-variant can be used as a CSS property. See the css font-variant property
//	     for more information.
//
// Portuguese
//
//	O atributo font-variant indica se o texto deve ser renderizado usando variações dos glifos da fonte.
//
//	 Entrada:
//	   value: indica onde o texto vai ser renderizado.
//	     const: KFontVariantRule... (ex. KFontVariantRuleSmallCaps)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, font-variant pode ser usado como uma propriedade CSS. Consulte a propriedade
//	     CSS font-variant para obter mais informações.
func (e *TagSvgAnimateMotion) FontVariant(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(FontVariantRule); ok {
		e.selfElement.Call("setAttribute", "font-variant", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-variant", value)
	return e
}

// FontWeight
//
// English:
//
//	The font-weight attribute refers to the boldness or lightness of the glyphs used to render the text, relative to
//	other fonts in the same font family.
//
//	 Input:
//	   value: refers to the boldness or lightness of the glyphs used to render the text
//	     const: KFontWeightRule... (e.g. KFontWeightRuleBold)
//	     any other type: interface{}
//
//	 Notes:
//	   * As a presentation attribute, font-weight can be used as a CSS property. See the css font-weight property for
//	     more information.
//
// Portuguese
//
//	O atributo font-weight refere-se ao negrito ou leveza dos glifos usados para renderizar o texto, em relação a
//	outras fontes na mesma família de fontes.
//
//	 Entrada:
//	   value: refere-se ao negrito ou leveza dos glifos usados para renderizar o texto
//	     const: KFontWeightRule... (ex. KFontWeightRuleBold)
//	     qualquer outro tipo: interface{}
//
//	 Notas:
//	   * Como atributo de apresentação, o peso da fonte pode ser usado como uma propriedade CSS. Consulte a propriedade
//	     CSS font-weight para obter mais informações.
func (e *TagSvgAnimateMotion) FontWeight(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(FontWeightRule); ok {
		e.selfElement.Call("setAttribute", "font-weight", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "font-weight", value)
	return e
}

// ImageRendering
//
// English:
//
//	The image-rendering attribute provides a hint to the browser about how to make speed vs. quality tradeoffs as it
//	performs image processing.
//
// The resampling is always done in a truecolor (e.g., 24-bit) color space even if the original data and/or the target
// device is indexed color.
//
//	Notes:
//	  * As a presentation attribute, image-rendering can be used as a CSS property. See the css image-rendering
//	    property for more information.
//
// Portuguese
//
//	O atributo de renderização de imagem fornece uma dica ao navegador sobre como fazer compensações de velocidade
//	versus qualidade enquanto executa o processamento de imagem.
//
// A reamostragem é sempre feita em um espaço de cores truecolor (por exemplo, 24 bits), mesmo que os dados originais e
// ou o dispositivo de destino sejam cores indexadas.
//
//	Notas:
//	  * Como um atributo de apresentação, a renderização de imagem pode ser usada como uma propriedade CSS. Consulte
//	    a propriedade de renderização de imagem css para obter mais informações.
func (e *TagSvgAnimateMotion) ImageRendering(imageRendering string) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "image-rendering", imageRendering)
	return e
}

// LetterSpacing
//
// English:
//
// The letter-spacing attribute controls spacing between text characters, in addition to any spacing from the kerning
// attribute.
//
//	Input:
//	  value: controls spacing between text characters
//
// If the attribute value is a unitless number (like 128), the browser processes it as a <length> in the current user
// coordinate system.
//
// If the attribute value has a unit identifier, such as .25em or 1%, then the browser converts the <length> into its
// corresponding value in the current user coordinate system.
//
// Notes:
//   - As a presentation attribute, letter-spacing can be used as a CSS property.
//     See the css letter-spacing property for more information.
//
// Português:
//
// O atributo letter-spacing controla o espaçamento entre caracteres de texto, além de qualquer espaçamento do atributo
// kerning.
//
//	Input:
//	  value: controla o espaçamento entre caracteres de texto
//
// Se o valor do atributo for um número sem unidade (como 128), o navegador o processará como um <comprimento> no
// sistema de coordenadas do usuário atual.
//
// Se o valor do atributo tiver um identificador de unidade, como .25em ou 1%, o navegador converterá o <comprimento>
// em seu valor correspondente no sistema de coordenadas do usuário atual.
//
// Notas:
//   - Como atributo de apresentação, o espaçamento entre letras pode ser usado como uma propriedade CSS.
//     Consulte a propriedade de espaçamento entre letras do CSS para obter mais informações.
func (e *TagSvgAnimateMotion) LetterSpacing(value float64) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "letter-spacing", value)
	return e
}

// LightingColor #presentation
//
// English:
//
// The lighting-color attribute defines the color of the light source for lighting filter primitives.
//
//	Input:
//	  value: defines the color of the light source
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// Português:
//
// O atributo lighting-color define a cor da fonte de luz para as primitivas do filtro de iluminação.
//
//	Input:
//	  value: define a cor da fonte de luz
//	    string: ex. "black"
//	    factory: ex. factoryColor.NewYellow()
//	    RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
func (e *TagSvgAnimateMotion) LightingColor(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "lighting-color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "lighting-color", value)
	return e
}

// MarkerEnd
//
// English:
//
// The marker-end attribute defines the arrowhead or polymarker that will be drawn at the final vertex of the given
// shape.
//
//	Input:
//	  value: the arrowhead or polymarker that will be drawn
//	    string: (e.g. "url(#triangle)")
//
// For all shape elements, except <polyline> and <path>, the last vertex is the same as the first vertex. In this case,
// if the value of marker-start and marker-end are both not none, then two markers will be rendered on that final
// vertex.
// For <path> elements, for each closed subpath, the last vertex is the same as the first vertex. marker-end is only
// rendered on the final vertex of the path data.
//
// Notes:
//   - As a presentation attribute, marker-end can be used as a CSS property.
//
// Português:
//
// O atributo marker-end define a ponta de seta ou polimarcador que será desenhado no vértice final da forma dada.
//
//	Entrada:
//	  value: a ponta de seta ou polimarcador que será desenhado
//	    string: (e.g. "url(#triangle)")
//
// Para todos os elementos de forma, exceto <polyline> e <path>, o último vértice é o mesmo que o primeiro vértice.
// Nesse caso, se o valor de marker-start e marker-end não for nenhum, então dois marcadores serão renderizados nesse
// vértice final.
// Para elementos <path>, para cada subcaminho fechado, o último vértice é igual ao primeiro vértice.
// O final do marcador é renderizado apenas no vértice final dos dados do caminho.
//
// Notas:
//   - Como atributo de apresentação, o marker-end pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) MarkerEnd(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "marker-end", value)
	return e
}

// MarkerMid
//
// English:
//
// The marker-mid attribute defines the arrowhead or polymarker that will be drawn at all interior vertices of the
// given shape.
//
//	Input:
//	  value: defines the arrowhead or polymarker that will be drawn
//	    string: e.g. "url(#circle)"
//
// The marker is rendered on every vertex other than the first and last vertices of the path data.
//
// Notes:
//   - As a presentation attribute, marker-mid can be used as a CSS property.
//
// Português:
//
// O atributo marker-mid define a ponta de seta ou polimarcador que será desenhado em todos os vértices internos da
// forma dada.
//
//	Input:
//	  value: define a ponta de seta ou polimarcador que será desenhado
//	    string: ex. "url(#circle)"
//
// O marcador é renderizado em todos os vértices, exceto no primeiro e no último vértice dos dados do caminho.
//
// Notas:
//   - Como atributo de apresentação, o marker-mid pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) MarkerMid(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "marker-mid", value)
	return e
}

// MarkerStart
//
// English:
//
// The marker-start attribute defines the arrowhead or polymarker that will be drawn at the first vertex of the given
// shape.
//
//	Input:
//	  value: defines the arrowhead or polymarker that will be drawn
//	    string: e.g. "url(#triangle)"
//
// For all shape elements, except <polyline> and <path>, the last vertex is the same as the first vertex. In this case,
// if the value of marker-start and marker-end are both not none, then two markers will be rendered on that final
// vertex.
// For <path> elements, for each closed subpath, the last vertex is the same as the first vertex. marker-start is only
// rendered on the first vertex of the path data.
//
// Notes:
//   - As a presentation attribute, marker-start can be used as a CSS property.
//
// Português:
//
// O atributo marker-start define a ponta de seta ou polimarcador que será desenhado no primeiro vértice da forma dada.
//
//	Entrada:
//	  value: define a ponta de seta ou polimarcador que será desenhado
//	    string: e.g. "url(#triangle)"
//
// Para todos os elementos de forma, exceto <polyline> e <path>, o último vértice é o mesmo que o primeiro vértice.
// Nesse caso, se o valor de marker-start e marker-end não for nenhum, então dois marcadores serão renderizados nesse
// vértice final.
// Para elementos <path>, para cada subcaminho fechado, o último vértice é igual ao primeiro vértice. O início do
// marcador é renderizado apenas no primeiro vértice dos dados do caminho.
//
// Notas:
//   - Como atributo de apresentação, o início do marcador pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) MarkerStart(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "marker-start", value)
	return e
}

// Mask
//
// English:
//
// The mask attribute is a presentation attribute mainly used to bind a given <mask> element with the element the
// attribute belongs to.
//
//	Input:
//	  value: attribute mainly used to bind a given <mask> element
//	    string: "url(#myMask)"
//
// Notes:
//   - As a presentation attribute mask can be used as a CSS property.
//
// Português:
//
// O atributo mask é um atributo de apresentação usado principalmente para vincular um determinado elemento <mask> ao
// elemento ao qual o atributo pertence.
//
//	Entrada:
//	  value: atributo usado principalmente para vincular um determinado elemento <mask>
//	    string: "url(#myMask)"
//
// Notas:
//   - Como uma máscara de atributo de apresentação pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateMotion) Mask(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "mask", value)
	return e
}

// Opacity
//
// English:
//
// The opacity attribute specifies the transparency of an object or of a group of objects, that is, the degree to which
// the background behind the element is overlaid.
//
//	Input:
//	  value: specifies the transparency of an object
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, opacity can be used as a CSS property. See the css opacity property for more
//	    information.
//
// Português:
//
// O atributo opacity especifica a transparência de um objeto ou de um grupo de objetos, ou seja, o grau em que o fundo
// atrás do elemento é sobreposto.
//
//	Entrada:
//	  value: especifica a transparência de um objeto
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
//	Notes:
//	  * Como atributo de apresentação, a opacidade pode ser usada como uma propriedade CSS. Consulte a propriedade de
//	    opacidade do CSS para obter mais informações.
func (e *TagSvgAnimateMotion) Opacity(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "opacity", value)
	return e
}

// Overflow
//
// English:
//
// The overflow attribute sets what to do when an element's content is too big to fit in its block formatting context.
//
//	Input:
//	  value: sets what to do when an element's content is too big to fit in its block formatting context
//	    const: KOverflow... (e.g. KOverflowHidden)
//	    any other type: interface{}
//
// This attribute has the same parameter values and meaning as the css overflow property, however, the following
// additional points apply:
//
//   - If it has a value of visible, the attribute has no effect (i.e., a clipping rectangle is not created).
//
//   - If the overflow property has the value hidden or scroll, a clip of the exact size of the SVG viewport is applied.
//
//   - When scroll is specified on an <svg> element, a scrollbar or panner is normally shown for the SVG viewport
//     whether or not any of its content is clipped.
//
//   - Within SVG content, the value auto implies that all rendered content for child elements must be visible, either
//     through a scrolling mechanism, or by rendering with no clip.
//
//     Notes:
//
//   - Although the initial value for overflow is auto, it is overwritten in the User Agent style sheet for the <svg>
//     element when it is not the root element of a stand-alone document, the <pattern> element, and the <marker>
//     element to be hidden by default.
//
//   - As a presentation attribute, overflow can be used as a CSS property. See the CSS overflow property for more
//     information.
//
// Português:
//
// O atributo overflow define o que fazer quando o conteúdo de um elemento é muito grande para caber em seu contexto
// de formatação de bloco.
//
//	Entrada:
//	  value: define o que fazer quando o conteúdo de um elemento é muito grande para caber em seu contexto de
//	      formatação de bloco
//	    const: KOverflow... (e.g. KOverflowHidden)
//	    qualquer outro tipo: interface{}
//
// Este atributo tem os mesmos valores de parâmetro e significado que a propriedade CSS overflow, no entanto, os
// seguintes pontos adicionais se aplicam:
//
//   - Se tiver um valor de visible, o atributo não terá efeito (ou seja, um retângulo de recorte não será criado).
//
//   - Se a propriedade overflow tiver o valor oculto ou rolar, um clipe do tamanho exato da janela de visualização SVG
//     será aplicado.
//
//   - Quando a rolagem é especificada em um elemento <svg>, uma barra de rolagem ou panner normalmente é mostrado para
//     a janela de visualização SVG, independentemente de seu conteúdo estar ou não recortado.
//
//   - No conteúdo SVG, o valor auto implica que o conteúdo renderizado para elementos filho deve ser visível por
//     completo, seja por meio de um mecanismo de rolagem ou renderizando sem clipe.
//
//     Notas:
//
//   - Embora o valor inicial para estouro seja auto, ele é substituído na folha de estilo do User Agent para o
//     elemento <svg> quando não é o elemento raiz de um documento autônomo, o elemento <pattern> e o elemento
//     <marker> para ser ocultado por padrão.
//
//   - Como atributo de apresentação, overflow pode ser usado como propriedade CSS. Consulte a propriedade CSS
//     overflow para obter mais informações.
func (e *TagSvgAnimateMotion) Overflow(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(Overflow); ok {
		e.selfElement.Call("setAttribute", "overflow", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "overflow", value)
	return e
}

// PointerEvents
//
// English:
//
// The pointer-events attribute is a presentation attribute that allows defining whether or when an element may be the
// target of a mouse event.
//
//	Input:
//	  value: defining whether or when an element may be the target of a mouse event
//	    const: KSvgPointerEvents... (e.g. KSvgPointerEventsVisibleStroke)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute pointer-events can be used as a CSS property.
//
// Português:
//
// O atributo pointer-events é um atributo de apresentação que permite definir se ou quando um elemento pode ser alvo
// de um evento de mouse.
//
//	Entrada:
//	  value: define se ou quando um elemento pode ser alvo de um evento de mouse.
//	    const: KSvgPointerEvents... (e.g. KSvgPointerEventsVisibleStroke)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um atributo de apresentação, os eventos de ponteiro podem ser usados como uma propriedade CSS.
func (e *TagSvgAnimateMotion) PointerEvents(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgPointerEvents); ok {
		e.selfElement.Call("setAttribute", "pointer-events", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "pointer-events", value)
	return e
}

// ShapeRendering
//
// English:
//
// The shape-rendering attribute provides hints to the renderer about what tradeoffs to make when rendering shapes like
// paths, circles, or rectangles.
//
//	Input:
//	  value: provides hints to the renderer
//	    const: KSvgShapeRendering... (e.g. KSvgShapeRenderingAuto)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, shape-rendering can be used as a CSS property.
//
// Português:
//
// O atributo shape-rendering fornece dicas ao renderizador sobre quais compensações fazer ao renderizar formas como
// caminhos, círculos ou retângulos.
//
//	Entrada:
//	  value: fornece dicas para o renderizador
//	    const: KSvgShapeRendering... (ex. KSvgShapeRenderingAuto)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um atributo de apresentação, a renderização de forma pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateMotion) ShapeRendering(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgShapeRendering); ok {
		e.selfElement.Call("setAttribute", "shape-rendering", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "shape-rendering", value)
	return e
}

// StopColor
//
// English:
//
//	The stop-color attribute indicates what color to use at a gradient stop.
//
//	 Input:
//	   value: indicates what color to use at a gradient stop
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//
//	 Notes:
//	   * With respect to gradients, SVG treats the transparent keyword differently than CSS. SVG does not calculate
//	     gradients in pre-multiplied space, so transparent really means transparent black. So, specifying a stop-color
//	     with the value transparent is equivalent to specifying a stop-color with the value black and a stop-opacity
//	     with the value 0.
//	   * As a presentation attribute, stop-color can be used as a CSS property.
//
// Português:
//
//	O atributo stop-color indica qual cor usar em uma parada de gradiente.
//
//	 Entrada:
//	   value: indica qual cor usar em um fim de gradiente
//	     string: e.g. "black"
//	     factory: e.g. factoryColor.NewYellow()
//	     RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//
//	 Notss:
//	   * Com relação aos gradientes, o SVG trata a palavra-chave transparente de maneira diferente do CSS. O SVG não
//	     calcula gradientes no espaço pré-multiplicado, portanto, transparente realmente significa preto transparente.
//	     Assim, especificar uma stop-color com o valor transparente é equivalente a especificar uma stop-color com o
//	     valor black e uma stop-opacity com o valor 0.
//	   * Como atributo de apresentação, stop-color pode ser usado como propriedade CSS.
func (e *TagSvgAnimateMotion) StopColor(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "stop-color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "stop-color", value)
	return e
}

// StopOpacity
//
// English:
//
// The stop-opacity attribute defines the opacity of a given color gradient stop.
//
//	Input:
//	  value: defines the opacity of a given color gradient stop
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// The opacity value used for the gradient calculation is the product of the value of stop-opacity and the opacity of
// the value of the stop-color attribute. For stop-color values that don't include explicit opacity information, the
// opacity is treated as 1.
//
//	Notes:
//	  * As a presentation attribute, stop-opacity can be used as a CSS property.
//
// Português:
//
// O atributo stop-opacity define a opacidade de uma determinada parada de gradiente de cor.
//
//	Entrada:
//	  value: define a opacidade de uma determinada parada de gradiente de cor
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// O valor de opacidade usado para o cálculo do gradiente é o produto do valor de stop-opacity e a opacidade do valor
// do atributo stop-color. Para valores de stop-color que não incluem informações explícitas de opacidade, a opacidade
// é tratada como 1.
//
//	Notas:
//	  * Como atributo de apresentação, stop-opacity pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) StopOpacity(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "stop-opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "stop-opacity", value)
	return e
}

// Stroke
//
// English:
//
// The stroke attribute is a presentation attribute defining the color (or any SVG paint servers like gradients or
// patterns) used to paint the outline of the shape
//
//	Input:
//	  value: presentation attribute defining the color
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute stroke can be used as a CSS property.
//
// Português:
//
// O atributo de traço é um atributo de apresentação que define a cor (ou qualquer servidor de pintura SVG, como
// gradientes ou padrões) usado para pintar o contorno da forma
//
//	Entrada:
//	  value: atributo de apresentação que define a cor
//	    string: e.g. "black"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como um traço de atributo de apresentação pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) Stroke(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "stroke", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "stroke", value)
	return e
}

// StrokeDasharray
//
// English:
//
// The stroke-dasharray attribute is a presentation attribute defining the pattern of dashes and gaps used to paint the
// outline of the shape
//
//	Input:
//	  value: presentation attribute defining the pattern of dashes
//	    []float64: (e.g. []float64{4, 1, 2}) = "4 1 2"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, stroke-dasharray can be used as a CSS property.
//
// Português:
//
// O atributo stroke-dasharray é um atributo de apresentação que define o padrão de traços e lacunas usados para pintar
// o contorno da forma
//
//	Entrada:
//	  value: atributo de apresentação que define o padrão de traços
//	    []float64: (e.g. []float64{4, 1, 2}) = "4 1 2"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o stroke-dasharray pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) StrokeDasharray(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.([]float64); ok {
		str := ""
		for _, v := range converted {
			str += strconv.FormatFloat(v, 'g', -1, 64) + " "
		}
		length := len(str) - 1

		e.selfElement.Call("setAttribute", "stroke-dasharray", str[:length])
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-dasharray", value)
	return e
}

// StrokeLineCap
//
// English:
//
// The stroke-linecap attribute is a presentation attribute defining the shape to be used at the end of open subpaths
// when they are stroked.
//
//	Input:
//	  value: presentation attribute defining the shape to be used at the end of open subpaths
//	    const: KSvgStrokeLinecap... (e.g. KSvgStrokeLinecapRound)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute stroke-linecap can be used as a CSS property.
//
// Português:
//
// O atributo stroke-linecap é um atributo de apresentação que define a forma a ser usada no final de subcaminhos
// abertos quando eles são traçados.
//
//	Input:
//	  value: atributo de apresentação que define a forma a ser usada no final de subcaminhos
//	    const: KSvgStrokeLinecap... (e.g. KSvgStrokeLinecapRound)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o traço-linecap pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) StrokeLineCap(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgStrokeLinecap); ok {
		e.selfElement.Call("setAttribute", "stroke-linecap", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-linecap", value)
	return e
}

// StrokeLineJoin
//
// English:
//
// The stroke-linejoin attribute is a presentation attribute defining the shape to be used at the corners of paths when
// they are stroked.
//
//	Notes:
//	  * As a presentation attribute stroke-linejoin can be used as a CSS property.
//
// Português:
//
// O atributo stroke-linejoin é um atributo de apresentação que define a forma a ser usada nos cantos dos caminhos
// quando eles são traçados.
//
//	Notas:
//	  * Como atributo de apresentação, stroke-linejoin pode ser usado como propriedade CSS.
func (e *TagSvgAnimateMotion) StrokeLineJoin(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgStrokeLinejoin); ok {
		e.selfElement.Call("setAttribute", "stroke-linejoin", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-linejoin", value)
	return e
}

// StrokeMiterLimit
//
// English:
//
// The stroke-miterlimit attribute is a presentation attribute defining a limit on the ratio of the miter length to the
// stroke-width used to draw a miter join. When the limit is exceeded, the join is converted from a miter to a bevel.
//
//	Input:
//	  value: defining a limit on the ratio of the miter length
//
//	Notes:
//	  * As a presentation attribute stroke-miterlimit can be used as a CSS property.
//
// Português:
//
// O atributo stroke-miterlimit é um atributo de apresentação que define um limite na proporção do comprimento da mitra
// para a largura do traço usado para desenhar uma junção de mitra. Quando o limite é excedido, a junção é convertida
// de uma mitra para um chanfro.
//
//	Entrada:
//	  value: definindo um limite na proporção do comprimento da mitra
//
//	Notas:
//	  * Como atributo de apresentação, stroke-miterlimit pode ser usado como propriedade CSS.
func (e *TagSvgAnimateMotion) StrokeMiterLimit(value float64) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "stroke-miterlimit", value)
	return e
}

// StrokeOpacity
//
// English:
//
// The stroke-opacity attribute is a presentation attribute defining the opacity of the paint server (color, gradient,
// pattern, etc) applied to the stroke of a shape.
//
//	Input:
//	  value: defining the opacity of the paint
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute stroke-opacity can be used as a CSS property.
//
// Português:
//
// O atributo de opacidade do traçado é um atributo de apresentação que define a opacidade do servidor de pintura (cor,
// gradiente, padrão etc.) aplicado ao traçado de uma forma.
//
//	Entrada:
//	  value: definindo a opacidade da tinta
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, a opacidade do traço pode ser usada como uma propriedade CSS.
func (e *TagSvgAnimateMotion) StrokeOpacity(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "stroke-opacity", p)
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-opacity", value)
	return e
}

// StrokeWidth
//
// English:
//
// The stroke-width attribute is a presentation attribute defining the width of the stroke to be applied to the shape.
//
//	Input:
//	  value: defining the width of the stroke
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// Português:
//
// O atributo stroke-width é um atributo de apresentação que define a largura do traço a ser aplicado à forma.
//
//	Entrada:
//	  value: definindo a largura do traço
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
func (e *TagSvgAnimateMotion) StrokeWidth(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "stroke-width", p)
		return e
	}

	e.selfElement.Call("setAttribute", "stroke-width", value)
	return e
}

// TextAnchor
//
// English:
//
// The text-anchor attribute is used to align (start-, middle- or end-alignment) a string of pre-formatted text or
// auto-wrapped text where the wrapping area is determined from the inline-size property relative to a given point.
//
//	Input:
//	  value: used to align a string
//	    const: KSvgTextAnchor... (e.g. KSvgTextAnchorStart)
//	    any other type: interface{}
//
// This attribute is not applicable to other types of auto-wrapped text. For those cases you should use text-align.
// For multi-line text, the alignment takes place for each line.
//
// The text-anchor attribute is applied to each individual text chunk within a given <text> element. Each text chunk
// has an initial current text position, which represents the point in the user coordinate system resulting from
// (depending on context) application of the x and y attributes on the <text> element, any x or y attribute values on a
// <tspan>, <tref> or <altGlyph> element assigned explicitly to the first rendered character in a text chunk, or
// determination of the initial current text position for a <textPath> element.
//
//	Notes:
//	  * As a presentation attribute, text-anchor can be used as a CSS property.
//
// Português:
//
// O atributo text-anchor é usado para alinhar (alinhamento inicial, intermediário ou final) uma string de texto
// pré-formatado ou texto com quebra automática onde a área de quebra é determinada a partir da propriedade inline-size
// relativa a um determinado ponto.
//
//	Entrada:
//	  value: usado para alinhar uma string
//	    const: KSvgTextAnchor... (e.g. KSvgTextAnchorStart)
//	    qualquer outro tipo: interface{}
//
// Este atributo não se aplica a outros tipos de texto com quebra automática. Para esses casos, você deve usar
// text-align. Para texto de várias linhas, o alinhamento ocorre para cada linha.
//
// O atributo text-anchor é aplicado a cada fragmento de texto individual dentro de um determinado elemento <text>.
// Cada pedaço de texto tem uma posição inicial de texto atual, que representa o ponto no sistema de coordenadas do
// usuário resultante (dependendo do contexto) da aplicação dos atributos x e y no elemento <text>, quaisquer valores
// de atributo x ou y em um <tspan >, elemento <tref> ou <altGlyph> atribuído explicitamente ao primeiro caractere
// renderizado em um pedaço de texto, ou determinação da posição inicial do texto atual para um elemento <textPath>.
//
//	Notes:
//	  * As a presentation attribute, text-anchor can be used as a CSS property.
func (e *TagSvgAnimateMotion) TextAnchor(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgTextAnchor); ok {
		e.selfElement.Call("setAttribute", "text-anchor", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "text-anchor", value)
	return e
}

// TextDecoration
//
// English:
//
// The text-decoration attribute defines whether text is decorated with an underline, overline and/or strike-through.
// It is a shorthand for the text-decoration-line and text-decoration-style properties.
//
//	Input:
//	  value: defines whether text is decorated
//	    const: KSvgTextDecorationLine... (e.g. KSvgTextDecorationLineUnderline)
//	    const: KSvgTextDecorationStyle... (e.g. KSvgTextDecorationStyleDouble)
//	    string: e.g. "black", "line-through"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    any other type: interface{}
//
// The fill and stroke of the text decoration are given by the fill and stroke of the text at the point where the text
// decoration is declared.
//
// The paint order of the text decoration, i.e. the fill and stroke, is determined by the value of the paint-order
// attribute at the point where the text decoration is declared.
//
//	Notes:
//	  * As a presentation attribute, text-decoration can be used as a CSS property. See the css text-decoration
//	    property for more information.
//
// Português:
//
// O atributo text-decoration define se o texto é decorado com sublinhado, overline e ou tachado.
// É um atalho para as propriedades text-decoration-line e text-decoration-style.
//
//	Entrada:
//	  value: define se o texto é decorado
//	    const: KSvgTextDecorationLine... (ex. KSvgTextDecorationLineUnderline)
//	    const: KSvgTextDecorationStyle... (ex. KSvgTextDecorationStyleDouble)
//	    string: e.g. "black", "line-through"
//	    factory: e.g. factoryColor.NewYellow()
//	    RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//	    qualquer outro tipo: interface{}
//
// O preenchimento e o traçado da decoração de texto são dados pelo preenchimento e traçado do texto no ponto em que a
// decoração de texto é declarada.
//
// A ordem de pintura da decoração do texto, ou seja, o preenchimento e o traço, é determinada pelo valor do atributo
// paint-order no ponto em que a decoração do texto é declarada.
//
//	Notas:
//	  * Como atributo de apresentação, a decoração de texto pode ser usada como uma propriedade CSS. Consulte a
//	    propriedade CSS text-decoration para obter mais informações.
func (e *TagSvgAnimateMotion) TextDecoration(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "text-decoration", RGBAToJs(converted))
		return e
	}

	if converted, ok := value.(SvgTextDecorationLine); ok {
		e.selfElement.Call("setAttribute", "text-decoration", converted.String())
		return e
	}

	if converted, ok := value.(SvgTextDecorationStyle); ok {
		e.selfElement.Call("setAttribute", "text-decoration", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "text-decoration", value)
	return e
}

// TextRendering
//
// English:
//
// The text-rendering attribute provides hints to the renderer about what tradeoffs to make when rendering text.
//
//	Notes:
//	  * As a presentation attribute, text-rendering can be used as a CSS property.
//	    See the css text-rendering property for more information.
//
// Português:
//
// O atributo text-rendering fornece dicas ao renderizador sobre quais compensações fazer ao renderizar o texto.
//
//	Notas:
//	  * Como um atributo de apresentação, a renderização de texto pode ser usada como uma propriedade CSS.
//	    Consulte a propriedade de renderização de texto css para obter mais informações.
func (e *TagSvgAnimateMotion) TextRendering(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgTextRendering); ok {
		e.selfElement.Call("setAttribute", "text-rendering", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "text-rendering", value)
	return e
}

// Transform
//
// English:
//
// The transform attribute defines a list of transform definitions that are applied to an element and the element's
// children.
//
//	Input:
//	  value: defines a list of transform definitions
//	    factory: e.g. factoryBrowser.NewTransform().Translate(100, 0).Scale(4, 1)
//	    string: e.g. "translate(300,0) scale(4,1)"
//	    any other type: interface{}
//
//	Notes:
//	  * As of SVG2, transform is a presentation attribute, meaning it can be used as a CSS property. However, be aware
//	    that there are some differences in syntax between the CSS property and the attribute. See the documentation for
//	    the CSS property transform for the specific syntax to use in that case.
//
// Português:
//
// O atributo transform define uma lista de definições de transformação que são aplicadas a um elemento e aos filhos do
// elemento.
//
//	Entrada:
//	  value: define uma lista de definições de transformação
//	    factory: ex. factoryBrowser.NewTransform().Translate(100, 0).Scale(4, 1)
//	    string: ex. "translate(300,0) scale(4,1)"
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * A partir do SVG2, transform é um atributo de apresentação, o que significa que pode ser usado como uma
//	    propriedade CSS. No entanto, esteja ciente de que existem algumas diferenças na sintaxe entre a propriedade CSS
//	    e o atributo. Consulte a documentação da transformação da propriedade CSS para obter a sintaxe específica a ser
//	    usada nesse caso.
func (e *TagSvgAnimateMotion) Transform(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(*TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "transform", converted.String())
		return e
	}

	if converted, ok := value.(TransformFunctions); ok {
		e.selfElement.Call("setAttribute", "transform", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "transform", value)
	return e
}

// UnicodeBidi
//
// English:
//
// The unicode-bidi attribute specifies how the accumulation of the background image is managed.
//
//	Input:
//	  value: specifies how the accumulation of the background image is managed
//	    const: KSvgTransformOrigin... (e.g. KSvgTransformOriginLeft)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, unicode-bidi can be used as a CSS property. See the CSS unicode-bidi property for
//	    more information.
//
// Português:
//
// O atributo unicode-bidi especifica como o acúmulo da imagem de fundo é gerenciado.
//
//	Entrada:
//	  value: especifica como o acúmulo da imagem de fundo é gerenciado
//	    const: KSvgTransformOrigin... (e.g. KSvgTransformOriginLeft)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o unicode-bidi pode ser usado como uma propriedade CSS. Consulte a propriedade
//	    CSS unicode-bidi para obter mais informações.
func (e *TagSvgAnimateMotion) UnicodeBidi(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgTransformOrigin); ok {
		e.selfElement.Call("setAttribute", "unicode-bidi", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "unicode-bidi", value)
	return e
}

// VectorEffect
//
// English:
//
// The vector-effect property specifies the vector effect to use when drawing an object.
//
//	Input:
//	  value: specifies the vector effect
//	    const: KSvgVectorEffect... (e.g. KSvgVectorEffectNonScalingStroke)
//
// Vector effects are applied before any of the other compositing operations, i.e. filters, masks and clips.
//
//	Notes:
//	  * As a presentation attribute, vector-effect can be used as a CSS property.
//
// Português:
//
// A propriedade vector-effect especifica o efeito vetorial a ser usado ao desenhar um objeto.
//
//	Entrada:
//	  value: especifica o efeito vetorial
//	    const: KSvgVectorEffect... (ex. KSvgVectorEffectNonScalingStroke)
//
// Os efeitos vetoriais são aplicados antes de qualquer outra operação de composição, ou seja, filtros, máscaras e
// clipes.
//
//	Notas:
//	  * Como atributo de apresentação, o efeito vetorial pode ser usado como uma propriedade CSS.
func (e *TagSvgAnimateMotion) VectorEffect(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgVectorEffect); ok {
		e.selfElement.Call("setAttribute", "vector-effect", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "vector-effect", value)
	return e
}

// Visibility
//
// English:
//
// The visibility attribute lets you control the visibility of graphical elements.
//
//	Input:
//	  value: lets you control the visibility
//	    const: KSvgVisibility... (e.g. KSvgVisibilityHidden)
//	    any other type: interface{}
//
// With a value of hidden or collapse the current graphics element is invisible.
//
// Depending on the value of attribute pointer-events, graphics elements which have their visibility attribute set to
// hidden still might receive events.
//
//	Notes:
//	  * If the visibility attribute is set to hidden on a text element, then the text is invisible but still takes up
//	    space in text layout calculations;
//	  * As a presentation attribute, visibility can be used as a CSS property. See the css visibility property for
//	    more information.
//
// Português:
//
// O atributo de visibilidade permite controlar a visibilidade dos elementos gráficos.
//
//	Entrada:
//	  value: permite controlar a visibilidade
//	    const: KSvgVisibility... (e.g. KSvgVisibilityHidden)
//	    qualquer outro tipo: interface{}
//
// Com um valor oculto ou recolhido, o elemento gráfico atual fica invisível.
//
// Dependendo do valor do atributo pointer-events, os elementos gráficos que têm seu atributo de visibilidade definido
// como oculto ainda podem receber eventos.
//
//	Notas:
//	  * Se o atributo de visibilidade estiver definido como oculto em um elemento de texto, o texto ficará invisível,
//	    mas ainda ocupará espaço nos cálculos de layout de texto;
//	  * Como atributo de apresentação, a visibilidade pode ser usada como propriedade CSS. Consulte a propriedade de
//	    visibilidade do CSS para obter mais informações.
func (e *TagSvgAnimateMotion) Visibility(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgVisibility); ok {
		e.selfElement.Call("setAttribute", "visibility", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "visibility", value)
	return e
}

// WordSpacing
//
// English:
//
// The word-spacing attribute specifies spacing behavior between words.
//
//	Input:
//	  value: specifies spacing behavior between words
//	    float32: 1.0 = "100%"
//	    any other type: interface{}
//
// If a <length> is provided without a unit identifier (e.g. an unqualified number such as 128), the browser processes
// the <length> as a width value in the current user coordinate system.
//
// If a <length> is provided with one of the unit identifiers (e.g. .25em or 1%), then the browser converts the <length>
// into a corresponding value in the current user coordinate system.
//
//	Notes:
//	  * As a presentation attribute, word-spacing can be used as a CSS property. See the css word-spacing property for
//	    more information.
//
// Português:
//
// O atributo word-spacing especifica o comportamento do espaçamento entre as palavras.
//
//	Entrada:
//	  value: especifica o comportamento de espaçamento entre palavras
//	    float32: 1.0 = "100%"
//	    qualquer outro tipo: interface{}
//
// Se um <comprimento> for fornecido sem um identificador de unidade (por exemplo, um número não qualificado como 128),
// o navegador processará o <comprimento> como um valor de largura no sistema de coordenadas do usuário atual.
//
// Se um <comprimento> for fornecido com um dos identificadores de unidade (por exemplo, .25em ou 1%), o navegador
// converterá o <comprimento> em um valor correspondente no sistema de coordenadas do usuário atual.
//
//	Notas:
//	  * Como atributo de apresentação, o espaçamento entre palavras pode ser usado como uma propriedade CSS.
//	    Consulte a propriedade de espaçamento entre palavras do CSS para obter mais informações.
func (e *TagSvgAnimateMotion) WordSpacing(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "word-spacing", p)
		return e
	}

	e.selfElement.Call("setAttribute", "word-spacing", value)
	return e
}

// WritingMode
//
// English:
//
// The writing-mode attribute specifies whether the initial inline-progression-direction for a <text> element shall be
// left-to-right, right-to-left, or top-to-bottom. The writing-mode attribute applies only to <text> elements;
// the attribute is ignored for <tspan>, <tref>, <altGlyph> and <textPath> sub-elements. (Note that the
// inline-progression-direction can change within a <text> element due to the Unicode bidirectional algorithm and
// properties direction and unicode-bidi.)
//
//	Input:
//	  value: specifies whether the initial inline-progression-direction
//	    const: KSvgWritingMode... (e.g. KSvgWritingModeHorizontalTb)
//	    any other type: interface{}
//
//	Notes:
//	  * As a presentation attribute, writing-mode can be used as a CSS property. See the CSS writing-mode property for
//	    more information.
//
// Português:
//
// O atributo write-mode especifica se a direção de progressão inline inicial para um elemento <text> deve ser da
// esquerda para a direita, da direita para a esquerda ou de cima para baixo. O atributo write-mode aplica-se apenas a
// elementos <text>; o atributo é ignorado para os subelementos <tspan>, <tref>, <altGlyph> e <textPath>.
// (Observe que a direção de progressão em linha pode mudar dentro de um elemento <text> devido ao algoritmo
// bidirecional Unicode e direção de propriedades e unicode-bidi.)
//
//	Entrada:
//	  value: especifica se a direção de progressão em linha inicial
//	    const: KSvgWritingMode... (ex. KSvgWritingModeHorizontalTb)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	  * Como atributo de apresentação, o modo de escrita pode ser usado como uma propriedade CSS. Consulte a
//	    propriedade do modo de gravação CSS para obter mais informações.
func (e *TagSvgAnimateMotion) WritingMode(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgWritingMode); ok {
		e.selfElement.Call("setAttribute", "writing-mode", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "writing-mode", value)
	return e
}

// #presentation end --------------------------------------------------------------------------------------------------

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
func (e *TagSvgAnimateMotion) Class(class string) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Style(value string) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "style", value)
	return e
}

// #styling end -------------------------------------------------------------------------------------------------------

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
func (e *TagSvgAnimateMotion) Begin(begin interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Dur(dur interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) End(end interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Min(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Max(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Restart(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) RepeatCount(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) RepeatDur(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "repeatDur", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "repeatDur", value)
	return e
}

// Rotate
//
// English:
//
// The rotate attribute specifies how the animated element rotates as it travels along a path specified in an
// <animateMotion> element.
//
//	Input:
//	  value: specifies how the animated element rotates
//	    const: KSvgRotate... (e.g. KSvgRotateAutoReverse)
//	    any other type: interface{}
//
// Português:
//
// O atributo de rotação especifica como o elemento animado gira enquanto percorre um caminho especificado em um
// elemento <animateMotion>.
//
//	Entrada:
//	  value: especifica como o elemento animado gira
//	    const: KSvgRotate... (e.g. KSvgRotateAutoReverse)
//	    qualquer outro tipo: interface{}
func (e *TagSvgAnimateMotion) Rotate(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgRotate); ok {
		e.selfElement.Call("setAttribute", "rotate", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "rotate", value)
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
func (e *TagSvgAnimateMotion) Fill(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) CalcMode(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Values(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) KeyTimes(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) KeySplines(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) From(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "from", TypeToString(value, ";", ";"))
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
func (e *TagSvgAnimateMotion) To(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "to", TypeToString(value, ";", ";"))
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
func (e *TagSvgAnimateMotion) By(by float64) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "by", by)
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
func (e *TagSvgAnimateMotion) AttributeName(attributeName string) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "attributeName", attributeName)
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
func (e *TagSvgAnimateMotion) Additive(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Accumulate(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(SvgAccumulate); ok {
		e.selfElement.Call("setAttribute", "accumulate", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "accumulate", value)
	return e
}

// KeyPoints
//
// English:
//
// The keyPoints attribute indicates the simple duration of an animation.
//
//	Input:
//	  keyPoints: indicates the simple duration of an animation.
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1);rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%;10%"
//	    []float64: []float64{0.0, 10.0} = "0;10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s;1s"
//	    any other type: interface{}
//
// This value defines a semicolon-separated list of floating point values between 0 and 1 and indicates how far along
// the motion path the object shall move at the moment in time specified by corresponding keyTimes value. The distance
// is calculated along the path specified by the path attribute. Each progress value in the list corresponds to a value
// in the keyTimes attribute list.
// If a list of key points is specified, there must be exactly as many values in the keyPoints list as in the keyTimes
// list.
// If there's a semicolon at the end of the value, optionally followed by white space, both the semicolon and the
// trailing white space are ignored.
// If there are any errors in the value specification (i.e. bad values, too many or too few values), then that's an
// error.
//
// Português:
//
// O atributo keyPoints indica a duração simples de uma animação.
//
//	Entrada:
//	  keyPoints: indica a duração simples de uma animação.
//	    []color.RGBA{factoryColor.NewBlack(),factoryColor.NewRed()} = "rgba(0,0,0,1);rgba(255,0,0,1)"
//	    []float32: []float64{0.0, 0.1} = "0%;10%"
//	    []float64: []float64{0.0, 10.0} = "0;10"
//	    []time.Duration: []time.Duration{0, time.Second} = "0s;1s"
//	    any other type: interface{}
//
// Este valor define uma lista separada por ponto e vírgula de valores de ponto flutuante entre 0 e 1 e indica a
// distância ao longo do caminho de movimento o objeto deve se mover no momento especificado pelo valor keyTimes
// correspondente. A distância é calculada ao longo do caminho especificado pelo atributo path. Cada valor de progresso
// na lista corresponde a um valor na lista de atributos keyTimes.
// Se uma lista de pontos-chave for especificada, deve haver exatamente tantos valores na lista keyPoints quanto na
// lista keyTimes.
// Se houver um ponto e vírgula no final do valor, opcionalmente seguido por espaço em branco, o ponto e vírgula e o
// espaço em branco à direita serão ignorados.
// Se houver algum erro na especificação do valor (ou seja, valores incorretos, muitos ou poucos valores), isso é um
// erro.
func (e *TagSvgAnimateMotion) KeyPoints(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "keyPoints", TypeToString(value, ";", ";"))
	return e
}

// Path
//
// English:
//
// The path attribute has two different meanings, either it defines a text path along which the characters of a text
// are rendered, or a motion path along which a referenced element is animated.
//
//	Input:
//	  value: defines a text path along which the characters of a text are rendered, or a motion path along which a
//	      referenced element is animated
//	    factory: factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()
//	    string: "M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z"
//	    any other type: interface{}
//
// Português:
//
// O atributo path tem dois significados diferentes: define um caminho de texto ao longo do qual os caracteres de um
// texto são renderizados ou um caminho de movimento ao longo do qual um elemento referenciado é animado.
//
//	Entrada:
//	  value: define um caminho de texto ao longo do qual os caracteres de um texto são renderizados ou um caminho de
//	      movimento ao longo do qual um elemento referenciado é animado
//	    factory: factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()
//	    string: "M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z"
//	    qualquer outro tipo: interface{}
func (e *TagSvgAnimateMotion) Path(value interface{}) (ref *TagSvgAnimateMotion) {
	if converted, ok := value.(*SvgPath); ok {
		e.selfElement.Call("setAttribute", "path", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "path", value)
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
func (e *TagSvgAnimateMotion) HRef(href string) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) XLinkHRef(value interface{}) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Text(value string) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) Html(value string) (ref *TagSvgAnimateMotion) {
	e.selfElement.Set("innerHTML", value)
	return e
}

// Origin
//
// English:
//
//	Input:
//	  value: specifies the origin of motion for an animation
//
// The origin attribute specifies the origin of motion for an animation. It has no effect in SVG.
//
// Português:
//
// O atributo origin especifica a origem do movimento de uma animação. Não tem efeito em SVG.
//
//	Entrada:
//	  value: especifica a origem do movimento de uma animação
func (e *TagSvgAnimateMotion) Origin(value interface{}) (ref *TagSvgAnimateMotion) {
	e.selfElement.Call("setAttribute", "origin", value)
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
func (e *TagSvgAnimateMotion) Reference(reference **TagSvgAnimateMotion) (ref *TagSvgAnimateMotion) {
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
func (e *TagSvgAnimateMotion) AddListenerBegin(animationEvent *chan animation.Data) (ref *TagSvgAnimateMotion) {
	var fn js.Func

	if e.fnBegin == nil {
		fn = js.FuncOf(func(this js.Value, _ []js.Value) interface{} {
			if this.IsNull() == true || this.IsUndefined() == true {
				return nil
			}

			*animationEvent <- animation.EventManager(animation.KEventBegin, this)
			return nil
		})
		e.fnBegin = &fn
	}
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
func (e *TagSvgAnimateMotion) RemoveListenerBegin() (ref *TagSvgAnimateMotion) {
	if e.fnBegin == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"beginEvent",
		*e.fnBegin,
	)
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
//	  * Em 07/2022, este evento não funciona no safari da apple
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
//	  * On 072022 this event doesn't work on apple safari
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
func (e *TagSvgAnimateMotion) AddListenerRepeat(animationEvent *chan animation.Data) (ref *TagSvgAnimateMotion) {
	var fn js.Func

	if e.fnRepeat == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if this.IsNull() == true || this.IsUndefined() == true {
				return nil
			}

			*animationEvent <- animation.EventManager(animation.KEventRepeat, this)
			return nil
		})
		e.fnRepeat = &fn
	}
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
func (e *TagSvgAnimateMotion) RemoveListenerRepeat() (ref *TagSvgAnimateMotion) {
	if e.fnRepeat == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"repeatEvent",
		*e.fnRepeat,
	)
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
func (e *TagSvgAnimateMotion) AddListenerEnd(animationEvent *chan animation.Data) (ref *TagSvgAnimateMotion) {
	var fn js.Func

	if e.fnEnd == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if this.IsNull() == true || this.IsUndefined() == true {
				return nil
			}

			*animationEvent <- animation.EventManager(animation.KEventEnd, this)
			return nil
		})
		e.fnEnd = &fn
	}

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
func (e *TagSvgAnimateMotion) RemoveListenerEnd() (ref *TagSvgAnimateMotion) {
	if e.fnEnd == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"endEvent",
		*e.fnEnd,
	)
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
func (e *TagSvgAnimateMotion) AddListenerMotion(animationEvent *chan animation.Data) (ref *TagSvgAnimateMotion) {
	var fn js.Func

	if e.fnMotion == nil {
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
					*animationEvent <- data
				},
			)
			return nil
		})
		e.fnMotion = &fn
	}
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
func (e *TagSvgAnimateMotion) RemoveListenerMotion() (ref *TagSvgAnimateMotion) {
	if e.fnMotion == nil {
		return
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
