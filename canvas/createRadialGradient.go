package canvas

// en: Creates a radial gradient (to use on canvas content). The parameters represent two circles, one with its center at (x0, y0) and a radius of r0, and the other with its center at (x1, y1) with a radius of r1.
//     x0: The x-coordinate of the starting circle of the gradient
//     y0: The y-coordinate of the starting circle of the gradient
//     r0: The radius of the starting circle. Must be non-negative and finite. (note: radius is a width, not a degrees angle)
//     x1: The x-coordinate of the ending circle of the gradient
//     y1: The y-coordinate of the ending circle of the gradient
//     r1: The radius of the ending circle. Must be non-negative and finite. (note: radius is a width, not a degrees angle)
//
// pt_br: Este método cria um gradiente radial (para ser usado com o canvas 2D). Os parâmetros representam dois círculos, um com o centro no ponto (x0, y0) e raio r0, e outro com centro no ponto (x1, y1) com raio r1
//     x0: Coordenada x do circulo inicial do gradiente
//     y0: Coordenada y do circulo inicial do gradiente
//     r0: Raio do círculo inicial. Deve ser um valor positivo e finito. (nota: o raio é um comprimento e não um ângulo)
//     x1: Coordenada x do circulo final do gradiente
//     y1: Coordenada y do circulo final do gradiente
//     r1: Raio do círculo final. Deve ser um valor positivo e finito. (nota: o raio é um comprimento e não um ângulo)
func (el *Canvas) CreateRadialGradient(x0, y0, r0, x1, y1, r1 interface{}) interface{} {
	return el.SelfContext.Call("createRadialGradient", x0, y0, r0, x1, y1, r1)
}
