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
// Because the aspect ratio of an SVG image is defined by the viewBox attribute, if this attribute isn't set, the
// preserveAspectRatio attribute has no effect (with one exception, the <image> element, as described below).
//
// Português:
//
//  O atributo preserveAspectRatio indica como um elemento com uma viewBox fornecendo uma determinada proporção deve
//  caber em uma viewport com uma proporção diferente.
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

func (e *TagSvg) AppendById(appendId string) (ref *TagSvg) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}
