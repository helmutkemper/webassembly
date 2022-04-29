package tween

import "math"

// en: quartic easing out - decelerating to zero velocity
var KEaseOutQuartic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	currentPercentage -= 1
	return -1*delta*(math.Pow(currentPercentage, 4.0)-1) + startValue
}
