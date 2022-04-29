package tween

import "math"

// en: Cubic easing in - accelerating from zero velocity
var KEaseInCubic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return delta*math.Pow(currentPercentage, 3.0) + startValue
}
