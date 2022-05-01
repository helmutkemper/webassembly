package easingTween

import "math"

// en: quadratic easing in - accelerating from zero velocity
var KEaseInQuadratic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return delta*math.Pow(currentPercentage, 2.0) + startValue
}
