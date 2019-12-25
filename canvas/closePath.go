package canvas

// en: Creates a path from the current point back to the starting point
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//     Tip: Use the fill() method to fill the drawing (black is default). Use the fillStyle property to fill with
//     another color/gradient.
//
// pt_br: cria um caminho entre o último ponto especificado e o primeiro ponto
//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
//     Dica: Use o método fill() para preencher o desenho (petro é a cor padrão). Use a propriedade fillStyle para mudar
//     a cor de preenchimento ou adicionar um gradiente
func (el *Canvas) ClosePath(x, y float64) {
	el.SelfContext.Call("closePath", x, y)
}
