package canvas

// en: Creates an arc/curve between two tangents
//     x0:     The x-axis coordinate of the first control point.
//     y0:     The y-axis coordinate of the first control point.
//     x1:     The x-axis coordinate of the second control point.
//     y1:     The y-axis coordinate of the second control point.
//     radius: The arc's radius. Must be non-negative.
//
// pt_br: Cria um arco/curva entre duas tangentes
//     x0:     Eixo x da primeira coordenada de controle
//     y0:     Eixo y da primeira coordenada de controle
//     x1:     Eixo x da segunda coordenada de controle
//     y1:     Eixo y da segunda coordenada de controle
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);              // Create a starting point
//     ctx.lineTo(100, 20);             // Create a horizontal line
//     ctx.arcTo(150, 20, 150, 70, 50); // Create an arc
//     ctx.lineTo(150, 120);            // Continue with vertical line
//     ctx.stroke();                    // Draw it
func (el *Canvas) ArcTo(x0, y0, x1, y1, radius interface{}) {
	el.SelfContext.Call("arcTo", x0, y0, x1, y1, radius)
}
