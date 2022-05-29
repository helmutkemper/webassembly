package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"sync"
	"syscall/js"
)

type TagSvgPath struct {

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
func (e *TagSvgPath) Init(id string) (ref *TagSvgPath) {
	e.listener = new(sync.Map)

	e.CreateElement()
	e.prepareStageReference()
	e.Id(id)

	return e
}

func (e *TagSvgPath) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagSvgPath) Id(id string) (ref *TagSvgPath) {
	e.id = id
	e.selfElement.Set("id", id)
	return e
}

func (e *TagSvgPath) CreateElement() (ref *TagSvgPath) {
	e.selfElement = js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "path")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	return e
}

// D
//
// English:
//
//  The d attribute defines a path to be drawn.
//
//   Input:
//     d: defines the shape of the path.
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
//   Entrada:
//     d: define a forma do caminho.
//
// Uma definição de caminho é uma lista de comandos de caminho em que cada comando é composto por uma letra de comando
// e números que representam os parâmetros do comando. Os comandos são detalhados abaixo.
//
// Você pode usar este atributo com os seguintes elementos SVG: <path>, <glyph>, <missing-glyph>.
//
// d é um atributo de apresentação e, portanto, também pode ser usado como uma propriedade CSS.
func (e *TagSvgPath) D(d *SvgPath) (ref *TagSvgPath) {
	e.selfElement.Call("setAttribute", "d", d.String())
	return e
}

func (e *TagSvgPath) AppendById(appendId string) (ref *TagSvgPath) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}
