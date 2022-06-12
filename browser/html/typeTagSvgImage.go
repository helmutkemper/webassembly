package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"sync"
	"syscall/js"
)

// TagSvgImage
//
// English:
//
// The <image> SVG element includes images inside SVG documents. It can display raster image files or other SVG files.
//
// The only image formats SVG software must support are JPEG, PNG, and other SVG files. Animated GIF behavior is
// undefined.
//
// SVG files displayed with <image> are treated as an image: external resources aren't loaded, :visited styles aren't
// applied, and they cannot be interactive. To include dynamic SVG elements, try <use> with an external URL. To include
// SVG files and run scripts inside them, try <object> inside of <foreignObject>.
//
//   Notes:
//     * The HTML spec defines <image> as a synonym for <img> while parsing HTML. This specific element and its
//       behavior only apply inside SVG documents or inline SVG.
//
// Português:
//
// O elemento SVG <image> inclui imagens dentro de documentos SVG. Ele pode exibir arquivos de imagem raster ou outros
// arquivos SVG.
//
// Os únicos formatos de imagem que o software SVG deve suportar são JPEG, PNG e outros arquivos SVG. O comportamento
// do GIF animado é indefinido.
//
// Arquivos SVG exibidos com <image> são tratados como uma imagem: recursos externos não são carregados, estilos
// :visited não são aplicados e não podem ser interativos. Para incluir elementos SVG dinâmicos, tente <use> com uma
// URL externa. Para incluir arquivos SVG e executar scripts dentro deles, tente <object> dentro de <foreignObject>.
//
//   Notes:
//     * The HTML spec defines <image> as a synonym for <img> while parsing HTML. This specific element and its
//       behavior only apply inside SVG documents or inline SVG.
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
func (e *TagSvgImage) ClipPath(clipPath string) (ref *TagSvgImage) {
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
func (e *TagSvgImage) ClipRule(value interface{}) (ref *TagSvgImage) {
	if converted, ok := value.(SvgClipRule); ok {
		e.selfElement.Call("setAttribute", "clip-rule", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "clip-rule", value)
	return e
}
