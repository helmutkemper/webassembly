package canvas

// CreateLinearGradient
// en: This method of the Canvas 2D API creates a gradient along the line connecting two given coordinates, starting at (x0, y0) point and ending at (x1, y1) point
//     x0: The x-coordinate of the start point of the gradient
//     y0: The y-coordinate of the start point of the gradient
//     x1: The x-coordinate of the end point of the gradient
//     y1: The y-coordinate of the end point of the gradient
//
//     The createLinearGradient() method creates a linear gradient object for be used with methods AddColorStopPosition(), SetFillStyle() and SetStrokeStyle().
//     The gradient can be used to fill rectangles, circles, lines, text, etc.
//     Tip: Use this object as the value to the strokeStyle() or fillStyle() methods
//     Tip: Use the addColorStopPosition() method to specify different colors, and where to position the colors in the gradient object.
//
// pt_br: Este método do canvas 2D cria um gradiente ao longo de uma linha conectando dois pontos, iniciando no ponto (x0, y0) e terminando no ponto (x1, y1)
//     x0: Coordenada x do ponto inicial do gradiente
//     y0: Coordenada y do ponto inicial do gradiente
//     x1: Coordenada x do ponto final do gradiente
//     y1: Coordenada y do ponto final do gradiente
//
//     O método CreateLinearGradient() cria um objeto de gradiente linear para ser usado em conjunto com os métodos AddColorStopPosition(), SetFillStyle() e SetStrokeStyle().
//     O gradiente pode ser usado para preencher retângulos, circulos, linhas, textos, etc.
//     Dica: Use este objeto como valor passados aos métodos strokeStyle() ou fillStyle()
//     Dica: Use o método addColorStopPosition() para especificar diferentes cores para o gradiente e a posição de cada cor
func (el *Canvas) CreateLinearGradient(x0, y0, x1, y1 interface{}) interface{} {
	return el.SelfContext.Call("createLinearGradient", x0, y0, x1, y1)
}
