package html

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"image/color"
	"log"
	"sync"
	"syscall/js"
)

type TagSvgRect struct {

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
func (e *TagSvgRect) Init(id string) (ref *TagSvgRect) {
	e.listener = new(sync.Map)

	e.CreateElement(KTagSvgRect)
	e.prepareStageReference()
	e.Id(id)

	return e
}

func (e *TagSvgRect) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgRect) Id(id string) (ref *TagSvgRect) {
	e.id = id
	e.selfElement.Set("id", id)
	return e
}

func (e *TagSvgRect) CreateElement(tag Tag) (ref *TagSvgRect) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	return e
}

func (e *TagSvgRect) ViewBox(minX, minY, width, height float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "viewBox", fmt.Sprintf("%v %v %v %v", minX, minY, width, height))
	return e
}

// X
//
// English:
//
//  The x attribute defines an x-axis coordinate in the user coordinate system.
//
// Português:
//
//  O atributo x define uma coordenada do eixo x no sistema de coordenadas do usuário.
func (e *TagSvgRect) X(x float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "x", x)
	return e
}

// Y
//
// English:
//
//  The y attribute defines a y-axis coordinate in the user coordinate system.
//
// Português:
//
//  O atributo y define uma coordenada do eixo y no sistema de coordenadas do usuário.
func (e *TagSvgRect) Y(y float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "y", y)
	return e
}

// Width
//
// English:
//
//  The width attribute defines the horizontal length of an element in the user coordinate system.
//
// Português:
//
//  O atributo largura define o comprimento horizontal de um elemento no sistema de coordenadas do usuário.
func (e *TagSvgRect) Width(width float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "width", width)
	return e
}

// Height
//
// English:
//
//  The height attribute defines the vertical length of an element in the user coordinate system.
//
// Português:
//
//  O atributo height define o comprimento vertical de um elemento no sistema de coordenadas do usuário.
func (e *TagSvgRect) Height(height float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "height", height)
	return e
}

// Rx
//
// English:
//
//  The rx attribute defines a radius on the x-axis.
//
// Português:
//
//  O atributo rx define um raio no eixo x.
func (e *TagSvgRect) Rx(rx float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "rx", rx)
	return e
}

// Ry
//
// English:
//
//  The ry attribute defines a radius on the y-axis.
//
// Português:
//
//  O atributo ry define um raio no eixo y.
func (e *TagSvgRect) Ry(ry float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "ry", ry)
	return e
}

// PathLength
//
// English:
//
//  The pathLength attribute lets authors specify a total length for the path, in user units. This value is then used
//  to calibrate the browser's distance calculations with those of the author, by scaling all distance computations
//  using the ratio pathLength/(computed value of path length).
//
// This can affect the actual rendered lengths of paths; including text paths, animation paths, and various stroke
// operations. Basically, all computations that require the length of the path. stroke-dasharray, for example, will
// assume the start of the path being 0 and the end point the value defined in the pathLength attribute.
//
//   Notes:
//     * Starting with SVG2, x, y, width, height, rx and ry are Geometry Properties, meaning those attributes can
//       also be used as CSS properties for that element.
//
// Português:
//
//  O atributo pathLength permite que os autores especifiquem um comprimento total para o caminho, em unidades de
//  usuário. Este valor é então usado para calibrar os cálculos de distância do navegador com os do autor, escalando
//  todos os cálculos de distância usando a razão pathLength (valor calculado do comprimento do caminho).
//
// Isso pode afetar os comprimentos reais dos caminhos renderizados; incluindo caminhos de texto, caminhos de animação
// e várias operações de traçado. Basicamente, todos os cálculos que exigem o comprimento do caminho. stroke-dasharray,
// por exemplo, assumirá o início do caminho sendo 0 e o ponto final o valor definido no atributo pathLength.
//
//   Notas:
//     * Começando com SVG2, x, y, largura, altura, rx e ry são Propriedades de Geometria, o que significa que esses
//       atributos também podem ser usados como propriedades CSS para aquele elemento.
func (e *TagSvgRect) PathLength(pathLength float64) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "pathLength", pathLength)
	return e
}

// BaselineShift
//
// English:
//
//  It allows repositioning of the dominant-baseline relative to the dominant-baseline of the parent text content
//  element.
//
// Português:
//
//  Ele permite o reposicionamento da linha de base dominante em relação à linha de base dominante do elemento de
//  conteúdo de texto pai.
func (e *TagSvgRect) BaselineShift(baselineShift interface{}) (ref *TagSvgRect) {
	if converted, ok := baselineShift.(SvgBaselineShift); ok {
		e.selfElement.Call("setAttribute", "baseline-shift", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "baseline-shift", baselineShift)
	return e
}

// ClipPath
//
// English:
//
//  It binds the element it is applied to with a given <clipPath> element.
//
// Português:
//
//  Ele associa o elemento ao qual é aplicado a um determinado elemento <clipPath>.
func (e *TagSvgRect) ClipPath(clipPath string) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "clip-path", clipPath)
	return e
}

// ClipRule
//
// English:
//
//  It indicates how to determine what side of a path is inside a shape in order to know how a <clipPath> should clip
//  its target.
//
// Português:
//
//  Ele indica como determinar qual lado de um caminho está dentro de uma forma para saber como um <clipPath> deve
//  recortar seu destino.
func (e *TagSvgRect) ClipRule(clipRule string) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "clip-rule", clipRule)
	return e
}

// Color
//
// English:
//
//  It provides a potential indirect value (currentcolor) for the fill, stroke, stop-color, flood-color and
//  lighting-color presentation attributes.
//
// Português:
//
//  Ele fornece um valor indireto potencial (currentcolor) para os atributos de apresentação de preenchimento, traçado,
//  cor de parada, cor de inundação e cor de iluminação.
func (e *TagSvgRect) Color(value interface{}) (ref *TagSvgRect) {
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
//  The color-interpolation attribute specifies the color space for gradient interpolations, color animations, and alpha
//  compositing.
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
//   Notes:
//     * For filter effects, the color-interpolation-filters property controls which color space is used.
//     * As a presentation attribute, color-interpolation can be used as a CSS property.
//
// Português:
//
//  O atributo color-interpolation especifica o espaço de cores para interpolações de gradiente, animações de cores e
//  composição alfa.
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
//   Notas:
//     * Para efeitos de filtro, a propriedade color-interpolation-filters controla qual espaço de cor é usado.
//     * Como atributo de apresentação, a interpolação de cores pode ser usada como uma propriedade CSS.
func (e *TagSvgRect) ColorInterpolation(value interface{}) (ref *TagSvgRect) {
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
//  The color-interpolation-filters attribute specifies the color space for imaging operations performed via filter
//  effects.
//
//   Notes:
//     * This property just has an affect on filter operations. Therefore, it has no effect on filter primitives like
//       <feOffset>, <feImage>, <feTile> or <feFlood>;
//     * color-interpolation-filters has a different initial value than color-interpolation. color-interpolation-filters
//       has an initial value of linearRGB, whereas color-interpolation has an initial value of sRGB. Thus, in the
//       default case, filter effects operations occur in the linearRGB color space, whereas all other color
//       interpolations occur by default in the sRGB color space;
//     * It has no affect on filter functions, which operate in the sRGB color space;
//     * As a presentation attribute, color-interpolation-filters can be used as a CSS property.
//
// Português:
//
//  O atributo color-interpolation-filters especifica o espaço de cores para operações de imagem realizadas por meio de
//  efeitos de filtro.
//
//   Notas:
//     * Esta propriedade afeta apenas as operações de filtro. Portanto, não tem efeito em primitivos de filtro como
//       <feOffset>, <feImage>, <feTile> ou <feFlood>.
//     * color-interpolation-filters tem um valor inicial diferente de color-interpolation. color-interpolation-filters
//       tem um valor inicial de linearRGB, enquanto color-interpolation tem um valor inicial de sRGB. Assim, no caso
//       padrão, as operações de efeitos de filtro ocorrem no espaço de cores linearRGB, enquanto todas as outras
//       interpolações de cores ocorrem por padrão no espaço de cores sRGB.
//     * Não afeta as funções de filtro, que operam no espaço de cores sRGB.
//     * Como atributo de apresentação, os filtros de interpolação de cores podem ser usados como uma propriedade CSS.
func (e *TagSvgRect) ColorInterpolationFilters(value interface{}) (ref *TagSvgRect) {
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
//  The cursor attribute specifies the mouse cursor displayed when the mouse pointer is over an element.
//
// This attribute behaves exactly like the css cursor property except that if the browser supports the <cursor> element,
// you should be able to use it with the <funciri> notation.
//
// As a presentation attribute, it also can be used as a property directly inside a CSS stylesheet, see css cursor for
// further information.
//
// Português:
//
//  O atributo cursor especifica o cursor do mouse exibido quando o ponteiro do mouse está sobre um elemento.
//
// Este atributo se comporta exatamente como a propriedade cursor css, exceto que, se o navegador suportar o elemento
// <cursor>, você poderá usá-lo com a notação <funciri>.
//
// Como atributo de apresentação, também pode ser usado como propriedade diretamente dentro de uma folha de estilo CSS,
// veja cursor css para mais informações.
func (e *TagSvgRect) Cursor(cursor SvgCursor) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "cursor", cursor.String())
	return e
}

// D
//
// English:
//
//  The d attribute defines a path to be drawn.
//
// A path definition is a list of path commands where each command is composed of a command letter and numbers that
// represent the command parameters. The commands are detailed below.
//
// You can use this attribute with the following SVG elements: <path>, <glyph>, <missing-glyph>.
//
// d is a presentation attribute, and hence can also be used as a CSS property.
//
// Português:
//
//  O atributo d define um caminho a ser desenhado.
//
// Uma definição de caminho é uma lista de comandos de caminho em que cada comando é composto por uma letra de comando
// e números que representam os parâmetros do comando. Os comandos são detalhados abaixo.
//
// Você pode usar este atributo com os seguintes elementos SVG: <path>, <glyph>, <missing-glyph>.
//
// d é um atributo de apresentação e, portanto, também pode ser usado como uma propriedade CSS.
func (e *TagSvgRect) D(d string) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "d", d)
	return e
}

// Direction
//
// English:
//
//  The direction attribute specifies the inline-base direction of a <text> or <tspan> element. It defines the start
//  and end points of a line of text as used by the text-anchor and inline-size properties. It also may affect the
//  direction in which characters are positioned if the unicode-bidi property's value is either embed or bidi-override.
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
//   Notes:
//     * As a presentation attribute, direction can be used as a CSS property. See css direction for further
//       information.
//
// Português:
//
//  O atributo direction especifica a direção da base embutida de um elemento <text> ou <tspan>. Ele define os pontos
//  inicial e final de uma linha de texto conforme usado pelas propriedades text-anchor e inline-size.
//  Também pode afetar a direção na qual os caracteres são posicionados se o valor da propriedade unicode-bidi for
//  incorporado ou substituído por bidi.
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
//   Notas:
//     * Como atributo de apresentação, a direção pode ser usada como uma propriedade CSS. Veja a direção do CSS para
//       mais informações.
func (e *TagSvgRect) Direction(direction SvgDirection) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "direction", direction.String())
	return e
}

// Display
//
// English:
//
//  The display attribute lets you control the rendering of graphical or container elements.
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
//   * If display is set to none on a <tspan>, <tref>, or <altGlyph> element, then the text string is ignored for the
//     purposes of text layout.
//   * Regarding events, if display is set to none, the element receives no events.
//   * The geometry of a graphics element with display set to none is not included in bounding box and clipping paths
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
//  Notes:
//    * As a presentation attribute, display can be used as a CSS property. See css display for further information.
//
// Português:
//
//  O atributo display permite controlar a renderização de elementos gráficos ou de contêiner.
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
//   * Se display for definido como none em um elemento <tspan>, <tref> ou <altGlyph>, a string de texto será ignorada
//     para fins de layout de texto.
//   * Com relação aos eventos, se display estiver definido como none, o elemento não recebe eventos.
//   * A geometria de um elemento gráfico com exibição definida como nenhum não é incluída nos cálculos da caixa
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
//  Notas:
//    * Como atributo de apresentação, display pode ser usado como propriedade CSS. Consulte a exibição css para obter
//      mais informações.
func (e *TagSvgRect) Display(display SvgDisplay) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "display", display.String())
	return e
}

// DominantBaseline
//
// English:
//
//  The dominant-baseline attribute specifies the dominant baseline, which is the baseline used to align the box's text and inline-level contents. It also indicates the default alignment baseline of any boxes participating in baseline alignment in the box's alignment context.
//
// It is used to determine or re-determine a scaled-baseline-table. A scaled-baseline-table is a compound value with three components:
//
//   1. a baseline-identifier for the dominant-baseline,
//   2. a baseline-table, and
//   3. a baseline-table font-size.
//
// Some values of the property re-determine all three values. Others only re-establish the baseline-table font-size. When the initial value, auto, would give an undesired result, this property can be used to explicitly set the desired scaled-baseline-table.
//
// If there is no baseline table in the nominal font, or if the baseline table lacks an entry for the desired baseline, then the browser may use heuristics to determine the position of the desired baseline.
//
//   Notes:
//     * As a presentation attribute, dominant-baseline can be used as a CSS property.
//
// Português:
//
//
func (e *TagSvgRect) DominantBaseline(dominantBaseline SvgDominantBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "dominant-baseline", dominantBaseline.String())
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// AlignmentBaseline
//
// English:
//
//
//
// Português:
//
//
func (e *TagSvgRect) AlignmentBaseline(alignmentBaseline SvgAlignmentBaseline) (ref *TagSvgRect) {
	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

func (e *TagSvgRect) Fill(value interface{}) (ref *TagSvgRect) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "fill", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "fill", value)
	return e
}

func (e *TagSvgRect) AppendToStage() (ref *TagSvgRect) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgRect) AppendById(appendId string) (ref *TagSvgRect) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}
