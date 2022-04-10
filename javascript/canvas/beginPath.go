package canvas

// BeginPath
//	en: Begins a path, or resets the current path
//      Tip: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), and arc(), to create paths.
//      Tip: Use the stroke() method to actually draw the path on the canvas.
//
// pt_br: Inicia ou reinicializa uma nova rota no desenho
//      Dica: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), e arc(), para criar uma nova rota no desenho
//      Dica: Use o método stroke() para desenhar a rota no elemento canvas
func (el *Canvas) BeginPath() {
	el.SelfContext.Call("beginPath")
}
