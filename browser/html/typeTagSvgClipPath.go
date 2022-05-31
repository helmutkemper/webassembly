package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"sync"
	"syscall/js"
)

// TagSvgClipPath
//
// English:
//
// The <clipPath> SVG element defines a clipping path, to be used by the clip-path property.
//
// A clipping path restricts the region to which paint can be applied. Conceptually, parts of the drawing that lie
// outside of the region bounded by the clipping path are not drawn.
//
// Português
//
// O elemento SVG <clipPath> define um caminho de recorte, a ser usado pela propriedade clip-path.
//
// Um traçado de recorte restringe a região na qual a tinta pode ser aplicada. Conceitualmente, as partes do desenho
// que estão fora da região delimitada pelo caminho de recorte não são desenhadas.
type TagSvgClipPath struct {

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

// ClipPathUnits
//
// English:
//
//  The clipPathUnits attribute indicates which coordinate system to use for the contents of the <clipPath> element.
//
//   Input:
//     clipPathUnits: indicates which coordinate system to used
//       KSvgClipPathUnits... (e.g. KSvgClipPathUnitsUserSpaceOnUse)
//
// Português:
//
//  O atributo clipPathUnits indica qual sistema de coordenadas deve ser usado para o conteúdo do elemento <clipPath>.
//
//   Input:
//     clipPathUnits: indica qual sistema de coordenadas deve ser usado
//       KSvgClipPathUnits... (ex. KSvgClipPathUnitsUserSpaceOnUse)
func (e *TagSvgClipPath) ClipPathUnits(clipPathUnits SvgClipPathUnits) (ref *TagSvgClipPath) {
	e.selfElement.Call("setAttribute", "clipPathUnits", clipPathUnits.String())
	return e
}
