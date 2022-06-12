package html

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"sync"
	"syscall/js"
)

type TagSvg struct {

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

// ClipPath
//
// English:
//
//  It binds the element it is applied to with a given <clipPath> element.
//
//   Input:
//     clipPath: the element it is applied
//       (e.g. "url(#myClip)", "circle() fill-box", "circle() stroke-box" or "circle() view-box")
//
// Português:
//
//  Ele associa o elemento ao qual é aplicado a um determinado elemento <clipPath>.
//
//   Entrada:
//     clipPath: elemento ao qual é aplicado
//       (ex. "url(#myClip)", "circle() fill-box", "circle() stroke-box" ou "circle() view-box")
func (e *TagSvg) ClipPath(clipPath string) (ref *TagSvg) {
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
//   Input:
//     value: side of a path
//       const: KSvgClipRule... (e.g. KSvgClipRuleNonzero)
//       any other type: interface{}
//
// Português:
//
//  Ele indica como determinar qual lado de um caminho está dentro de uma forma para saber como um <clipPath> deve
//  recortar seu destino.
//
//   Input:
//     value: lado de um caminho
//       const: KSvgClipRule... (e.g. KSvgClipRuleNonzero)
//       qualquer outro tipo: interface{}
func (e *TagSvg) ClipRule(value interface{}) (ref *TagSvg) {
	if converted, ok := value.(SvgClipRule); ok {
		e.selfElement.Call("setAttribute", "clip-rule", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "clip-rule", value)
	return e
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

// Init
//
// English:
//
//  Initializes the object correctly.
//
// Português:
//
//  Inicializa o objeto corretamente.
func (e *TagSvg) Init(id string) (ref *TagSvg) {
	e.listener = new(sync.Map)

	e.CreateElement(KTagSvg)
	e.prepareStageReference()
	//e.selfElement.Set("version", "1.1")
	//e.ViewBox(minX, minY, width, height)
	e.Id(id)

	//e.Width(width)
	//e.Height(height)

	return e
}

// PreserveAspectRatio
//
// English:
//
//  The preserveAspectRatio attribute indicates how an element with a viewBox providing a given aspect ratio must fit
//  into a viewport with a different aspect ratio.
//
//   Input:
//     ratio: Indicates how an element with a viewBox providing a given aspect ratio.
//       KRatioNone: Do not force uniform scaling. Scale the graphic content of the given element non-uniformly if
//         necessary such that the element's bounding box exactly matches the viewport rectangle. Note that if <align>
//         is none, then the optional <meetOrSlice> value is ignored.
//       KRatioXMinYMin: Force uniform scaling. Align the <min-x> of the element's viewBox with the smallest X value of
//         the viewport. Align the <min-y> of the element's viewBox with the smallest Y value of the viewport.
//       KRatioXMidYMin: Force uniform scaling. Align the midpoint X value of the element's viewBox with the midpoint X
//         value of the viewport. Align the <min-y> of the element's viewBox with the smallest Y value of the viewport.
//       KRatioXMaxYMin: Force uniform scaling. Align the <min-x>+<width> of the element's viewBox with the maximum X
//         value of the viewport. Align the <min-y> of the element's viewBox with the smallest Y value of the viewport.
//       KRatioXMinYMid: Force uniform scaling. Align the <min-x> of the element's viewBox with the smallest X value of
//         the viewport. Align the midpoint Y value of the element's viewBox with the midpoint Y value of the viewport.
//       KRatioXMidYMid: (default) - Force uniform scaling. Align the midpoint X value of the element's viewBox with the
//         midpoint X value of the viewport. Align the midpoint Y value of the element's viewBox with the midpoint Y
//         value of the viewport.
//       KRatioXMaxYMid: Force uniform scaling. Align the <min-x>+<width> of the element's viewBox with the maximum X
//         value of the viewport. Align the midpoint Y value of the element's viewBox with the midpoint Y value of the
//         viewport.
//       KRatioXMinYMax: Force uniform scaling. Align the <min-x> of the element's viewBox with the smallest X value of
//         the viewport. Align the <min-y>+<height> of the element's viewBox with the maximum Y value of the viewport.
//       KRatioXMidYMax: Force uniform scaling. Align the midpoint X value of the element's viewBox with the midpoint X
//         value of the viewport. Align the <min-y>+<height> of the element's viewBox with the maximum Y value of the
//         viewport.
//       KRatioXMaxYMax: Force uniform scaling. Align the <min-x>+<width> of the element's viewBox with the maximum X
//         value of the viewport. Align the <min-y>+<height> of the element's viewBox with the maximum Y value of the
//         viewport.
//     meet: The meet or slice reference
//       KMeetOrSliceReferenceMeet: Scale the graphic such that: Aspect ratio is preserved; The entire viewBox is
//         visible within the viewport; The viewBox is scaled up as much as possible, while still meeting the other
//         criteria.
//       KMeetOrSliceReferenceSlice: Scale the graphic such that: Aspect ratio is preserved; The entire viewport is
//         covered by the viewBox; The viewBox is scaled down as much as possible, while still meeting the other
//         criteria.
//
// Because the aspect ratio of an SVG image is defined by the viewBox attribute, if this attribute isn't set, the
// preserveAspectRatio attribute has no effect (with one exception, the <image> element, as described below).
//
// Português:
//
//  O atributo preserveAspectRatio indica como um elemento com uma viewBox fornecendo uma determinada proporção deve
//  caber em uma viewport com uma proporção diferente.
//
//   Input:
//     ratio: Indica como um elemento com uma viewBox fornece uma determinada proporção.
//       KRatioNone: Não force a escala uniforme. Dimensione o conteúdo gráfico do elemento fornecido de forma não
//         uniforme, se necessário, de modo que a caixa delimitadora do elemento corresponda exatamente ao retângulo da
//         janela de visualização. Observe que se <align> for none, o valor opcional <meetOrSlice> será ignorado.
//       KRatioXMinYMin: Forçar escala uniforme. Alinhe o <min-x> da viewBox do elemento com o menor valor X da
//         viewport. Alinhe o <min-y> da viewBox do elemento com o menor valor Y da viewport.
//       KRatioXMidYMin: Forçar escala uniforme. Alinhe o valor X do ponto médio da viewBox do elemento com o valor X do
//         ponto médio da viewport. Alinhe o <min-y> da viewBox do elemento com o menor valor Y da viewport.
//       KRatioXMaxYMin: Forçar escala uniforme. Alinhe o <min-x>+<width> da viewBox do elemento com o valor X máximo da
//         viewport. Alinhe o <min-y> da viewBox do elemento com o menor valor Y da viewport.
//       KRatioXMinYMid: Forçar escala uniforme. Alinhe o <min-x> da viewBox do elemento com o menor valor X da
//         viewport. Alinhe o valor Y do ponto médio da viewBox do elemento com o valor Y do ponto médio da viewport.
//       KRatioXMidYMid: (default) - Força o dimensionamento uniforme. Alinhe o valor X do ponto médio da viewBox do
//         elemento com o valor X do ponto médio da viewport. Alinhe o valor Y do ponto médio da viewBox do elemento com
//         o valor Y do ponto médio da viewport.
//       KRatioXMaxYMid: Forçar escala uniforme. Alinhe o <min-x>+<width> da viewBox do elemento com o valor X máximo da
//         viewport. Alinhe o valor Y do ponto médio da viewBox do elemento com o valor Y do ponto médio da viewport.
//       KRatioXMinYMax: Forçar escala uniforme. Alinhe o <min-x> da viewBox do elemento com o menor valor X da
//         viewport. Alinhe o <min-y>+<height> da viewBox do elemento com o valor Y máximo da viewport.
//       KRatioXMidYMax: Forçar escala uniforme. Alinhe o valor X do ponto médio da viewBox do elemento com o valor X do
//         ponto médio da viewport. Alinhe o <min-y>+<height> da viewBox do elemento com o valor Y máximo da viewport.
//       KRatioXMaxYMax: Forçar escala uniforme. Alinhe o <min-x>+<width> da viewBox do elemento com o valor X máximo da
//         viewport. Alinhe o <min-y>+<height> da viewBox do elemento com o valor Y máximo da viewport.
//     meet: A referência de encontro ou fatia
//       KMeetOrSliceReferenceMeet: Dimensione o gráfico de forma que: A proporção seja preservada; A viewBox inteira é
//         visível dentro da viewport; A viewBox é ampliada o máximo possível, enquanto ainda atende aos outros
//         critérios.
//       KMeetOrSliceReferenceSlice: Dimensione o gráfico de forma que: A proporção seja preservada; A viewport inteira
//         é coberta pela viewBox; A viewBox é reduzida o máximo possível, enquanto ainda atende aos outros critérios.
//
// Como a proporção de uma imagem SVG é definida pelo atributo viewBox, se esse atributo não estiver definido, o
// atributo preserveAspectRatio não terá efeito (com uma exceção, o elemento <image>, conforme descrito abaixo).
func (e *TagSvg) PreserveAspectRatio(ratio Ratio, meet MeetOrSliceReference) (ref *TagSvg) {
	e.selfElement.Call("setAttribute", "preserveAspectRatio", ratio.String()+" "+meet.String())
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
func (e *TagSvg) Width(width float64) (ref *TagSvg) {
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
func (e *TagSvg) Height(height float64) (ref *TagSvg) {
	e.selfElement.Call("setAttribute", "height", height)
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
func (e *TagSvg) X(x float64) (ref *TagSvg) {
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
func (e *TagSvg) Y(y float64) (ref *TagSvg) {
	e.selfElement.Call("setAttribute", "y", y)
	return e
}

func (e *TagSvg) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvg) Id(id string) (ref *TagSvg) {
	e.id = id
	e.selfElement.Set("id", id)

	return e
}

func (e *TagSvg) CreateElement(tag Tag) (ref *TagSvg) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	e.selfElement.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")

	return e
}

// ViewBox
//
// English:
//
//  The viewBox attribute defines the position and dimension, in user space, of an SVG viewport.
//
// The value of the viewBox attribute is a list of four numbers: min-x, min-y, width and height. The numbers, which are
// separated by whitespace and/or a comma, specify a rectangle in user space which is mapped to the bounds of the
// viewport established for the associated SVG element (not the browser viewport).
//
// Português:
//
//  The viewBox attribute defines the position and dimension, in user space, of an SVG viewport.
//
// O valor do atributo viewBox é uma lista de quatro números: min-x, min-y, largura e altura. Os números, que são
// separados por espaço em branco e ou vírgula, especificam um retângulo no espaço do usuário que é mapeado para os
// limites da janela de visualização estabelecida para o elemento SVG associado (não a janela de visualização do
// navegador).
func (e *TagSvg) ViewBox(minX, minY, width, height float64) (ref *TagSvg) {
	e.selfElement.Call("setAttribute", "viewBox", fmt.Sprintf("%v %v %v %v", minX, minY, width, height))
	return e
}

func (e *TagSvg) AppendToStage() (ref *TagSvg) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvg) AppendToElement(el js.Value) (ref *TagSvg) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvg) Get() (el js.Value) {
	return e.selfElement
}

func (e *TagSvg) AppendById(appendId string) (ref *TagSvg) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}
