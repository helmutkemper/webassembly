package basic

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.platform.textMetrics"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"image/color"
	"time"
)

type Sprite struct {
	Platform iotmaker_platform_IDraw.IDraw
}

func (el *Sprite) SetPlatform(platform iotmaker_platform_IDraw.IDraw) {
	el.Platform = platform
}

func (el *Sprite) GetPlatform() iotmaker_platform_IDraw.IDraw {
	return el.Platform
}

// en: Begins a path, or resets the current path
//     Tip: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), and
//     arc(), to create paths.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//
// pt_br: Inicia ou reinicializa uma nova rota no desenho
//     Dica: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), e
//     arc(), para criar uma nova rota no desenho
//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
func (el *Sprite) BeginPath() {
	el.Platform.BeginPath()
}

// en: Moves the path to the specified point in the canvas, without creating a line
//     x: The x-coordinate of where to move the path to
//     y: The y-coordinate of where to move the path to
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//
// pt_br: Move o caminho do desenho para o ponto dentro do elemento canvas, sem
// inicializar uma linha
//     X: Coordenada x para onde o ponto vai ser deslocado
//     Y: Coordenada y para onde o ponto vai ser deslocado
//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
func (el *Sprite) MoveTo(x, y int) {
	el.Platform.MoveTo(x, y)
}

// en: Creates an arc/curve between two tangents
//     x0:     The x-axis coordinate of the first control point.
//     y0:     The y-axis coordinate of the first control point.
//     x1:     The x-axis coordinate of the second control point.
//     y1:     The y-axis coordinate of the second control point.
//     radius: The arc's radius. Must be non-negative.
//
// pt_br: Cria um arco/curva entre duas tangentes
//     x0:     Eixo x da primeira coordenada de controle
//     y0:     Eixo y da primeira coordenada de controle
//     x1:     Eixo x da segunda coordenada de controle
//     y1:     Eixo y da segunda coordenada de controle
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);              // Create a starting point
//     ctx.lineTo(100, 20);             // Create a horizontal line
//     ctx.arcTo(150, 20, 150, 70, 50); // Create an arc
//     ctx.lineTo(150, 120);            // Continue with vertical line
//     ctx.stroke();                    // Draw it
func (el *Sprite) ArcTo(x, y, radius, startAngle, endAngle int) {
	el.Platform.ArcTo(x, y, radius, startAngle, endAngle)
}

// en: Adds a new point and creates a line from that point to the last specified
// point in the canvas. (this method does not draw the line).
//     x: The x-coordinate of where to create the line to
//     y: The y-coordinate of where to create the line to
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//
// pt_br: Adiciona um novo ponto e cria uma linha ligando o ponto ao último ponto
// especificado no elemento canvas. (este método não desenha uma linha).
//     x: coordenada x para a criação da linha
//     y: coordenada y para a criação da linha
//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
func (el *Sprite) LineTo(x, y int) {
	el.Platform.LineTo(x, y)
}

// en: Creates a path from the current point back to the starting point
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//     Tip: Use the fill() method to fill the drawing (black is default). Use the
//          fillStyle property to fill with another color/gradient.
//
// pt_br: cria um caminho entre o último ponto especificado e o primeiro ponto
//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
//     Dica: Use o método fill() para preencher o desenho (petro é a cor padrão).
//           Use a propriedade fillStyle para mudar a cor de preenchimento ou
//           adicionar um gradiente
func (el *Sprite) ClosePath(x, y int) {
	el.Platform.ClosePath(x, y)
}

// en: The stroke() method actually draws the path you have defined with all those
//     moveTo() and lineTo() methods. The default color is black.
//     Tip: Use the strokeStyle property to draw with another color/gradient.
//
// pt_br: O método stroke() desenha o caminho definido com os métodos moveTo() e
//     lineTo() usando a cor padrão, preta.
//     Dica: Use a propriedade strokeStyle para desenhar com outra cor ou usar um
//     gradiente
func (el *Sprite) Stroke() {
	el.Platform.Stroke()
}

// en: Sets the current line width in pixels
//     Default value: 1
//     JavaScript syntax: context.lineWidth = number;
//
// pt_br: Define a espessura da linha em pixels
//     Valor padrão: 1
//     Sintaxe JavaScript: context.lineWidth = número
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.strokeRect(20, 20, 80, 100);
func (el *Sprite) SetLineWidth(value interface{}) {
	el.Platform.SetLineWidth(value)
}

// en: Return the current line width in pixels
//     Default value: 1
//     JavaScript syntax: var l = context.lineWidth;
//
// pt_br: Retorna a espessura da linha em pixels
//     Valor padrão: 1
//     Sintaxe JavaScript: var l = context.lineWidth;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.strokeRect(20, 20, 80, 100);
//     var l = ctx.lineWidth;
func (el *Sprite) GetLineWidth() int {
	return el.Platform.GetLineWidth()
}

// en: Sets the blur level for shadows
//     Default value: 0
//
// pt_br: Define o valor de borrão da sombra
//     Valor padrão: 0
func (el *Sprite) SetShadowBlur(value int) {
	el.Platform.SetShadowBlur(value)
}

// en: Return the blur level for shadows
//     Default value: 0
//
// pt_br: Retorna o valor de borrão da sombra
//     Valor padrão: 0
func (el *Sprite) GetShadowBlur() int {
	return el.Platform.GetShadowBlur()
}

// en: Sets the color to use for shadows
//     Note: Use the shadowColor property together with the shadowBlur property to
//           create a shadow.
//     Tip: Adjust the shadow by using the shadowOffsetX and shadowOffsetY
//          properties.
//     Default value: #000000
//
// pt_br: Define a cor da sombra
//     Nota: Use a propriedade shadowColor em conjunto com a propriedade shadowBlur
//           para criar a sombra
//     Dica: Ajuste o local da sombra usando as propriedades shadowOffsetX e
//           shadowOffsetY
//     Valor padrão: #000000
func (el *Sprite) SetShadowColor(value color.RGBA) {
	el.Platform.SetShadowColor(value)
}

// en: Sets the horizontal distance of the shadow from the shape
//     shadowOffsetX = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetX = 20 indicates that the shadow starts 20 pixels to the right
//     (from the shape's left position).
//     shadowOffsetX = -20 indicates that the shadow starts 20 pixels to the left
//     (from the shape's left position).
//     Tip: To adjust the vertical distance of the shadow from the shape, use the
//     shadowOffsetY property.
//     Default value: 0
//
// pt_br: Define a distância horizontal entre a forma e a sua sombra
//     shadowOffsetX = 0 indica que a forma e sua sombra estão alinhadas uma em
//     cima da outra.
//     shadowOffsetX = 20 indica que a forma e a sua sombra estão 20 pixels
//     afastadas a direita (em relação a parte mais a esquerda da forma)
//     shadowOffsetX = -20 indica que a forma e a sua sombra estão 20 pixels
//     afastadas a esquerda (em relação a parte mais a esquerda da forma)
//     Dica: Para ajustar a distância vertical, use a propriedade shadowOffsetY
//     Valor padrão: 0
func (el *Sprite) ShadowOffsetX(value int) {
	el.Platform.ShadowOffsetX(value)
}

// en: Sets or returns the vertical distance of the shadow from the shape
//     The shadowOffsetY property sets or returns the vertical distance of the
//     shadow from the shape.
//     shadowOffsetY = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetY = 20 indicates that the shadow starts 20 pixels below the
//     shape's top position.
//     shadowOffsetY = -20 indicates that the shadow starts 20 pixels above the
//     shape's top position.
//     Tip: To adjust the horizontal distance of the shadow from the shape, use the
//     shadowOffsetX property.
//     Default value: 0
//
// pt_br: Define a distância vertical entre a forma e a sua sombra
//     shadowOffsetY = 0 indica que a forma e sua sombra estão alinhadas uma em
//     cima da outra.
//     shadowOffsetY = 20 indica que a forma e a sua sombra estão 20 pixels
//     afastadas para baixo (em relação a parte mais elevada da forma)
//     shadowOffsetY = -20 indica que a forma e a sua sombra estão 20 pixels
//     afastadas para cima (em relação a parte mais elevada da forma)
//     Dica: Para ajustar a distância horizontal, use a propriedade shadowOffsetX
//     Valor padrão: 0
func (el *Sprite) ShadowOffsetY(value int) {
	el.Platform.ShadowOffsetY(value)
}

// en: Specifies the colors and stop positions in a gradient object
//     gradient: A gradient object created by CreateLinearGradient() or
//     CreateRadialGradient() methods
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
func (el *Sprite) AddColorStopPosition(gradient interface{}, stop float64, color color.RGBA) {
	el.Platform.AddColorStopPosition(gradient, stop, color)
}

// en: The fill() method fills the current drawing (path). The default color is
//     black.
//     Tip: Use the fillStyle property to fill with another color/gradient.
//     Note: If the path is not closed, the fill() method will add a line from the
//     last point to the start point of the path to close the path (like
//     closePath()), and then fill the path.
//
// pt_br: O método fill() preenche e pinta do desenho (caminho). A cor padrão é
//     preto.
//     Dica: Use a propriedade fillStyle para adicionar uma cor ou gradient.
//     Nota: Se o caminho não estiver fechado, o método fill() irá adicioná uma
//     linha do último ao primeiro ponto do caminho para fechar o caminho
//     (semelhante ao método closePath()) e só então irá pintar
func (el *Sprite) Fill() {
	el.Platform.Fill()
}

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
func (el *Sprite) CreateLinearGradient(x0, y0, x1, y1 interface{}) interface{} {
	return el.Platform.CreateLinearGradient(x0, y0, x1, y1)
}

// en: Creates a radial gradient (to use on canvas content). The parameters
// represent two circles, one with its center at (x0, y0) and a radius of r0, and
// the other with its center at (x1, y1) with a radius of r1.
//     x0: The x-coordinate of the starting circle of the gradient
//     y0: The y-coordinate of the starting circle of the gradient
//     r0: The radius of the starting circle. Must be non-negative and finite.
//     (note: radius is a width, not a degrees angle)
//     x1: The x-coordinate of the ending circle of the gradient
//     y1: The y-coordinate of the ending circle of the gradient
//     r1: The radius of the ending circle. Must be non-negative and finite.
//     (note: radius is a width, not a degrees angle)
//
// pt_br: Este método cria um gradiente radial (para ser usado com o canvas 2D). Os
// parâmetros representam dois círculos, um com o centro no ponto (x0, y0) e raio
// r0, e outro com centro no ponto (x1, y1) com raio r1
//     x0: Coordenada x do circulo inicial do gradiente
//     y0: Coordenada y do circulo inicial do gradiente
//     r0: Raio do círculo inicial. Deve ser um valor positivo e finito. (nota: o
//     raio é um comprimento e não um ângulo)
//     x1: Coordenada x do circulo final do gradiente
//     y1: Coordenada y do circulo final do gradiente
//     r1: Raio do círculo final. Deve ser um valor positivo e finito. (nota: o
//     raio é um comprimento e não um ângulo)
func (el *Sprite) CreateRadialGradient(x0, y0, r0, x1, y1, r1 interface{}) interface{} {
	return el.Platform.CreateRadialGradient(x0, y0, r0, x1, y1, r1)
}

// en: Sets the color, gradient, or pattern used to fill the drawing
//     value: a valid JavaScript value or a color.RGBA{} struct
//     Default value:	#000000
//
// pt_br: Define a cor, gradiente ou padrão usado para preencher o desenho
//     value: um valor JavaScript valido ou um struct color.RGBA{}
//     Valor padrão: #000000
func (el *Sprite) SetFillStyle(value interface{}) {
	el.Platform.SetFillStyle(value)
}

// en: Sets the color, gradient, or pattern used for strokes
//     value: a valid JavaScript value or a color.RGBA{} struct
//     Default value: #000000
//
// pt_br: Define a cor, gradiente ou padrão usado para o contorno
//     value: um valor JavaScript valido ou um struct color.RGBA{}
//     Valor padrão: #000000
func (el *Sprite) SetStrokeStyle(value interface{}) {
	el.Platform.SetStrokeStyle(value)
}

// en: Returns an ImageData map[x][y]color.RGBA that copies the pixel data for the
// specified rectangle on a canvas
//     x: The x coordinate (in pixels) of the upper-left corner to start copy from
//     y: The y coordinate (in pixels) of the upper-left corner to start copy from
//     width: The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//     return: map[x(int)][y(int)]color.RGBA
//             Note: return x and y are NOT relative to the coordinate (0,0) on the
//             image, are relative to the coordinate (0,0) on the canvas
//
//     Note: The ImageData object is not a picture, it specifies a part (rectangle)
//     on the canvas, and holds information of every pixel inside that rectangle.
//
//     For every pixel in the map[x][y] there are four pieces of information, the
//     color.RGBA values:
//     R - The color red (from 0-255)
//     G - The color green (from 0-255)
//     B - The color blue (from 0-255)
//     A - The alpha channel (from 0-255; 0 is transparent and 255 is fully
//     visible)
//
//     Tip: After you have manipulated the color/alpha information in the
//     map[x][y], you can copy the image data back onto the canvas with the
//     putImageData() method.
//
// pr_br: Retorna um mapa map[x][y]color.RGBA com parte dos dados da imagem contida
// no retângulo especificado.
//     x: Coordenada x (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     y: Coordenada y (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     width: comprimento do retângulo a ser copiado
//     height: altura do retângulo a ser copiado
//     return: map[x(int)][y(int)]color.RGBA
//             Nota: x e y do retorno não são relativos a coordenada (0,0) da
//             imagem, são relativos a coordenada (0,0) do canvas
//
//     Nota: Os dados da imagem não são uma figura, eles representam uma parte
//     retangular do canvas e guardam informações de cada pixel contido nessa área
//
//     Para cada pixel contido no mapa há quatro peças de informação com valores no
//     formato de color.RGBA:
//     R - Cor vermelha (de 0-255)
//     G - Cor verde (de 0-255)
//     B - Cor azul (de 0-255)
//     A - Canal alpha (de 0-255; onde, 0 é transparente e 255 é totalmente
//     visível)
//
//     Dica: Depois de manipular as informações de cor/alpha contidas no map[x][y],
//     elas podem ser colocadas de volta no canvas com o método putImageData().
func (el *Sprite) GetImageData(x, y, width, height int) interface{} {
	return el.Platform.GetImageData(x, y, width, height)
}

// todo: documentation
func (el *Sprite) GetImageDataAlphaChannelByCoordinate(data interface{}, x, y, width int) uint8 {
	return el.Platform.GetImageDataAlphaChannelByCoordinate(data, x, y, width)
}

func (el *Sprite) GetImageDataPixelByCoordinate(data interface{}, x, y, width int) color.RGBA {
	return el.Platform.GetImageDataPixelByCoordinate(data, x, y, width)
}

// en: Returns an ImageData map[x][y]uint8 that copies the pixel alpha channel for
// the specified rectangle on a canvas
//     x: The x coordinate (in pixels) of the upper-left corner to start copy from
//     y: The y coordinate (in pixels) of the upper-left corner to start copy from
//     width: The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//     return: map[x(int)][y(int)]uint8
//             Note: return x and y are NOT relative to the coordinate (0,0) on the
//             image, are relative to the coordinate (0,0) on the canvas
//
//     Note: The ImageData object is not a picture, it specifies a part (rectangle)
//     on the canvas, and holds information only for alpha channel of every pixel
//     inside that rectangle.
//
//     For every pixel in the map[x][y] there are one piece of information, the
//     alpha channel uint8 value (from 0-255; 0 is transparent and 255 is fully
//     visible)
//
//     Tip: After you have manipulated the color/alpha information in the
//     map[x][y], you can copy the image data back onto the canvas with the
//     putImageData() method.
//
// pr_br: Retorna um mapa map[x][y]uint8 com parte dos dados da imagem contida
// no retângulo especificado.
//     x: Coordenada x (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     y: Coordenada y (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     width: comprimento do retângulo a ser copiado
//     height: altura do retângulo a ser copiado
//     return: map[x(int)][y(int)]uint8
//             Nota: x e y do retorno não são relativos a coordenada (0,0) da
//             imagem, são relativos a coordenada (0,0) do canvas
//
//     Nota: Os dados da imagem não são uma figura, eles representam uma parte
//     retangular do canvas e guardam informações apenas do canal alpha de cada
//     pixel contido nessa área.
//
//     Para cada pixel contido no mapa há apenas uma peça da informação do canal
//     alpha com valores no formato uint8, com valoes de 0-255; onde, 0 é
//     transparente e 255 é totalmente visível
//
//     Dica: Depois de manipular as informações de cor/alpha contidas no map[x][y],
//     elas podem ser colocadas de volta no canvas com o método putImageData().
func (el *Sprite) GetImageDataAlphaChannelOnly(x, y, width, height int) map[int]map[int]uint8 {
	return el.Platform.GetImageDataAlphaChannelOnly(x, y, width, height)
}

// en: Returns an ImageData map[x][y]bool that copies the pixel alpha channel for
// the specified rectangle on a canvas
//     x: The x coordinate (in pixels) of the upper-left corner to start copy from
//     y: The y coordinate (in pixels) of the upper-left corner to start copy from
//     width: The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//     minimumAcceptableValue: (alpha channel < minimumAcceptableValue) true:false
//     return: map[x(int)][y(int)]bool
//             Note: return x and y are NOT relative to the coordinate (0,0) on the
//             image, are relative to the coordinate (0,0) on the canvas
//
//     Note: The ImageData object is not a picture, it specifies a part (rectangle)
//     on the canvas, and holds information only for alpha channel of every pixel
//     inside that rectangle.
//
//     For every pixel in the map[x][y] there are one piece of information, the
//     alpha channel bool value, visible or invisible
//
//     Tip: After you have manipulated the color/alpha information in the
//     map[x][y], you can copy the image data back onto the canvas with the
//     putImageData() method.
//
// pr_br: Retorna um mapa map[x][y]bool com parte dos dados da imagem contida
// no retângulo especificado.
//     x: Coordenada x (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     y: Coordenada y (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     width: comprimento do retângulo a ser copiado
//     height: altura do retângulo a ser copiado
//     minimumAcceptableValue: (canal alpha < minimumAcceptableValue) true : false
//     return: map[x(int)][y(int)]bool
//             Nota: x e y do retorno não são relativos a coordenada (0,0) da
//             imagem, são relativos a coordenada (0,0) do canvas
//
//     Nota: Os dados da imagem não são uma figura, eles representam uma parte
//     retangular do canvas e guardam informações booleanas apenas do canal alpha
//     de cada pixel contido nessa área.
//
//     Para cada pixel contido no mapa há apenas uma peça da informação do canal
//     alpha com valores no formato bool, visível ou invisível.
//
//     Dica: Depois de manipular as informações de cor/alpha contidas no map[x][y],
//     elas podem ser colocadas de volta no canvas com o método putImageData().
func (el *Sprite) GetImageDataCollisionByAlphaChannelValue(x, y, width, height int, minimumAcceptableValue uint8) map[int]map[int]bool {
	return el.Platform.GetImageDataCollisionByAlphaChannelValue(x, y, width, height, minimumAcceptableValue)
}

// en: Clears the specified pixels within a given rectangle
//     x: The x-coordinate of the upper-left corner of the rectangle to clear
//     y: The y-coordinate of the upper-left corner of the rectangle to clear
//     width: The width of the rectangle to clear, in pixels
//     height: The height of the rectangle to clear, in pixels
//
// pt_br: Limpa todos os pixels de um determinado retângulo
//     x: Coordenada x da parte superior esquerda do retângulo a ser limpo
//     y: Coordenada y da parte superior esquerda do retângulo a ser limpo
//     width: Comprimento do retângulo a ser limpo
//     height: Altura do retângulo a ser limpo
//
func (el *Sprite) ClearRect(x, y, width, height int) {
	el.Platform.ClearRect(x, y, width, height)
}

// en: Draws a "filled" rectangle
//     x: The x-coordinate of the upper-left corner of the rectangle
//     y: The y-coordinate of the upper-left corner of the rectangle
//     width: The width of the rectangle, in pixels
//     height: The height of the rectangle, in pixels
//
//     Tip: Use the fillStyle property to set a color, gradient, or pattern used to
//     fill the drawing.
//
// pt_br: Desenha um retângulo preenchido com "tinta"
//     x: Coordenada x da parte superior esquerda do retângulo
//     y: Coordenada y da parte superior esquerda do retângulo
//     width: Comprimento do retângulo
//     height: Altura do retângulo
//
//     Dica: Use a propriedade fillStile() para determinar a cor, gradiente ou
//     padrão a ser usado no reenchimento.
func (el *Sprite) FillRect(x, y, width, height int) {
	el.Platform.FillRect(x, y, width, height)
}

// en: Draws an image, canvas, or video onto the canvas
//     image: Specifies the image, canvas, or video element to use
//     sx: [optional] The x coordinate where to start clipping
//     sy: [optional] The y coordinate where to start clipping
//     sWidth: [optional] The width of the clipped image
//     sHeight: [optional] The height of the clipped image
//     x: The x coordinate where to place the image on the canvas
//     y: The y coordinate where to place the image on the canvas
//     width: [optional] The width of the image to use (stretch or reduce the
//            image)
//     height: [optional] The height of the image to use (stretch or reduce the
//     image)
//
//     Position the image on the canvas:
//     Golang Syntax: platform.DrawImage(img, x, y)
//
//     Position the image on the canvas, and specify width and height of the image:
//     Golang Syntax: platform.DrawImage(img, x, y, width, height)
//
//     Clip the image and position the clipped part on the canvas:
//     Golang Syntax: platform.drawImage(img, sx, sy, sWidth, sHeight, x, y, width,
//                    height)
//
// pt_br: Desenha uma imagem, canvas ou vídeo no elemento canvas
//     image: Especifica a imagem, canvas ou vídeo a ser usado
//     sx: [opcional] Coordenada x de onde o corte vai começar
//     sy: [opcional] Coordenada y de onde o corte vai começar
//     sWidth: [opcional] largura do corte
//     sHeight: [opcional] altura do corte
//     x: Coordenada x do canvas de onde o corte vai ser colocado
//     y: Coordenada y do canvas de onde o corte vai ser colocado
//     width: [opcional] Novo comprimento da imagem
//     height: [opcional] Nova largura da imagem
//
//     Posiciona a imagem no canvas
//     Golang Sintaxe: platform.DrawImage(img, x, y)
//
//     Posiciona a imagem no canvas e determina um novo tamanho da imagem final
//     Golang Sintaxe: platform.DrawImage(img, x, y, width, height)
//
//     Corta um pedaço da imagem e determina uma nova posição e tamanho para a
//     imagem final
//     Golang Sintaxe: platform.drawImage(img, sx, sy, sWidth, sHeight, x, y,
//                     width, height)
func (el *Sprite) DrawImage(image interface{}, value ...int) {
	el.Platform.DrawImage(image, value...)
}

// todo: descrição aqui
func (el *Sprite) DrawImageMultiplesSprites(image interface{}, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex int, spriteChangeInterval time.Duration, x, y, width, height, clearRectX, clearRectY, clearRectWidth, clearRectHeight, lifeCycleLimit, lifeCycleRepeatLimit int, lifeCycleRepeatInterval time.Duration) {
	el.Platform.DrawImageMultiplesSprites(image, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex, spriteChangeInterval, x, y, width, height, clearRectX, clearRectY, clearRectWidth, clearRectHeight, lifeCycleLimit, lifeCycleRepeatLimit, lifeCycleRepeatInterval)
}

// en: Draws "filled" text on the canvas
//     text: Specifies the text that will be written on the canvas
//     x: The x coordinate where to start painting the text (relative to the
//     canvas)
//     y: The y coordinate where to start painting the text (relative to the
//     canvas)
//     maxWidth: [Optional] The maximum allowed width of the text, in pixels
//
// pt_br: Desenha um texto "preenchido" no elemento canvas
//     text: Especifica o texto a ser escrito
//     x: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     y: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     maxWidth: [Opcional] Comprimento máximo do texto em pixels
func (el *Sprite) FillText(text string, x, y int, maxWidth ...int) {
	el.Platform.FillText(text, x, y, maxWidth...)
}

// en: Draws text on the canvas with no fill
//     text: Specifies the text that will be written on the canvas
//     x: The x coordinate where to start painting the text (relative to the
//     canvas)
//     y: The y coordinate where to start painting the text (relative to the
//     canvas)
//     maxWidth: [Optional] The maximum allowed width of the text, in pixels
//
// pt_br: Desenha um texto no elemento canvas sem preenchimento
//     text: Especifica o texto a ser escrito
//     x: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     y: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     maxWidth: [Opcional] Comprimento máximo do texto em pixels
func (el *Sprite) StrokeText(text string, x, y int, maxWidth ...int) {
	el.Platform.StrokeText(text, x, y, maxWidth...)
}

// en: Sets the current font properties for text content
//
// pt_br: Define as propriedades da fonte atual
func (el *Sprite) Font(font font.Font) {
	el.Platform.Font(font)
}

// en: Returns a struct TextMetrics that contains the width of the specified text
//     text: The text to be measured
//
// pt_br: Retorna o struct TextMetrics com os dados de comprimento do texto
//     text: Texto a ser medido
func (el *Sprite) MeasureText(text string) iotmaker_platform_textMetrics.TextMetrics {
	return el.Platform.MeasureText(text)
}

func (el *Sprite) ResetFillStyle() {
	el.Platform.ResetFillStyle()
}

func (el *Sprite) ResetStrokeStyle() {
	el.Platform.ResetStrokeStyle()
}

func (el *Sprite) ResetShadow() {
	el.Platform.ResetShadow()
}

func (el *Sprite) ResetLineWidth() {
	el.Platform.ResetLineWidth()
}

func (el *Sprite) SetMouseCursor(cursor mouse.CursorType) {
	el.Platform.SetMouseCursor(cursor)
}

func (el *Sprite) AddEventListener(eventType interface{}, mouseMoveEvt interface{}) {
	el.Platform.AddEventListener(eventType, mouseMoveEvt)
}

func (el *Sprite) SetPixel(x, y int, pixel interface{}) {
	el.Platform.SetPixel(x, y, pixel)
}

func (el *Sprite) MakePixel(pixelColor color.RGBA) interface{} {
	return el.Platform.MakePixel(pixelColor)
}

func (el *Sprite) CreateImageData(width, height int, pixelColor color.RGBA) interface{} {
	return el.Platform.CreateImageData(width, height, pixelColor)
}
