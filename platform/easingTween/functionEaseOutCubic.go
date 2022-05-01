package easingTween

import "math"

// en: cubic easing out - decelerating to zero velocity
var KEaseOutCubic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	currentPercentage--
	return delta*(math.Pow(currentPercentage, 3.0)+1) + startValue
}
