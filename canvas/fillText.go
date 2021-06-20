package canvas

// FillText
// en: Draws "filled" text on the canvas
//     text: Specifies the text that will be written on the canvas
//     x: The x coordinate where to start painting the text (relative to the canvas)
//     y: The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: [Optional] The maximum allowed width of the text, in pixels
//
// pt_br: Desenha um texto "preenchido" no elemento canvas
//     text: Especifica o texto a ser escrito
//     x: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     y: coordenada x do texto a ser escrito (relativo ao elemento canvas)
//     maxWidth: [Opcional] Comprimento máximo do texto em pixels
func (el *Canvas) FillText(text string, x, y int, maxWidth ...int) {
	if maxWidth == nil {
		el.SelfContext.Call("fillText", text, x, y)
	} else {
		el.SelfContext.Call("fillText", text, x, y, maxWidth[0])
	}
}
