package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"image/color"
	"strconv"
	"sync"
	"syscall/js"
)

// TagSvgTextPath
//
// English:
//
// To render text along the shape of a <path>, enclose the text in a <textPath> element that has an href attribute with
// a reference to the <path> element.
//
// Português
//
// Para renderizar o texto ao longo da forma de um <path>, coloque o texto em um elemento <textPath> que tenha um
// atributo href com uma referência ao elemento <path>.
type TagSvgTextPath struct {

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

// AlignmentBaseline
//
// English:
//
//  The alignment-baseline attribute specifies how an object is aligned with respect to its parent. This property specifies which baseline of this element is to be aligned with the corresponding baseline of the parent. For example, this allows alphabetic baselines in Roman text to stay aligned across font size changes. It defaults to the baseline with the same name as the computed value of the alignment-baseline property.
//
//   Input:
//     alignmentBaseline: specifies how an object is aligned with respect to its parent.
//       string: url(#myClip)
//       consts KSvgAlignmentBaseline... (e.g. KSvgAlignmentBaselineTextBeforeEdge)
//
//   Notes:
//     * As a presentation attribute alignment-baseline can be used as a CSS property.
//     * SVG 2 introduces some changes to the definition of this property. In particular: the values auto, before-edge, and after-edge have been removed. For backwards compatibility, text-before-edge may be mapped to text-top and text-after-edge to text-bottom. Neither text-before-edge nor text-after-edge should be used with the vertical-align property.
//
// Português:
//
//  O atributo alinhamento-base especifica como um objeto é alinhado em relação ao seu pai. Esta propriedade especifica qual linha de base deste elemento deve ser alinhada com a linha de base correspondente do pai. Por exemplo, isso permite que as linhas de base alfabéticas em texto romano permaneçam alinhadas nas alterações de tamanho da fonte. O padrão é a linha de base com o mesmo nome que o valor calculado da propriedade de linha de base de alinhamento.
//
//   Input:
//     alignmentBaseline: especifica como um objeto é alinhado em relação ao seu pai.
//       string: url(#myClip)
//       consts KSvgAlignmentBaseline...  (ex. KSvgAlignmentBaselineTextBeforeEdge)
//
//   Notas:
//     * Como um atributo de apresentação, a linha de base de alinhamento pode ser usada como uma propriedade CSS.
//     * O SVG 2 introduz algumas mudanças na definição desta propriedade. Em particular: os valores auto, before-edge e after-edge foram removidos. Para compatibilidade com versões anteriores, text-before-edge pode ser mapeado para text-top e text-after-edge para text-bottom. Nem text-before-edge nem text-after-edge devem ser usados com a propriedade vertical-align.
func (e *TagSvgTextPath) AlignmentBaseline(alignmentBaseline interface{}) (ref *TagSvgTextPath) {
	if converted, ok := alignmentBaseline.(SvgAlignmentBaseline); ok {
		e.selfElement.Call("setAttribute", "alignment-baseline", converted.String())
	}

	e.selfElement.Call("setAttribute", "alignment-baseline", alignmentBaseline)
	return e
}

// BaselineShift
//
// English:
//
//  The baseline-shift attribute allows repositioning of the dominant-baseline relative to the dominant-baseline of the parent text content element. The shifted object might be a sub- or superscript.
//
//   Input:
//     baselineShift: allows repositioning of the dominant-baseline relative to the dominant-baseline of the parent text content element.
//       float32: 0.05 = "5%"
//       string: "5%"
//       consts KSvgBaselineShift... (e.g. KSvgBaselineShiftAuto)
//
//   Notes:
//     * As a presentation attribute baseline-shift can be used as a CSS property.
//     * This property is going to be deprecated and authors are advised to use vertical-align instead.
//
// Português:
//
//  O atributo baseline-shift permite o reposicionamento da linha de base dominante em relação à linha de base dominante do elemento de conteúdo de texto pai. O objeto deslocado pode ser um sub ou sobrescrito.
//
//   Input:
//     baselineShift: permite o reposicionamento da linha de base dominante em relação à linha de base dominante do elemento de conteúdo de texto pai.
//       float32: 0.05 = "5%"
//       string: "5%"
//       consts KSvgBaselineShift... (ex. KSvgBaselineShiftAuto)
//
//   Notas:
//     * Como atributo de apresentação, baseline-shift pode ser usado como propriedade CSS.
//     * Essa propriedade será preterida e os autores são aconselhados a usar alinhamento vertical.
func (e *TagSvgTextPath) BaselineShift(baselineShift interface{}) (ref *TagSvgTextPath) {
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

// Color
//
// English:
//
//  It provides a potential indirect value (currentcolor) for the fill, stroke, stop-color, flood-color and
//  lighting-color presentation attributes.
//
//   Notes:
//     * As a presentation attribute, color can be used as a CSS property. See CSS color for further information.
//
// Português:
//
//  Ele fornece um valor indireto potencial (currentcolor) para os atributos de apresentação de preenchimento, traçado,
//  cor de parada, cor de inundação e cor de iluminação.
//
//   Notas:
//     * Como atributo de apresentação, a cor pode ser usada como propriedade CSS. Veja cor CSS para mais informações.
func (e *TagSvgTextPath) Color(value interface{}) (ref *TagSvgTextPath) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color", value)
	return e
}
