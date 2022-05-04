package iotmaker_platform_IDraw

import "image/color"

type ICanvasGradient interface {

	// AddColorStopPosition
	// en: Specifies the colors and stop positions in a gradient object
	//     gradient: A gradient object created by CreateLinearGradient() or
	//    CreateRadialGradient() methods
	//     stopPosition: A value between 0.0 and 1.0 that represents the position
	//     between start (0%) and end (100%) in a gradient
	//     color: A color RGBA value to display at the stop position
	//
	//     Note: You can call the addColorStopPosition() method multiple times to
	//     change a gradient. If you omit this method for gradient objects, the
	//     gradient will not be visible. You need to create at least one color stop to
	//     have a visible gradient.
	//
	// pt_br: Especifica a cor e a posição final para a cor dentro do gradiente
	//     gradient: Objeto de gradiente criado pelos métodos CreateLinearGradient() ou
	//     CreateRadialGradient()
	//     stopPosition: Um valor entre 0.0 e 1.0 que representa a posição entre o
	//     início (0%) e o fim (100%) dentro do gradiente
	//     color: Uma cor no formato RGBA para ser mostrada na posição determinada
	//
	//     Nota: Você pode chamar o método AddColorStopPosition() várias vezes para
	//     adicionar várias cores ao gradiente, porém, se você omitir o método, o
	//     gradiente não será visivel. Você tem a obrigação de chamar o método pelo
	//     menos uma vez com uma cor para que o gradiente seja visível.
	AddColorStopPosition(gradient interface{}, stop float64, color color.RGBA)

	// CreateLinearGradient
	// en: This method of the Canvas 2D API creates a gradient along the line
	// connecting two given coordinates, starting at (x0, y0) point and ending at
	// (x1, y1) point
	//     x0: The x-coordinate of the start point of the gradient
	//     y0: The y-coordinate of the start point of the gradient
	//     x1: The x-coordinate of the end point of the gradient
	//     y1: The y-coordinate of the end point of the gradient
	//
	//     The createLinearGradient() method creates a linear gradient object for be
	//     used with methods AddColorStopPosition(), SetFillStyle() and
	//     SetStrokeStyle().
	//     The gradient can be used to fill rectangles, circles, lines, text, etc.
	//     Tip: Use this object as the value to the strokeStyle() or fillStyle()
	//     methods
	//     Tip: Use the addColorStopPosition() method to specify different colors, and
	//     where to position the colors in the gradient object.
	//
	// pt_br: Este método do canvas 2D cria um gradiente ao longo de uma linha
	// conectando dois pontos, iniciando no ponto (x0, y0) e terminando no ponto
	// (x1, y1)
	//     x0: Coordenada x do ponto inicial do gradiente
	//     y0: Coordenada y do ponto inicial do gradiente
	//     x1: Coordenada x do ponto final do gradiente
	//     y1: Coordenada y do ponto final do gradiente
	//
	//     O método CreateLinearGradient() cria um objeto de gradiente linear para ser
	//     usado em conjunto com os métodos AddColorStopPosition(), SetFillStyle() e
	//     SetStrokeStyle().
	//     O gradiente pode ser usado para preencher retângulos, circulos, linhas,
	//     textos, etc.
	//     Dica: Use este objeto como valor passados aos métodos strokeStyle() ou
	//     fillStyle()
	//     Dica: Use o método addColorStopPosition() para especificar diferentes cores
	//     para o gradiente e a posição de cada cor
	CreateLinearGradient(x0, y0, x1, y1 interface{}) interface{}

	// CreateRadialGradient
	// en: Creates a radial gradient (to use on canvas content). The parameters
	// represent two circles, one with its center at (x0, y0) and a radius of r0, and
	// the other with its center at (x1, y1) with a radius of r1.
	//     x0: The x-coordinate of the starting circle of the gradient
	//     y0: The y-coordinate of the starting circle of the gradient
	//     r0: The radius of the starting circle. Must be non-negative and finite.
	//         (note: radius is a width, not a degrees angle)
	//     x1: The x-coordinate of the ending circle of the gradient
	//     y1: The y-coordinate of the ending circle of the gradient
	//     r1: The radius of the ending circle. Must be non-negative and finite.
	//         (note: radius is a width, not a degrees angle)
	//
	// pt_br: Este método cria um gradiente radial (para ser usado com o canvas 2D). Os
	// parâmetros representam dois círculos, um com o centro no ponto (x0, y0) e raio
	// r0, e outro com centro no ponto (x1, y1) com raio r1
	//     x0: Coordenada x do circulo inicial do gradiente
	//     y0: Coordenada y do circulo inicial do gradiente
	//     r0: Raio do círculo inicial. Deve ser um valor positivo e finito.
	//     (nota: o raio é um comprimento e não um ângulo)
	//     x1: Coordenada x do circulo final do gradiente
	//     y1: Coordenada y do circulo final do gradiente
	//     r1: Raio do círculo final. Deve ser um valor positivo e finito.
	//     (nota: o raio é um comprimento e não um ângulo)
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 interface{}) interface{}

	// SetFillStyle
	// en: Sets the color, gradient, or pattern used to fill the drawing
	//     value: a valid JavaScript value or a color.RGBA{} struct
	//     Default value:	#000000
	//
	// pt_br: Define a cor, gradiente ou padrão usado para preencher o desenho
	//     value: um valor JavaScript valido ou um struct color.RGBA{}
	//     Valor padrão: #000000
	SetFillStyle(value interface{})

	// SetStrokeStyle
	// en: Sets the color, gradient, or pattern used for strokes
	//     value: a valid JavaScript value or a color.RGBA{} struct
	//     Default value: #000000
	//
	// pt_br: Define a cor, gradiente ou padrão usado para o contorno
	//     value: um valor JavaScript valido ou um struct color.RGBA{}
	//     Valor padrão: #000000
	SetStrokeStyle(value interface{})

	// Fill
	// en: The fill() method fills the current drawing (path). The default color is
	// black.
	//     Tip: Use the fillStyle property to fill with another color/gradient.
	//     Note: If the path is not closed, the fill() method will add a line from the
	//     last point to the start point of the
	//     path to close the path (like closePath()), and then fill the path.
	// pt_br: O método fill() preenche e pinta do desenho (caminho). A cor padrão é
	// preto.
	//     Dica: Use a propriedade fillStyle para adicionar uma cor ou gradient.
	//     Nota: Se o caminho não estiver fechado, o método fill() irá adicioná uma
	//     linha do último ao primeiro ponto do
	//     caminho para fechar o caminho (semelhante ao método closePath()) e só então
	//     irá pintar
	Fill()

	// Stroke
	// en: The stroke() method actually draws the path you have defined with all those
	// moveTo() and lineTo() methods. The default color is black.
	//     Tip: Use the strokeStyle property to draw with another color/gradient.
	//
	// pt_br: O método stroke() desenha o caminho definido com os métodos moveTo() e
	// lineTo() usando a cor padrão, preta.
	//     Dica: Use a propriedade strokeStyle para desenhar com outra cor ou usar um
	//     gradiente
	Stroke()
}
