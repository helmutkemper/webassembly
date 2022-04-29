package tween

import (
	"math"
)

// en: exponential easing in - accelerating from zero velocity
var KEaseInExponential = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return delta*math.Pow(2, 10*(currentPercentage-1)) + startValue
}
