package tween

import "math"

// en: circular easing in - accelerating from zero velocity

// KEaseInCircular
//
// English:
//
//  Circular easing in, accelerating from zero velocity
//
// Português:
//
//  Circular easing in, acelerando do zero até a velocidade
var KEaseInCircular = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return -delta*(math.Sqrt(math.Abs(1-math.Pow(currentPercentage, 2.0)))-1) + startValue
}
