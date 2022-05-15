package algorithm

import "math"

type Geometry struct {
	original  []Point
	processed []Point
}

// DistanceBetweenTwoPoints
//
// English:
//
//  Calculates the distance between two points.
//
//   Input:
//     p1, p2: Start and end points of the line.
//
//   Output:
//     d: distance between p1 and p2 in pixels.
//
// Português:
//
//  Calcula a distância entre dois pontos.
//
//   Entrada:
//     p1, p2: Pontos inicial e final da reta.
//
//   Saída:
//     d: distância entre p1 e p2 em pixels.
func (e *Geometry) DistanceBetweenTwoPoints(p0, p1 Point) (d float64) {
	return math.Sqrt(math.Abs(math.Pow(p1.X-p0.X, 2) + math.Pow(p1.Y-p0.Y, 2)))
}

func (e *Geometry) NewPointByDistance(point Point, distance, angle float64) Point {
	point.X = point.X + (distance * math.Cos(angle))
	point.Y = point.Y + (distance * math.Sin(angle))

	return point
}

func (e *Geometry) AngleBetweenTwoPoints(p1, p2 Point) (angle float64) {
	angle = math.Atan2(p1.Y-p2.Y, p1.X-p2.X)
	return
}

// Round
//
// English:
//
//
//
//
//
// Português:
//
//  Arredonda a quantidade máxima de casas decimais de um número de ponto flutuante.
//
//   Entrada:
//     value: número a ser arredondado;
//     places: quantidade de dígitos decimais.
//
//   Notes:
//     * Se o dígito seguinte a places for menor que 5, o dígito place será arredondado para baixo.
//       Exemplo: Round(3.1415, 2) = 3.14
//                Round(3.1415, 3) = 3.142
func (e *Geometry) Round(value float64, places int) float64 {
	var roundOn = 0.5

	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * value
	_, div := math.Modf(digit)

	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	return round / pow
}

// PointBetweenTwoPoints
//
// English:
//
//  Returns a point between two points from the distance of P1.
//
//   Input:
//     p1, p2: start and end point of the line;
//     distance: distance of the new point from p1.
//
//   Output:
//     p: midpoint between p1 and p2.
//
// Português:
//
//  Devolve um ponto entre dois pontos a partir da distância de P1.
//
//   Entrada:
//     p1, p2: ponto inicial e final da reta;
//     distance: distância do novo ponto em relação a p1.
//
//   Saída:
//     p: ponto intermediário entre p1 e p2.
func (e *Geometry) PointBetweenTwoPoints(p1, p2 Point, distance float64) (p Point) {
	a := math.Atan2(p1.Y-p2.Y, p1.X-p2.X)
	x := e.Round(p1.X-math.Cos(a)*distance, 2)
	y := e.Round(p1.Y-math.Sin(a)*distance, 2)
	p = Point{X: x, Y: y}
	return
}
