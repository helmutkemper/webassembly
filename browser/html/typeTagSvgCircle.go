package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"image/color"
	"sync"
	"syscall/js"
)

// TagSvgCircle
//
// English:
//
// The <circle> SVG element is an SVG basic shape, used to draw circles based on a center point and a radius.
//
// Português:
//
// O elemento SVG <circle> é uma forma básica SVG, usada para desenhar círculos com base em um ponto central e um raio.
type TagSvgCircle struct {

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

// Cx
//
// English:
//
//  The cx attribute define the x-axis coordinate of a center point.
//
// Português:
//
//  O atributo cx define a coordenada do eixo x de um ponto central.
func (e *TagSvgCircle) Cx(cx float64) (ref *TagSvgCircle) {
	e.selfElement.Call("setAttribute", "cx", cx)
	return e
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
func (e *TagSvgCircle) ClipPath(clipPath string) (ref *TagSvgCircle) {
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
func (e *TagSvgCircle) ClipRule(clipRule SvgClipRule) (ref *TagSvgCircle) {
	e.selfElement.Call("setAttribute", "clip-rule", clipRule.String())
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
func (e *TagSvgCircle) Color(value interface{}) (ref *TagSvgCircle) {
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
func (e *TagSvgCircle) ColorInterpolation(value interface{}) (ref *TagSvgCircle) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color-interpolation", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color-interpolation", value)
	return e
}
