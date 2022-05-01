package easingTween

import "math"

// en: circular easing out - decelerating to zero velocity
var KEaseOutCircular = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	currentPercentage--
	return delta*math.Sqrt(1-math.Pow(currentPercentage, 2.0)) + startValue
}
