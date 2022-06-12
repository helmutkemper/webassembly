package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"image/color"
	"sync"
	"syscall/js"
)

// TagSvgFeSpecularLighting
//
// English:
//
// The <feSpecularLighting> SVG filter primitive lights a source graphic using the alpha channel as a bump map. The
// resulting image is an RGBA image based on the light color. The lighting calculation follows the standard specular
// component of the Phong lighting model. The resulting image depends on the light color, light position and surface
// geometry of the input bump map. The result of the lighting calculation is added. The filter primitive assumes that
// the viewer is at infinity in the z direction.
//
// This filter primitive produces an image which contains the specular reflection part of the lighting calculation.
// Such a map is intended to be combined with a texture using the add term of the arithmetic <feComposite> method.
// Multiple light sources can be simulated by adding several of these light maps before applying it to the texture
// image.
//
// Português:
//
// A primitiva de filtro SVG <feSpecularLighting> ilumina um gráfico de origem usando o canal alfa como um mapa de
// relevo.
// A imagem resultante é uma imagem RGBA baseada na cor clara. O cálculo da iluminação segue o componente especular
// padrão do modelo de iluminação Phong. A imagem resultante depende da cor da luz, posição da luz e geometria da
// superfície do mapa de relevo de entrada. O resultado do cálculo de iluminação é adicionado. A primitiva de filtro
// assume que o visualizador está no infinito na direção z.
//
// Esta primitiva de filtro produz uma imagem que contém a parte de reflexão especular do cálculo de iluminação.
// Tal mapa deve ser combinado com uma textura usando o termo add do método aritmético <feComposite>.
// Várias fontes de luz podem ser simuladas adicionando vários desses mapas de luz antes de aplicá-los à imagem de
// textura.
type TagSvgFeSpecularLighting struct {

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

// Color
//
// English:
//
//  It provides a potential indirect value (currentcolor) for the fill, stroke, stop-color, flood-color and
//  lighting-color presentation attributes.
//
//   Input:
//     value: potential indirect value of color
//       string: e.g. "black"
//       factory: e.g. factoryColor.NewYellow()
//       RGBA: e.g. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//
//   Notes:
//     * As a presentation attribute, color can be used as a CSS property. See CSS color for further information.
//
// Português:
//
//  Ele fornece um valor indireto potencial (currentcolor) para os atributos de apresentação de preenchimento, traçado,
//  cor de parada, cor de inundação e cor de iluminação.
//
//   Entrada:
//     value: valor indireto potencial da cor
//       string: ex. "black"
//       factory: ex. factoryColor.NewYellow()
//       RGBA: ex. color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
//
//   Notas:
//     * Como atributo de apresentação, a cor pode ser usada como propriedade CSS. Veja cor CSS para mais informações.
func (e *TagSvgFeSpecularLighting) Color(value interface{}) (ref *TagSvgFeSpecularLighting) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "color", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "color", value)
	return e
}
