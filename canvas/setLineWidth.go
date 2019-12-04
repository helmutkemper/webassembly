package canvas

// en: Sets the current line width in pixels
//     Default value: 1
//     JavaScript syntax: context.lineWidth = number;
//
// pt_br: Define a espessura da linha em pixels
//     Valor padrão: 1
//     Sintaxe JavaScript: context.lineWidth = número
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.strokeRect(20, 20, 80, 100);
func (el *Canvas) SetLineWidth(value int) {
	el.SelfContext.Set("lineWidth", value)
}
