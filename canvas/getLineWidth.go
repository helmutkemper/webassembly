package canvas

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
func (el *Canvas) GetLineWidth() float64 {
	return el.SelfContext.Get("lineWidth").Float()
}
