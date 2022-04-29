package tween

import "math"

// en: sinusoidal easing out - decelerating to zero velocity
var KEaseOutSine = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return delta*math.Sin(currentPercentage*1.570796) + startValue
}
