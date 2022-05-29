package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"strconv"
	"sync"
	"syscall/js"
)

type TagSvgImage struct {

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
func (e *TagSvgImage) Init(id string) (ref *TagSvgImage) {
	e.listener = new(sync.Map)

	e.CreateElement(KTagSvgImage)
	e.prepareStageReference()
	e.Id(id)

	return e
}

func (e *TagSvgImage) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgImage) Id(id string) (ref *TagSvgImage) {
	e.id = id
	e.selfElement.Set("id", id)
	return e
}

func (e *TagSvgImage) CreateElement(tag Tag) (ref *TagSvgImage) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "image")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	return e
}

func (e *TagSvgImage) X(x float64) (ref *TagSvgImage) {
	e.selfElement.Call("setAttribute", "x", x)
	return e
}

func (e *TagSvgImage) Y(y float64) (ref *TagSvgImage) {
	e.selfElement.Call("setAttribute", "y", y)
	return e
}

func (e *TagSvgImage) Width(width interface{}) (ref *TagSvgImage) {
	if converted, ok := width.(float64); ok {
		p := strconv.FormatFloat(100*converted, 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "width", p)
		return e
	}

	e.selfElement.Call("setAttribute", "width", width)
	return e
}

func (e *TagSvgImage) Height(height interface{}) (ref *TagSvgImage) {
	if converted, ok := height.(float64); ok {
		p := strconv.FormatFloat(100*converted, 'g', -1, 64) + "%"
		e.selfElement.Call("setAttribute", "height", p)
		return e
	}

	e.selfElement.Call("setAttribute", "height", height)
	return e
}

func (e *TagSvgImage) XLinkHRef(xLinkHRef string) (ref *TagSvgImage) {
	e.selfElement.Call("setAttribute", "href", xLinkHRef)
	return e
}

func (e *TagSvgImage) Style(style string) (ref *TagSvgImage) {
	e.selfElement.Call("setAttribute", "style", style)
	return e
}

func (e *TagSvgImage) AppendToElement(el js.Value) (ref *TagSvgImage) {
	e.selfElement.Call("appendChild", el)
	return e
}

func (e *TagSvgImage) AppendToStage() (ref *TagSvgImage) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgImage) AppendById(appendId string) (ref *TagSvgImage) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSvgImage) Get() (element js.Value) {
	return e.selfElement
}
