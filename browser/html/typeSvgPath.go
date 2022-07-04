package html

import "fmt"

type SvgPath struct {
	path string
}

// M (MoveTo)
//
// English:
//
//  Move the current point to the coordinate x,y. Any subsequent coordinate pair(s) are interpreted as parameter(s) for
//  implicit absolute LineTo (L) command(s) (see below).
//
// Formula: Pn = {x, y}
//
// Português:
//
//  Mova o ponto atual para a coordenada x,y. Quaisquer pares de coordenadas subsequentes são interpretados como
//  parâmetro(s) para comando(s) LineTo (L) absoluto implícito (veja abaixo).
//
// Fórmula: Pn = {x, y}
func (e *SvgPath) M(x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("M %v,%v ", x, y)
	return e
}

// MoveTo
//
// English:
//
//  Move the current point to the coordinate x,y. Any subsequent coordinate pair(s) are interpreted as parameter(s) for
//  implicit absolute LineTo (L) command(s) (see below).
//
// Formula: Pn = {x, y}
//
// Português:
//
//  Mova o ponto atual para a coordenada x,y. Quaisquer pares de coordenadas subsequentes são interpretados como
//  parâmetro(s) para comando(s) LineTo (L) absoluto implícito (veja abaixo).
//
// Fórmula: Pn = {x, y}
func (e *SvgPath) MoveTo(x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("M %v,%v ", x, y)
	return e
}

// Md (MoveToDelta)
//
// English:
//
//  Move the current point by shifting the last known position of the path by dx along the x-axis and by dy along the
//  y-axis. Any subsequent coordinate pair(s) are interpreted as parameter(s) for implicit relative LineTo (l)
//  command(s) (see below).
//
// Formula: Pn = {xo + dx, yo + dy}
//
// Português:
//
//  Mova o ponto atual deslocando a última posição conhecida do caminho por dx ao longo do eixo x e por dy ao longo do
//  eixo y. Quaisquer pares de coordenadas subsequentes são interpretados como parâmetro(s) para o(s) comando(s) LineTo
//  (l) relativo implícito (veja abaixo).
//
// Fórmula: Pn = {xo + dx, yo + dy}
func (e *SvgPath) Md(dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("m %v,%v ", dx, dy)
	return e
}

// MoveToDelta
//
// English:
//
//  Move the current point by shifting the last known position of the path by dx along the x-axis and by dy along the
//  y-axis. Any subsequent coordinate pair(s) are interpreted as parameter(s) for implicit relative LineTo (l)
//  command(s) (see below).
//
// Formula: Pn = {xo + dx, yo + dy}
//
// Português:
//
//  Mova o ponto atual deslocando a última posição conhecida do caminho por dx ao longo do eixo x e por dy ao longo do
//  eixo y. Quaisquer pares de coordenadas subsequentes são interpretados como parâmetro(s) para o(s) comando(s) LineTo
//  (l) relativo implícito (veja abaixo).
//
// Fórmula: Pn = {xo + dx, yo + dy}
func (e *SvgPath) MoveToDelta(dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("m %v,%v ", dx, dy)
	return e
}

// L (LineTo)
//
// English:
//
//  Draw a line from the current point to the end point specified by x,y. Any subsequent coordinate pair(s) are
//  interpreted as parameter(s) for implicit absolute LineTo (L) command(s).
//
// Formula: Po′ = Pn = {x, y}
//
// Português:
//
//  Desenhe uma linha do ponto atual até o ponto final especificado por x,y. Quaisquer pares de coordenadas subsequentes
//  são interpretados como parâmetros para comandos LineTo (L) absolutos implícitos.
//
// Fórmula: Po′ = Pn = {x, y}
func (e *SvgPath) L(x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("L %v,%v ", x, y)
	return e
}

// LineTo
//
// English:
//
//  Draw a line from the current point to the end point specified by x,y. Any subsequent coordinate pair(s) are
//  interpreted as parameter(s) for implicit absolute LineTo (L) command(s).
//
// Formula: Po′ = Pn = {x, y}
//
// Português:
//
//  Desenhe uma linha do ponto atual até o ponto final especificado por x,y. Quaisquer pares de coordenadas subsequentes
//  são interpretados como parâmetros para comandos LineTo (L) absolutos implícitos.
//
// Fórmula: Po′ = Pn = {x, y}
func (e *SvgPath) LineTo(x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("L %v,%v ", x, y)
	return e
}

// Ld (LineToDelta)
//
// English:
//
//  Draw a line from the current point to the end point, which is the current point shifted by dx along the x-axis and
//  dy along the y-axis. Any subsequent coordinate pair(s) are interpreted as parameter(s) for implicit relative LineTo
//  (l) command(s) (see below).
//
// Formula: Po′ = Pn = {xo + dx, yo + dy}
//
// Português:
//
//  Desenhe uma linha do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao longo do eixo x e dy ao
//  longo do eixo y. Quaisquer pares de coordenadas subsequentes são interpretados como parâmetro(s) para o(s)
//  comando(s) LineTo (l) relativo implícito (veja abaixo).
//
// Fórmula: Po′ = Pn = {xo + dx, yo + dy}
func (e *SvgPath) Ld(dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("l %v,%v ", dx, dy)
	return e
}

// LineToDelta
//
// English:
//
//  Draw a line from the current point to the end point, which is the current point shifted by dx along the x-axis and
//  dy along the y-axis. Any subsequent coordinate pair(s) are interpreted as parameter(s) for implicit relative LineTo
//  (l) command(s) (see below).
//
// Formula: Po′ = Pn = {xo + dx, yo + dy}
//
// Português:
//
//  Desenhe uma linha do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao longo do eixo x e dy ao
//  longo do eixo y. Quaisquer pares de coordenadas subsequentes são interpretados como parâmetro(s) para o(s)
//  comando(s) LineTo (l) relativo implícito (veja abaixo).
//
// Fórmula: Po′ = Pn = {xo + dx, yo + dy}
func (e *SvgPath) LineToDelta(dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("l %v,%v ", dx, dy)
	return e
}

// H (HorizontalLine)
//
// English:
//
//  Draw a horizontal line from the current point to the end point, which is specified by the x parameter and the
//  current point's y coordinate. Any subsequent value(s) are interpreted as parameter(s) for implicit absolute
//  horizontal LineTo (H) command(s).
//
// Formula: Po′ = Pn = {x, yo}
//
// Português:
//
//  Desenhe uma linha horizontal do ponto atual até o ponto final, que é especificado pelo parâmetro x e pela coordenada
//  y do ponto atual. Quaisquer valores subsequentes são interpretados como parâmetro(s) para comandos LineTo (H)
//  horizontais absolutos implícitos.
//
// Formula: Po′ = Pn = {x, yo}
func (e *SvgPath) H(x float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("H %v ", x)
	return e
}

// HorizontalLine
//
// English:
//
//  Draw a horizontal line from the current point to the end point, which is specified by the x parameter and the
//  current point's y coordinate. Any subsequent value(s) are interpreted as parameter(s) for implicit absolute
//  horizontal LineTo (H) command(s).
//
// Formula: Po′ = Pn = {x, yo}
//
// Português:
//
//  Desenhe uma linha horizontal do ponto atual até o ponto final, que é especificado pelo parâmetro x e pela coordenada
//  y do ponto atual. Quaisquer valores subsequentes são interpretados como parâmetro(s) para comandos LineTo (H)
//  horizontais absolutos implícitos.
//
// Formula: Po′ = Pn = {x, yo}
func (e *SvgPath) HorizontalLine(x float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("H %v ", x)
	return e
}

// Hd (HorizontalLineDelta)
//
// English:
//
//  Draw a horizontal line from the current point to the end point, which is specified by the current point shifted by
//  dx along the x-axis and the current point's y coordinate. Any subsequent value(s) are interpreted as parameter(s)
//  for implicit relative horizontal LineTo (h) command(s).
//
// Formula: Po′ = Pn = {xo + dx, yo}
//
// Português:
//
//  Draw a horizontal line from the current point to the end point, which is specified by the current point shifted by
//  dx along the x-axis and the current point's y coordinate. Any subsequent value(s) are interpreted as parameter(s)
//  for implicit relative horizontal LineTo (h) command(s).
//
// Fórmula: Po′ = Pn = {xo + dx, yo}
func (e *SvgPath) Hd(dx float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("h %v ", dx)
	return e
}

// HorizontalLineDelta
//
// English:
//
//  Draw a horizontal line from the current point to the end point, which is specified by the current point shifted by
//  dx along the x-axis and the current point's y coordinate. Any subsequent value(s) are interpreted as parameter(s)
//  for implicit relative horizontal LineTo (h) command(s).
//
// Formula: Po′ = Pn = {xo + dx, yo}
//
// Português:
//
//  Draw a horizontal line from the current point to the end point, which is specified by the current point shifted by
//  dx along the x-axis and the current point's y coordinate. Any subsequent value(s) are interpreted as parameter(s)
//  for implicit relative horizontal LineTo (h) command(s).
//
// Fórmula: Po′ = Pn = {xo + dx, yo}
func (e *SvgPath) HorizontalLineDelta(dx float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("h %v ", dx)
	return e
}

// V (VerticalLine)
//
// English:
//
//  Draw a vertical line from the current point to the end point, which is specified by the y parameter and the current
//  point's x coordinate. Any subsequent values are interpreted as parameters for implicit absolute vertical LineTo (V)
//  command(s).
//
// Formula: Po′ = Pn = {xo, y}
//
// Português:
//
//  Draw a vertical line from the current point to the end point, which is specified by the y parameter and the current
//  point's x coordinate. Any subsequent values are interpreted as parameters for implicit absolute vertical LineTo (V)
//  command(s).
//
// Fórmula: Po′ = Pn = {xo, y}
func (e *SvgPath) V(y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("V %v ", y)
	return e
}

// VerticalLine
//
// English:
//
//  Draw a vertical line from the current point to the end point, which is specified by the y parameter and the current
//  point's x coordinate. Any subsequent values are interpreted as parameters for implicit absolute vertical LineTo (V)
//  command(s).
//
// Formula: Po′ = Pn = {xo, y}
//
// Português:
//
//  Draw a vertical line from the current point to the end point, which is specified by the y parameter and the current
//  point's x coordinate. Any subsequent values are interpreted as parameters for implicit absolute vertical LineTo (V)
//  command(s).
//
// Fórmula: Po′ = Pn = {xo, y}
func (e *SvgPath) VerticalLine(y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("V %v ", y)
	return e
}

// Vd (VerticalLineDelta)
//
// English:
//
//  Draw a vertical line from the current point to the end point, which is specified by the current point shifted by dy
//  along the y-axis and the current point's x coordinate. Any subsequent value(s) are interpreted as parameter(s) for
//  implicit relative vertical LineTo (v) command(s).
//
// Formula: Po′ = Pn = {xo, yo + dy}
//
//  Draw a vertical line from the current point to the end point, which is specified by the current point shifted by dy
//  along the y-axis and the current point's x coordinate. Any subsequent value(s) are interpreted as parameter(s) for
//  implicit relative vertical LineTo (v) command(s).
//
// Fórmula: Po′ = Pn = {xo, yo + dy}
//
// Português:
//
//  Desenhe uma linha vertical do ponto atual até o ponto final, que é especificado pelo ponto atual deslocado por dy
//  ao longo do eixo y e a coordenada x do ponto atual. Quaisquer valores subsequentes são interpretados como
//  parâmetro(s) para comando(s) vertical(is) LineTo (v) relativo implícito.
//
// Formula: Po′ = Pn = {xo, yo + dy}
func (e *SvgPath) Vd(dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("v %v ", dy)
	return e
}

// VerticalLineDelta
//
// English:
//
//  Draw a vertical line from the current point to the end point, which is specified by the current point shifted by dy
//  along the y-axis and the current point's x coordinate. Any subsequent value(s) are interpreted as parameter(s) for
//  implicit relative vertical LineTo (v) command(s).
//
// Formula: Po′ = Pn = {xo, yo + dy}
//
//  Draw a vertical line from the current point to the end point, which is specified by the current point shifted by dy
//  along the y-axis and the current point's x coordinate. Any subsequent value(s) are interpreted as parameter(s) for
//  implicit relative vertical LineTo (v) command(s).
//
// Fórmula: Po′ = Pn = {xo, yo + dy}
//
// Português:
//
//  Desenhe uma linha vertical do ponto atual até o ponto final, que é especificado pelo ponto atual deslocado por dy
//  ao longo do eixo y e a coordenada x do ponto atual. Quaisquer valores subsequentes são interpretados como
//  parâmetro(s) para comando(s) vertical(is) LineTo (v) relativo implícito.
//
// Formula: Po′ = Pn = {xo, yo + dy}
func (e *SvgPath) VerticalLineDelta(dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("v %v ", dy)
	return e
}

// C (CubicBezierCurve)
//
// English:
//
//  Draw a cubic Bézier curve from the current point to the end point specified by x,y. The start control point is
//  specified by x1,y1 and the end control point is specified by x2,y2. Any subsequent triplet(s) of coordinate pairs
//  are interpreted as parameter(s) for implicit absolute cubic Bézier curve (C) command(s).
//
// Formula:
//   Po′ = Pn = {x, y} ;
//   Pcs = {x1, y1} ;
//   Pce = {x2, y2}
//
// Português:
//
//  Desenhe uma curva Bézier cúbica do ponto atual até o ponto final especificado por x,y. O ponto de controle inicial
//  é especificado por x1,y1 e o ponto de controle final é especificado por x2,y2. Quaisquer tripletos subsequentes de
//  pares de coordenadas são interpretados como parâmetro(s) para comando(s) implícito(s) da curva Bézier cúbica
//  absoluta (C).
//
// Fórmula:
//  Po′ = Pn = {x, y} ;
//  Pcs = {x1, y1} ;
//  Pce = {x2, y2}
func (e *SvgPath) C(x1, y1, x2, y2, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("C %v,%v %v,%v %v,%v ", x1, y1, x2, y2, x, y)
	return e
}

// CubicBezierCurve
//
// English:
//
//  Draw a cubic Bézier curve from the current point to the end point specified by x,y. The start control point is
//  specified by x1,y1 and the end control point is specified by x2,y2. Any subsequent triplet(s) of coordinate pairs
//  are interpreted as parameter(s) for implicit absolute cubic Bézier curve (C) command(s).
//
// Formula:
//   Po′ = Pn = {x, y} ;
//   Pcs = {x1, y1} ;
//   Pce = {x2, y2}
//
// Português:
//
//  Desenhe uma curva Bézier cúbica do ponto atual até o ponto final especificado por x,y. O ponto de controle inicial
//  é especificado por x1,y1 e o ponto de controle final é especificado por x2,y2. Quaisquer tripletos subsequentes de
//  pares de coordenadas são interpretados como parâmetro(s) para comando(s) implícito(s) da curva Bézier cúbica
//  absoluta (C).
//
// Fórmula:
//  Po′ = Pn = {x, y} ;
//  Pcs = {x1, y1} ;
//  Pce = {x2, y2}
func (e *SvgPath) CubicBezierCurve(x1, y1, x2, y2, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("C %v,%v %v,%v %v,%v ", x1, y1, x2, y2, x, y)
	return e
}

// Cd (CubicBezierCurveDelta)
//
// English:
//
//  Draw a cubic Bézier curve from the current point to the end point, which is the current point shifted by dx along
//  the x-axis and dy along the y-axis. The start control point is the current point (starting point of the curve)
//  shifted by dx1 along the x-axis and dy1 along the y-axis. The end control point is the current point (starting point
//  of the curve) shifted by dx2 along the x-axis and dy2 along the y-axis. Any subsequent triplet(s) of coordinate
//  pairs are interpreted as parameter(s) for implicit relative cubic Bézier curve (c) command(s).
//
// Formula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pcs = {xo + dx1, yo + dy1} ;
//  Pce = {xo + dx2, yo + dy2}
//
// Português:
//
//  Desenhe uma curva de Bézier cúbica do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao longo
//  do eixo x e dy ao longo do eixo y. O ponto de controle inicial é o ponto atual (ponto inicial da curva) deslocado
//  por dx1 ao longo do eixo x e dy1 ao longo do eixo y. O ponto de controle final é o ponto atual (ponto inicial da
//  curva) deslocado por dx2 ao longo do eixo x e dy2 ao longo do eixo y. Quaisquer tripletos subsequentes de pares de coordenadas são interpretados como parâmetro(s) para o(s) comando(s) implícito(s) da curva de Bézier cúbica relativa (c).
//
// Fórmula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pcs = {xo + dx1, yo + dy1} ;
//  Pce = {xo + dx2, yo + dy2}
func (e *SvgPath) Cd(dx1, dy1, dx2, dy2, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("c %v,%v %v,%v %v,%v ", dx1, dy1, dx2, dy2, dx, dy)
	return e
}

// CubicBezierCurveDelta
//
// English:
//
//  Draw a cubic Bézier curve from the current point to the end point, which is the current point shifted by dx along
//  the x-axis and dy along the y-axis. The start control point is the current point (starting point of the curve)
//  shifted by dx1 along the x-axis and dy1 along the y-axis. The end control point is the current point (starting point
//  of the curve) shifted by dx2 along the x-axis and dy2 along the y-axis. Any subsequent triplet(s) of coordinate
//  pairs are interpreted as parameter(s) for implicit relative cubic Bézier curve (c) command(s).
//
// Formula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pcs = {xo + dx1, yo + dy1} ;
//  Pce = {xo + dx2, yo + dy2}
//
// Português:
//
//  Desenhe uma curva de Bézier cúbica do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao longo
//  do eixo x e dy ao longo do eixo y. O ponto de controle inicial é o ponto atual (ponto inicial da curva) deslocado
//  por dx1 ao longo do eixo x e dy1 ao longo do eixo y. O ponto de controle final é o ponto atual (ponto inicial da
//  curva) deslocado por dx2 ao longo do eixo x e dy2 ao longo do eixo y. Quaisquer tripletos subsequentes de pares de coordenadas são interpretados como parâmetro(s) para o(s) comando(s) implícito(s) da curva de Bézier cúbica relativa (c).
//
// Fórmula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pcs = {xo + dx1, yo + dy1} ;
//  Pce = {xo + dx2, yo + dy2}
func (e *SvgPath) CubicBezierCurveDelta(dx1, dy1, dx2, dy2, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("c %v,%v %v,%v %v,%v ", dx1, dy1, dx2, dy2, dx, dy)
	return e
}

// S (SmoothCubicBezier)
//
// English:
//
//  Draw a smooth cubic Bézier curve from the current point to the end point specified by x,y. The end control point is
//  specified by x2,y2. The start control point is a reflection of the end control point of the previous curve command.
//  If the previous command wasn't a cubic Bézier curve, the start control point is the same as the curve starting point
//  (current point). Any subsequent pair(s) of coordinate pairs are interpreted as parameter(s) for implicit absolute
//  smooth cubic Bézier curve (S) commands.
//
// Português:
//
//  Desenhe uma curva Bézier cúbica suave do ponto atual até o ponto final especificado por x,y. O ponto de controle
//  final é especificado por x2,y2. O ponto de controle inicial é um reflexo do ponto de controle final do comando de
//  curva anterior. Se o comando anterior não era uma curva Bézier cúbica, o ponto de controle inicial é o mesmo que o
//  ponto inicial da curva (ponto atual). Quaisquer pares subsequentes de pares de coordenadas são interpretados como
//  parâmetros para comandos implícitos de curva Bézier cúbica suave absoluta (S).
func (e *SvgPath) S(x2, y2, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("S %v,%v %v,%v ", x2, y2, x, y)
	return e
}

// SmoothCubicBezier
//
// English:
//
//  Draw a smooth cubic Bézier curve from the current point to the end point specified by x,y. The end control point is
//  specified by x2,y2. The start control point is a reflection of the end control point of the previous curve command.
//  If the previous command wasn't a cubic Bézier curve, the start control point is the same as the curve starting point
//  (current point). Any subsequent pair(s) of coordinate pairs are interpreted as parameter(s) for implicit absolute
//  smooth cubic Bézier curve (S) commands.
//
// Português:
//
//  Desenhe uma curva Bézier cúbica suave do ponto atual até o ponto final especificado por x,y. O ponto de controle
//  final é especificado por x2,y2. O ponto de controle inicial é um reflexo do ponto de controle final do comando de
//  curva anterior. Se o comando anterior não era uma curva Bézier cúbica, o ponto de controle inicial é o mesmo que o
//  ponto inicial da curva (ponto atual). Quaisquer pares subsequentes de pares de coordenadas são interpretados como
//  parâmetros para comandos implícitos de curva Bézier cúbica suave absoluta (S).
func (e *SvgPath) SmoothCubicBezier(x2, y2, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("S %v,%v %v,%v ", x2, y2, x, y)
	return e
}

// Sd (SmoothCubicBezierDelta)
//
// English:
//
//  Draw a smooth cubic Bézier curve from the current point to the end point, which is the current point shifted by dx
//  along the x-axis and dy along the y-axis. The end control point is the current point (starting point of the curve)
//  shifted by dx2 along the x-axis and dy2 along the y-axis. The start control point is a reflection of the end control
//  point of the previous curve command. If the previous command wasn't a cubic Bézier curve, the start control point is
//  the same as the curve starting point (current point). Any subsequent pair(s) of coordinate pairs are interpreted as
//  parameter(s) for implicit relative smooth cubic Bézier curve (s) commands.
//
// Português:
//
//  Desenhe uma curva de Bézier cúbica suave do ponto atual ao ponto final, que é o ponto atual deslocado por dx ao
//  longo do eixo x e dy ao longo do eixo y. O ponto de controle final é o ponto atual (ponto inicial da curva)
//  deslocado por dx2 ao longo do eixo x e dy2 ao longo do eixo y. O ponto de controle inicial é um reflexo do ponto de
//  controle final do comando de curva anterior. Se o comando anterior não era uma curva Bézier cúbica, o ponto de
//  controle inicial é o mesmo que o ponto inicial da curva (ponto atual). Quaisquer pares subsequentes de pares de
//  coordenadas são interpretados como parâmetro(s) para comandos implícitos de curva(s) Bézier cúbica suave relativa.
func (e *SvgPath) Sd(dx2, dy2, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("s %v,%v %v,%v ", dx2, dy2, dx, dy)
	return e
}

// SmoothCubicBezierDelta
//
// English:
//
//  Draw a smooth cubic Bézier curve from the current point to the end point, which is the current point shifted by dx
//  along the x-axis and dy along the y-axis. The end control point is the current point (starting point of the curve)
//  shifted by dx2 along the x-axis and dy2 along the y-axis. The start control point is a reflection of the end control
//  point of the previous curve command. If the previous command wasn't a cubic Bézier curve, the start control point is
//  the same as the curve starting point (current point). Any subsequent pair(s) of coordinate pairs are interpreted as
//  parameter(s) for implicit relative smooth cubic Bézier curve (s) commands.
//
// Português:
//
//  Desenhe uma curva de Bézier cúbica suave do ponto atual ao ponto final, que é o ponto atual deslocado por dx ao
//  longo do eixo x e dy ao longo do eixo y. O ponto de controle final é o ponto atual (ponto inicial da curva)
//  deslocado por dx2 ao longo do eixo x e dy2 ao longo do eixo y. O ponto de controle inicial é um reflexo do ponto de
//  controle final do comando de curva anterior. Se o comando anterior não era uma curva Bézier cúbica, o ponto de
//  controle inicial é o mesmo que o ponto inicial da curva (ponto atual). Quaisquer pares subsequentes de pares de
//  coordenadas são interpretados como parâmetro(s) para comandos implícitos de curva(s) Bézier cúbica suave relativa.
func (e *SvgPath) SmoothCubicBezierDelta(dx2, dy2, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("S %v,%v %v,%v ", dx2, dy2, dx, dy)
	return e
}

// Q (QuadraticBezierCurve)
//
// English:
//
//  Draw a quadratic Bézier curve from the current point to the end point specified by x,y. The control point is
//  specified by x1,y1. Any subsequent pair(s) of coordinate pairs are interpreted as parameter(s) for implicit absolute
//  quadratic Bézier curve (Q) command(s).
//
//  Formula:
//   Po′ = Pn = {x, y} ;
//   Pc = {x1, y1}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática do ponto atual até o ponto final especificado por x,y. O ponto de controle é
//  especificado por x1,y1. Quaisquer pares subsequentes de pares de coordenadas são interpretados como parâmetro(s)
//  para comando(s) de curva de Bézier (Q) quadrática absoluta implícita.
//
//  Fórmula:
//   Po′ = Pn = {x, y} ;
//   Pc = {x1, y1}
func (e *SvgPath) Q(x1, y1, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("Q %v,%v %v,%v ", x1, y1, x, y)
	return e
}

// QuadraticBezierCurve
//
// English:
//
//  Draw a quadratic Bézier curve from the current point to the end point specified by x,y. The control point is
//  specified by x1,y1. Any subsequent pair(s) of coordinate pairs are interpreted as parameter(s) for implicit absolute
//  quadratic Bézier curve (Q) command(s).
//
//  Formula:
//   Po′ = Pn = {x, y} ;
//   Pc = {x1, y1}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática do ponto atual até o ponto final especificado por x,y. O ponto de controle é
//  especificado por x1,y1. Quaisquer pares subsequentes de pares de coordenadas são interpretados como parâmetro(s)
//  para comando(s) de curva de Bézier (Q) quadrática absoluta implícita.
//
//  Fórmula:
//   Po′ = Pn = {x, y} ;
//   Pc = {x1, y1}
func (e *SvgPath) QuadraticBezierCurve(x1, y1, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("Q %v,%v %v,%v ", x1, y1, x, y)
	return e
}

// Qd (QuadraticBezierCurveDelta)
//
// English:
//
//  Draw a quadratic Bézier curve from the current point to the end point, which is the current point shifted by dx
//  along the x-axis and dy along the y-axis. The control point is the current point (starting point of the curve)
//  shifted by dx1 along the x-axis and dy1 along the y-axis. Any subsequent pair(s) of coordinate pairs are interpreted
//  as parameter(s) for implicit relative quadratic Bézier curve (q) command(s).
//
// Formula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pc = {xo + dx1, yo + dy1}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao longo
//  do eixo x e dy ao longo do eixo y. O ponto de controle é o ponto atual (ponto inicial da curva) deslocado por dx1
//  ao longo do eixo x e dy1 ao longo do eixo y. Quaisquer pares subsequentes de pares de coordenadas são interpretados
//  como parâmetro(s) para o(s) comando(s) da curva de Bézier quadrática relativa implícita (q).
//
// Fórmula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pc = {xo + dx1, yo + dy1}
func (e *SvgPath) Qd(dx1, dy1, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("q %v,%v %v,%v ", dx1, dy1, dx, dy)
	return e
}

// QuadraticBezierCurveDelta
//
// English:
//
//  Draw a quadratic Bézier curve from the current point to the end point, which is the current point shifted by dx
//  along the x-axis and dy along the y-axis. The control point is the current point (starting point of the curve)
//  shifted by dx1 along the x-axis and dy1 along the y-axis. Any subsequent pair(s) of coordinate pairs are interpreted
//  as parameter(s) for implicit relative quadratic Bézier curve (q) command(s).
//
// Formula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pc = {xo + dx1, yo + dy1}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao longo
//  do eixo x e dy ao longo do eixo y. O ponto de controle é o ponto atual (ponto inicial da curva) deslocado por dx1
//  ao longo do eixo x e dy1 ao longo do eixo y. Quaisquer pares subsequentes de pares de coordenadas são interpretados
//  como parâmetro(s) para o(s) comando(s) da curva de Bézier quadrática relativa implícita (q).
//
// Fórmula:
//  Po′ = Pn = {xo + dx, yo + dy} ;
//  Pc = {xo + dx1, yo + dy1}
func (e *SvgPath) QuadraticBezierCurveDelta(dx1, dy1, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("q %v,%v %v,%v ", dx1, dy1, dx, dy)
	return e
}

// T (SmoothQuadraticBezierCurve)
//
// English:
//
//  Draw a smooth quadratic Bézier curve from the current point to the end point specified by x,y. The control point is
//  a reflection of the control point of the previous curve command. If the previous command wasn't a quadratic Bézier
//  curve, the control point is the same as the curve starting point (current point). Any subsequent coordinate pair(s)
//  are interpreted as parameter(s) for implicit absolute smooth quadratic Bézier curve (T) command(s).
//
// Formula:
//  Po′ = Pn = {x, y}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática suave do ponto atual até o ponto final especificado por x,y. O ponto de controle
//  é um reflexo do ponto de controle do comando de curva anterior. Se o comando anterior não for uma curva Bézier
//  quadrática, o ponto de controle é o mesmo que o ponto inicial da curva (ponto atual). Quaisquer pares de coordenadas
//  subsequentes são interpretados como parâmetro(s) para comando(s) de curva de Bézier (T) quadrática suave absoluta
//  implícita.
//
// Formula:
//  Po′ = Pn = {x, y}
func (e *SvgPath) T(x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("T %v,%v ", x, y)
	return e
}

// SmoothQuadraticBezierCurve
//
// English:
//
//  Draw a smooth quadratic Bézier curve from the current point to the end point specified by x,y. The control point is
//  a reflection of the control point of the previous curve command. If the previous command wasn't a quadratic Bézier
//  curve, the control point is the same as the curve starting point (current point). Any subsequent coordinate pair(s)
//  are interpreted as parameter(s) for implicit absolute smooth quadratic Bézier curve (T) command(s).
//
// Formula:
//  Po′ = Pn = {x, y}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática suave do ponto atual até o ponto final especificado por x,y. O ponto de controle
//  é um reflexo do ponto de controle do comando de curva anterior. Se o comando anterior não for uma curva Bézier
//  quadrática, o ponto de controle é o mesmo que o ponto inicial da curva (ponto atual). Quaisquer pares de coordenadas
//  subsequentes são interpretados como parâmetro(s) para comando(s) de curva de Bézier (T) quadrática suave absoluta
//  implícita.
//
// Formula:
//  Po′ = Pn = {x, y}
func (e *SvgPath) SmoothQuadraticBezierCurve(x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("T %v,%v ", x, y)
	return e
}

// Td (SmoothQuadraticBezierCurveDelta)
//
// English:
//
//  Draw a smooth quadratic Bézier curve from the current point to the end point, which is the current point shifted by
//  dx along the x-axis and dy along the y-axis. The control point is a reflection of the control point of the previous
//  curve command. If the previous command wasn't a quadratic Bézier curve, the control point is the same as the curve
//  starting point (current point). Any subsequent coordinate pair(s) are interpreted as parameter(s) for implicit
//  relative smooth quadratic Bézier curve (t) command(s).
//
// Formulae:
//   Po′ = Pn = {xo + dx, yo + dy}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática suave do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao
//  longo do eixo x e dy ao longo do eixo y. O ponto de controle é um reflexo do ponto de controle do comando de curva
//  anterior. Se o comando anterior não for uma curva Bézier quadrática, o ponto de controle é o mesmo que o ponto
//  inicial da curva (ponto atual). Quaisquer pares de coordenadas subsequentes são interpretados como parâmetro(s) para
//  comando(s) de curva de Bézier (t) quadrática suave relativa implícita.
//
// Formulae:
//  Po′ = Pn = {xo + dx, yo + dy}
func (e *SvgPath) Td(dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("t %v,%v ", dx, dy)
	return e
}

// SmoothQuadraticBezierCurveDelta
//
// English:
//
//  Draw a smooth quadratic Bézier curve from the current point to the end point, which is the current point shifted by
//  dx along the x-axis and dy along the y-axis. The control point is a reflection of the control point of the previous
//  curve command. If the previous command wasn't a quadratic Bézier curve, the control point is the same as the curve
//  starting point (current point). Any subsequent coordinate pair(s) are interpreted as parameter(s) for implicit
//  relative smooth quadratic Bézier curve (t) command(s).
//
// Formulae:
//   Po′ = Pn = {xo + dx, yo + dy}
//
// Português:
//
//  Desenhe uma curva Bézier quadrática suave do ponto atual até o ponto final, que é o ponto atual deslocado por dx ao
//  longo do eixo x e dy ao longo do eixo y. O ponto de controle é um reflexo do ponto de controle do comando de curva
//  anterior. Se o comando anterior não for uma curva Bézier quadrática, o ponto de controle é o mesmo que o ponto
//  inicial da curva (ponto atual). Quaisquer pares de coordenadas subsequentes são interpretados como parâmetro(s) para
//  comando(s) de curva de Bézier (t) quadrática suave relativa implícita.
//
// Formulae:
//  Po′ = Pn = {xo + dx, yo + dy}
func (e *SvgPath) SmoothQuadraticBezierCurveDelta(dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("t %v,%v ", dx, dy)
	return e
}

// A (ArcCurve)
//
// English:
//
//  Draw an Arc curve from the current point to the coordinate x,y. The center of the ellipse used to draw the arc is
//  determined automatically based on the other parameters of the command:
//
//   * rx and ry are the two radii of the ellipse;
//   * angle represents a rotation (in degrees) of the ellipse relative to the x-axis;
//   * large-arc-flag and sweep-flag allows to chose which arc must be drawn as 4 possible arcs can be drawn out of the
//     other parameters.
//   * large-arc-flag allows to chose one of the large arc (1) or small arc (0),
//   * sweep-flag allows to chose one of the clockwise turning arc (1) or counterclockwise turning arc (0)
//
// The coordinate x,y becomes the new current point for the next command. All subsequent sets of parameters are
// considered implicit absolute arc curve (A) commands.
//
// Português:
//
//  Desenhe uma curva de arco do ponto atual até a coordenada x,y. O centro da elipse usada para desenhar o arco é
//  determinado automaticamente com base nos outros parâmetros do comando:
//
//   * rx e ry são os dois raios da elipse;
//   * ângulo representa uma rotação (em graus) da elipse em relação ao eixo x;
//   * large-arc-flag e sweep-flag permitem escolher qual arco deve ser desenhado, pois 4 arcos possíveis podem ser
//     desenhados a partir dos outros parâmetros.
//   * large-arc-flag permite escolher um arco grande (1) ou arco pequeno (0),
//   * sweep-flag permite escolher um arco de giro no sentido horário (1) ou arco de giro no sentido anti-horário (0)
//
// A coordenada x,y torna-se o novo ponto atual para o próximo comando. Todos os conjuntos de parâmetros subsequentes
// são considerados comandos implícitos de curva de arco absoluto (A).
func (e *SvgPath) A(rx, ry, angle, largeArcFlag, sweepFlag, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("A %v %v %v %v %v %v,%v ", rx, ry, angle, largeArcFlag, sweepFlag, x, y)
	return e
}

// ArcCurve
//
// English:
//
//  Draw an Arc curve from the current point to the coordinate x,y. The center of the ellipse used to draw the arc is
//  determined automatically based on the other parameters of the command:
//
//   * rx and ry are the two radii of the ellipse;
//   * angle represents a rotation (in degrees) of the ellipse relative to the x-axis;
//   * large-arc-flag and sweep-flag allows to chose which arc must be drawn as 4 possible arcs can be drawn out of the
//     other parameters.
//   * large-arc-flag allows to chose one of the large arc (1) or small arc (0),
//   * sweep-flag allows to chose one of the clockwise turning arc (1) or counterclockwise turning arc (0)
//
// The coordinate x,y becomes the new current point for the next command. All subsequent sets of parameters are
// considered implicit absolute arc curve (A) commands.
//
// Português:
//
//  Desenhe uma curva de arco do ponto atual até a coordenada x,y. O centro da elipse usada para desenhar o arco é
//  determinado automaticamente com base nos outros parâmetros do comando:
//
//   * rx e ry são os dois raios da elipse;
//   * ângulo representa uma rotação (em graus) da elipse em relação ao eixo x;
//   * large-arc-flag e sweep-flag permitem escolher qual arco deve ser desenhado, pois 4 arcos possíveis podem ser
//     desenhados a partir dos outros parâmetros.
//   * large-arc-flag permite escolher um arco grande (1) ou arco pequeno (0),
//   * sweep-flag permite escolher um arco de giro no sentido horário (1) ou arco de giro no sentido anti-horário (0)
//
// A coordenada x,y torna-se o novo ponto atual para o próximo comando. Todos os conjuntos de parâmetros subsequentes
// são considerados comandos implícitos de curva de arco absoluto (A).
func (e *SvgPath) ArcCurve(rx, ry, angle, largeArcFlag, sweepFlag, x, y float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("A %v %v %v %v %v %v,%v ", rx, ry, angle, largeArcFlag, sweepFlag, x, y)
	return e
}

// Ad (ArcCurveDelta)
//
// English:
//
//  Draw an Arc curve from the current point to a point for which coordinates are those of the current point shifted by
//  dx along the x-axis and dy along the y-axis. The center of the ellipse used to draw the arc is determined
//  automatically based on the other parameters of the command:
//
//   * rx and ry are the two radii of the ellipse;
//   * angle represents a rotation (in degrees) of the ellipse relative to the x-axis;
//   * large-arc-flag and sweep-flag allows to chose which arc must be drawn as 4 possible arcs can be drawn out of the
//     other parameters.
//   * large-arc-flag allows a choice of large arc (1) or small arc (0),
//   * sweep-flag allows a choice of a clockwise arc (1) or counterclockwise arc (0)
//
// The current point gets its X and Y coordinates shifted by dx and dy for the next command. All subsequent sets of
// parameters are considered implicit relative arc curve (a) commands.
//
// Português:
//
//  Desenhe uma curva de arco do ponto atual para um ponto para o qual as coordenadas são as do ponto atual deslocado
//  por dx ao longo do eixo x e dy ao longo do eixo y. O centro da elipse usada para desenhar o arco é determinado
//  automaticamente com base nos outros parâmetros do comando:
//
//   * rx e ry são os dois raios da elipse;
//   * ângulo representa uma rotação (em graus) da elipse em relação ao eixo x;
//   * large-arc-flag e sweep-flag permitem escolher qual arco deve ser desenhado, pois 4 arcos possíveis podem ser
//     desenhados a partir dos outros parâmetros.
//   * large-arc-flag permite a escolha de arco grande (1) ou arco pequeno (0),
//   * A bandeira de varredura permite a escolha de um arco no sentido horário (1) ou no sentido anti-horário (0)
//
// O ponto atual obtém suas coordenadas X e Y deslocadas por dx e dy para o próximo comando. Todos os conjuntos de
// parâmetros subsequentes são considerados comandos implícitos de curva de arco relativo (a).
func (e *SvgPath) Ad(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("a %v %v %v %v %v %v,%v ", rx, ry, angle, largeArcFlag, sweepFlag, dx, dy)
	return e
}

// ArcCurveDelta
//
// English:
//
//  Draw an Arc curve from the current point to a point for which coordinates are those of the current point shifted by
//  dx along the x-axis and dy along the y-axis. The center of the ellipse used to draw the arc is determined
//  automatically based on the other parameters of the command:
//
//   * rx and ry are the two radii of the ellipse;
//   * angle represents a rotation (in degrees) of the ellipse relative to the x-axis;
//   * large-arc-flag and sweep-flag allows to chose which arc must be drawn as 4 possible arcs can be drawn out of the
//     other parameters.
//   * large-arc-flag allows a choice of large arc (1) or small arc (0),
//   * sweep-flag allows a choice of a clockwise arc (1) or counterclockwise arc (0)
//
// The current point gets its X and Y coordinates shifted by dx and dy for the next command. All subsequent sets of
// parameters are considered implicit relative arc curve (a) commands.
//
// Português:
//
//  Desenhe uma curva de arco do ponto atual para um ponto para o qual as coordenadas são as do ponto atual deslocado
//  por dx ao longo do eixo x e dy ao longo do eixo y. O centro da elipse usada para desenhar o arco é determinado
//  automaticamente com base nos outros parâmetros do comando:
//
//   * rx e ry são os dois raios da elipse;
//   * ângulo representa uma rotação (em graus) da elipse em relação ao eixo x;
//   * large-arc-flag e sweep-flag permitem escolher qual arco deve ser desenhado, pois 4 arcos possíveis podem ser
//     desenhados a partir dos outros parâmetros.
//   * large-arc-flag permite a escolha de arco grande (1) ou arco pequeno (0),
//   * A bandeira de varredura permite a escolha de um arco no sentido horário (1) ou no sentido anti-horário (0)
//
// O ponto atual obtém suas coordenadas X e Y deslocadas por dx e dy para o próximo comando. Todos os conjuntos de
// parâmetros subsequentes são considerados comandos implícitos de curva de arco relativo (a).
func (e *SvgPath) ArcCurveDelta(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy float64) (ref *SvgPath) {
	e.path += fmt.Sprintf("a %v %v %v %v %v %v,%v ", rx, ry, angle, largeArcFlag, sweepFlag, dx, dy)
	return e
}

// Z (Close)
//
// English:
//
//  Close the current subpath by connecting the last point of the path with its initial point. If the two points are at
//  different coordinates, a straight line is drawn between those two points.
//
//   Notes:
//     * The appearance of a shape closed with ClosePath may be different to that of one closed by drawing a line to
//       the origin, using one of the other commands, because the line ends are joined together (according to the
//       stroke-linejoin setting), rather than just being placed at the same coordinates.
//
// Português:
//
//  Feche o subcaminho atual conectando o último ponto do caminho com seu ponto inicial. Se os dois pontos estiverem em
//  coordenadas diferentes, uma linha reta será traçada entre esses dois pontos.
//
//   Notas:
//     * A aparência de uma forma fechada com ClosePath pode ser diferente daquela fechada desenhando uma linha até a
//       origem, usando um dos outros comandos, porque as extremidades da linha são unidas (de acordo com a
//       configuração stroke-linejoin), em vez de apenas sendo colocado nas mesmas coordenadas.
func (e *SvgPath) Z() (ref *SvgPath) {
	e.path += fmt.Sprintf("Z ")
	return e
}

// Close
//
// English:
//
//  Close the current subpath by connecting the last point of the path with its initial point. If the two points are at
//  different coordinates, a straight line is drawn between those two points.
//
//   Notes:
//     * The appearance of a shape closed with ClosePath may be different to that of one closed by drawing a line to
//       the origin, using one of the other commands, because the line ends are joined together (according to the
//       stroke-linejoin setting), rather than just being placed at the same coordinates.
//
// Português:
//
//  Feche o subcaminho atual conectando o último ponto do caminho com seu ponto inicial. Se os dois pontos estiverem em
//  coordenadas diferentes, uma linha reta será traçada entre esses dois pontos.
//
//   Notas:
//     * A aparência de uma forma fechada com ClosePath pode ser diferente daquela fechada desenhando uma linha até a
//       origem, usando um dos outros comandos, porque as extremidades da linha são unidas (de acordo com a
//       configuração stroke-linejoin), em vez de apenas sendo colocado nas mesmas coordenadas.
func (e *SvgPath) Close() (ref *SvgPath) {
	e.path += fmt.Sprintf("Z ")
	return e
}

// String
//
// English:
//
//  Returns the string formed with the SVG path
//
// Português:
//
//  Retorna a string formada com o path SVG
func (e SvgPath) String() (path string) {
	return e.path
}
