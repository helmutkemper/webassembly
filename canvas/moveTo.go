package canvas

// en: Moves the path to the specified point in the canvas, without creating a line
//     x: The x-coordinate of where to move the path to
//     y: The y-coordinate of where to move the path to
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//
// pt_br: Move o caminho do desenho para o ponto dentro do elemento canvas, sem inicializar uma linha
//     X: Coordenada x para onde o ponto vai ser deslocado
//     Y: Coordenada y para onde o ponto vai ser deslocado
//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
func (el *Canvas) MoveTo(x, y float64) {
	el.SelfContext.Call("moveTo", x, y)
}
