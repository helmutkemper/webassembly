package html

import (
	"fmt"
	"strings"
)

// TransformFunctions
//
// English:
//
// The transform functions defines a list of transform definitions that are applied to an element and the element's
// children.
//
//   Notes:
//     * As of SVG2, transform is a presentation attribute, meaning it can be used as a CSS property. However, be aware
//       that there are some differences in syntax between the CSS property and the attribute. See the documentation for
//       the CSS property transform for the specific syntax to use in that case.
//
// Português:
//
// As funções de transformação definem uma lista de definições de transformação que são aplicadas a um elemento e aos
// filhos do elemento.
//
//   Notes:
//     * A partir do SVG2, transform é um atributo de apresentação, o que significa que pode ser usado como uma
//       propriedade CSS. No entanto, esteja ciente de que existem algumas diferenças na sintaxe entre a propriedade CSS
//       e o atributo. Consulte a documentação da transformação da propriedade CSS para obter a sintaxe específica a ser
//       usada nesse caso.
type TransformFunctions struct {
	data []string
}

// Matrix
//
// English:
//
// The matrix(<a> <b> <c> <d> <e> <f>) transform function specifies a transformation in the form of a transformation
// matrix of six values. matrix(a,b,c,d,e,f) is equivalent to applying the transformation matrix:
//
//     a c e
//   ( b d f )
//     0 0 1
//
// which maps coordinates from a previous coordinate system into a new coordinate system by the following matrix
// equalities:
//
//     x_newCoordSys       a c e       x_prevCoordSys       a*x_prevCoordSys + c*y_prevCoordSys + e
//   ( y_newCoordSys ) = ( b d f )   ( y_prevCoordSys )   ( b*x_prevCoordSys + d*y_prevCoordSys + f )
//          1              0 0 1             1                                 1
//
//   Example:
//     <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
//       <rect x="10" y="10" width="30" height="20" fill="green" />
//
//       <!--
//       In the following example we are applying the matrix:
//       [a c e]    [3 -1 30]
//       [b d f] => [1  3 40]
//       [0 0 1]    [0  0  1]
//
//       which transform the rectangle as such:
//
//       top left corner: oldX=10 oldY=10
//       newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 10 + 30 = 50
//       newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 10 + 40 = 80
//
//       top right corner: oldX=40 oldY=10
//       newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 10 + 30 = 140
//       newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 10 + 40 = 110
//
//       bottom left corner: oldX=10 oldY=30
//       newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 30 + 30 = 30
//       newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 30 + 40 = 140
//
//       bottom right corner: oldX=40 oldY=30
//       newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 30 + 30 = 120
//       newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 30 + 40 = 170
//       -->
//       <rect x="10" y="10" width="30" height="20" fill="red" transform="matrix(3 1 -1 3 30 40)" />
//     </svg>
//
// Português:
//
// A função de transformação matrix(<a> <b> <c> <d> <e> <f>) especifica uma transformação na forma de uma matriz de
// transformação de seis valores. matrix(a,b,c,d,e,f) é equivalente a aplicar a matriz de transformação:
//
//     a c e
//   ( b d f )
//     0 0 1
//
// que mapeia as coordenadas de um sistema de coordenadas anterior em um novo sistema de coordenadas pelas seguintes
// igualdades de matriz:
//
//     x_newCoordSys       a c e       x_prevCoordSys       a*x_prevCoordSys + c*y_prevCoordSys + e
//   ( y_newCoordSys ) = ( b d f )   ( y_prevCoordSys )   ( b*x_prevCoordSys + d*y_prevCoordSys + f )
//          1              0 0 1             1                                 1
//
//   Example:
//     <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
//       <rect x="10" y="10" width="30" height="20" fill="green" />
//
//       <!--
//       In the following example we are applying the matrix:
//       [a c e]    [3 -1 30]
//       [b d f] => [1  3 40]
//       [0 0 1]    [0  0  1]
//
//       which transform the rectangle as such:
//
//       top left corner: oldX=10 oldY=10
//       newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 10 + 30 = 50
//       newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 10 + 40 = 80
//
//       top right corner: oldX=40 oldY=10
//       newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 10 + 30 = 140
//       newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 10 + 40 = 110
//
//       bottom left corner: oldX=10 oldY=30
//       newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 30 + 30 = 30
//       newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 30 + 40 = 140
//
//       bottom right corner: oldX=40 oldY=30
//       newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 30 + 30 = 120
//       newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 30 + 40 = 170
//       -->
//       <rect x="10" y="10" width="30" height="20" fill="red" transform="matrix(3 1 -1 3 30 40)" />
//     </svg>
func (el *TransformFunctions) Matrix(a, b, c, d, e, f float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("matrix(%v %v %v %v %v %v)", a, b, c, d, e, f))
}

// Translate
//
// English:
//
// The translate(<x> [<y>]) transform function moves the object by x and y. If y is not provided, it is assumed to be 0.
//
// In other words:
//
//   xnew = xold + <x>
//   ynew = yold + <y>
//
// Português:
//
// A função transform translate(<x> [<y>]) move o objeto em x e y. Se y não for fornecido, assume-se que é 0.
//
// Em outras palavras:
//
//   xnew = xold + <x>
//   ynew = yold + <y>
func (el *TransformFunctions) Translate(x, y float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("translate(%v %v)", x, y))
}

// Scale
//
// English:
//
// The scale(<x> [<y>]) transform function specifies a scale operation by x and y. If y is not provided, it is assumed
// to be equal to x.
//
// Português:
//
// A função de transformação scale(<x> [<y>]) especifica uma operação de escala por x e y. Se y não for fornecido,
// assume-se que é igual a x.
func (el *TransformFunctions) Scale(x, y float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("scale(%v %v)", x, y))
}

// RotateAngle
//
// English:
//
// The rotate(<a>) transform function specifies a rotation by a degrees about a given point.
//
// The rotation is about the origin of the current user coordinate system.
//
// Português:
//
// A função de transformação rotate(<a>) especifica uma rotação de um grau em torno de um determinado ponto.
//
// A rotação é sobre a origem do sistema de coordenadas do usuário atual.
func (el *TransformFunctions) RotateAngle(a float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("rotate(%v)", a))
}

// RotateAngleXCoord
//
// English:
//
// The rotate(<a>) transform function specifies a rotation by a degrees about a given point.
//
// The rotation is about the supplied x and the origin of the current point y.
//
// Português:
//
// A função de transformação rotate(<a>) especifica uma rotação de um grau em torno de um determinado ponto.
//
// A rotação é sobre o x fornecido e a origem do ponto atual y.
func (el *TransformFunctions) RotateAngleXCoord(a, x float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("rotate(%v,%v)", a, x))
}

// Rotate
//
// English:
//
// The rotate(<a>) transform function specifies a rotation by a degrees about a given point.
//
// The rotation is about the point (x, y).
//
// Português:
//
// A função de transformação rotate(<a>) especifica uma rotação de um grau em torno de um determinado ponto.
//
// A rotação é em torno do ponto (x, y).
func (el *TransformFunctions) Rotate(a, x, y float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("rotate(%v,%v,%v)", a, x, y))
}

// SkewX
//
// English:
//
// The skewX(<a>) transform function specifies a skew transformation along the x axis by a degrees.
//
// Português:
//
// A função de transformação skewX(<a>) especifica uma transformação de inclinação ao longo do eixo x em um grau.
func (el *TransformFunctions) SkewX(a float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("skewX(%v)", a))
}

// SkewY
//
// English:
//
// The skewY(<a>) transform function specifies a skew transformation along the y axis by a degrees.
//
// Português:
//
// A função de transformação skewY(<a>) especifica uma transformação de inclinação ao longo do eixo y em um grau.
func (el *TransformFunctions) SkewY(a float64) {
	if el.data == nil {
		el.data = make([]string, 0)
	}

	el.data = append(el.data, fmt.Sprintf("skewY(%v)", a))
}

func (el TransformFunctions) String() string {
	return strings.Join(el.data, "\n")
}
