package canvas

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
func (el *Canvas) ClearRect(x, y, width, height interface{}) {
	el.SelfContext.Call("clearRect", x, y, width, height)
}
