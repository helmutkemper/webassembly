package canvas

// en: Draws text on the canvas with no fill
//     text: Specifies the text that will be written on the canvas
//     x: The x coordinate where to start painting the text (relative to the canvas)
//     y: The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: [Optional] The maximum allowed width of the text, in pixels
//
// pt_br: Desenha um texto no elemento canvas sem preenchimento
//     text: Especifica o texto a ser escrito
//     x: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     y: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     maxWidth: [Opcional] Comprimento máximo do texto em pixels
func (el *Canvas) StrokeText(text string, x, y float64, maxWidth ...float64) {
	if maxWidth == nil {
		el.SelfContext.Call("strokeText", text, x, y)
	} else {
		el.SelfContext.Call("strokeText", text, x, y, maxWidth[0])
	}
}
