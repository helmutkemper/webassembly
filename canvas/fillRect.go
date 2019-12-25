package canvas

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
//     Dica: Use a propriedade fillStile() para determinar a cor, gradiente ou padrão
//     a ser usado no reenchimento.
func (el *Canvas) FillRect(x, y, width, height float64) {
	el.SelfContext.Call("fillRect", x, y, width, height)
}
