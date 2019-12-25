package canvas

// en: Adds a new point and creates a line from that point to the last specified point in the canvas. (this method does not draw the line).
//     x: The x-coordinate of where to create the line to
//     y: The y-coordinate of where to create the line to
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//
// pt_br: Adiciona um novo ponto e cria uma linha ligando o ponto ao último ponto especificado no elemento canvas. (este método não desenha uma linha).
//     x: coordenada x para a criação da linha
//     y: coordenada y para a criação da linha
//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
func (el *Canvas) LineTo(x, y float64) {
	el.SelfContext.Call("lineTo", x, y)
}
